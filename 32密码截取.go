package main

import "fmt"

func main() {
	s := ""
	fmt.Scan(&s)
	max := ""
	res := ""
	for i := 0; i < len(s); i++ {
		res = long(i, i, s)
		//res = long(i, i+1, s)
		if len(res) > len(max) {
			max = res
		}
	}
	for i := 0; i < len(s); i++ {
		//res = long(i, i, s)
		res = long(i, i+1, s)
		if len(res) > len(max) {
			max = res
		}
	}
	fmt.Println(len(max))
}
func long(left, right int, s string) string {
	res := ""
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			res = s[left : right+1]
		} else {
			break
		}
		left--
		right++
	}
	return res
}
