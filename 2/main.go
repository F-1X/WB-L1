package main

import (
	"fmt"
	"sync"
)

// Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
func main() {
	// go run -race main.go

	// входные данные
	array := []int{2, 4, 6, 8, 10}

	// для каждого значения из массива запускается горутина которая выводит квадрат в stdout
	first(array)

	// синхронизация обмена данными мьютексом
	second(array)

	// обмен данными через канал
	third(array)
}

func first(array []int) {
	var wg sync.WaitGroup
	for i := 0; i < len(array); i++ {
		wg.Add(1)
		go func() {
			fmt.Printf("%d ", array[i]*array[i])
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println()
}

func second(arr []int) {
	array := make([]int, len(arr)) // нужно создать копию так как меняется значения слайса
	copy(array, arr)

	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < len(array); i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			array[i] *= array[i]
			mu.Unlock()
			wg.Done()
		}()

	}
	wg.Wait()

	// вывод измененного слайса
	for i := range array {
		fmt.Printf("%d ", array[i])
	}
	fmt.Println()
}

func third(array []int) {
	var wg sync.WaitGroup
	ch := make(chan int)

	// вывод
	go func() {
		for square := range ch {
			fmt.Printf("%d ", square)
		}
		fmt.Println()
	}()

	// запуск горутин с вводом в канал
	for i := 0; i < len(array); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- array[i] * array[i]
		}()
	}

	wg.Wait()
	close(ch)
}
