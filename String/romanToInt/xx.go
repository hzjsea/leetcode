package main

import (
	"fmt"
)

func romanToInt(s string)(r int) {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	pre:=0
	// 拆解字符串 然后匹配map中对应的值 然后加起来就OK了
	for i:=0;i<len(s);i++{
		cur := m[s[i]]
		if cur >= pre {
			r += cur
		} else {
			r -= cur
		}
		pre = cur
	}
	return r
}

func main() {
	res := romanToInt("IV")
	fmt.Println(res)
	//s:= "你好"
	//fmt.Println(len(s))  // 6
	//fmt.Println(utf8.RuneCountInString(s)) //2
	//fmt.Println([]rune(s)) // [20320 22909]
	//fmt.Println(len([]rune(s))) // 2
}

