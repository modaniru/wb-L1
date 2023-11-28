package main

import (
	"fmt"
	"strconv"
)

func main(){
	//Переменная, которую будем менять
	var num int64 = 0
	//Ввод пользователя
	var input string
	fmt.Println("что-то кроме числа - выход.")
	fmt.Printf("num: %064b\n", num)
	for {
		fmt.Scan(&input)
		//Если строка - выход
		inputNum, err := strconv.Atoi(input)
		if err != nil{
			fmt.Println("Выход...")
			return
		}
		/*
		Cдивигаем 1 на n шагов и xor'им наше число;
		1 xor 1 = 0
		0 xor 1 = 1
		*/
		v := int64(1<<inputNum)
		num = num ^  v
		fmt.Printf("num: %064b = %d\n", uint64(num), num)
	}
}