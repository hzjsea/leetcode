package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 判空
	if head == nil{
		return head
	}

	// 同起点
	fast:=head
	last:=head
	for n>0{
		fast = fast.Next
		n--
	}

	for fast != nil{
		last = last.Next
		fast = fast.Next
	}

	last.Val = last.Next.Val
	last.Next = last.Next.Next
	return head

}
