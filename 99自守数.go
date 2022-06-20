package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Scan(&n)
	count := 2
	base := 10
	for i := 2; i < n; i++ {
		if i == base {
			base = base * 10
		}
		if (i*i)%base == i {
			count++
		}
	}
	fmt.Println(count)
}
