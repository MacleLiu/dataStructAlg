package mysort

import (
	mysort "dataStructAlg/mysort/insertionSort"
)

// 快速排序，使用三数中值分割法确定枢纽元
func QuickSort(datas []int) {
	quickSort(datas, 0, len(datas)-1)
}

func quickSort(datas []int, left, right int) {
	//对长度小于20的数组使用插入排序
	if right-left < 20 {
		mysort.InsertionSort(datas)
	} else {
		pivot := median3(datas, left, right)

		//median3的附加作用对三个元素进行了排序
		//所以i,j的初始值比他们的正确值超出1，而不存在需要考虑的特殊情况
		i := left
		j := right - 1
		for {
			/* 当datas[i]=datas[j]=pivot时，出现无限循环
			for ; datas[i] < pivot; i++ { //i右移，移过小于枢纽元的元素
			}
			for ; datas[j] > pivot; j-- { //j左移，移过大于枢纽元的元素
			}
			*/
			for { //i右移，移过小于枢纽元的元素
				if i++; datas[i] < pivot {
				} else {
					break
				}
			}
			for { //j左移，移过大于枢纽元的元素
				if j--; datas[j] > pivot {
				} else {
					break
				}
			}
			//i,j移动停止时，如果i,j未交错，交换i,j指向的元素值
			if i < j {
				swap(datas, i, j)
			} else {
				break
			}
		}
		swap(datas, i, right-1) //复位枢纽元

		quickSort(datas, left, i-1)  //对小于枢纽元的部分进行快排
		quickSort(datas, i+1, right) //对大于枢纽元的部分进行快排
	}

}

// 三数中值分割法，并将三个元素放置到合适的位置
func median3(datas []int, left, right int) (median int) {
	center := (left + right) / 2
	if datas[0] > datas[center] {
		swap(datas, 0, center)
	}
	if datas[0] > datas[right] {
		swap(datas, 0, right)
	}
	if datas[center] > datas[right] {
		swap(datas, center, right)
	}

	swap(datas, center, right-1)
	return datas[right-1]
}

func swap(datas []int, i, j int) {
	datas[i], datas[j] = datas[j], datas[i]
}
