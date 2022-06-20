package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s := input.Text()
	slice := strings.Split(s, " ")
	for i := len(slice) - 1; i >= 0; i-- {
		fmt.Print(slice[i], " ")
	}
}
