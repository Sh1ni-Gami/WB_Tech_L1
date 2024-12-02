package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	chan1 := make(chan int)
	chan2 := make(chan int)

	// Горутина для отправки чисел в первый канал
	go func() {
		for _, num := range numbers {
			chan1 <- num
		}
		close(chan1) // Закрыл первый канал после отправки всех чисел
	}()

	// Горутина для чтения из первого канала, умножения на 2 и отправки во второй канал
	go func() {
		for num := range chan1 {
			chan2 <- num * 2
		}
		close(chan2) // Закрыл второй канал после обработки всех чисел
	}()

	for result := range chan2 {
		fmt.Println(result)
	}
}
