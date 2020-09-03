package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 反转链表
// 利用临时项一组一组反转
func reverseList(head *ListNode) *ListNode {
	var curr,pre *ListNode = head,nil
	// 去除空链表
	if head  == nil{
		return head
	}

	// 开始反转
	for curr != nil{
		tmp:=curr.Next
		curr.Next = pre
		pre = curr
		curr = tmp
	}
	return pre
}


