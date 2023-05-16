package mysort

import (
	"reflect"
	"sort"
	"testing"
)

func TestHeapSort(t *testing.T) {
	data := []int{34, 37, 2, 58, 67, 34, 567, 67, 88, 54, 1, 4, 8, 34, 0, 67, 3, 6, 434}
	databak := []int{0, 1, 2, 3, 4, 6, 8, 34, 34, 34, 37, 54, 58, 67, 67, 67, 88, 434, 567}
	data2 := []int{33, 6, 8, 23, 4, 89, 5, 555, 21, 66}
	databak2 := []int{4, 5, 6, 8, 21, 23, 33, 66, 89, 555}
	sort.Slice(databak, func(i, j int) bool {
		return i < j
	})
	HeapSort(data)
	if !reflect.DeepEqual(data, databak) {
		t.Errorf("sort result is %v, want %v", data, databak)
	}
	HeapSort(data2)
	if !reflect.DeepEqual(data2, databak2) {
		t.Errorf("sort result is %v, want %v", data2, databak2)
	}
}
