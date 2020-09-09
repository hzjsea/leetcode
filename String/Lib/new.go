package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 看看切片是什么类型
	var ss = make([]string,10,10)
	ss = append(ss, "1")
	ss = append(ss,"2")
	fmt.Println(ss)
	for _,i :=range ss{
		fmt.Println(reflect.TypeOf(i))
	}


}