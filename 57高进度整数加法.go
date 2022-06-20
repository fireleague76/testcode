package main

import (
	"fmt"
	"math/big"
)

func main() {
	m, n := "", ""
	fmt.Scan(&m, &n)
	num1 := new(big.Int)
	num1.SetString(m, 10)
	num2 := new(big.Int)
	num2.SetString(n, 10)
	num1 = num1.Add(num1, num2)
	fmt.Println(num1.String())
}
