package queue

import "testing"

func TestNewQueue(t *testing.T) {
	q := NewQueue()
	if q == nil {
		t.Error("New queue error")
	}
}

func TestEmptyQ(t *testing.T) {
	q := NewQueue()
	if !q.EmpytQ() {
		t.Error("This queue is not empty")
	}
}

func TestEnQueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(5)
	if q.EmpytQ() {
		t.Error("EnQueue fail. The queue is still empty!")
	}
	if q.rear.data != 5 {
		t.Errorf("q.rear.data = %d, want %d", q.rear.data, 5)
	}
	q.Enqueue(8)
	if q.rear.data != 8 {
		t.Errorf("q.rear.data = %d, want %d", q.rear.data, 8)
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(5)
	q.Enqueue(8)
	if d := q.Dequeue(); d != 5 {
		t.Errorf("q.Dequeue() = %d, want %d", d, 5)
	}
	if d := q.Dequeue(); d != 8 {
		t.Errorf("q.Dequeue() = %d, want %d", d, 8)
	}
	if !q.EmpytQ() {
		t.Error("All Dequeue, but is not empty!")
	}
}
