package main

import (
	"fmt"
	"strings"
	"unicode"
)
// strings字符串操作

//字符串长度；
//求子串；
//是否存在某个字符或子串；
//子串出现的次数（字符串匹配）；
//字符串分割（切分）为[]string；
//字符串是否有某个前缀或后缀；
//字符或子串在字符串中首次出现的位置或最后一次出现的位置；
//通过某个字符串将[]string 连接起来；
//字符串重复几次；
//字符串中子串替换；
//大小写转换；
//Trim 操作；
//...

func main()  {
	//ss := "hello"
	//fmt.Println(getStringLen(ss))
	//返回一个空格
	fmt.Println(unicode.IsSpace)

	ss:="hello world"
	fmt.Println(strings.Fields(ss))

	ss="hello - world"
	fmt.Println(strings.FieldsFunc(ss,unicode.IsSpace))


	// 分割字符串
	ss = "hello,world"
	fmt.Println(strings.Split(ss,",")) // ["hello" "world"]
	fmt.Println(strings.SplitAfter(ss,",")) // ["hello," "world"]
	// 分割前面几个
	fmt.Println(strings.SplitAfterN(ss,",",1))  // [hello,world]


	// 字符串是否有某个前缀(prefix)或后缀(suffix)
	if strings.HasPrefix(ss,"hello"){
		fmt.Println("hello is head")
	}
	if strings.HasSuffix(ss,"world"){
		fmt.Println("world is tail")
	}


	qq := []string{"one","two","three"}
 	fmt.Println(JoinString(qq))

	for k,v := range("ddsdsds"){
		fmt.Println(k,v)
	}



}

// 获取字符串长度
func getStringLen(ss string)  int {
	return  len(ss)
}

// 字符串比较  不忽略大小写 返回int类型
func CompareString(ss string, ss2 string) int{
	return strings.Compare(ss,ss2)
}

// 字符串 忽略大小写 返回bool类型
func EqualFoldString(ss string, ss2 string) bool{
	return strings.EqualFold(ss,ss2)
}

// 子串在字符串中返回true
func ContainsSubstr(ss string , sub string) bool{
	return strings.Contains(ss,sub)
}
// 字符在字符串中返回true
func ContainsAnyChar(ss string, char string) bool{
	return strings.ContainsAny(ss,char)
}
// rune在字符串中返回true
func ContainsRuneChar(ss string,char rune) bool{
	return strings.ContainsRune(ss,char)
}

// 子串出现次数 ( 字符串匹配 )
//func Count(s, substr string) int {
//	// special case
//	if len(substr) == 0 {
//		return utf8.RuneCountInString(s) + 1
//	}
//	if len(substr) == 1 {
//		return bytealg.CountString(s, substr[0])
//	}
//	n := 0
//	for {
//		i := Index(s, substr)
//		if i == -1 {
//			return n
//		}
//		n++
//		s = s[i+len(substr):]
//	}
//}
// 计算某一个字符char 在字符串中出现的次数
func CountString(ss string,char string ) int{
	return strings.Count(ss,char)
}

// 字符串根据空格剪切
func FieldsString(ss string) []string  {
	return FieldsString(ss)
}

// 字符或子串在字符串中首次出现的位置
func IndexChar(ss string) int {
	return strings.Index(ss,"h")
}

// 将字符串数组（或 slice）连接起来可以通过 Join 实现，函数签名如下：
func JoinString(ss []string) string{
	return strings.Join(ss,"-")
}

// 字符串重复几次
func RepeatString(ss string) string{
	return strings.Repeat(ss,2)
}

//https://books.studygolang.com/The-Golang-Standard-Library-by-Example/chapter02/02.1.html