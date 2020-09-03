package  main

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



