package main

type ListNode struct {
	Val int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	// 定义快慢指针，初始都指向head
	fast, slow := head, head
	// 快指针每次前进2步，慢指针每次前进1步，循环完慢指针极为mid
	// 利用慢指针来取中间点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 使用mid进行后半段的链表反转
	var pre, current *ListNode = nil, slow
	for current != nil {
		temp := current.Next
		current.Next = pre
		pre = current
		current = temp
	}

	for pre != nil {
		if pre.Val != head.Val {
			return false
		}
		pre = pre.Next
		head = head.Next
	}
	return true

}