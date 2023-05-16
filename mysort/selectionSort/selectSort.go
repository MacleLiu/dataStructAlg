package mysort

//选择排序，时间复杂度O(N²)
func SelectionSort(datas []int) {
	length := len(datas)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i + 1; j < length; j++ {
			if datas[j] < datas[min] {
				min = j
			}
		}
		datas[i], datas[min] = datas[min], datas[i]
	}
}
