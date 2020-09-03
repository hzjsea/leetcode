package main

import "fmt"

func MinWindows(string1 string, string2 string ) string{
	if len(string1) == 0 || len(string1) == 1 || len(string1)<len(string2){
		return string1
	}

	// 定义窗口
	minwindows := make(map[string]int)
	// 初始化窗口
	for i:=0 ;i<len(string2);i++{
		minwindows[string(string2[i])]++
	}
	windowsize := len(minwindows)

	// 移动窗口定义边框
	moveleft,moveright,c := 0,0,0
	ans := ""



	for moveright < len(string1){
		// 不断的移动右边框，匹配到相同的元素之后 减少数量
		minwindows[string(string1[moveright])] --
		fmt.Println(minwindows,c)
		if minwindows[string(string1[moveright])] == 0{
			c ++
		}
		for c == windowsize && minwindows[string(string1[moveleft])] < 0 {
			minwindows[string(string1[moveleft])]++
			moveleft ++
		}
		if c == windowsize {
			if len(ans) == 0 || moveright-moveleft+1 < len(ans) {
				ans = string1[moveleft : moveright+1]
			}
		}
		moveright ++
	}
	return ans
}

func minWindow(s string, t string) string {
	ls, lt := len(s), len(t)
	if ls < lt {
		return ""
	}

	// 窗口里存的是t中字符应该出现的次数
	// 正数表示该字符还缺的出现次数，0表示刚好出现，负数表示s中字符出现的次数多于t中字符出现次数
	windowMap := map[byte]int{}
	// 初始化窗口
	for i := 0; i < lt; i++ {
		windowMap[t[i]]++
	}
	windowSize := len(windowMap)
	// 其实在go语言里map有零值的概念，这块代码可以不要
	// 在其他语言，比如Java的HashMap没有零值概念，需要先初始化一下所有s中的字符出现次数
	// for i := 0; i < ls; i++ {
	// 	if _, ok := windowMap[s[i]]; !ok {
	// 		windowMap[s[i]] = 0
	// 	}
	// }

	left, right := 0, 0
	// 窗口中已经包含T的不同字符的种类
	c := 0
	ans := ""

	for right < ls {
		// 窗口右边界移动，扩大窗口
		windowMap[s[right]]--
		fmt.Println(windowMap,c)
		// 统计窗口中已经包含的T中的不同字符的种类
		if windowMap[s[right]] == 0 {
			c++
		}

		// c==windowSize说明窗口已经包含所有T中的字符
		for c == windowSize && windowMap[s[left]] < 0 {
			windowMap[s[left]]++
			left++
		}
		if c == windowSize {
			if len(ans) == 0 || right-left+1 < len(ans) {
				ans = s[left : right+1]
			}
		}
		right++
	}
	return ans
}

func main() {
	var S = "ADOBECODEBANC"
	var T = "ABC"

	fmt.Println(MinWindows(S,T))
	fmt.Println(minWindow(S,T))
}
