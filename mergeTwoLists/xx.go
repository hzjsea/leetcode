package main

type ListNode struct {
	Val int
	Next *ListNode
}


// 给定两个升序链表，合并成一个升序链表
// 使用递归法解决，递归当前函数
// 比较出两者中小的node节点，让小的node节点指向大的node节点所在链表
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 判断链表是否为空

	if l1 == nil{
		return l2
	} else if l2 == nil{
		return l1
	}

	if (l1.Val < l2.Val) {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}