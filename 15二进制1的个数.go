package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := 0
	n := 0
	fmt.Scan(&num)
	s := strconv.FormatInt(int64(num), 2)
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			n++
		}
	}
	fmt.Println(n)
}
