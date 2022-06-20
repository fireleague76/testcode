package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	str := ""
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			str += string(s[i])
		} else {
			if i == 0 || (s[i-1] < '0' || s[i-1] > '9') {
				str += "*"
			}
			str += string(s[i])
			if i == len(s)-1 || (s[i+1] < '0' || s[i+1] > '9') {
				str += "*"
			}
		}
	}
	fmt.Println(str)
}
