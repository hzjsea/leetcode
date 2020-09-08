package main

import "fmt"

func main() {
	var ss []byte = []byte("hello")
	reverseString(ss)

}

func reverseString(s []byte)  {
	// 反转字符串
	// 获取字符串长度
	slen := len(s)
	// 获取中间值 取商
	smid := len(s)/2 // 2
	// 前后替换
	for i:=0;i<smid;i++{
		s[i], s[slen-i-1] = s[slen-i-1], s[i] // 前后替换位置
		fmt.Println(string(s[:]))
	}

}
