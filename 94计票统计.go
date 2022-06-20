package main

import "fmt"

func main() {
	invalid := 0
	m := make(map[string]int)
	n := 0
	fmt.Scan(&n)
	slice := make([]string, 0)
	for i := 0; i < n; i++ {
		s := ""
		fmt.Scan(&s)
		slice = append(slice, s)
		m[s] = 0
	}
	num := 0
	fmt.Scan(&num)
	for i := 0; i < num; i++ {
		s1 := ""
		fmt.Scan(&s1)
		_, ok := m[s1]
		if ok {
			m[s1]++
		} else {
			invalid++
		}
	}
	for _, v := range slice {
		fmt.Printf("%s : %d\n", v, m[v])
	}
	fmt.Printf("invalid : %d\n", invalid)
}
