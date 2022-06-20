package main

import "fmt"

func main() {
	s1 := ""
	fmt.Scan(&s1)
	s2 := ""
	fmt.Scan(&s2)
	max := 1
	for i := 0; i < len(s2)-len(s1); i++ {
		n := 0
		if s2[i] == s1[0] {
			for j := 0; j < len(s1); j++ {
				if s2[i+j] == s1[j] {
					n++
				} else {
					break
				}
			}

		}
		max = big(max, n)
	}
	fmt.Println(max)
}
func big(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
