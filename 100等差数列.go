package main

import "fmt"

func main() {
	a := 0
	fmt.Scan(&a)
	sum := 2*a + a*(a-1)/2*3
	fmt.Println(sum)
}
