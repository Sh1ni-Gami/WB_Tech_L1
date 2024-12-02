package main

import "fmt"

func main() {

	// Создали слайс с переменными типа interface{}
	// interface{} может хранить значения любого типа
	variables := []interface{}{42, "Hello", true, make(chan int)}

	// Цикл по каждой переменной и определяем её тип
	for i, variable := range variables {
		// конструкция switch-type для определения типа
		switch v := variable.(type) {
		case int:
			fmt.Printf("Значение: %v, имеет тип int\n", v)
		case string:
			fmt.Printf("Значение: %v, имеет тип string\n", v)
		case bool:
			fmt.Printf("Значение: %v, имеет тип bool\n", v)
		case chan int:
			fmt.Printf("Значение: %v, имеет тип channel\n", v)
		default:
			fmt.Printf("%d имеет неизвестный тип\n", i)
		}
	}
}
