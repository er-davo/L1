package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(0)
	b := big.NewInt(0)

	a.SetString("1000000000000000000000000000000000", 10)
	b.SetString("2000000000000000000000000000000000", 10)

	sum := new(big.Int).Add(a, b)
	sub := new(big.Int).Sub(a, b)
	mul := new(big.Int).Mul(a, b)
	div := new(big.Int).Div(a, b)

	fmt.Println("Сложение:", sum)
	fmt.Println("Вычитание:", sub)
	fmt.Println("Умножение:", mul)
	fmt.Println("Деление:", div)
}
