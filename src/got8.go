package main

//链表实现
import (
	"fmt"
)


//定义元素类型
type Element interface {}

//定义节点
type linkNode struct {
	Data Element   //数据域
	Nest *linkNode //指针域，指向下一个节点
}

func NewLinkNode() *linkNode{
	return &linkNode{Data:"head"}
}

//函数接口
type LinkNode interface {
	Add(data Element)              //后面添加
	Delete(index int) Element              //删除指定index位置元素
	Insert(index int, data Element) //在指定index位置插入元素
	GetLength() int                   //获取元素长度
	Search(data Element)   int         //查询元素的位置
	GetData(index int) Element      //获取指定index位置的元素
	GetAll()  []Element    //获取所有的元素
}

//添加 头结点，数据
func (head *linkNode)Add(data Element) {
	point := head //临时指针，为最后一个节点
	for point.Nest != nil {
		point = point.Nest //移位
	}

	newNode := linkNode{Data:data} //需要添加的新节点
	point.Nest = &newNode

}

//删除 头结点 index 位置
func (head *linkNode)Delete(index int) Element {
	//判断index合法性
	if index < 0 || index > head.GetLength() {
		fmt.Println("please check index")
		return fmt.Errorf("check index error")
	} else {
		point := head
		for i := 0; i < index; i++ {
			point = point.Nest //移位
		}
		point.Nest = point.Nest.Nest //赋值
		data := point.Nest.Data
		return data
	}
}

//插入 头结点 index位置 data元素
func (head *linkNode)Insert( index int, data Element) {
	//检验index合法性
	if index < 0 || index > head.GetLength() {
		fmt.Println("please check index")
	} else {
		point := head
		for i := 0; i < index-1; i++ {
			point = point.Nest //移位
		}
		newNode := &linkNode{Data:data} //初始化新加的元素
		newNode.Nest = point.Nest  //指定新加元素的下一个元素
		point.Nest = newNode
	}
}

//获取长度 头结点
func (head *linkNode)  GetLength()  int {
	point := head
	var length int
	for point.Nest != nil {
		length++
		point = point.Nest
	}
	length++
	return length
}

//搜索 头结点 data元素
func (head *linkNode)Search(data Element) int{
	point := head
	index := 0
	for point.Nest != nil {
		if point.Data == data {
			fmt.Println(data, "exist at", index, "th")
			return index
		} else {
			index++
			point = point.Nest
			if index > head.GetLength()-1 {
				fmt.Println(data, "not exist at")
				break
			}
			continue
		}
	}
	if point.Data == data {   //判断最后一个元素是否匹配
		fmt.Println(data, "exist at", index, "th")
		return index
	}

	return -1
}

//获取data 头结点 index位置
func (head *linkNode)GetData( index int) Element {
	point := head
	if index < 0 || index > head.GetLength() {
		return 	fmt.Errorf("check index error")
	} else {
		for i := 0; i < index-1; i++ {
			point = point.Nest
		}
		return point.Data
	}
}

//获取所有的data
func  (head *linkNode)GetAll() []Element{
	dataList := make([]Element,0,head.GetLength())
	point := head
	for point.Nest != nil {
		dataList = append(dataList,point.Data)
		point = point.Nest
	}
	dataList = append(dataList,point.Data)
	return dataList
}


//主函数测试
func main() {
	linkNode := NewLinkNode()
	for i:=0 ; i <50 ;i++ {
		linkNode.Add(2*i+1)
	}
	fmt.Println(linkNode.GetLength(),linkNode.GetAll())
	linkNode.Add("new add data")
	fmt.Println(linkNode.GetLength(),linkNode.GetAll())
	linkNode.Insert(40,"new insert data")
	fmt.Println(linkNode.GetLength(),linkNode.GetAll())


	d := linkNode.Delete(140)
	if _,ok := d.(error);ok {
		fmt.Println("delete fail",d)
	}else {
		fmt.Println("delete success")
	}

	fmt.Println(linkNode.GetLength(),linkNode.GetAll())

	i := linkNode.Search("new add data")
	fmt.Println(i)
}
