package main

import "fmt"

func main(){
	fmt.Println(remove([]int{1,2,3,4}, 3))
}

func remove(arr []int, index int) []int{
	if index > len(arr) || index < 0{
		return nil
	}
	//Когда мы делаем arr[:index], его капасити равно index, затем мы пытаемся добавить еще элементы из-за этого выделяется новый участок памяти, поэтому изначальный массив не меняется 
	return append(arr[:index], arr[index+1:]...)
}