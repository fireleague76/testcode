package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Scan(&a)
	s := fmt.Sprintf("%b", a)
	num := 1
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			j := i + 1
			for j < len(s) && s[j] == '1' {
				j++
				num = max(num, j-i)
			}
		}
	}
	fmt.Println(num)
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
