package main

import (
	"fmt"
	"strings"
)

// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».
func main() {

	testCases := []struct {
		input string
		want  string
	}{
		{"snow dog sun", "sun dog snow"},
		{"сссч ячс фыв", "фыв ячс сссч"},
	}

	for i := 0; i < len(testCases); i++ {
		// 1 вариант (алгоритм проходится посимвольно с конца строки и записывает слова в новую строку)
		got := reverse1(testCases[i].input)
		if got != testCases[i].want {
			fmt.Printf("testCase(%d) failed: reverse1(%q) = %q (expected: %q)\n", i, testCases[i].input, got, testCases[i].want)
		} else {
			fmt.Printf("testCase(%d) passed: reverse1(%q) == %q (result: %t)\n", i, testCases[i].input, got, testCases[i].want == got)
		}

		// 2 вариант
		got = reverse2(testCases[i].input)
		if got != testCases[i].want {
			fmt.Printf("testCase(%d) failed: reverse1(%q) = %q (expected: %q)\n", i, testCases[i].input, got, testCases[i].want)
		} else {
			fmt.Printf("testCase(%d) passed: reverse1(%q) == %q (result: %t)\n", i, testCases[i].input, got, testCases[i].want == got)
		}

		// 3 вариант (strings.Split + перестановка местами + strings.Join)
		got = reverse3(testCases[i].input)
		if got != testCases[i].want {
			fmt.Printf("testCase(%d) failed: reverse1(%q) = %q (expected: %q)\n", i, testCases[i].input, got, testCases[i].want)
		} else {
			fmt.Printf("testCase(%d) passed: reverse1(%q) == %q (result: %t)\n", i, testCases[i].input, got, testCases[i].want == got)
		}
	}
}

func reverse1(word string) string {
	runes := []rune(word)

	newWord := ""

	// счетчик длины слова
	cW := 0
	for i := len(runes) - 1; i >= 0; i-- {
		cW++
		// проверка пробела или начала строки
		if runes[i] == 32 || i == 0 {
			// тут нужно увеличиить счетчик тк мы еще находимся на границе слова и уменьшить интератор чтобы удовлетворить нижележащему циклу, а именно стартовой позиции j
			if i == 0 {
				cW++
				i = -1
			}

			for j := i + 1; j < i+cW; j++ {
				newWord += string(runes[j])
			}

			cW = 0
			if i != -1 {
				newWord += " "
			}
		}
	}
	return newWord
}

func reverse2(word string) string {

	answer := ""

	// разделяем на слова
	split := strings.Split(word, " ")
	
	// правого слова сдвигаемся влево и записывает слово в новую строку answer
	for i := len(split) - 1; i >= 0; i-- {
		answer += split[i] + " "
	}

	// удаляем последний пробел
	return answer[:len(answer)-1]
}

func reverse3(word string) string {
	// разделяем на слова
	split := strings.Split(word, " ")

	// меняем местами противолежащие элементы
	for i, j := 0, len(split)-1; i < j; i, j = i+1, j-1 {
		split[i], split[j] = split[j], split[i]
	}

	// объединяем слова в 1 строку
	return strings.Join(split, " ")
}
