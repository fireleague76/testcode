package main

import "fmt"

func main() {
	slice := make([]byte, 0, 26)
	for i := 0; i < 26; i++ {
		slice = append(slice, byte('a'+i))
	}
	fmt.Println(string(slice))
}
