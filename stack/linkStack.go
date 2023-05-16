package linkStack

//栈的链式实现，带有头节点的
type Node struct {
	Data int
	Next *Node
}

//新建一个栈
func NewStack() *Node {
	return &Node{}
}

//判断栈是否为空
func IsEmpty(s *Node) bool {
	return s.Next == nil
}

//进栈操作，链表的前端作为栈顶
func Push(s *Node, data int) {
	n := &Node{}
	n.Data = data
	n.Next = s.Next
	s.Next = n
}

//出栈操作
func Pop(s *Node) int {
	if IsEmpty(s) {
		return 0
	}
	n := s.Next.Data
	s.Next = s.Next.Next
	return n
}

//获取栈顶元素，但不出栈
func Top(s *Node) int {
	if IsEmpty(s) {
		return 0
	}
	return s.Next.Data
}
