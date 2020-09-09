package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 拼接字符串

// 因此，这种拼接字符串的方式会导致大量的string创建、销毁和内存分配 这样的写法是错误的
func func1(ss []string) string{
	var str string
	for _,v :=range ss{
		str +=v
	}
	fmt.Println(str)
	return str
}

func func2(ss []string) string{
	var str bytes.Buffer
	for _,v := range ss {
		fmt.Fprint(&str,v)
	}
	fmt.Println(str.String())
	return str.String()
}


func func3(ss []string)string{
	var str strings.Builder
	for _,v := range ss {
		fmt.Fprint(&str,v) // 这个io stdin
	}
	fmt.Println(str.String())
	return str.String()
}

func main() {
	ss := []string{
		"A","B","C",
	}
	func1(ss)
	func2(ss)
	func3(ss)
}

