package main

import "fmt"

func main() {
	str1, str2 := "", ""
	fmt.Scan(&str1, &str2)
	fmt.Println(len1(str1, str2))
}

func len1(str1, str2 string) int {
	col := len(str1) + 1
	val := len(str2) + 1
	res := make([][]int, 0)
	for i := 0; i < col; i++ {
		//res[i] = make([]int, val)
		res = append(res, make([]int, val))
	}
	for i := 0; i < col; i++ {
		res[i][0] = i
	}
	for i := 0; i < val; i++ {
		res[0][i] = i
	}

	for i := 1; i < col; i++ {
		for j := 1; j < val; j++ {
			if str1[i-1] == str2[j-1] {
				res[i][j] = res[i-1][j-1]
			} else {
				res[i][j] = min(res[i-1][j], res[i-1][j-1], res[i][j-1]) + 1
			}
		}
	}
	return res[col-1][val-1]
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	} else {
		return c
	}
}
