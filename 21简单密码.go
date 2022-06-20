package main

import (
	"fmt"
)

func main() {
	s := ""
	fmt.Scan(&s)
	slice := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		slice[i] = s[i]
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == 'a' || slice[i] == 'b' || slice[i] == 'c' {
			slice[i] = '2'
		} else if slice[i] == 'd' || slice[i] == 'e' || slice[i] == 'f' {
			slice[i] = '3'
		} else if slice[i] == 'g' || slice[i] == 'h' || slice[i] == 'i' {
			slice[i] = '4'
		} else if slice[i] == 'j' || slice[i] == 'k' || slice[i] == 'l' {
			slice[i] = '5'
		} else if slice[i] == 'm' || slice[i] == 'n' || slice[i] == 'o' {
			slice[i] = '6'
		} else if slice[i] == 'p' || slice[i] == 'q' || slice[i] == 'r' || slice[i] == 's' {
			slice[i] = '7'
		} else if slice[i] == 't' || slice[i] == 'u' || slice[i] == 'v' {
			slice[i] = '8'
		} else if slice[i] == 'w' || slice[i] == 'x' || slice[i] == 'y' || slice[i] == 'z' {
			slice[i] = '9'
		} else if slice[i] >= 'A' && slice[i] < 'Z' {
			slice[i] += 33
		} else if slice[i] == 'Z' {
			slice[i] = 'a'
		}
	}
	fmt.Println(string(slice))
}
