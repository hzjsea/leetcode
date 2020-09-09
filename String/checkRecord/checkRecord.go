package main

import (
	"fmt"
	"strings"
)

func main() {

	ss:="PPALLP"
	ss= "PPALLL"
	flag:=checkRecord(ss)
	fmt.Println(flag)

}
func checkRecord(s string) bool {
	var dictmap = make(map[string]int)
	for i:=0;i<len(s);i++{
		dictmap[string(s[i])]++
	}
	if dictmap["A"] >2 || strings.Contains(s,"LLL"){
		return false
	}
	return true
}
