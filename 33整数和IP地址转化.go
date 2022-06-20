package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	ip := ""
	fmt.Scan(&ip)
	ips := strings.Split(ip, ".")
	res := 0
	for _, i := range ips {
		s1, _ := strconv.Atoi(i)
		res = res*256 + s1
	}
	//s1,_:=strconv.Atoi(s)
	fmt.Println(res)

	s := 0
	fmt.Scan(&s)
	s1 := ""
	for i := 0; i < 4; i++ {
		if i == 0 {
			s1 = strconv.Itoa(s % 256)
		} else {
			s1 = strconv.Itoa(s%256) + "." + s1
		}
		s = s / 256
	}
	fmt.Println(s1)
}
