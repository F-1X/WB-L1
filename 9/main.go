package main

import (
	"sync"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
// массива, во второй — результат операции x*2, после чего данные из второго
// канала должны выводиться в stdout.
func main() {

	// создадим последовательность чисел
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	// исходящий канал
	inbound := make(chan int)

	// выходящий канал
	outbound := make(chan int)

	var wg sync.WaitGroup

	// запустим конвеер
	wg.Add(1)
	go func() {
		defer wg.Done()
		// прием чисел из пайплайна
		for i := range second(first(inbound)) {
			println(i)
		}
		close(outbound)
	}()

	// передадим числа в канал
	for i := 0; i < len(array); i++ {
		inbound <- array[i]
	}

	// после отправки всех чисел можно закрыть канал
	close(inbound)
	wg.Wait()
}

func first(ch <-chan int) <-chan int {

	out := make(chan int)
	go func() {
		for i := range ch {
			out <- i
		}
		// при закрытии канала ch цикл остановится и мы закроем исходящий канал out
		close(out)
	}()

	// передадим канал далее (в следующий stage)
	return out
}

func second(ch <-chan int) <-chan int {

	out := make(chan int)
	go func() {
		for i := range ch {
			out <- i * 2
		}
		close(out)
	}()

	return out
}
