package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Реализовать структуру-счетчик, которая будет инкрементироваться в
// конкурентной среде. По завершению программа должна выводить итоговоезначение счетчика.
func main() {

	// считать до
	target := 1000

	result := exampleMutex(target)
	if result != target {
		fmt.Println("exampleMutex failed: result =", result, "expected =", target)
	} else {
		fmt.Printf("exampleMutex success: result(%d) == expected(%d): %t\n", result, target, result == target)
	}

	result = exampleCond(target)
	if result != target {
		fmt.Println("exampleCond failed: result =", result, "expected =", target)
	} else {
		fmt.Printf("exampleCond success: result(%d) == expected(%d): %t\n", result, target, result == target)
	}

	result = exampleAtomic(target)
	if result != target {
		fmt.Println("exampleAtomic failed: result =", result, "expected =", target)
	} else {
		fmt.Printf("exampleAtomic success: result(%d) == expected(%d): %t\n", result, target, result == target)
	}

	result = exampleCAS(target)
	if result != target {
		fmt.Println("exampleCAS failed: result =", result, "expected =", target)
	} else {
		fmt.Printf("exampleCAS success: result(%d) == expected(%d): %t\n", result, target, result == target)
	}
}

func exampleMutex(target int) int {
	counter := 0

	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	for i := 0; i < target; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			counter++
			mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	return counter
}

func exampleCond(target int) int {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	counter := 0

	for i := 0; i < 10; i++ {
		go func() {
			for {
				cond.L.Lock()
				for counter >= target {
					cond.L.Unlock()
					return
				}
				counter++
				cond.Signal()
				cond.L.Unlock()
			}
		}()
	}

	cond.L.Lock()
	defer cond.L.Unlock()
	for counter < target {
		cond.Wait()
	}

	return counter

}

func exampleAtomic(target int) int {
	var counter int32

	wg := sync.WaitGroup{}
	for i := 0; i < target; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}

	wg.Wait()
	return int(atomic.LoadInt32(&counter))
}

func exampleCAS(target int) int {
	var counter int32
	var wg sync.WaitGroup

	target32 := int32(target)

	incrementor := func() {
		defer wg.Done()
		for {
			// выгружаем переменную
			current := atomic.LoadInt32(&counter)

			// сравниваем с целевый значение
			if current < target32 {
				newValue := current + 1
				// записываем новое значение
				if atomic.CompareAndSwapInt32(&counter, current, newValue) {
					if newValue == target32 {
						break
					}
				}
			} else {
				break
			}
		}
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go incrementor()
	}

	wg.Wait()
	return int(counter)
}
