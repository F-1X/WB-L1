package main

import "fmt"

// Дана переменная int64. Разработать программу которая устанавливает i-й бит в
// 1 или 0.
func main() {
	var n int64
	var i int
	var b int

	n = 10 // 10 = 1010  (вывод числа в двоичн. системе) fmt.Printf("%b\n", n)
	i = 0  // порядок отсчитывается с 0
	b = 1  // значение для бита

	switch b {
	case 0:
		// fmt.Printf("%b\n", 1<<i)
		n = n &^ (1 << i) // оператор И НЕ - (x &^ y) ставит 0 если в у 1, если в y 0 то ставит бит из х
	case 1:
		// fmt.Printf("%b\n", 1<<i)
		n = n | (1 << i) // оператор ИЛИ - (x | y) ставит 1 в бит 1 если в x или y стоит 1
	}

	// вывод результата
	fmt.Println(n)


	// для использования тест-кейсов закоментируйте код выше и раскоментируйте RunTests()
	// RunTests()
}

func RunTests(){
	var n int64
	var i int
	var b int

	testCases := []struct {
		n int64
		i int
		b int
	}{
		{10, 0, 1},
	}

	for _, tc := range testCases {
		n = tc.n
		i = tc.i
		b = tc.b
		switch b {
		case 0:
			// fmt.Printf("%b\n", 1<<i)
			n = n &^ (1 << i) // оператор И НЕ - (x &^ y) ставит 0 если в у 1, если в y 0 то ставит бит из х
		case 1:
			// fmt.Printf("%b\n", 1<<i)
			n = n | (1 << i) // оператор ИЛИ - (x | y) ставит 1 в бит 1 если в x или y стоит 1
		}

		fmt.Println(n)
	}
}
