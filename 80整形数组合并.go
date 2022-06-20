package main

import (
	"fmt"
	"sort"
)

func main() {
	m := 0
	fmt.Scan(&m)
	slice1 := make([]int, 0)
	for i := 0; i < m; i++ {
		a := 0
		fmt.Scan(&a)
		slice1 = append(slice1, a)
	}
	n := 0
	fmt.Scan(&n)
	slice2 := make([]int, 0)
	for i := 0; i < n; i++ {
		a := 0
		fmt.Scan(&a)
		slice2 = append(slice2, a)
	}
	slice1 = append(slice1, slice2...)
	sort.Ints(slice1)
	if slice1[0] == slice1[1] {

	} else {
		fmt.Print(slice1[0])
	}
	for i := 1; i < len(slice1); i++ {
		if slice1[i] == slice1[i-1] {
			i = i + 1
			fmt.Print(slice1[i])
		} else {
			fmt.Print(slice1[i])
		}
	}
}
