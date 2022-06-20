package main

import "fmt"

func main() {
	num := 0
	sum := 0
	slice := make([]int, 0)
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		a := 0
		fmt.Scan(&a)
		slice = append(slice, a)
	}
	dp := make([]int, len(slice))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if slice[i] > slice[j] && dp[j]+1 > dp[j] {
				dp[i] = dp[j] + 1
			}
		}
		sum = max(sum, dp[i])
	}
	fmt.Println(sum)
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
