package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	// 判空
	if len(strs) == 0{
		return ""
	}
	prefix := strs[0]
	count := len(strs)

	for i:=1;i<count;i++{
		prefix = lcp(prefix,strs[i])
		if len(prefix) == 0{
			break
		}
	}
	return prefix
}




// 返回两个字符串中的最长前缀
func lcp(ls,ls1 string) string{
	// 因为是前缀比较，所以 公用部分肯定最短的元素为主
	length := min(ls,ls1)
	index:=0
	for index<length && ls[index] == ls1[index]{
		index ++
	}
	return ls1[:index]
}


// 取最小的长度的字符串的长度
func min(s string ,s1 string) int {
	if len(s)>len(s1){
		return len(s1)
	}else {
		return len(s)
	}
}



func main()  {
	var slist = []string{"flow","fl","fffl"}
	ss := longestCommonPrefix(slist)
	fmt.Println(ss)
}
