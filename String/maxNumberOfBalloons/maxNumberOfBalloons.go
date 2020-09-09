package main

import (
	"fmt"
)

func main(){
	var text  string = "nlaebolko"
	nusm := maxNumberOfBalloons(text)
	fmt.Println(nusm)

}
// 方法1
func maxNumberOfBalloons1(text string) int {
	// balloon
	// 创建一个map 把所有的字符都放进去
	mapdict := make(map[string]int)
	for i:=0;i<len(text);i++{
		mapdict[string(text[i])]++
	}

	// 整体去处理
	count :=0
	for mapdict["b"] >=1 && mapdict["a"] >= 1 && mapdict["l"] >= 2 && mapdict["o"] >= 2 && mapdict["n"] >= 1{
		count++
		mapdict["b"] --
		mapdict["a"] --
		mapdict["l"] -=2
		mapdict["n"] --
		mapdict["o"] -=2
	}
	return count
}

// 方法2 使用字节数组看看速度会不会有提升
func maxNumberOfBalloons2(text string)int{
	mapdict :=  make(map[byte]int)
	for i:=0;i<len(text);i++{
		mapdict[text[i]]++
	}

	count :=0
	for mapdict[byte('b')] >1 && mapdict[byte('a')] >=1 && mapdict[byte('l')] >= 2 && mapdict[byte('o')] >=2 && mapdict[byte('n')] >=1{
		count++
		mapdict[byte('b')] --
		mapdict[byte('a')] --
		mapdict[byte('l')] -=2
		mapdict[byte('n')] --
		mapdict[byte('o')] -=2
	}
	return count
}

// 方法3 使用字节
//将字母遍历存储到数组中统计字母数量，然后按balloon遍历数组，边遍历边递减字母的数量，遍历一次即代表可以拼成一个balloon，记录次数，当某个字母的数量小于0时，返回总数
func maxNumberOfBalloons(text string) int {
	charCount := [26]int{}
	charArray := []byte(text)
	for _, v := range charArray {
		charCount[v - 'a']++
	}
	charCount['l' - 'a'] /= 2
	charCount['o' - 'a'] /= 2
	min := len(charArray)
	for _, v := range []rune("balloon") {
		if charCount[v - 'a'] < min {
			min = charCount[v - 'a']
		}
	}
	return min
}

