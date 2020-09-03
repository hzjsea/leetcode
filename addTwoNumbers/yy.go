package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode,l2 *ListNode) *ListNode{
	//  创建新的链表用来存储结果
	var newHead *ListNode = new(ListNode)
	// 遍历两个链表，每进一位就进行相加得到新的链表，并且需要用到tmp来保留进位内容
	sum:=0
	for l1 != nil || l2 !=nil  {
		if l1 !=nil{
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil{
			sum += l2.Val
			l2 = l2.Next
		}
		newHead.Next = &ListNode{Val: sum%10}
		sum = sum / 10
		newHead = newHead.Next

	}
	return newHead
}