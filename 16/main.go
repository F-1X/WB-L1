package main

import "fmt"

// Реализовать быструю сортировку массива (quicksort) встроенными методами
// языка.
func main() {
	array := []int{8, 3, 1, 10, 4, 2, 6, 7, 9, 5}
	
	QuickSort(array, 0, len(array)-1)
	fmt.Println(array)
}

// Алгоритм быстрой сортировки n*log(n), вернее его смысл реализации, в том, чтобы установить опорный элемент, в нашем случае - срединный, взять левую и правую границу и перемещать
// левую границу в случае если элемент меньше опорного, а правую границу j если элемент больше опорного, и рекурсирно вызывать сортировку для левого и правого подмассива.
// Поторять пока левая граница меньше правой
func QuickSort(array []int, left int, right int) {
	if left > right {
		return
	}
	if len(array) == 0 {
		return
	}

	mid := left + (right-left)/2
	pivot := array[mid]

	i := left
	j := right

	for i <= j {
		for array[i] < pivot {
			i++
		}

		for array[j] > pivot {
			j--
		}

		if i <= j {
			array[i], array[j] = array[j], array[i]
			i++
			j--
		}

	}

	if left < j {
		QuickSort(array, left, j)
	}
	if i < right {
		QuickSort(array, i, right)
	}

}
