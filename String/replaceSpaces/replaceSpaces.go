package main

import (
	"fmt"
	"strings"
)

func main() {
	ss:="               "
	ss = replaceSpaces(ss,len(ss))
	fmt.Println(ss)



}


func replaceSpaces(S string, length int) string {
	return strings.Replace(S[:length]," ","%20",-1)
}
