package main

import "fmt"

// Поменять местами два числа без создания временной переменной.
func main() {
	var A int = 10
	var B int = 20

	testCases := []struct {
		input Pair
		want  Pair
	}{
		{Pair{A, B}, Pair{B, A}},
	}

	for i := range testCases {

		A, B = swapArithmetic(testCases[i].input.A, testCases[i].input.B)
		if A != testCases[i].want.A || B != testCases[i].want.B {
			fmt.Printf("testCase(%d) failed: swapArithmetic(%d, %d) = %d, %d (expected: %d, %d)\n", i, testCases[i].input.A, testCases[i].input.B, A, B, testCases[i].want.A, testCases[i].want.B)
		}	else {	
			fmt.Printf("testCase(%d) passed: swapArithmetic(%d, %d) = %d, %d (result: %t)\n", i, testCases[i].input.A, testCases[i].input.B, A, B, testCases[i].want.A == A && testCases[i].want.B == B)
		}

		A, B = swapMultipleAssignments(testCases[i].input.A, testCases[i].input.B)
		if A != testCases[i].want.A || B != testCases[i].want.B {
			fmt.Printf("testCase(%d) failed: swapMultipleAssignments(%d, %d) = %d, %d (expected: %d, %d)\n", i, testCases[i].input.A, testCases[i].input.B, A, B, testCases[i].want.A, testCases[i].want.B)
		} else {
			fmt.Printf("testCase(%d) passed: swapMultipleAssignments(%d, %d) = %d, %d (result: %t)\n", i, testCases[i].input.A, testCases[i].input.B, A, B, testCases[i].want.A == A && testCases[i].want.B == B)
		}

		A, B = swapXOR(testCases[i].input.A, testCases[i].input.B)
		if A != testCases[i].want.A || B != testCases[i].want.B {
			fmt.Printf("testCase(%d) failed: swapXOR(%d, %d) = %d, %d (expected: %d, %d)\n", i, testCases[i].input.A, testCases[i].input.B, A, B, testCases[i].want.A, testCases[i].want.B)
		} else {
			fmt.Printf("testCase(%d) passed: swapXOR(%d, %d) = %d, %d (result: %t)\n", i, testCases[i].input.A, testCases[i].input.B, A, B, testCases[i].want.A == A && testCases[i].want.B == B)
		}

		A, B = swap(testCases[i].input.A, testCases[i].input.B)
		if A != testCases[i].want.A || B != testCases[i].want.B {
			fmt.Printf("testCase(%d) failed: swap(%d, %d) = %d, %d (expected: %d, %d)\n", i, testCases[i].input.A, testCases[i].input.B, A, B, testCases[i].want.A, testCases[i].want.B)
		} else {
			fmt.Printf("testCase(%d) passed: swap(%d, %d) = %d, %d (result: %t)\n", i, testCases[i].input.A, testCases[i].input.B, A, B, testCases[i].want.A == A && testCases[i].want.B == B)
		}
	}

}

type Pair struct {
	A int
	B int
}

func swap(a, b int) (int, int) {
	return b, a
}

func swapArithmetic(A, B int) (int, int) {
	A = A + B
	B = A - B
	A = A - B
	return A, B
	
}

func swapMultipleAssignments(A, B int) (int, int) {
	// multiple assignments
	A, B = B, A
	return A, B
}

func swapXOR(A, B int) (int, int) {
	A = A ^ B
	B = A ^ B
	A = A ^ B
	return A, B
}
