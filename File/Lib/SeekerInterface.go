package main

import (
	"fmt"
	"io"
	"strings"
)

//type Seeker interface {
//	Seek(offset int64, whence int) (ret int64, err error)
//}

func main() {
	reader := strings.NewReader("Go语言中文网")
	reader.Seek(-6, io.SeekEnd)
	r, _, _ := reader.ReadRune()
	fmt.Printf("%c\n", r)
}
