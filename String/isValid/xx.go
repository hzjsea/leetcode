package main

import "strings"

// 括号问题 ，用栈来解决


// 暴力解决，只要是true 就会呈现这种结构    ([{}]) 只要替换就行了

func isValid(s string) bool {
	for  {
		old := s
		s = strings.Replace(s,"()","",-1)
		s = strings.Replace(s,"[]","",-1)
		s = strings.Replace(s,"{}","",-1)
		if s == "" {
			return true
		}
		if s == old {
			return false
		}
	}
	return false

}

// 栈解决，现将前半部分压入栈中， 设定一个条件，再依次出栈进行匹配

func isValid2(s string)bool{
	hash := map[byte]byte{')':'(', ']':'[', '}':'{'}
	stack := make([]byte, 0)

	if len(s) == 0 {
		return true
	}

	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else if len(stack) > 0 && stack[len(stack)-1] == hash[s[i]] {
			stack = stack[:len(stack)-1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}