package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := ""
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	s = input.Text()
	str := make([]byte, len(s))
	ans := make([]byte, 0)
	copy(str, s)
	for i := 0; i < 26; i++ {
		for k, v := range str {
			if int(v) == 'a'+i || int(v) == 'A'+i {
				ans = append(ans, v)
				str[k] = 'a'
			}
		}
	}
	k := 0
	for i := 0; i < len(str); i++ {
		if str[i] == 'a' {
			str[i] = ans[k]
			k++
		}
	}
	fmt.Println(string(str))
}
