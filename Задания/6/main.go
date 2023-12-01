package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)


func withChannel(done <-chan struct{}){
	for {
		select {
		case <- done:
			fmt.Println("withChannel exit")
			return
		default:
			fmt.Println("withChannel...")
			time.Sleep(1 * time.Second)
		}
	}
}

func withTimeout(){
	select {
	case <-time.After(time.Second * 5):
		fmt.Println("withTimeout exit...")
	default:
		fmt.Println("Долгая операция...")
		time.Sleep(6 * time.Second)
	}
}

func withContext(ctx context.Context){
	for {
		select {
		case <- ctx.Done():
			fmt.Println("withContext exit")
			return
		default:
			fmt.Println("withContext...")
			time.Sleep(1 * time.Second)
		}
	}
}

func main(){
	{
		ch := make(chan struct{})
		wg := sync.WaitGroup{}

		wg.Add(1)
		go func(){
			withChannel(ch)
			wg.Done()
		}()

		time.Sleep(5 * time.Second)
		ch <- struct{}{}

		wg.Wait()
		fmt.Println("Выход...")
	}

	{
		context, cancel := context.WithTimeout(context.Background(), time.Second * 5)
		defer cancel()
		wg := sync.WaitGroup{}

		wg.Add(1)
		go func(){
			withContext(context)
			wg.Done()
		}()

		wg.Wait()
		fmt.Println("Выход...")
	}
	{
		wg := sync.WaitGroup{}

		wg.Add(1)
		go func(){
			withTimeout()
			wg.Done()
		}()

		wg.Wait()
		fmt.Println("Выход...")
	}
}