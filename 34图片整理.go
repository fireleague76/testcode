package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	s := ""
	res := ""
	fmt.Scan(&s)
	s1 := strings.Split(s, "")
	sort.Strings(s1)
	for i := 0; i < len(s1); i++ {
		res = res + s1[i]
	}
	fmt.Println(res)
}
