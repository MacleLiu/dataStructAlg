package linkList

import (
	"testing"
)

func TestCreatLiatF(t *testing.T) {
	data := []int{3, 4, 1, 6, 7, 5, 4, 1}
	head := CreatListF(data)
	p := head.Next
	i := len(data)
	for {
		if p != nil && i > 0 {
			if p.Data == data[i-1] {
				p = p.Next
				i--
			} else {
				t.Errorf("p.data = %d, want %d", p.Data, data[i-1])
			}
		} else {
			return
		}
	}
}
func TestCreatLiatR(t *testing.T) {
	data := []int{3, 4, 1, 6, 7, 5, 4, 1}
	head := CreatListR(data)
	p := head.Next
	i := 0
	for {
		if p != nil && i < len(data) {
			if p.Data == data[i] {
				p = p.Next
				i++
			} else {
				t.Errorf("p.data = %d, want %d", p.Data, data[i-1])
			}
		} else {
			return
		}
	}
}

func TestLocate(t *testing.T) {
	data := []int{3, 4, 1, 6, 7, 5, 4, 1}
	head := CreatListR(data)
	p := Locate(*head, 4)
	if p.Data != 4 {
		t.Errorf("p.data = %d, want %d", p.Data, 4)
	}
}

func TestGet(t *testing.T) {
	data := []int{3, 4, 1, 6, 7, 5, 4, 1}
	head := CreatListR(data)
	p := Get(*head, 7)
	if p.Data != data[7-1] {
		t.Errorf("p.data = %d, want %d", p.Data, data[7-1])
	}
}

func TestInsert(t *testing.T) {
	data := []int{3, 4, 1, 6, 7, 5, 4, 1}
	head := CreatListR(data)
	p := Locate(*head, 4)
	if !Insert(p, 9) {
		t.Error("Insert fail!")
	}
	if p.Next.Data != 9 {
		t.Errorf("p.next.data = %d, want %d", p.Next.Data, 9)
	}
	if p.Next.Next.Data != 1 {
		t.Errorf("p.next.next.data = %d, want %d", p.Next.Next.Data, 1)
	}
}

func TestFindPrevious(t *testing.T) {
	data := []int{3, 4, 1, 6, 7, 5, 4, 1}
	head := CreatListR(data)
	p := Locate(*head, 7)
	pre := FindPrevious(*head, p)
	if pre.Data != 6 {
		t.Errorf("pre.data = %d, want %d", pre.Data, 6)
	}
}

func TestDelete(t *testing.T) {
	data := []int{3, 4, 1, 6, 7, 5, 4, 1}
	head := CreatListR(data)
	if !Delete(*head, Locate(*head, 4)) {
		t.Error("Delete fail!")
	}
	result := []int{3, 1, 6, 7, 5, 4, 1}
	p := head.Next
	i := 0
	for {
		if p != nil && i < len(result) {
			if p.Data == result[i] {
				p = p.Next
				i++
			} else {
				t.Errorf("p.data = %d, want %d", p.Data, result[i-1])
			}
		} else {
			return
		}
	}
}
