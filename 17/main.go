package main

import "fmt"

// Реализовать бинарный поиск встроенными методами языка.
func main() {
	//array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	array := []int{8, 3, 1, 10, 4, 2, 6, 7, 9, 5}

	// для предварительного использования бинарного поиска необходимо отсортировать массив
	QuickSort(array, 0, len(array)-1)
	// fmt.Println(array)


	// попробуем найти значение 2, должен вернуться индекс 1
	find := 2
	BinSearch(array, find)

}


// Двоичный поиск O(log n), заключается в том, чтобы взять срединный элемент и сравнить элемент с искомым, если искомый элемент меньше, то правую границу установить в среднюю позицию mid-1
// и продолжать брать срединный элемент из массива используя новые границы аналогичным образом сравнивая с искомым значением. Аналогично если элемент больше, то установить левую границу 
// на позицию mid + 1. Продолжать пока границы не сойдуться
func BinSearch(array []int, find int) {
	left := 0
	right := len(array) - 1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if array[mid] == find {
			fmt.Println("found:", mid)
			return
		}
		if array[mid] > find {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Println("not found")
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
