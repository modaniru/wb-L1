package main

import "fmt"

//В качестве значения struct, потому что он занимает 0 байт
func main(){
	s1 := map[int]struct{}{
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
		6: {},
	}
	s2 := map[int]struct{}{
		4: {},
		5: {},
		6: {},
		7: {},
		8: {},
		9: {},
	}
	fmt.Println(setIntersection(
		s1, s2,
	))
}

func setIntersection(set1, set2 map[int]struct{}) map[int]struct{}{
	result := make(map[int]struct{})
	//Если во втором сете есть елемент этого сета, то пишем в результирующее множество
	//Сложность
	//Время: О(n), обращение к мапе по ключу = O(1)
	//Память: O(n)
	for k := range set1{
		if _, ok := set2[k]; ok{
			result[k] = struct{}{}
		}
	}
	return result
}