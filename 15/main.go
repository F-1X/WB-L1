package main

import (
	"fmt"
)

// К каким негативным последствиям может привести данный фрагмент кода, и как
// это исправить? Приведите корректный пример реализации.
func main() {
	someFunc()
}

func createHugeString(n int) string {
	var s string
	for i := 0; i < n; i++ {
		s += "f"
	}
	return s
}

func createUTF8String(n int) string {
	var s string
	for i := 0; i < n; i++ {
		s += "щ"
	}
	return s
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)

	// так как len подсчитывает количество байт в строке, то необходимо использовать тип rune для под счета длины строки
	fmt.Println(1<<10, "длина в байтах:", len(v), ",длина в рунах:", len([]rune(v)), ",длина в байтах:", len([]byte(v)))

	// 1. нужно проверить что длина строки больше 100, тк беря срез от меньшего кол-ва байт получим панику
	if len(v) > 100 {
		// 2. срез также берется по байтово
		justString = v[:100]
		justString = string([]rune(v)[:100])
	}

	fmt.Println(len(justString), "длина символов:", len([]rune(justString)), justString)



	v2 := createUTF8String(1 << 10)
	fmt.Println(1<<10, "длина в байтах:", len(v2), ",длина в рунах:", len([]rune(v2)), ",длина в байтах:", len([]byte(v2)))

	if len(v2) > 100 {
		justString = v2[:100]
		justString = string([]rune(v2)[:100])
	}

	fmt.Println(len(justString), "длина символов:", len([]rune(justString)), justString)

}
