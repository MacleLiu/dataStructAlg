package linkStack

import "testing"

func TestNewStack(t *testing.T) {
	stack := NewStack()
	if stack == nil {
		t.Error("stack is nil!")
	}
}

func TestPush(t *testing.T) {
	stack := NewStack()
	stack.Push(7)
	stack.Push(4)
	if stack.Next.Data != 4 {
		t.Errorf("stack.Next.Data = %d, want %d", stack.Next.Data, 4)
	}
}

func TestPop(t *testing.T) {
	stack := NewStack()
	stack.Push(7)
	stack.Push(4)
	if stack.Pop() != 4 {
		t.Errorf("Pop(stack) = %d, want %d", stack.Pop(), 4)
	}
	if stack.Pop() != 7 {
		t.Errorf("Pop(stack) = %d, want %d", stack.Pop(), 7)
	}
}

func TestTop(t *testing.T) {
	stack := NewStack()
	stack.Push(7)
	stack.Push(4)
	stack.Push(10)
	stack.Pop()
	stack.Push(17)
	stack.Push(66)
	if stack.Next.Data != 66 {
		t.Errorf("stack.Next.Data = %d, want %d", stack.Next.Data, 66)
	}
}
