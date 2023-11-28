package main

import (
	"fmt"
	"sync"
)

func main(){
	fmt.Println("With WaitGroup")
	WithWG()
	fmt.Println("\nWith channels")
	WithChannels()
	fmt.Println("\nWith buffered channels")
	WithChannels()
}

//Выводит квадраты в хаотичном порядке. Используется WaitGroup, чтобы ждать выполнение всех горутин.
func WithWG(){
	arr := []int{2,4,6,8,10}
	//Создаем WG
	wg := sync.WaitGroup{}
	for _, n := range arr{
		//+1 к счетчику у wg
		wg.Add(1)
		/*
		Создаем отдельную копию v, потому что иначе горутину будут обращаться к одному участку памяти v(н.р. цикл пройдет v станет 10).
		*/
		value := n
		go func ()  {
			fmt.Println(value * value)
			//-1 к счетчику у wg
			wg.Done()
		}()
	}
	//Ждем, когда счетчик у wg станет равным 0
	wg.Wait()
}

func WithChannels(){
	arr := []int{2,4,6,8,10}
	//Создаем канал, в который будем складывать тип int
	ch := make(chan int)
	for _, v := range arr{
		//Либо можно так предоставить v
		go func (value int)  {
			//Записываем канал, блокируем пока не прочитает
			ch <- value * value
		}(v)
	}
	for i := 0; i < len(arr); i++{
		//Ждем пока кинут что-то в канал, снимаем блокировку
		fmt.Println(<-ch)
	}
}

func WithBufferedChannels(){
	arr := []int{2,4,6,8,10}
	//Создаем буферизированный канал, в который будем складывать тип int
	ch := make(chan int, len(arr))
	for _, v := range arr{
		//Либо можно так предоставить v
		go func (value int)  {
			//Записываем канал, если к-во записей больше чем длина буфера канала, блокируем
			ch <- value * value
		}(v)
	}
	for i := 0; i < len(arr); i++{
		//Ждем пока кинут что-то в канал
		fmt.Println(<-ch)
	}
	//Тут закрывать канал нигде не получится, потому что у нас в циклах горутины, а не цикл в горутине
}