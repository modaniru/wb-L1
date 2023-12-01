package main

import (
	"fmt"
)

func main(){
	//Канал для записи
	writer := make(chan int, 5)
	//Канал для обработанных данных
	reciever := make(chan int, 5)

	//Горутина, которая записывает данные в канал для записи
	go func(writer chan<- int){
		for i := 0; i < 15; i++{
			fmt.Printf("Конвеер записи: %d\n", i)
			// time.Sleep(500 * time.Millisecond)
			writer <- i
		}
		//Закрывает канал
		close(writer)
		fmt.Println("Записывающая горутина, и канал для записи - остановлены")
	}(writer)

	//Горутина, которая обрабатывает данные от канала для записи. Обработанные данные записывает в канал для резултатов
	go func(sender <-chan int, result chan<- int){
		for v := range sender{
			fmt.Printf("Конвеер обработки: %d\n", v * v)
			result <- v * v
		}
		//Закрывает канал
		close(result)
		fmt.Println("Обрабатывающая горутина, и канал для записи результата - остановлены")
	}(writer, reciever)

	//Проходимся по всем значениям в канале, пока он открыт
	for v := range reciever{
		fmt.Printf("Выход: %d\n", v)
	}
}