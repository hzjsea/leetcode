package main

import "strings"

func main() {

}

func detectCapitalUse(word string) bool {
	//  输出三种类型比较一下就可以了
	allUpper := strings.ToUpper(word)
	alllower := strings.ToLower(word)
	TitleString := strings.ToTitle(word)

	if word == alllower || word == allUpper || word == TitleString {
		return true
	}
	return false
}