package main

import "fmt"

func main(){
	array := []string{"cat", "cat", "dog", "cat", "tree"}

	//Наше множество
	res := map[string]struct{}{}

	//Достаточно записать все данные в наш аналог сета, после чего на выходе получаем наше множество значений
	/*
	Время: O(n)
	Память: O(n)
	*/
	for _, v := range array {
		res[v] = struct{}{}
	}
	fmt.Println(res)
}