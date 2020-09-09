package main

import (
	"fmt"
	"strings"
)

func main() {
	ss := "Hello, my name is John"
	fmt.Println(countSegments(ss))
}


func countSegments(s string) int {
	return len(strings.Split(s," "))
}