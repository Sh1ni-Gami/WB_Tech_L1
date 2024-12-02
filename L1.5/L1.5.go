package main

import (
	"fmt"
	"time"
)

func main() {
	// время выполнения программы в секундах
	duration := 5 * time.Second

	// канал для передачи целых чисел
	ch := make(chan int)

	// Запуск горутины для отправки значений в канал
	go func() {
		for i := 0; ; i++ {
			ch <- i
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Запуск горутины для чтения значений из канала
	go func() {
		for {
			value := <-ch
			fmt.Println("Получено значение:", value)
		}
	}()

	<-time.After(duration)
	fmt.Println("Время истекло. Программа завершена.")
}
