package main

import (
	"fmt"
	"unicode/utf8"
)

// переворачивает строку, поддерживая символы Unicode
func reverseString(s string) string {
	// срез для хранения перевернутых символов
	runes := make([]rune, 0, len(s))

	// цикл по строке
	for i := len(s) - 1; i >= 0; {
		// Получаем следующий символ и его байтовую длину
		r, size := utf8.DecodeLastRuneInString(s[:i+1])
		runes = append(runes, r) // Добавляем символ в срез перевернутых символов
		i -= size
	}

	// Преобразовали срез runes обратно в строку
	return string(runes)
}

func main() {
	testStrings := []string{
		"главрыба",
		"абырвалг",
		"こんにちは",
		"Hello, 世界",
	}

	for _, str := range testStrings {
		fmt.Printf("Исходная строка: %s\n", str)
		fmt.Printf("Перевернутая строка: %s\n\n", reverseString(str))
	}
}
