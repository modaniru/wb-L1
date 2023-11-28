package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	var input string
	in := bufio.NewReader(os.Stdin)
	//Читаем, пока не введут '\n' символ
	fmt.Print("Введите строку: ")
	input, _ = in.ReadString('\n')
	//Убираем '\n' символ в конце
	fmt.Printf("result: %s\n", reverseWords(input[:len(input) - 1]))
}

//Время O(n)
//Память O(2n)
func reverseWords(str string) string{
	//Делим строку, на массив строк.
	strs := strings.Split(str, " ")
	//Переворачиваем аналогично 19 задаче
	for i := 0; i < len(strs) / 2; i++{
		strs[i], strs[len(strs) - 1 - i] = strs[len(strs) - 1 - i], strs[i]
	}
	//Соединяем
	return strings.Join(strs, " ")
}