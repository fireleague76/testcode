package main

import (
	"fmt"
	"strings"
)

func main() {
	m := make(map[byte]byte)
	slice1 := make([]byte, 0, 26)
	keys := ""
	fmt.Scan(keys)
	ss := ""
	fmt.Scan(ss)
	key := strings.Split(keys, "")
	s := strings.Split(ss, "")
	for i := 0; i < 26; i++ {
		slice1 = append(slice1, byte('a'+i))
	}
	slice2 := make([]byte, 0, 26)
	/* 	for i :=key{

	   	} */
}
