package main

import (
	"fmt"
	"strings"
)

func main() {
	s := ""
	fmt.Scan(&s)
	s1 := strings.Split(s, "")
	for i := 0; i < len(s); i++ {
		if strings.Count(s, s1[i]) == 1 {
			fmt.Println(s1[i])
			break
		}

		if i == len(s1)-1 && strings.Count(s, s1[i]) != 1 {
			fmt.Println(-1)
		}
	}
}
