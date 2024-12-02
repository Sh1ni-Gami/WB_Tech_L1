package main

import "fmt"

func intersection(set1, set2 []int) []int {
	// Создали map для хранения элементов первого множества
	set1Map := make(map[int]bool)
	for _, item := range set1 {
		set1Map[item] = true
	}

	// Создали слайс для хранения пересечения
	var result []int

	// Проверка элементов второго множества
	for _, item := range set2 {
		if set1Map[item] { // Если элемент есть в первом множестве
			result = append(result, item) // Добавили его в пересечение
		}
	}

	return result
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 5, 6, 7}

	result := intersection(set1, set2)
	fmt.Println("Пересечение:", result)
}
