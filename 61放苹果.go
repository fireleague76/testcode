package main

import "fmt"

func main() {
	m, n := 0, 0
	fmt.Scan(&m, &n)
	fmt.Println(f1(m, n))
}
func f1(m, n int) int {
	if m <= 1 || n == 1 {
		return 1
	} else if m < n {
		return f1(m, m)
	} else {
		return f1(m, n-1) + f1(m-n, n)
	}
}
