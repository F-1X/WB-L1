package main

import (
	"fmt"
	"syscall"
	"time"
)

// Реализовать собственную функцию sleep.
func main() {

	sleep(time.Second * 2)

}

func sleep(t time.Duration) {
	// преобразуем время в формат Timespec
	tpc := syscall.NsecToTimespec(t.Nanoseconds())

	// второй аргумент отражает сколько времени действительно прошло в случае прерывания
	var tpc2 syscall.Timespec

	start := time.Now()

	err := syscall.Nanosleep(&tpc, &tpc2)

	fmt.Println("elapsed:", time.Since(start))

	if err != nil {
		fmt.Println(err)
	}

}
