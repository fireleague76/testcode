package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := ""
	num := 0
	a := 0
	kongge := 0
	other := 0
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s = input.Text()
	slice := strings.Split(s, "")
	s1 := make([]byte, len(s))
	for i := 0; i < len(s1); i++ {
		s1[i] = s[i]
	}
	for i := 0; i < len(slice); i++ {
		if s1[i] >= '0' && s1[i] <= '9' {
			num++
		} else if s1[i] >= 'a' && s1[i] <= 'z' {
			a++
		} else if s1[i] == ' ' {
			kongge++
		} else {
			other++
		}
	}
	fmt.Println(a)
	fmt.Println(kongge)
	fmt.Println(num)
	fmt.Println(other)
}
