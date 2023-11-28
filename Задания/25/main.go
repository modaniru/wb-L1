package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println(time.Now())
	MySleep(time.Second * 1)
	fmt.Println(time.Now())
	fmt.Println(time.Now())
	time.Sleep(time.Second * 1)
	fmt.Println(time.Now())
}

//TODO
//мб какое-то еще решение есть
//Погрешность 0,000200 мс
//Эта идея пришла самой первой в голову
func MySleep(duration time.Duration){
	//Берем текущее время, прибавляем duration.
	start := time.Now().Add(duration)
	for {
		//Ждем пока текущее время не обгонит наш start
		if time.Now().After(start){
			break
		}
	}
}