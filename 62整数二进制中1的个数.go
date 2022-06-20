package main

import "fmt"

func main() {
	for {
		n := 0
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}

		count := 0
		for n != 0 {
			if n&1 == 1 {
				count++
			}
			n >>= 1
		}
		fmt.Println(count)
	}
}
