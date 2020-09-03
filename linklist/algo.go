package main

import (
	"fmt"
)

//Element
type Element interface {}

// 创建Node
type Node struct {
	Val Element
	Next *Node  // 指向Node类型的指针，  这里你用 &Node 就错了
}


// 创建头结点
func New(data Element) *Node{
	return &Node{Val:"head",Next: nil}
}



//创建接口,规定一组行为
type nodeAction interface {
	Add(data Element)
	InsertTail(data Element)              //后面添加
	InsertHead(data Element)  			//前面添加
	Insert(index int, data Element) //在指定index位置插入元素
	Delete(index int) Element              //删除指定index位置元素
	Length() int                   //获取元素长度
	Search(data Element)   int         //查询元素的位置
	GetData(index int) Element      //获取指定index位置的元素
	Traverse()  []Element    //获取所有的元素
	Traverse2()      //获取所有的元素
	IsEmpty() bool 	// 判断是否为空

}

func (head *Node) IsEmpty () bool {
	if head.Next == nil {
		return true
	}

	return false
}


func (head *Node) Length () int {
	current := head
	count := 1
	for current.Next != nil{
		current = current.Next
		count ++
	}
	return count
}


func (head *Node) InsertTail(data Element) {
	current := head
	for current.Next != nil{
		current = current.Next
	}
	newNode := &Node{Val: data}
	current.Next = newNode
}


func (head *Node) InsertHead(data Element){
	// 虚拟头结点
	//dummy_head := &Node{Val:data}
	//dummy_head.Next  = head
	//
	//// 创建新的节点地址
	//newNode := &Node{Val:data}
	//newNode.Next = head
	//dummy_head.Next = newNode

	newNode := &Node{Val:data}
	newNode.Next = head.Next
	head.Next = newNode
}


func (head *Node) Insert(index int,data Element) {
	//for index > 1{
	//	index --
	//	current = current.Next
	//}
	current :=head
	if index < 0{
		head.InsertHead(data)
	}else if index > head.Length(){
		head.InsertTail(data)
	}else {
		for count:=0;index > count;index--{
			current = current.Next
		}
	}
	newNode := &Node{data,nil}
	newNode.Next = current.Next
	current.Next = newNode
}

func (head *Node) Delete(index int) Element {
	current :=head
	if index < 0  {
		return 0
	}else if index > head.Length(){
		return 0
	}else {
		for count:=0;(index-1) > count;index--{
			current = current.Next
		}
	}
	val := current.Next.Val
	current.Next = current.Next.Next
	return val
}

func(head *Node ) Search(data Element) int {
	current := head
	count := 0
	for current.Next != nil{
		if current.Val != data {
			count ++
		}else {
			return count
		}
		current = current.Next
	}
	return 0

}

func (head *Node) Traverse2() {

	current := head
	for current.Next != nil{
		current = current.Next
		fmt.Println(current)
	}
}

func (head *Node) Traverse()  []Element  {
	// 创建切片
	//SliceN:=make([]Element,0,head.Length()) // 创建数组
	slice := make([]Element,0,head.Length())
	curr:= head
	for curr.Next !=nil{
		slice = append(slice,curr.Val)
		curr = curr.Next
	}
	slice = append(slice,curr.Val)
	return slice

}

//反转链表
func (head *Node) reverseList() []Element{
	// 反转链表就是创建一个新的头结点，然后遍历链表A 将链表A中的所有结点脱离 并且连接到新的头结点上
	var curr,newHead *Node = head,nil
	// 使用tmp结点来临时保存curr.Next的结点
	for curr != nil{
		tmp:=curr.Next
		curr.Next = newHead
		newHead = curr
		curr = tmp
	}
	return newHead.Traverse()

}

// 递归法
// 递归法在链表题目中经常使用，因为链表中的每一个节点都可以作为链表头来使用,所以递归法的参数就很有用了

// 比如有一道算法题 交换链表中的两两节点互相反转
// 1. 这里两两反转是一道子程序
// 2. 多组两两反转是一个循环
// 3. 循环的终止节点是 节点为空，或者节点的下一位为空







func main() {
	linkNode := New(1)
	fmt.Println(*linkNode)

	// 判断是否为空
	flag:=linkNode.IsEmpty()
	fmt.Println(flag)

	// 计算长度
	len := linkNode.Length()
	fmt.Println(len)

	// 头部插入数据
	fmt.Println("---")

	for i:=0 ;i<3;i++{
		linkNode.InsertTail(i*2+1)
	}
	linkNode.InsertHead(20)
	linkNode.Traverse2()
	fmt.Println("---")

	linkNode.Insert(2,10)
	linkNode.Traverse2()
	fmt.Println("---")

	linkNode.Delete(2)
	linkNode.Traverse2()

	index := linkNode.Search(10)
	fmt.Println(index)

	slice := linkNode.Traverse()
	fmt.Println(slice)

	reverseSlice:= linkNode.reverseList()
	fmt.Println(reverseSlice)

}