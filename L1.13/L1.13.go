package main

import "fmt"

func main() {
	// Исходные числа
	a := 104
	b := 205

	fmt.Println("До замены:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)

	// 1. Добавляем к первому числу второе
	a = a + b
	// 2. Получаем новое значение для b, вычитая из суммы старое значение b
	b = a - b
	// 3. Получаем новое значение для a, вычитая из суммы новое значение b
	a = a - b

	fmt.Println("\nПосле замены:")
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}
