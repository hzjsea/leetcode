package main

import (
	"fmt"
)

// 定义元素类型
type Element interface {}

// 定义节点
type Node struct {
	Val Element
	Next *Node
}


func New() *Node{
	return &Node{0,nil}
}

// 函数接口
type node interface {
	Add(data Element)              //后面添加
	Delete(index int) Element              //删除指定index位置元素
	Insert(index int, data Element) //在指定index位置插入元素
	GetLength() int                   //获取元素长度
	Search(data Element)   int         //查询元素的位置
	GetData(index int) Element      //获取指定index位置的元素
	GetAll()  []Element    //获取所有的元素
}




// 获取指定index位置的元素
func (head *Node) GetData(index int ) Element {
	current:=head
	// 判断Index是否合理
	if index <0 || index > current.GetLength(){
		fmt.Println("error")
		return 0
	}
	for i:=0;i<index-1;i++ {
		current = current.Next
	}
	return current.Val
}

// 查询元素的位置
func (head *Node) Search(data Element) int{
	current:=head
	j :=0
	for current.Next != nil{
		j ++
		if current.Val == data{
			return j
		}
		current = current.Next
	}
	return 0

}

// 获取长度
func (head *Node) GetLength() int{
	current:=head
	var lenght int
	for current.Next!=nil{
		lenght++
		current = current.Next
	}
	lenght++
	return lenght
}
func (head *Node) Delete(index int ) Element{
	// 判断index的合法性
	if index < 0 || index > head.GetLength() {
		fmt.Println("error")
		return 0
	}

	current := head
	for i:=0;i<index;i++{
		current = current.Next
	}
	current.Next = current.Next.Next
	current.Val = current.Next.Val
	return 1

}

// 添加尾结点 数据
func (head *Node)  Add(data Element) {
	// 获取游标
	current :=head
	if current == nil{
		return
	}

	for current.Next != nil{
		current = current.Next
	}
	newNode := Node{data,nil}
	current.Next = &newNode
}


func(head *Node) Transerve(){
	curr := head.Next
	for curr != nil{
		fmt.Println(curr.Val)
		curr = curr.Next
	}
	fmt.Println("---done---")
}

func(head *Node) GetAll()[]Element{
	dataList := make([]Element,0,head.GetLength())
	point := head
	for point.Next != nil {
		dataList = append(dataList,point.Val)
		point = point.Next
	}
	dataList = append(dataList,point.Val)
	return dataList
}

func main() {
	linkNode := New()
	for i:=0 ; i <3 ;i++ {
		linkNode.Add(2*i+1)
	}
	fmt.Println(linkNode.GetLength(),linkNode.GetAll())
	linkNode.Add("new add data")
	fmt.Println(linkNode.GetLength(),linkNode.GetAll())
	//linkNode.Insert(40,"new insert data")
	//fmt.Println(linkNode.GetLength(),linkNode.GetAll())


	d := linkNode.Delete(1)
	if _,ok := d.(error);ok {
		fmt.Println("delete fail",d)
	}else {
		fmt.Println("delete success")
	}

	fmt.Println(linkNode.GetLength(),linkNode.GetAll())

	i := linkNode.Search("new add data")
	fmt.Println(i)
}