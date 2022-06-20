package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := ""
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s = input.Text()
	l1 := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			l1 = append(l1, s[i])
		} else if s[i] >= 'A' && s[i] <= 'Z' {
			l1 = append(l1, s[i])
		} else {
			l1 = append(l1, ' ')
		}
	}
	s1 := strings.Split(string(l1), " ")
	for i := len(s1) - 1; i >= 0; i-- {
		fmt.Print(s1[i], " ")
	}
	//fmt.Println(s1)
}

/* package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {

		s := input.Text()
		l1 := make([]byte, 0)
		for i := 0; i < len(s); i++ {
			if s[i] >= 'a' && s[i] <= 'z' {
				l1 = append(l1, s[i])
			} else if s[i] >= 'A' && s[i] <= 'Z' {
				l1 = append(l1, s[i])
			} else {
				l1 = append(l1, ' ')
			}
		}
		s = string(l1)
		li := strings.Split(s, " ")
		l := len(li)

		for i := 0; i < len(li); i++ {

			fmt.Printf("%s ", li[l-i-1])
		}
		fmt.Printf(string(l1))

		fmt.Printf("\n")

	}
} */
