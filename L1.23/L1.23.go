package main

import "fmt"

// Функция для удаления элемента из слайса по индексу
func removeElement(slice []int, index int) []int {
	// Проверка, что индекс находится в допустимых пределах
	if index < 0 || index >= len(slice) {
		fmt.Println("Индекс за пределами слайса")
		return slice
	}

	// соединяем части слайса до и после индекса
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	index := 2

	fmt.Println("Исходный слайс:", numbers)

	numbers = removeElement(numbers, index)
	fmt.Printf("После удаления элемента с индексом %d: %v\n", index, numbers)
}
