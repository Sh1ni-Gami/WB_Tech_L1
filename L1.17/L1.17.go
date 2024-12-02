package main

import "fmt"

// binarySearch выполняет бинарный поиск в отсортированном срезе arr
func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2 // Находим средний индекс, чтобы избежать переполнения

		if arr[mid] == target {
			return mid // Возвращаем индекс найденного элемента
		} else if arr[mid] < target {
			left = mid + 1 // Ищем в правой части
		} else {
			right = mid - 1 // Ищем в левой части
		}
	}

	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	target := 7

	fmt.Println("Исходный массив:", arr)
	index := binarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Элемент %d найден на индексе %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}
}
