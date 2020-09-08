package main

import (
	"strings"
)

func strStr(haystack string, needle string) int {
	if haystack == "" || len(haystack) < len(needle) {
		return -1
	}

	if needle == "" {
		return 0
	}

	return strings.Index(haystack,needle)
}
func main() {
	haystack := "aaaaa"
	needle := "bba"
	strStr(haystack,needle)
}