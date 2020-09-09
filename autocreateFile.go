package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
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

// 检测文件或者文件夹是否存在
func Exists(path string) bool{
	_,err := os.Stat(path) // 获取文件信息
	if err != nil{
		if  os.IsExist(err){ // 查询这个文件错误是不是系统提示的文件不存在错误
			return true
		}
		return false
	}
	return true
}

// 判断给定路径是否为文件夹
func IsDir(path string) bool{
	s,err := os.Stat(path)
	if err != nil{
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}


func WirteThingsToFile(path string,gofilename string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil{
		panic(nil)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("package main\n")
	writer.WriteString("func main(){}\n")
	var funcstring string
	if strings.Contains(gofilename,".go"){
		funcstring = "func "+strings.ReplaceAll(gofilename,".go","")+"(){\n\n"
	} else {
		funcstring = "func "+gofilename+"(){\n\n"
	}
	writer.WriteString(funcstring)
	writer.WriteString("\n\n")
	writer.WriteString("}")
	writer.Flush()
}


func leetcodeStringFileTemplateCreate(filename string,gofilename string){
	filePath,_:=os.Getwd()
	newFilePath := filePath + "/" +filename
	goFilePath := newFilePath + "/" + gofilename
	if !Exists(newFilePath) &&  !Exists(goFilePath){
		if strings.Contains(filename,"/"){
			os.MkdirAll(newFilePath,os.ModePerm)
			os.Create(newFilePath+"/"+gofilename)
			WirteThingsToFile(newFilePath+"/"+gofilename,gofilename)
		} else {
			os.Create(newFilePath +"/"+gofilename)
		}
	} else {
		fmt.Printf("%s 已经存在,请重新输入\n%s 已经存在,请重新输入\n",newFilePath,goFilePath)
	}
}

var (
	h bool
	q *bool
	f string
	t string
	b string
)

func init() {
	q = flag.Bool("q", false, "Exit")
	flag.BoolVar(&h, "h", false, "Show help")
	flag.StringVar(&f,"f", "s", "direct file")
	flag.StringVar(&t,"t","zz","go file template")
	flag.StringVar(&b,"b","xx","set basefile")
}


func main()  {
	flag.Parse()
	if h {
		flag.Usage()
	} else {
		if *q {
			fmt.Println("is use -q the *q result is true, example like this ", *q)
			os.Exit(0)
		}
		if  t == "zz"{
			list := strings.Split(f,"/")
			 t = list[len(list)-1:][0]+".go"
		}
		leetcodeStringFileTemplateCreate(f,t)
	}
}