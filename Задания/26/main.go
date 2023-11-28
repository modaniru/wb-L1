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
	fmt.Println("result:", isUniqueLetters(input[:len(input) - 1]))
}

//Время O(n)
//память O(2n)
//Сделал преобразование в нижний регистр для учета и английских и русских символов.
func isUniqueLetters(str string) bool{
	str = strings.ToLower(str)
	dict := map[rune]bool{}
	for _, r := range str{
		if dict[r]{
			return false
		}
		dict[r] = true
	}
	return true
}