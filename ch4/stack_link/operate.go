// Author: xufei
// Date: 2019-07-26

package stack_link

func InitStack() *LinkStack {
	link := new(LinkStack)
	return link
}

func Push(link *LinkStack, e int) bool {
	node := new(StackNode)
	node.data = e

	node.next = link.top
	link.top = node
	link.count++

	return true
}

func (link *LinkStack) Push(e int) bool {
	node := new(StackNode)
	node.data = e

	node.next = link.top
	link.top = node

	return true
}

func (link *LinkStack) Pop(e *int) bool {
	*e = link.top.data
	//p := link.top
	link.top = link.top.next
	link.count--
	// free(p) // go 语言是自动垃圾回收（GC），不需要像 c 那样需要自己去释放内存

	return true
}
