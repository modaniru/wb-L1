package main

import "fmt"

//Интерфейс собаки
type Dog interface{
	Bark()
}

//Интерфейс кошки
type Cat interface{
	Meow()
}

//Структура(класс) манула
type Manul struct{

}

func(m *Manul) Meow(){
	fmt.Println("Manul: meow-meow")
}

//Структура(Класс) адаптера кошки под собаку
type DogAdapter struct{
	//содержит экземпляр какой-то кошки
	cat Cat
}

//Метод, чтобы сделать нашу структуру "Cобакой"
//Таким образом мы сделаем любую кошку - собакой
func (d *DogAdapter) Bark(){
	fmt.Println("\"dog\":")
	fmt.Printf("\t")
	d.cat.Meow()
}

func main(){
	//Манул
	Manul := Manul{}
	//Встраиваем в адаптер
	var dog Dog = &DogAdapter{&Manul}
	Manul.Meow()
	//Можем вызвать у даптера метод интерфейса Dog
	dog.Bark()
}