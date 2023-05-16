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
	Push(stack, 7)
	Push(stack, 4)
	if stack.Next.Data != 4 {
		t.Errorf("stack.Next.Data = %d, want %d", stack.Next.Data, 4)
	}
}

func TestPop(t *testing.T) {
	stack := NewStack()
	Push(stack, 7)
	Push(stack, 4)
	if Pop(stack) != 4 {
		t.Errorf("Pop(stack) = %d, want %d", Pop(stack), 4)
	}
	if Pop(stack) != 7 {
		t.Errorf("Pop(stack) = %d, want %d", Pop(stack), 7)
	}
}

func TestTop(t *testing.T) {
	stack := NewStack()
	Push(stack, 7)
	Push(stack, 4)
	Push(stack, 10)
	Pop(stack)
	Push(stack, 17)
	Push(stack, 66)
	if stack.Next.Data != 66 {
		t.Errorf("stack.Next.Data = %d, want %d", stack.Next.Data, 66)
	}
}
