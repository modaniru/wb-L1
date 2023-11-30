package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//TODO
func main(){
	second()
}

//Первый способ, не самый лучший, потому что программа завершится резко
func first(){
	ch := make(chan string, 5)
	go func(ch chan<- string){
		for i := 0; i <= 100; i++{
			ch <- fmt.Sprintf("Message: %d", i)
			time.Sleep(500 * time.Millisecond)
		}
		close(ch)
	}(ch)

	go func(ch <-chan string){
		for value := range ch{
			fmt.Printf("GetMessage: %s\n", value)
		}
	}(ch)

	time.Sleep(10 * time.Second)
}

//Второй вариант
func writer(ctx context.Context, writeChannel chan<- string){
	//Будем посылвать в канал 100 строк
	for i := 0; i <= 100; i++{
		select{
		//Если контекст сверху был отменен, то завершаем работу горутины
		case <-ctx.Done():
			fmt.Println("context close")
			return
		//Иначе записываем в канал
		default:
			writeChannel <- fmt.Sprintf("Message: %d", i)
			time.Sleep(500 * time.Millisecond)
		}
	}
	fmt.Println("writer done")
}

func reciever(ctx context.Context, ch <-chan string, done chan<- bool){
	for {
		select{
		//Если контекст сверху был отменен, то завершаем работу горутины.
		case <-ctx.Done():
			fmt.Println("context close")
			return
		//Если все еще есть элементы в канале, обрабатываем. (Возник вопрос по поводу того, что если контекст отменилсяя, а данные в канале все еще есть) В моем примере мы это проигнорируем и не будем обрабатывать те данные. Можно обрабатвать все, а потом уже по закрытому каналу закрывать горутину.
		case value, ok := <- ch:
			if !ok{
				fmt.Println("channel was close")
				done <- true
				return
			}
			fmt.Printf("GetMessage: %s\n", value)
		}
	}
}

func second(){
	//Контекст с таймаутов в 10 секунд
	context, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	//Отменяем
	defer cancel()
	//Записывающий канал
	writerChannel := make(chan string, 10)
	//Канал о том, что reciever завершил свою работу досрочно
	done := make(chan bool)
	//Отслеживаем чтобы все горутины выполнились до конца
	wg := sync.WaitGroup{}

	go func ()  {
		wg.Add(1)
		writer(context, writerChannel)
		close(writerChannel)
		wg.Done()
	}()
	go func(){
		wg.Add(1)
		reciever(context, writerChannel, done)
		wg.Done()
	}()

	//Если время проходит, контекст закрывается, закрывается канал, ждем довыполнение всех горутин
	select{
	case <-context.Done():
		wg.Wait()
	//Если получатель завершает работу до таймера, отменяем контекст
	case <-done:
		cancel()
	}
}