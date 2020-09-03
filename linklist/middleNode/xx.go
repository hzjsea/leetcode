package main

type ListNode struct {
	Val int
	Next *ListNode
}

// 双指针法
func middleNode(head *ListNode) *ListNode {
	fast, slow := head, head
	for {
		if fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		} else if fast.Next != nil {
			return slow.Next
		} else {
			return slow
		}
	}
}