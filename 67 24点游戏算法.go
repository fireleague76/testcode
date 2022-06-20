package main

import (
	"fmt"
)

var flag = make([]int, 4)
var num = make([]int, 4)

func main() {
	for {
		n, _ := fmt.Scan(&num[0], &num[1], num[2], num[3])
		if n == 0 {
			break
		}
		flag = make([]int, 4)
		fmt.Println(dfs(0))
	}
}

func dfs(sum int) bool {
	if sum == 24 {
		return true
	}
	for i := 0; i < 4; i++ {
		if flag[i] == 0 {
			flag[i] = 1
			if dfs(sum + num[i]) {
				return true
			} else if dfs(sum - num[i]) {
				return true
			} else if dfs(sum * num[i]) {
				return true
			} else if dfs(sum/num[i]) && sum%num[i] == 0 {
				return true
			}
			flag[i] = 0
		}

	}
	return false
}
