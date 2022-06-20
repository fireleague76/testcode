package main

import "fmt"

func main() {
	for {
		var m, n int
		_, err := fmt.Scan(&m, &n)
		if err != nil {
			return
		}
		res := getRes(m, n)
		fmt.Println(res)
	}

}
func getRes(m, n int) int {
	if m == 0 || n == 0 {
		return 1
	}
	return getRes(m-1, n) + getRes(m, n-1)
}
