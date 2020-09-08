package main

import "fmt"

// golang判断元素是否存在数组当中

func IsContain(items []string,item string) bool{
	for _,eachItem := range items{
		if eachItem  == item{
			return true
		}
	}
	return false
}
func main() {
	ss:="hello"
	var sentence = []string{"my","words","hello"}
	if IsContain(sentence,ss){
		fmt.Println("true")
	}else {
		fmt.Println("false")
	}

}