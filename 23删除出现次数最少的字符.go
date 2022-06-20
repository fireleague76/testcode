package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := ""
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s = input.Text()
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}
	min := len(m)
	for _, v := range m {
		if v < min {
			min = v
		}
	}
	for i := 0; i < len(s); i++ {
		if m[rune(s[i])] == min {

		} else {
			fmt.Print(string(s[i]))
		}
	}
}
