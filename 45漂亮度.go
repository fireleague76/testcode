package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		s := ""
		fmt.Scan(&s)
		m := make(map[byte]int)
		for j := 0; j < len(s); j++ {
			m[s[j]]++
		}
		slice := make([]int, 0)
		for _, v := range m {
			slice = append(slice, v)
		}
		sort.Ints(slice)
		k := 26
		sum := 0
		for i := 0; i < len(slice); i++ {
			n := slice[len(slice)-1-i]
			sum += n * k
			k--
		}
		fmt.Println(sum)
	}
}
