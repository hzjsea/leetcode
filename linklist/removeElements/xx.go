package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

// 创建头结点
func New(data int) *ListNode{
    return &ListNode{Val:data,Next: nil}
}

type node interface {
    InsertHead(data int)
    Traverse2()
    removeElements(head *ListNode, val int) *ListNode
}

func (head *ListNode) InsertHead(data int){
    // 虚拟头结点
    //dummy_head := &Node{Val:data}
    //dummy_head.Next  = head
    //
    //// 创建新的节点地址
    //newNode := &Node{Val:data}
    //newNode.Next = head
    //dummy_head.Next = newNode

    newNode := &ListNode{Val:data}
    newNode.Next = head.Next
    head.Next = newNode
}

func (head *ListNode) Traverse2() {

    current := head
    for current.Next != nil{
        current = current.Next
        fmt.Println(current)
    }
}

func main()  {
    linkNode := New(1)
    linkNode.InsertHead(6)
    linkNode.InsertHead(5)
    linkNode.InsertHead(4)
    linkNode.InsertHead(3)
    linkNode.InsertHead(6)
    linkNode.InsertHead(2)
    linkNode.InsertHead(1)
    linkNode.Traverse2()
    head := linkNode.removeElements(3)
    fmt.Println("----")
    head.Traverse2()



}

// -------

func (head *ListNode) removeElements( val int) *ListNode {
    current:=head
    for current.Next != nil{
        if current.Next.Val == val{
            current.Next = current.Next.Next
        }else {
            current = current.Next
        }
    }
    return head
}



