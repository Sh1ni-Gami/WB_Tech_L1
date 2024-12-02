package main

import (
	"fmt"
	"sync"
)

func main() {
	var m = make(map[int]int) // создаем мапу
	var mutex sync.Mutex      // объявляем мьютекс
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mutex.Lock()
			m[i] = i * 2
			mutex.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Println(m)
}
