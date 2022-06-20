package main

import "fmt"

func main() {
	for {
		s := ""
		_, err := fmt.Scan(&s)
		if err != nil {
			break
		}
		if len(s) < 8 {
			fmt.Println("NG")
		} else if abc(s) == false {
			fmt.Println("NG")
		} else if chongfu(s) == false {
			fmt.Println("NG")
		} else {
			fmt.Println("OK")
		}
	}
}
func abc(s string) bool {
	s1 := []byte(s)
	a, b, c, d := 0, 0, 0, 0
	for i := 0; i < len(s1); i++ {
		if s1[i] >= 'a' && s1[i] <= 'z' {
			a = 1
		} else if s1[i] >= '0' && s1[i] <= '9' {
			b = 1
		} else if s1[i] >= 'A' && s1[i] <= 'Z' {
			c = 1
		} else {
			d = 1
		}
	}
	if a+b+c+d >= 3 {
		return true
	} else {
		return false
	}
}
func chongfu(s string) bool {
	m := make(map[string]bool)
	for i := 0; i < len(s)-2; i++ {
		if _, ok := m[s[i:i+3]]; !ok {
			m[s[i:i+3]] = true
		} else {
			return false
		}
	}
	return true
}
