package main

import "fmt"

func main() {
	n := 0
	count := 0
	var sum float64 = 0
	var num float64 = 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		a := 0
		fmt.Scan(&a)
		if a < 0 {
			count++
		}
		if a > 0 {
			sum = sum + float64(a)
			num++
		}
	}
	if num == 0 {
		fmt.Printf("%d 0.0", count)
	} else {
		fmt.Printf("%d %.1f", count, sum/num)
	}

}
