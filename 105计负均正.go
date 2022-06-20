package main

import "fmt"

func main() {
	count, n := 0, 0
	var sum float64 = 0
	for {
		var num float64
		_, err := fmt.Scan(&num)
		if err != nil {
			break
		}
		if num < 0 {
			n++
		} else {
			sum += num
			count++
		}
	}

	if count > 0 {
		avg := sum / float64(count)
		fmt.Println(n)
		fmt.Printf("%.1f", avg)
	} else {
		fmt.Println(n)
		fmt.Println("0.0")
	}
}
