package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 旋转链表
// 1. 定位要旋转的内容 用快慢指针
//func rotateRight(head *ListNode, k int) *ListNode {
//	fast:=head
//	slow:=head
//	for k>0{
//		fast = fast.Next
//		k--
//	}
//	for fast.Next!=nil {
//		slow = slow.Next
//		fast = fast.Next
//	}
//	//// 创建一个新的链表来接纳
//	slow.Next = nil
//	// 这个时候newlist就是前缀 slow 就是尾巴  newlist = [1,2,3]  slow = [4,5]
//	// 之前一直想着把[4,5]反转拼接到netlist前面去
//	// 弄了好久 ，但其实可以吧newlist放到slowh后面去啊！！
//	curr:=slow
//	for curr.Next !=nil{
//		curr = curr.Next
//	}
//	fmt.Println(curr)
//	curr.Next = head
//	return slow
//}



func rotateRight(head *ListNode, k int) *ListNode {
	//if head == nil || head.Next == nil || k == 0 {
	//	return head
	//}
	//// 先变成首尾相连的环状链表
	//len, cycle := 1, head
	//for cycle.Next != nil {
	//	len++
	//	cycle = cycle.Next
	//}
	//cycle.Next = head
	//
	//// 找到倒数第k+1个结点
	//p := head
	//for i := 0; i < len-k%len-1; i++ {
	//	p = p.Next
	//}
	//
	//// 分割
	//ans := p.Next
	//p.Next = nil
	//return ans

	// 判空u
	if head == nil || head.Next == nil || k == 0{
		return head
	}
	// 创建环链
	curr:=head
	len := 1
	for curr.Next != nil{
		len++
		curr = curr.Next
	}
	curr.Next = head

	// 找到段点
	p := head
	for i:=0;i<len-k%len-1;i++{
		p = p.Next
	}
	temp := p.Next
	p.Next = nil
	return  temp
}
