package main

import (
	"fmt"
	"strings"
)

func main() {
	for {
		str1, str2 := "", ""
		_, err := fmt.Scan(&str1)
		if err != nil || str1 == "" {
			break
		}
		fmt.Scan(&str2)
		if len(str1) > len(str2) {
			str1, str2 = str2, str1
		}
		fmt.Println(long(str1, str2))
	}
}

func long(str1, str2 string) string {
	l := len(str1)
	for i := 0; i < l; i++ {
		for j := 0; j <= i; j++ {
			new := string(str1[j : l-i+j])
			if strings.Index(str2, new) != -1 {
				return new
			}
		}
	}
	return ""
}
