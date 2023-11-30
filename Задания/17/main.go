package main

import (
	"fmt"
	"sort"
)

func main(){
	//Чтобы искать бинарным поиском, массив должен быть отсортирован
	array := []int{1,2,2,3,3,3,4,5,6,6,6,6,7,8,9,10}
	sort.Ints(array)
	for i := -2; i < 12; i++{
		res, ok := binarySort(array, i)
		fmt.Println(i, "->", res, ok)
	}
}

//Сложность Время: O(logn) Память: O(1)
//Эта реализцаия бинарного поиска выдает индекс первого вхождения элемента в массиве
func binarySort(array []int, num int) (int, bool){
	left := 0
	right := len(array)
	for left < right{
		mid := (left + right) / 2
		if array[mid] >= num{
			right = mid
		} else if array[mid] < num{
			left = mid + 1
		}
	}
	//Крайний случай, когда элемент больше максимального, а left > len(array)
	if left < len(array) && array[left] == num{
		return left, true
	}
	return 0, false
}