package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//Функция, которая будет принимать данные с канала.
func worker(ch <-chan int, i int){
	for v := range ch{
		//Обрабатывает значение, полученное с канала
		fmt.Printf("%d worker, get %d\n", i, v)
	}
	//Если канал закроется, то воркер завершает свою работу
	fmt.Printf("%d worker завершил работу\n", i)
}

func worker2(ctx context.Context, ch <-chan int, i int){
	for {
		//Ждем, когда какой-нибудь канал не будет заблокирован
		select{
			//Если пришло значение с какого-нибудь канала
		case value, ok := <- ch:
			if !ok{
				fmt.Printf("%d worker2 завершил работу\n", i) 
				return
			}
			fmt.Printf("%d worker2, get %d\n", i, value)
			//Если отменился контекст
		case <-ctx.Done():
			fmt.Printf("%d worker2 завершил работу\n", i) 
			return
		}
	}
}

func main(){
	//К-во воркеров
	workerCount := 5
	//Cоздаю канал, который будут прослушивать воркеры
	ch1 := make(chan int)
	ch2 := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}
	for i := 0; i < workerCount; i++{
		wg.Add(2)
		//Горутины с воркерами, которые будут выполняться
		//Здесь нужны WaitGroup, чтобы успели завершиться все воркеры
		go func(i int){
			worker(ch1, i)
			wg.Done()
		}(i)

		go func(i int){
			worker2(ctx, ch2, i)
			wg.Done()
		}(i)
	}

	//Запускаю горутину, которая будет записывать данные в канал
	go func(){
		for i := 0; i < 100; i++{
			ch1 <- i
			ch2 <- i
			time.Sleep(1 * time.Second)
		}
	}()

	sig := make(chan os.Signal, 1)
    signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	//Блокируем главную горутину, пока не будет получен сигнал. В нашем случае ctrl+c
	/*
	Это нужно для того, чтобы программа завершила все дополнительные горутины(не прерывать же работу воркеров неожиданно), а только затем завершилась основная.
	К тому же в воркерах у нас могут быть открыты файлы, соединения, которые требуют их закрытия, а если не 
	*/
    <-sig

	close(ch1)
	close(ch2)
	cancel()
	
	wg.Wait()

	fmt.Println("Все воркеры завершили работу")
}