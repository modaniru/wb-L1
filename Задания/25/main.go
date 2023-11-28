package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println(time.Now())
	MySleep(time.Second * 1)
	fmt.Println(time.Now())
	fmt.Println("-----------")
	fmt.Println(time.Now())
	MySleep2(time.Second * 1)
	fmt.Println(time.Now())
	fmt.Println("-----------")
	fmt.Println(time.Now())
	time.Sleep(time.Second * 1)
	fmt.Println(time.Now())

	/*
	РЕЗЮМЕ:
	вывод:
	2023-11-29 00:11:33.781312 +0500 +05 m=+0.000343960
	2023-11-29 00:11:34.781582 +0500 +05 m=+1.000620585
	-----------
	2023-11-29 00:11:34.781601 +0500 +05 m=+1.000639876
	2023-11-29 00:11:35.782687 +0500 +05 m=+2.001732751
	-----------
	2023-11-29 00:11:35.782858 +0500 +05 m=+2.001904001
	2023-11-29 00:11:36.783984 +0500 +05 m=+3.003036793

	функция через обычный for, работает точнее, чем функция через time.After и обычный time.Sleep.
	Есть догадки, что это происходит из-за использования каналов
	Но через for, думаю, хуже в производительности
	*/
}

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

//Нашел одно решение, возвращает канал, который блочится на duration.
//До этого не знал о функции After в пакете(функция структуры не в счет)
//Это решение выглядит куда более красивее, чем первое
func MySleep2(duration time.Duration){
	<- time.After(duration)
}