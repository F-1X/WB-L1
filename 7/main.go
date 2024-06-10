package main

import (
	"fmt"
	"sync"
)

// Реализовать конкурентную запись данных в map.
func main() {
	// go run -race main.go

	exampleMutex()

	exampleChannels()

	exampleAtomic()
}

func exampleMutex() {
	store := make(map[int]int)
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	// запустим 10 воркеров (10 горутин) "Worker Pool"
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			// каждая из которых обращается по тем-же ключам в цикле
			for i := 0; i < 10; i++ {
				mu.Lock() // мьютекс заблокирует другие горутины которые также захотят заблокироваться
				store[i] = i
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(store)
}

func exampleChannels() {
	store := make(map[int]int)

	// канал для синхронизации (буферизированный или нет в данной ситуации не важно, так как читает из кала только 1 горутина, это и обеспечивает синхронизацию)
	ch := make(chan int)

	// в слуйчаном порядке 20 ключей
	data := []int{7, 8, 0, 15, 8, 0, 15, 12, 18, 3, 14, 19, 5, 17, 18, 3, 4, 6, 11, 14, 9, 10, 16, 4, 6, 11, 1, 2, 13}

	// в случае одних и тех же ключей
	// data := []int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1}

	// синхронизация
	done := make(chan struct{})

	go func() {
		for i := range ch {
			store[i] = i
		}
		close(done)
	}()

	// так будет DATA RACE
	// done1 := make(chan struct{})
	// go func() {
	// 	for i := range ch {
	// 		store[i] = i
	// 	}
	// 	close(done1)
	// }()

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < len(data); i++ {
				ch <- data[i]
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(ch)
	<-done
	// <-done1

	fmt.Println(store)
}

func exampleAtomic() {
	// store := make(map[int]int)
	var wg sync.WaitGroup
	var store sync.Map // в структуре мьютекс

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				store.Store(i, i)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Printf("map[")
	for i := 0; i < 10; i++ {
		x, _ := store.Load(i)
		fmt.Printf("%d:%d ", i, x)
	}
	fmt.Printf("]\n")
}
