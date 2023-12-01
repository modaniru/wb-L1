package main

import "fmt"

func main(){
	//Используем int64, потому что иначе произойдет переполнение
	var n1 int64 = 2 << 22
	var n2 int64 = 1 << 20

	fmt.Println(n1 * n2)

	fmt.Println(n1 + n2)

	fmt.Println(n1 - n2)

	fmt.Println(n1 / n2)
}