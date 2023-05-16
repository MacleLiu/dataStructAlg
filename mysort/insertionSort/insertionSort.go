package mysort

// 插入排序算法，时间复杂度O(N²)
func InsertionSort(datas []int) {
	var tmp, p, j int
	for p = 1; p < len(datas); p++ {
		tmp = datas[p]
		for j = p; j > 0 && datas[j-1] > tmp; j-- {
			datas[j] = datas[j-1]
		}
		datas[j] = tmp
	}
}
