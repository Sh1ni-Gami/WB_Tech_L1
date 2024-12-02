package main

import (
	"fmt"
	"math"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	for _, temp := range temperatures {
		// Определили нижнюю границу диапазона (с шагом 10)
		key := int(math.Floor(temp/10) * 10)

		// Добавили в соответствующую группу
		groups[key] = append(groups[key], temp)
	}

	for key, values := range groups {
		fmt.Printf("%d: %v\n", key, values)
	}
}
