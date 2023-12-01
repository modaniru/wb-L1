package main

import (
	"fmt"
	"math/rand"
)

func main(){
	arr := []int{2, 10, 3, 2, 4, 6, 5, 7}
	mySort(arr)
	fmt.Println(arr)
}

//Функция сортировка, которая выполняется рекурсивно
func mySort(arr []int){
	if len(arr) <= 1{
		return
	}
	pivotIndex := rand.Intn(len(arr))
	pivotIndex = move(arr, pivotIndex)
	mySort(arr[:pivotIndex])
	mySort(arr[pivotIndex + 1:])
}

//Функция которая перемещает элементы относительно pivot
func move(arr []int, pivotIndex int) int {
	i := -1
	j := 0
	//Перемещаю его в конец, чтобы было удобнее перемещать
	arr[pivotIndex], arr[len(arr) - 1] = arr[len(arr) - 1], arr[pivotIndex]
	pivotIndex = len(arr) - 1
	for ;j < pivotIndex; j++{
		if !(arr[j] > arr[pivotIndex]){
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	i++
	arr[i], arr[j] = arr[j], arr[i]
	return i
}