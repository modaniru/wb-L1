package main

import (
	"fmt"
	"sync"
)

type Counter struct{
	count int
	//Mutex, чтобы не возникало ситуации "гонки"
	mutex sync.Mutex
}

func (c *Counter) Increment(){
	//В момент выполнения этой функции, только эта функция будет иметь доступ к переменной
	c.mutex.Lock()
	//Снимаем блокировку в конце выполнения функции
	defer c.mutex.Unlock()
	c.count++
}

func main(){
	wg := sync.WaitGroup{}
	counter := Counter{count: 0, mutex: sync.Mutex{}}

	for i := 0; i < 10000; i++{
		wg.Add(1)
		go func ()  {
			counter.Increment()
			wg.Done()
		}()
	}
	//Ждем выполнение всех горутин
	wg.Wait()
	fmt.Println(counter.count)
}