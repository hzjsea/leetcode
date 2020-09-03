package main

import (
	"flag"
	"fmt"
	"os"
	"reflect"
)

func fileAction() {
	// 文件夹名 文件权限 0777
	os.Mkdir("abc",os.ModePerm)
	// 创建多级目录
	os.Mkdir("dir1/dir2/dir3",os.ModePerm)


	// 上面两个动作错误的时候都会返回一个error值
	//err:=os.Mkdir（“dir1/dir2/dir3”，os.ModePerm)
	//if err!=nil{
	//	fmt.Println(err)
	//}
	//result:mkdir ./dir1/dir2/dir3: The system cannot find the path specified.

	// 获取当前目录 以及 error
	os.Getwd()

	// 创建文件夹  返回一个 *file 和一个error
	f1,_ := os.Create("./name.txt")
	fmt.Println(f1)
	defer f1.Close()

	//以读写方式打开文件，如果不存在则创建文件，等同于上面os.Create
	f4, _ := os.OpenFile("./4.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	defer f4.Close()

	// 删除文件夹
	os.Remove("abc/d/e/f")

	// 删除指定目录
	os.RemoveAll("abc")

	// 重命名
	os.Rename("./2.txt", "./2_new.txt")
}





func leetcodeStringFileTemplateCreate(filename string,gofilename string){
	if gofilename == "None"{
		gofilename = "xx.go"
	}
	filePath,_:=os.Getwd()
	filePath  =  filePath + "/" + "String"
	newFilePath := filePath +"/"+filename
	if os.Mkdir(newFilePath,os.ModePerm) == nil{
		os.Create(newFilePath +"/"+gofilename)
	} else {
		os.RemoveAll(newFilePath)
		os.Mkdir(newFilePath,os.ModePerm)
		os.Create(newFilePath+"/"+gofilename)
	}
}

var (
	n int
	h bool
	q *bool
	f string
	t string
	c string
)

func init() {
	q = flag.Bool("q", false, "Exit")
	flag.BoolVar(&h, "h", false, "Show help")
	flag.IntVar(&n, "n", 0, "Set number")
	flag.StringVar(&f,"f", "s", "Default string")
	flag.StringVar(&t,"t","xx.go","default string")
	flag.StringVar(&c,"c","xx","set number")

}


func main()  {
	flag.Parse()
	if h {
		flag.Usage()
	} else {
		if *q {
			fmt.Println("q is ", *q)
			os.Exit(0)
		}
		//fmt.Println("Number is ", n)
		//fmt.Println("filename is ", f)
		//fmt.Println("gofilename is ",t)
		//fmt.Println("cfilename is ", c)
		leetcodeStringFileTemplateCreate(f,t)
	}
}