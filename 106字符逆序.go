package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := ""
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s = scanner.Text()
	slice := strings.Split(s, "")
	for i := len(slice) - 1; i >= 0; i-- {
		fmt.Print(slice[i])
	}
}
