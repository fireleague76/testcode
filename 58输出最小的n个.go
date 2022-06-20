package main

import (
	"fmt"
	"sort"
)

func main() {
	s := make([]int, 0)
	m := 0
	n := 0
	fmt.Scan(&m)
	fmt.Scan(&n)
	for i := 0; i < m; i++ {
		a := 0
		fmt.Scan(&a)
		s = append(s, a)
	}
	sort.Ints(s)
	for i := 0; i < n; i++ {
		fmt.Print(s[i], " ")
	}
}
