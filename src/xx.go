package main

import "fmt"

func main() {
	Slice := make([]int,0,10)
	fmt.Println(Slice)
	fmt.Println(len(Slice))
	fmt.Println(cap(Slice))
	Slice = append(Slice, 2)
	fmt.Println(Slice)


	listN := []int{1,2,34}
	listN = append(listN, 2)
	fmt.Println(listN)

	name := []int{1,2,3,45}
	name = append(name,1)
	fmt.Println(name)

	fmt.Println(*&name)
}

