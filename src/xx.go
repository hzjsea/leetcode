package main

import (
	"fmt"
	"reflect"
)

func main() {
	s:="hello wolrd"
	for _,v:= range ([]byte(s)){
		fmt.Println(reflect.TypeOf(v),v)
	}
}


