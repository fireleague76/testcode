package main

import (
	"fmt"
	"strings"
)

func main() {
	s := ""
	fmt.Scan(&s)
	count := 0
	//slice:=make([]string,0)
	slice := strings.Split(s, "")
	for i := 0; i < len(slice); i++ {
		if slice[i] >= "A" && slice[i] <= "Z" {
			count++
		}
	}
	fmt.Println(count)
}
