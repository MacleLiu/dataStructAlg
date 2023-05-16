package mysort

//堆排序
func HeapSort(datas []int) {
	length := len(datas)
	//构建Max堆，完全二叉树的最后一个非叶子节点的序号为N/2-1
	for i := length/2 - 1; i >= 0; i-- {
		percDown(datas, i, length)
	}

	//DelectMax操作
	for i := length - 1; i > 0; i-- {
		datas[0], datas[i] = datas[i], datas[0]
		percDown(datas, 0, i)
	}
}

func percDown(datas []int, i, length int) {
	var child, tmp int
	for tmp = datas[i]; leftChild(i) < length; i = child {
		child = leftChild(i)
		//节点i的左孩子不是最后一个节点
		//如果右孩子比左孩子大，则将索引加一指向右孩子
		if child != length-1 && datas[child+1] > datas[child] {
			child++
		}
		//比较节点i与i的最大孩子的大小
		if tmp < datas[child] {
			datas[i] = datas[child]
		} else {
			break
		}
	}
	datas[i] = tmp
}

func leftChild(n int) int {
	return 2*n + 1
}
