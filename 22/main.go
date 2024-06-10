package main

import (
	"fmt"
	"math/big"
)

// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20.
func main() {

	// используем пакет math/big

	// сосздадим число big.Int из строки
	a, _ := new(big.Int).SetString("1111111111111111111111111111111111111", 10)
	b, _ := new(big.Int).SetString("8888888888888888888888888888888888888", 10)

	fmt.Printf("mul: %d\n", mul(*a, *b)) // умножение
	fmt.Printf("div: %d\n", div(*a, *b)) // деление
	fmt.Printf("sub: %d\n", sub(*a, *b)) // вычитание
	fmt.Printf("sum: %d\n", sum(*a, *b)) // сложение

}

func mul(a, b big.Int) *big.Int {
	return new(big.Int).Mul(&a, &b)
}

func div(a, b big.Int) *big.Int {
	return new(big.Int).Div(&a, &b)
}

func sub(a, b big.Int) *big.Int {
	return new(big.Int).Sub(&a, &b)
}

func sum(a, b big.Int) *big.Int {
	return new(big.Int).Add(&a, &b)
}
