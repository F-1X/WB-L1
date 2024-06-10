package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.
func main() {
	testCases := []struct {
		input string
		want  string
	}{
		{"главрыба", "абырвалг"},
		{"dsajkbщзш", "шзщbkjasd"},
	}

	for i := 0; i < len(testCases); i++ {
		// 1 вариант (используя руны)
		got := reverse1(testCases[i].input)
		if got != testCases[i].want {
			fmt.Printf("testCase(%d) failed: reverse1(%q) = %q (expected: %q)\n", i, testCases[i].input, got, testCases[i].want)
		} else {
			fmt.Printf("testCase(%d) passed: reverse1(%q) == %q (result: %t)\n", i, testCases[i].input, got, testCases[i].want == got)
		}

		// 2 вариант (используя strings.Split)
		got = reverse2(testCases[i].input)
		if got != testCases[i].want {
			fmt.Printf("testCase(%d) failed: reverse1(%q) = %q (expected: %q)\n", i, testCases[i].input, got, testCases[i].want)
		} else {
			fmt.Printf("testCase(%d) passed: reverse1(%q) == %q (result: %t)\n", i, testCases[i].input, got, testCases[i].want == got)
		}
	}

}

func reverse1(word string) string {
	runes := []rune(word)

	// меняем местами протиположные буквы
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func reverse2(word string) string {

	// используя сплит, разделяем слово побуквенно
	split := strings.Split(word, "")

	// меняем местами протиположные буквы
	for i, j := 0, len(split)-1; i < j; i, j = i+1, j-1 {
		split[i], split[j] = split[j], split[i]
	}

	// объединяем в строку
	return strings.Join(split, "")
}
