package main

import (
	"fmt"
	"strings"
)

// проверяет, все ли символы в строке уникальные
func hasUniqueChars(s string) bool {
	// Привили строку к нижнему регистру
	s = strings.ToLower(s)

	// Создали множество для хранения уникальных символов
	charSet := make(map[rune]struct{})

	for _, char := range s {
		if _, exists := charSet[char]; exists {
			return false
		}
		charSet[char] = struct{}{}
	}

	return true
}

func main() {
	stringsToTest := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
		"golang_ninja",
		"Python",
		"golang",
	}

	for _, str := range stringsToTest {
		fmt.Printf("Строка \"%s\" имеет уникальные символы: %t\n", str, hasUniqueChars(str))
	}
}
