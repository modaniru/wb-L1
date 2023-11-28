package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	var input string
	in := bufio.NewReader(os.Stdin)
	//Читаем, пока не введут '\n' символ
	fmt.Print("Введите строку: ")
	input, _ = in.ReadString('\n')
	//Убираем '\n' символ в конце
	fmt.Printf("result: %s\n", reverse(input[:len(input) -1]))
}

func reverse(str string) string{
	//C рунами удобнее работать, и лучше. Напрямую с строками лучше не работать, так как они immutable
	runes := []rune(str)
	/*
	Асимтотика по времени O(n/2)
	по памяти O(2n), так как создаем массив рун, а после возвращаем строку.
	Можно использовать Slices.Reverse(), но решил написать логику 'переворачивания' сам
	*/
	for i := 0; i < len(runes) / 2; i++{
		runes[i], runes[len(runes) - 1 - i] = runes[len(runes) - 1 - i], runes[i]
	}
	return string(runes)
}