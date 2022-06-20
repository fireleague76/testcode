package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Scan(&num)
	precise := 1e-3
	t := num
	for math.Abs(t*t*t-num) > precise {
		t = t - (t*t*t-num)/(3*t*t)
	}
	fmt.Printf("%.1f", t)
}
