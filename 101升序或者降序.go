package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 0
	slice := make([]int, 0)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		num := 0
		fmt.Scan(&num)
		slice = append(slice, num)
	}
	s := 0
	fmt.Scan(&s)
	if s == 0 {
		sort.Ints(slice)
		for i := 0; i < len(slice)-1; i++ {
			fmt.Print(slice[i], " ")
		}
		fmt.Print(slice[len(slice)-1])
	} else if s == 1 {
		sort.Ints(slice)
		for i := len(slice) - 1; i > 0; i-- {
			fmt.Print(slice[i], " ")
		}
		fmt.Print(slice[0])
	}
}
