package main

import (
	"fmt"
	"strings"
)

// переворачивает слова в строке
func reverseWords(s string) string {
	// Разбили строку на слова
	words := strings.Fields(s)

	reversed := make([]string, len(words))

	// Заполняем новый срез словами в обратном порядке
	for i := 0; i < len(words); i++ {
		reversed[i] = words[len(words)-1-i]
	}

	// Объединили слова в строку
	return strings.Join(reversed, " ")
}

func main() {
	input := "snow dog sun"
	reversed := reverseWords(input)
	fmt.Printf("Исходная строка: %s\n", input)
	fmt.Printf("Перевернутая строка: %s\n", reversed)
}
