package main

import "fmt"

type Element interface {
}

type Stack interface {
	// 添加若干元素到栈顶
	Push(e ...Element)

	// 弹出栈顶元素
	Pop() Element

	// 只返回不弹出栈顶元素
	Peek() Element

	// 栈是否为空
	Empty() bool

	// 清空栈元素
	Clear()

	// 返回栈的大小
	Size() int
}

type ArrayStack []Element

// 向栈顶添加元素
func (s *ArrayStack) Push(e ...Element) {
	*s = append(*s, e...)
}

// 弹出栈顶元素
func (s *ArrayStack) Pop() Element {
	if s.Empty() {
		return nil
	}
	ln := len(*s)
	// 获取栈顶元素
	val := (*s)[ln-1]

	// 弹出栈顶元素
	*s = (*s)[:ln-1]
	return val
}

// 查看栈顶元素
func (s *ArrayStack) Peek() Element {
	if s.Empty() {
		return nil
	}
	return (*s)[len(*s)-1]
}

// 栈是否为空
func (s *ArrayStack) Empty() bool {
	return s.Size() == 0
}

// 清空栈
func (s *ArrayStack) Clear() {
	*s = ArrayStack{}
}

// 返回栈大小
func (s *ArrayStack) Size() int {
	return len(*s)
}

// 基于Slice实现构造栈
func NewArrayStack() Stack {
	return &ArrayStack{}
}



// 链表实现栈节点定义
type node struct {
	Value Element
	Next  *node
}

// 链表实现栈
type LinkedStack struct {
	Head *node
	size int
}

// 向栈顶添加元素
func (s *LinkedStack) Push(e ...Element) {
	for _, v := range e {
		n := s.Head.Next
		s.Head.Next = &node{
			Value: v,
			Next:  n,
		}
		s.size++
	}
}

// 弹出栈顶元素
func (s *LinkedStack) Pop() Element {
	if s.Empty() {
		return nil
	}
	first := s.Head.Next
	s.Head.Next = first.Next
	s.size--
	return first.Value
}

// 查看栈顶元素
func (s *LinkedStack) Peek() Element {
	if s.Empty() {
		return nil
	}
	return s.Head.Next.Value
}

// 栈是否为空
func (s *LinkedStack) Empty() bool {
	return s.Head.Next == nil
}

// 清空栈
func (s *LinkedStack) Clear() {
	s.Head.Next = nil
}

// 返回栈大小
// 由于使用了一个计数器，直接返回即可，时间复杂度O(1)
func (s *LinkedStack) Size() int {
	return s.size
}

// 链表实现栈构造
func NewLinkedStack() Stack {
	return &LinkedStack{
		Head: &node{},
	}
}

func main() {
	//s := NewStack()
	s := NewLinkedStack()
	fmt.Println("Empty", s.Empty())

	s.Push(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("Push 1,2,3,4,5")
	fmt.Println("Size:", s.Size())

	fmt.Println("Peek:", s.Peek())

	fmt.Println("Pop:", s.Pop(), s.Pop(), s.Pop())
	s.Clear()
	fmt.Println("Clear Empty:", s.Empty())

}


//https://www.jianshu.com/p/e2978dc93ef3
//https://www.cnblogs.com/TimLiuDream/p/9902496.html