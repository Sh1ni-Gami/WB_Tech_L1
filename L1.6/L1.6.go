package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var wg sync.WaitGroup

	// 1. Остановка горутины с помощью канала
	stopChan := make(chan bool)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stopChan:
				fmt.Println("Горутина 1 завершена (канал).")
				return
			default:
				fmt.Println("Горутина 1 работает (канал)...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	// 2. Остановка горутины с использованием context.Context
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина 2 завершена (контекст).")
				return
			default:
				fmt.Println("Горутина 2 работает (контекст)...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	// 3. Остановка горутины с использованием обычного флага
	var stopFlag bool
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if stopFlag {
				fmt.Println("Горутина 3 завершена (флаг).")
				return
			}
			fmt.Println("Горутина 3 работает (флаг)...")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// 4. Остановка горутины с использованием атомарного флага
	var atomicFlag int32
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if atomic.LoadInt32(&atomicFlag) == 1 {
				fmt.Println("Горутина 4 завершена (атомарный флаг).")
				return
			}
			fmt.Println("Горутина 4 работает (атомарный флаг)...")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	time.Sleep(2 * time.Second)

	// Останавливаем горутины
	stopChan <- true
	cancel()
	stopFlag = true
	atomic.StoreInt32(&atomicFlag, 1)

	wg.Wait()

	fmt.Println("Все горутины завершены.")
}
