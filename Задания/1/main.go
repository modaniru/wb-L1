package main

import "fmt"

//Родительская структура
type Human struct{
	name string
	surname string
	age int
}
//Ниже написаны функции для Human'a
func (h *Human) getName() string{
	return h.name
}

func (h *Human) getSurname() string{
	return h.surname
}

func (h *Human) getAge() int{
	return h.age
}

func (h *Human) getInfo() string{
	return fmt.Sprintf("%s %s %d y.o", h.name, h.surname, h.age)
}

//Дочерняя структура
type Action struct{
	//Встраивание (аналог наследования) в go происходит следующим образом
	Human
	//Можем дополнять дочернюю структуру
	catName string
}

// 'Перекрывать(думаю тут этот термин больше подходит, чем переопределять)' родительские методы
func (a *Action) getInfo() string{
	return fmt.Sprintf("%s %s %d y.o, i have cat, his name is %s", a.name, a.surname, a.age, a.catName)
}

func main(){
	human := Human{name: "Данил", surname: "Мокшанцев", age: 20}
	fmt.Println(human.getInfo()) // stdout "Данил Мокшанцев 20 y.o"
	fmt.Println(human.getName()) // stdout "Данил"

	//Создаем экземпляр дочерней структуры.
	action := Action{Human: human, catName: "Зефирка"}

	fmt.Println(action.getName()) //Благодаря встраиванию мы обращаемся напрямую к структуре Human.

	//Именно тут и происходит перекрытие методов. Обращаясь к методу getInfo(), мы сначала проверяем есть ли у дочерней структуры
	//этот метод, если нет, обращаемся к родителю. В нашем случае он есть
	fmt.Println(action.getInfo()) // stdout "Данил Мокшанцев 20 y.o, i have cat, his name is Зефирка"
	//Но всегда напрямую можно обратиться к перекрытому, родительскому методу
	fmt.Println(action.Human.getInfo()) // stdout "Данил Мокшанцев 20 y.o"

	//Можем обратиться к полям родительской структуры. Тут работает логика 'перекрытия' как и с методами.
	fmt.Println(action.name, action.surname, action.age, action.catName)

	// human = action так сделать не можем, как в той же java. Обобщить "дочерние" структуры мы можем только при помощи интерфейсов.
}