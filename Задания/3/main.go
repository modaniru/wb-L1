package main

import (
	"fmt"
	"sync"
)

func main(){
	fmt.Println("With wg")
	WithWaitGroup()
	fmt.Println("With channel")
	WithChannel()
}

func WithWaitGroup(){
	arr := []int{2,4,6,8,10}
	count := 0
	wg := sync.WaitGroup{}
	for _, v := range arr{
		wg.Add(1)
		go func(value int){
			defer wg.Done()
			count += value * value
		}(v)
	}
	//Ждем когда счетчик у wg будет равен 0
	wg.Wait()
	fmt.Println(count)
}

func WithChannel(){
	arr := []int{2,4,6,8,10}
	count := 0
	ch := make(chan int, len(arr))
	go func ()  {
		for _, v := range arr{
			ch <- v * v
		}
		close(ch)
	}()
	for v := range ch{
		count+=v
	}
	fmt.Println(count)
}