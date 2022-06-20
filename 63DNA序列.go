package main

import "fmt"

func main() {
	s := ""
	fmt.Scan(&s)
	n := 0
	var max float64 = 0
	m := make(map[string]float64)
	fmt.Scan(&n)
	for i := 0; i < len(s)-n; i++ {
		new := s[i : i+n]
		num := 0
		for j := 0; j < len(new); j++ {
			if new[j] == 'C' || new[j] == 'G' {
				num++
			}
		}
		var res float64
		res = float64(float64(num) / float64(len(s)))
		m[new] = res
		if res > max {
			max = res
		}
	}
	for i := 0; i < len(m); i++ {
		if m[s[i:i+n]] == max {
			fmt.Println(s[i : i+n])
			break
		}
	}
	if n == len(s) {
		fmt.Println(s)
	}

}
