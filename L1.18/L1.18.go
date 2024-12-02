package main

import (
	"fmt"
	"sync"
)

// структура для хранения счетчика и мьютекса
type Counter struct {
	value int
	mu    sync.Mutex // Мьютекс для защиты состояния счетчика
}

// увеличивает значение счетчика на 1
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// возвращает текущее значение счетчика
func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	// Запускаем 100 горутин, каждая из которых инкрементирует счетчик 1000 раз
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()

	fmt.Println("Итоговое значение счетчика:", counter.Value())
}
