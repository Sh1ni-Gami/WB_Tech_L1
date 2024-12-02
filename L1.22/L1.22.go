package main

import "fmt"

func main() {
	var a int64 = 2 << 20
	var b int64 = 3 << 20

	// Проверка
	const threshold = 1 << 20
	if a <= threshold || b <= threshold {
		fmt.Println("Оба числа должны быть больше 2^20 (1048576).")
		return
	}

	sum := a + b
	difference := a - b
	product := a * b
	var quotient float64
	if b != 0 {
		quotient = float64(a) / float64(b)
	} else {
		fmt.Println("Ошибка: деление на ноль.")
		return
	}

	fmt.Printf("Сумма: %d\n", sum)
	fmt.Printf("Разность: %d\n", difference)
	fmt.Printf("Произведение: %d\n", product)
	fmt.Printf("Частное: %.2f\n", quotient)
}
