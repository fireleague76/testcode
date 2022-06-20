package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 0
	input := make([]string, 0)
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		s := ""
		fmt.Scan(&s)
		input = append(input, s)
	}
	key := ""
	fmt.Scan(&key)
	map1 := make(map[string]int)
	slice := make([]string, 0)
	for _, v := range key {
		map1[string(v)]++
	}
	for _, v := range input {
		map2 := make(map[string]int)
		ta := 0
		if len(v) == len(key) && v != key {
			for _, a := range v {
				map2[string(a)]++
			}
			for _, v := range key {
				if map1[string(v)] != map2[string(v)] {
					ta = 1
				}
			}
			if ta == 0 {
				slice = append(slice, v)
			}
		}
	}

	sort.Strings(slice)
	num := 0
	fmt.Scan(&num)
	fmt.Println(len(slice))
	if num-1 > len(slice)-1 {
		return
	} else {
		fmt.Println(slice[num-1])
	}
}
