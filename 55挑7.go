package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 0
	num := 0
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		s := strconv.Itoa(i)
		if strings.Contains(s, "7") {
			num++
			continue
		} else if i%7 == 0 {
			num++
		}
	}
	fmt.Println(num)
}
