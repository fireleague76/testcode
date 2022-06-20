package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
		a := ""
		slice := make([]byte, 0)
		_, err := fmt.Scan(&a)
		if err != nil {
			break
		}
		l := 0
		res := ""
		for i := 0; i < len(a); i++ {
			if a[i] < '0' || a[i] > '9' {
				slice = append(slice, ' ')
			} else {
				slice = append(slice, a[i])
			}
		}
		str := strings.Split(string(slice), " ")
		for _, q := range str {
			if len(q) > l {
				l = len(q)
			}
		}
		for _, q := range str {
			if len(q) == l {
				res = res + q
			}
		}
		fmt.Printf("%s,%d\n", res, l)
	}
}
