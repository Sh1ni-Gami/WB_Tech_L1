package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var summ int
	var wg sync.WaitGroup // Использую WaitGroup для ожидания завершения горутин
	// Проходим по всем числам массива
	for _, num := range numbers {
		wg.Add(1) // Увеличиваем счетчик WaitGroup

		go func(n int) {
			defer wg.Done()
			square := n * n
			summ += square
		}(num)
	}
	wg.Wait()
	fmt.Println("Сумма квадратов равна", summ)
}
