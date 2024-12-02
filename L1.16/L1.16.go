package main

import "fmt"

// quicksort выполняет быструю сортировку на срезе чисел
func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2] // опорный элемент
	left := []int{}          // Срез для элементов меньше опорного
	right := []int{}         // Срез для элементов больше опорного

	for _, v := range arr {
		if v < pivot {
			left = append(left, v) // Элементы меньше опорного
		} else if v > pivot {
			right = append(right, v) // Элементы больше опорного
		}
	}

	// Рекурсивно сортируем левые и правые части и объединяем их с опорным элементом
	return append(append(quicksort(left), pivot), quicksort(right)...)
}

func main() {
	arr := []int{3, 6, 8, 10, 1, 2, 1}
	fmt.Println("Исходный массив:", arr)

	sortedArr := quicksort(arr)
	fmt.Println("Отсортированный массив:", sortedArr)
}
