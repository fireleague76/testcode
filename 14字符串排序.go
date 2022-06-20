package main

import (
	"fmt"
	"sort"
)

func main() {
	num := 0
	fmt.Scan(&num)
	slice := make([]string, num)
	for i := 0; i < num; i++ {
		s := ""
		fmt.Scan(&s)
		slice[i] = s
	}
	sort.Strings(slice)
	for i := 0; i < num; i++ {
		fmt.Println(slice[i])
	}
}
