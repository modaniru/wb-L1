package main

var justBytes []byte

func someFunc() {
	//получаем строку
	v := createHugeString(1 << 10)
	//Создаем массив байтов, длинны которую хотим получить
	justBytes = make([]byte, 100)
	//Копируем первые 100 байтов с строки v
	//Если бы мы присволи строку к строке, то это привело бы к потенциальной утечке памяти
	//Т.к строки неизменяемые мы так бы ссылались на первые 100 символов строки v.
	//Присвоили бы их justString и v бы не удалилась gb.
	copy(justBytes, []byte(v))
}

func main() {
	someFunc()
}

func createHugeString(len int) string{
	return "test"
}