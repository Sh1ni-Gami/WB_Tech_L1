package main

import (
	"fmt"
)

// структура Human с полями Name и Age
type Human struct {
	Name string
	Age  int
}

// Метод для структуры Human
func (h Human) Introduce() string {
	return "Привет, меня зовут " + h.Name + ", мне " + fmt.Sprint(h.Age) + " года."
}

// структура Action, которая включает в себя Human
type Action struct {
	Human
}

// Метод для выполнения действия, который использует метод Introduce
func (a Action) PerformAction() {
	fmt.Println(a.Introduce())
}

func main() {
	// Создаем экземпляр Human
	human := Human{Name: "Руслан", Age: 23}
	// Создаем экземпляр Action
	action := Action{Human: human}
	action.PerformAction()
}
