package main

import (
	"bufio"
	"bytes"
	"os"
)

func main() {
	// 读取文件
	ReaderFromInterface()
	WriterToInterface()
}


//type ReaderFrom interface {
//	ReadFrom(r Reader) (n int64, err error)
//}


func ReaderFromInterface(){
	// 打开一个文件夹
	file,err := os.Open("./xx.txt")
	if err != nil{
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(os.Stdout)
	//writer.ReadFrom(file)
	//writer.Flush()
	writer.ReadFrom(file)
	writer.Flush()
}

func WriterToInterface() {
	// 创建一个标准输入
	reader := bytes.NewReader([]byte("Go语言中文网\n"))
	reader.WriteTo(os.Stdout)
}