package main

import "fmt"

func main() {
	a := 0
	fmt.Scan(&a)
	n := a*(a-1) + 1
	for i := 0; i < a-1; i++ {
		fmt.Print(n, "+")
		n += 2
	}
	fmt.Print(n)
}
