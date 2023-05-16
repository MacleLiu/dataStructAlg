package mysort

import (
	"reflect"
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	data := []int{34, 37, 2, 58, 67, 34, 567, 67, 88, 54, 1, 4, 8, 34, 0, 67, 3, 6, 434, 999, 11111, 56, 32, 343, 54, 29}
	databak := []int{0, 1, 2, 3, 4, 6, 8, 29, 32, 34, 34, 34, 37, 54, 54, 56, 58, 67, 67, 67, 88, 343, 434, 567, 999, 11111}
	data2 := []int{33, 6, 8, 23, 4, 89, 5, 555, 21, 66}
	databak2 := []int{4, 5, 6, 8, 21, 23, 33, 66, 89, 555}
	specialData := []int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7}
	specialDatabak := []int{7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7, 7}
	sort.Slice(databak, func(i, j int) bool {
		return i < j
	})
	QuickSort(data)
	if !reflect.DeepEqual(data, databak) {
		t.Errorf("sort result is %v, want %v", data, databak)
	}
	QuickSort(data2)
	if !reflect.DeepEqual(data2, databak2) {
		t.Errorf("sort result is %v, want %v", data2, databak2)
	}
	QuickSort(specialData)
	if !reflect.DeepEqual(specialData, specialDatabak) {
		t.Errorf("sort result is %v, want %v", specialData, specialDatabak)
	}
}
