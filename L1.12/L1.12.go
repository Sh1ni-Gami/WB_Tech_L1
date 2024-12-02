package main

import "fmt"

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создали map для хранения уникальных значений
	uniqueSet := make(map[string]struct{})

	// Цикл по всем строкам и добавляем их в множество
	for _, str := range strings {
		uniqueSet[str] = struct{}{}
	}

	// Вывод
	fmt.Println("Уникальные строки:")
	for key := range uniqueSet {
		fmt.Println(key)
	}
}
