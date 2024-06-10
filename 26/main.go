package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна быть
// регистронезависимой.
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false
func main() {

	testCases := []struct {
		input string
		want  bool
	}{
		{"abcd", true},
		{"abCdefAaf", false},
		{"aabcd", false},
		{"asdfghjkjhgfdsdfghjklkjdsvsdvsd", false},
		{"абвгд", true},
	}
	for _, test := range testCases {
		got := IsUniqCharsMap(test.input)
		if got != test.want {
			// вывод если тест не прошел
			fmt.Printf("allUniqueChars(%q) = %v (expected: %v)\n", test.input, got, test.want)
		}

		// для вывода из условия
		if got == true {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	}
}

func IsUniqCharsMap(input string) bool {
	// опускаем регистр
	input = strings.ToLower(input)

	// преобразуем в руны чтобы работать можно было с ютф
	s := []rune(input)

	memory := make(map[rune]struct{}, 0)
	for i := 0; i < len(s); i++ {
		if _, ok := memory[rune(s[i])]; ok {
			return false
		}
		memory[rune(s[i])] = struct{}{}
	}

	return true
}
