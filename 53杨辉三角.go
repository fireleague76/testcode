package main

import "fmt"

func main() {
	n := 0
	fmt.Scanln(&n)

	if n < 3 {
		fmt.Println(-1)
		return
	}

	switch n % 4 {
	case 1, 3:
		fmt.Println(2)
	case 2:
		fmt.Println(4)
	case 0:
		fmt.Println(3)
	}
}
