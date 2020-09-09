package main

import (
	"fmt"
	"strings"
)

func main() {
	ss := "aabcaabc"
	fmt.Println(repeatedSubstringPattern1(ss))
}

// 扩展成2个字符串
func repeatedSubstringPattern1(s string) bool {
	new_str := s+s
	// 去头去尾
	match_str := new_str[1:len(new_str)-1]
	if strings.Contains(match_str,s){
		return true
	}
	return false
}
