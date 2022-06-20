package main

import "fmt"

func main() {
	num := 0
	slice := make([]int, 0)
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		n := 0
		fmt.Scan(&n)
		slice = append(slice, n)
	}
	min := len(slice)
	for i := 0; i < len(slice); i++ {
		res := 0
		for j := 1; j < i; j++ {
			if slice[j] >= slice[i] || slice[j] >= slice[j-1] {
				res++
			}
		}
		for k := i; k < len(slice); k++ {
			if slice[k] >= slice[i] || slice[k] >= slice[k-1] {
				res++
			}
		}
		if res < min {
			min = res
		}
	}
	fmt.Println(min)
}

func a(slice []int) int {
	return 1
}
