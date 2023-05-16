// linkList包是关于链表相关实现的学习示例代码
package linkList

//本例是带有头节点的单链表实现
type Node struct {
	Data int
	Next *Node
}

func newNode() *Node {
	return new(Node)
}

//新建链表，返回一个空的头节点
func NewLinkList() *Node {
	return newNode()
}

//判断链表是否为空
func IsEmpty(head Node) bool {
	return head.Next == nil
}

//判断当前位置是否是链表的尾部
func IsLast(position *Node) bool {
	return position.Next == nil
}

//按值查找，查找指定数据在链表中的节点位置，返回节点地址
func Locate(head Node, data int) *Node {
	position := head.Next
	for {
		if position == nil {
			return nil
		} else if position.Data == data {
			return position
		}
		position = position.Next
	}
}

//按序号查找，在带头节点的单链表中查找第i个节点，返回该节点的地址
func Get(head Node, index int) (position *Node) {
	var j int
	position = head.Next
	for {
		j++
		if position != nil && j < index {
			position = position.Next
		} else if index == j {
			return position
		} else {
			return nil
		}
	}
}

//插入节点，在指定位置之后插入值为data的新节点
func Insert(position *Node, data int) bool {
	if position == nil {
		return false
	}
	n := newNode()
	if n == nil {
		return false
	}
	n.Data = data
	n.Next = position.Next
	position.Next = n
	return true
}

//查找前驱节点，查找指定位置的前驱节点
func FindPrevious(head Node, position *Node) *Node {
	if position == nil {
		return nil
	}
	q := head.Next
	for {
		if q.Next != position {
			q = q.Next
		} else {
			return q
		}
	}
}

//删除节点
func Delete(head Node, position *Node) bool {
	if position == nil {
		return false
	}
	pre := FindPrevious(head, position)
	if pre == nil {
		return false
	}
	pre.Next = position.Next
	return true
}

//头插法建表，接收一个int型的slice，以头插法的方式建立链表
func CreatListF(data []int) *Node {
	head := newNode()
	for _, v := range data {
		n := newNode()
		n.Data = v
		n.Next = head.Next
		head.Next = n
	}
	return head
}

//尾插法建表，接收一个int型的slice，以尾插法的方式建立链表
func CreatListR(data []int) *Node {
	head := newNode()
	r := head //尾指针
	for _, v := range data {
		n := newNode()
		n.Data = v
		r.Next = n
		r = n
	}
	r.Next = nil
	return head
}
