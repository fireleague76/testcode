package main

import (
	"fmt"
	"sort"
)

func main() {
	num := 0
	fmt.Scan(&num)
	slice := make([]int, num)
	for i := 0; i < num; i++ {
		n := 0
		fmt.Scan(&n)
		slice[i] = n
	}
	res := make([]int, 0)
	m := make(map[int]bool)
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			res = append(res, v)
			m[v] = true
		}
	}
	sort.Ints(res)
	for _, v := range res {
		fmt.Println(v)
	}
}
