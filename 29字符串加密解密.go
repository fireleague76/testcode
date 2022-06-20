package main

import "fmt"

func main() {
	s := ""
	fmt.Scan(&s)
	slice := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		slice[i] = s[i]
	}
	for i := 0; i < len(s); i++ {
		if slice[i] >= 'A' && slice[i] < 'Z' {
			slice[i] += 33
		} else if slice[i] >= 'a' && slice[i] < 'z' {
			slice[i] -= 31
		} else if slice[i] == 'z' {
			slice[i] = 'A'
		} else if slice[i] == 'Z' {
			slice[i] = 'a'
		} else if slice[i] >= '0' && slice[i] < '9' {
			slice[i] += 1
		} else if slice[i] == '9' {
			slice[i] = '0'
		}
	}
	fmt.Println(string(slice))

	s1 := ""
	fmt.Scan(&s1)
	slice1 := make([]byte, len(s1))
	for i := 0; i < len(s1); i++ {
		slice1[i] = s1[i]
	}
	for i := 0; i < len(s1); i++ {
		if slice1[i] > 'A' && slice1[i] <= 'Z' {
			slice1[i] += 31
		} else if slice1[i] > 'a' && slice1[i] <= 'z' {
			slice1[i] -= 33
		} else if slice1[i] == 'a' {
			slice1[i] = 'Z'
		} else if slice1[i] == 'A' {
			slice1[i] = 'z'
		} else if slice1[i] > '0' && slice1[i] <= '9' {
			slice1[i] -= 1
		} else if slice1[i] == '0' {
			slice1[i] = '9'
		}
	}
	fmt.Println(string(slice1))
}
