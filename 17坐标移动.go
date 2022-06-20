package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := ""
	x := 0
	y := 0
	fmt.Scan(&s)
	s1 := strings.Split(s, ";")
	for i := 0; i < len(s1); i++ {

		if len(s1[i]) >= 2 {
			s2 := s1[i]
			if s2[0] == 'A' || s2[0] == 'S' || s2[0] == 'D' || s2[0] == 'W' {
				s3 := s2[1:]
				num, err := strconv.Atoi(s3)
				if err == nil {
					switch s2[0] {
					case 'A':
						x -= num
					case 'S':
						y -= num
					case 'D':
						x += num
					case 'W':
						y += num
					}
				}
			}
		}
	}
	fmt.Printf("%d,%d", x, y)
}
