package main

import (
	"fmt"
	"math"
	"strings"
)

// 滑动窗口算法
// 滑动窗口算法可以用以解决数组/字符串的子元素问题，它可以将嵌套的循环问题，转换为单循环问题，降低时间复杂度

// 字符串暴力解法 多个循环

//给定一个整数数组，计算长度为 'k' 的连续子数组的最大总和。
//输入：arr [] = {100,200,300,400}
//k = 2
//
//输出：700
//
//解释：300 + 400 = 700

func main() {
	max := getmax()
	fmt.Println(max)
	max = getmax2()
	fmt.Println(max)


	res := getmax3()
	fmt.Println(res)
}

func getmax() int{
	arr := [4]int{100,200,300,400}
	max := arr[0]+arr[1]
	for i:=0;i<len(arr);i++{
		for j:=i+1;j<len(arr);j++{
			temp := arr[i]+arr[j]
			if max<temp{
				max = temp
			}
		}
	}
	return max
}


// 滑动窗口算法，规定一个窗口向后移动 减少循环嵌套  O(n)
func getmax2() int {
	arr := [4]int{100,200,300,400}
	// 定义一个窗口
	maxSum := 0
	maxSum = arr[0]+arr[1]

	sum := maxSum
	for i:=2;i<len(arr);i++{
		sum+=arr[i]-arr[i-2] // 新窗口的和等于新加入的值-新移除的值
	}
	return maxSum
}

// 滑动窗口算法2 字符串最大长度计算
// 给定一个字符串S 和一个字符串T 请在 S 中找出包含 T 所有字母的最小子串
// 输入: S = "ADOBECODEBANC", T = "ABC"
// 输出: "BANC"  O（s + t）


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

// 滑动窗口3
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

func lengthOfLongestSubstring(s string) int {
	var Length int
	var s1 string
	left := 0
	right := 0
	s1 = s[left:right]
	for ; right < len(s); right++ {

		if index := strings.IndexByte(s1, s[right]); index != -1 {
			left += index + 1
		}
		s1 = s[left : right+1]
		if len(s1) > Length {
			Length = len(s1)
		}
	}
	return Length
}


// 滑动窗口
//给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引
