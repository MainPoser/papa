package str

import (
	"fmt"
	"strings"
)

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func reverseStr() {
	// 字符串倒序
	str := "192.168.19.100,2022::19:100"
	fmt.Println(strings.Join(reverse(strings.Split(str, ",")), ","))
}
