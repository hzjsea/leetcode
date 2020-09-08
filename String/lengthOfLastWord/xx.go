package main

import (
	"strings"
)

func lengthOfLastWord(s string) int {
	// 判空
	if s == "" {
		return 0
	}
	// 先按照空格剪切
	slsit := strings.Split(s," ")
	slsit = slsit[len(slsit)-1:]
	return len(slsit[0])
}

func main() {

	ss := "hello world"
	lengthOfLastWord(ss)

}