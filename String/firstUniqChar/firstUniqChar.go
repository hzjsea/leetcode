package main

import (
	"fmt"
	"strings"
)

func main() {
	ss := "cc"
	fmt.Println(firstUniqChar(ss))
}


// 这种时间上更快
//func firstUniqChar(s string) int {
//	last := [26]int{}
//
//	for i, v := range s {
//		last[v-'a'] = i
//	}
//
//	for i, v := range s {
//		if last[v-'a'] == i {
//			return i
//		} else {
//			last[v-'a'] = -1
//		}
//	}
//	return -1
//}


func firstUniqChar(s string) int {
	if len(s) <= 0{
		return -1
	}
	// go里面可以把它们放到一个字典里面，然后同样的字符 累加
	// 创建数组
	maplist := map[byte]int{}
	indexlist := []int{}
	minIndex :=0
	// 循环进行累加
	for i:=0;i<len(s);i++{
		maplist[s[i]]++
	}
	// 读取一个值
	for k,v := range maplist{
		// 因为map是无序的 所以还要比较一下大小
		if v == 1{
			indexlist = append(indexlist,strings.IndexByte(s,k))
		}
	}
	if len(indexlist) > 0{
		minIndex =indexlist[0]
	}else {
		return 0
	}

	fmt.Println(indexlist)
	for i:=0;i<len(indexlist);i++{
		if indexlist[i] < minIndex{
			minIndex = indexlist[i]
		}
	}
	return minIndex
}