package main

import (
	"fmt"
)

// Удалить i-ый элемент из слайса.
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	N := 4 // значение 5

	// вариант 1 (через создание нового слайса)
	slice = append(slice[:N:N], slice[N+1:]...)
	fmt.Println(slice)

	// вариант 2 (через перемещение элемента на край вправо и отсечение последнего элемента)
	slice = []int{1, 2, 3, 4, 5, 6, 7, 8}

	slice[N] = slice[len(slice)-1]

	slice = slice[:len(slice)-1]
	fmt.Println(slice)

	// вариант 3 (через корирование)
	slice = []int{1, 2, 3, 4, 5, 6, 7, 8}
	copy(slice[N:], slice[N+1:])

	slice = slice[:len(slice)-1]

	fmt.Println(slice)
}
