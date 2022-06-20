package main

import "fmt"

func main() {
	for {
		emptynum := 0
		fmt.Scan(&emptynum)
		if emptynum == 0 {
			break
		}
		fmt.Println(emptynum / 2)

	}
}
