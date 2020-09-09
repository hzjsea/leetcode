package main

//type Closer interface {
//	Close() error
//}


// 用于关闭数据流， 文件 (os.File)、归档（压缩包）、数据库连接、Socket 等需要手动关闭的资源都实现了 Closer 接口。

// !警告
//当文件 studygolang.txt 不存在或找不到时，file.Close() 会 panic，因为 file 是 nil。因此，应该将 defer file.Close() 放在错误检查之后。
//file, err := os.Open("studygolang.txt")
//	defer file.Close()
//	if err != nil {
//		...
//	}
// 上面的这块是错误的，首先如果文件不存在 defer file.Close会直接报错，所以应该先判断文件是否存在，不存在则直接抛出异常
// Like this
//file,err := os.Open("xx.txt")
//if err != nil{
//	panic(err)
//}
//defer file.Close()

func (f *File) Close() error {
	if f == nil {
		return ErrInvalid
	}
	return f.file.close()
}

