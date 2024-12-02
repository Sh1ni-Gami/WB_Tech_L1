package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

// Функция воркера, который читает данные из канала
func worker(id int, jobs <-chan string) {
	for job := range jobs {
		fmt.Printf("Воркер %d обработал %s\n", id, job)
	}
}

func main() {
	var numWorkers int
	fmt.Print("Введите количество воркеров: ")
	fmt.Scan(&numWorkers)

	// канал для передачи данных
	jobs := make(chan string)

	// Запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		go worker(i, jobs)
	}

	// перехват сигнала для завершения программы
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Горутина для записи данных в канал
	go func() {
		for {
			jobs <- "данные: " + strconv.Itoa(time.Now().Second())
			time.Sleep(1 * time.Second) // Пауза перед следующей записью
		}
	}()

	// Ожидания завершения
	<-signalChan

	// завершаем воркеры
	close(jobs)
	fmt.Println("\nЗавершение работы программы...")
}
