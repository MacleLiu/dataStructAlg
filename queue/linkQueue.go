package queue

//链式队列，不含头节点

//数据节点
type node struct {
	data int
	next *node
}

//头尾指针
type LinkQueue struct {
	front, rear *node
}

//新建队列
func NewQueue() *LinkQueue {
	q := new(LinkQueue)
	return q
}

//判断队列是否为空
func (q *LinkQueue) EmpytQ() bool {
	return q.front == nil && q.rear == nil
}

//入队
func (q *LinkQueue) Enqueue(data int) {
	n := new(node)
	n.data = data
	n.next = nil
	if q.EmpytQ() {
		q.rear = n
		q.front = n
		return
	}
	q.rear.next = n
	q.rear = n
}

// 出队
func (q *LinkQueue) Dequeue() int {
	//队列为空
	if q.EmpytQ() {
		return 0
	}
	n := q.front
	//队列长度为一
	if q.front.next == nil {
		q.front = nil
		q.rear = nil
		return n.data
	}
	q.front = q.front.next
	return n.data
}
