package main

import (
	"fmt"
	"strings"
)

func main() {
	s := ""
	fmt.Scan(&s)
	s1 := ""
	fmt.Scan(&s1)
	output := "true"
	for _, str := range s {
		if !strings.Contains(s1, string(str)) {
			output = "false"
			break
		}
	}
	fmt.Println(output)
}
