package main

import (
	"fmt"
	"time"
)

// приостанавливает выполнение на заданное количество миллисекунд
func Sleep(milliseconds int) {
	// Преобразование миллисекунд в время.Duration
	duration := time.Duration(milliseconds) * time.Millisecond
	time.Sleep(duration)
}

func main() {
	fmt.Println("Начало задержки...")
	Sleep(2000) // Задержка на 2 секунды
	fmt.Println("Задержка завершена.")
}
