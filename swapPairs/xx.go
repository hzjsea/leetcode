package main

type ListNode struct {
	Val  int
	Next *ListNode
}


func swapPairs(head *ListNode) *ListNode {
	// 判空
	if head == nil || head.Next == nil{
		return head
	}
	curr, pre := head, head.Next
	// 当当前节点被交换到前置节点的位置时，当前节点的next应该是前置节点的Next
	curr.Next = swapPairs(pre.Next)
	// 前后对掉后，前置节点的next应该指向当前节点
	pre.Next = curr
	return pre

}

//// 先写一个常用方法 就是前后调换节点位置
//func reverseNode(nodelist *ListNode) *ListNode{
//	nodelist.Val = nodelist.Next.Val
//	nodelist.Next = nodelist.Next.Next
//}