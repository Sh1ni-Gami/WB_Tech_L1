package main

import "fmt"

// Функция для установки i-го бита в 1
func setBitTo1(num int64, i uint) int64 {
	// Используем побитовое ИЛИ с маской, где i-й бит равен 1
	return num | (1 << i)
}

// Функция для установки i-го бита в 0
func setBitTo0(num int64, i uint) int64 {
	// Используем побитовое И с инверсией маски, где i-й бит равен 1
	return num &^ (1 << i)
}

func main() {
	var num int64 = 42
	var i uint = 3 // Установил 3-й бит

	fmt.Printf("Число до изменения: %064b\n", num)

	// Установил i-й бит в 1
	num = setBitTo1(num, i)
	fmt.Printf("После установки %d-го бита в 1: %064b\n", i, num)

	// Установил i-й бит в 0
	num = setBitTo0(num, i)
	fmt.Printf("После установки %d-го бита в 0: %064b\n", i, num)
}
