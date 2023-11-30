package main

import (
	"fmt"
	"reflect"
)

func main(){
	whatIs(int64(2))
	whatIs("2")
	whatIs(true)
	whatIs(make(chan int))
	//Выведет unsupported
	whatIs(2.9)
}

func whatIs(value interface{}){
	t := reflect.ValueOf(value)
	//Проходим по всем типам, которые нам нужны
	switch t.Kind(){
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Println("This is int!")
	case reflect.String:
		fmt.Println("This is string!")
	case reflect.Bool:
		fmt.Println("This is bool!")
	case reflect.Chan:
		fmt.Println("This is channel!")
	default:
		fmt.Printf("unsupported %s", t.String())
	}
}