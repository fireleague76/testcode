package main

import "fmt"

func main() {
	a := 0
	b := 0
	fmt.Scan(&a)
	fmt.Scan(&b)
	c := max(a, b)
	for {
		if c%a == 0 && c%b == 0 {
			break
		}
		c++
	}
	fmt.Println(c)

}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
