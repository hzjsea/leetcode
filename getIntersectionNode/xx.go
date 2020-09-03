package main

type ListNode struct {
	Val int
	Next *ListNode
}


// 内部消化，狸猫换太子
func deleteNode(node *ListNode) {
	node.Val,node.Next = node.Next.Val,node.Next.Next
}

