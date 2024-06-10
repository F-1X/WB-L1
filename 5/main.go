package main

import (
	"fmt"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в
// канал, а с другой стороны канала — читать. По истечению N секунд программа
// должна завершаться.
func main() {
	// N секунд чтобы завершиться
	N := 3

	ch := make(chan any, 1)

	// получатель
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()

	// отправитель
	go func() {
		for {
			ch <- 1
		}
	}()

	<-time.After(time.Second * time.Duration(N))
	close(ch)
}
