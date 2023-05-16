package mysort

//归并排序
func MergeSort(datas []int) []int {
	length := len(datas)
	if length < 2 {
		return datas
	}
	middle := length / 2
	return merge(MergeSort(datas[:middle]), MergeSort(datas[middle:]))
}

//归并例程
func merge(left, right []int) (result []int) {
	for len(left) != 0 && len(right) != 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	if len(left) != 0 {
		result = append(result, left...)
	}
	if len(right) != 0 {
		result = append(result, right...)
	}
	return
}
