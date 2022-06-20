package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := ""
	n := 0
	fmt.Scan(&s)
	slice := strings.Split(s, ".")
	if len(slice) != 4 {
		fmt.Println("NO")
	} else {
		for i := 0; i < len(slice); i++ {

			m, _ := strconv.Atoi(slice[i])
			if m < 0 || m > 255 {
				n++
			}
		}
		if n == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
