package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	slice := make([]int, n, n)
	slice[0] = 1
	for i := 2; i <= n; i++ {
		slice[i-1] = slice[i-2] + i
	}
	for i := n; i >= 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print(slice[j+n-i]+i-n, " ")
		}
		fmt.Printf("\n")
	}

}
