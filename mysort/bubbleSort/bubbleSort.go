package mysort

//冒泡排序
func BubbleSort(datas []int) {
	length := len(datas)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if datas[j+1] < datas[j] {
				datas[j], datas[j+1] = datas[j+1], datas[j]
			}
		}
	}
}
