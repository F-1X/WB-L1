package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
// квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
func main() {

	array := []int{2, 4, 6, 8, 10}
	total := 0

	// c использованием мьютекса
	total = ConcurrentSumOfSquares(array)
	fmt.Println("1:", total) // 220 = 2^2 + 4^2 + 6^2 + 8^2 + 10^2 = 4 + 16 + 36 + 64 + 100

	// c использованием каналов
	total = ConcurrentSumOfSquaresChannels(array)
	fmt.Println("2:", total)

	// на атомике AddInt64
	total = ConcurrentSumOfSquaresAtomic(array)
	fmt.Println("3:", total)

	// с использованием каналов, но чтение в отдельной горутине
	total = ConcurrentSumOfSquaresChannelsExtended(array)
	fmt.Println("4:", total)

	total = ConcurrentSumOfSquaresChannelsReturnReader(array)
	fmt.Println("5:", total)

	total = ConcurrentSumOfSquaresChannelsReturnChan(array)
	fmt.Println("6:", total)
}

// ConcurrentSumOfSquares - данной функций реалируется конкурентный доступ к переменной total из нескольких горутин, потокобезопасность основывается на блокировке мьютекса
func ConcurrentSumOfSquares(array []int) int {

	// результат суммируется сюда
	total := 0
	wg := sync.WaitGroup{}

	mu := sync.Mutex{}

	for i := 0; i < len(array); i++ {
		wg.Add(1)
		go func() {
			mu.Lock() // блокировка мьютекса
			total += array[i] * array[i]
			mu.Unlock()
			wg.Done()
		}()

	}
	wg.Wait() // ждем когда все горутины закончат выполнение

	return total
}

// ConcurrentSumOfSquaresChannels - использует примитив синхронизации такой как канал
func ConcurrentSumOfSquaresChannels(array []int) int {

	total := 0

	// буферизированный канал или нет, может быть любой
	ch := make(chan int)

	// для синхронизации
	var wg sync.WaitGroup

	for i := 0; i < len(array); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- array[i] * array[i]
		}()
	}

	go func() {
		wg.Wait() // группа также служит для ожидания и закрытия канала, который заблокировал main ниже
		close(ch)
	}()

	for i := range ch {
		total += i
	}
	return total
}

func ConcurrentSumOfSquaresAtomic(array []int) int {
	var total int64

	// Создаем канал для передачи результатов квадратов
	ch := make(chan int64)

	// для синхронизации
	var wg sync.WaitGroup

	for i := 0; i < len(array); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			atomic.AddInt64(&total, int64(array[i]*array[i]))
		}()
	}

	wg.Wait()
	close(ch)

	return int(total)
}

func ConcurrentSumOfSquaresChannelsExtended(array []int) int {
	// буферизированный канал или нет, может быть любой
	ch := make(chan int)

	// для синхронизации
	var wg sync.WaitGroup

	for i := 0; i < len(array); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- array[i] * array[i]
		}()
	}

	// шаренная переменная,
	total := 0

	wg2 := sync.WaitGroup{}
	wg2.Add(1)
	go reader(&total, ch, &wg2) // запускаем чтение в отдельной горутине передавая указатель на переменую. счетчика

	// ждем писателей
	wg.Wait()
	close(ch)

	// ждем читателя
	wg2.Wait()

	return total
}

func reader(total *int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range ch {
		*total += i
	}
}

func ConcurrentSumOfSquaresChannelsReturnReader(array []int) int {
	// буферизированный канал или нет, может быть любой
	ch := make(chan int)

	// для синхронизации
	var wg sync.WaitGroup

	for i := 0; i < len(array); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- array[i] * array[i]
		}()
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	return reader2(ch)
}

func reader2(ch chan int) int {
	total := 0
	for i := range ch {
		total += i
	}

	return total
}

func ConcurrentSumOfSquaresChannelsReturnChan(array []int) int {
	// буферизированный канал или нет, может быть любой
	ch := make(chan int)

	// для синхронизации
	var wg sync.WaitGroup

	for i := 0; i < len(array); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- array[i] * array[i]
		}()
	}

	out := make(chan int)
	go reader3(ch, out)

	wg.Wait()
	close(ch)

	return <-out
}

func reader3(ch, out chan int) {
	total := 0
	for i := range ch {
		total += i
	}
	out <- total
}
