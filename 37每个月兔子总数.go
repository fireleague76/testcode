package main

import "fmt"

func main() {
	a1 := 1
	b2 := 0
	c3 := 0
	n := 0
	fmt.Scan(&n)
	for i := 1; i < n; i++ {
		c3 = a1 + b2
		a1 = b2
		b2 = c3
	}
	fmt.Println(c3)
}
