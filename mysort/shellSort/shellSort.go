package mysort

import "fmt"

// 希尔排序
func ShellSort(datas []int) {
	i, j, k := 0, 0, 0
	length := len(datas)
	stepArr := getSedgewickStepArr(length)
	for i = len(stepArr) - 1; i >= 0; i-- {
		for j = stepArr[i]; j < length; j++ {
			tmp := datas[j]
			for k = j; k >= stepArr[i]; k -= stepArr[i] {
				if tmp < datas[k-stepArr[i]] {
					datas[k] = datas[k-stepArr[i]]
				} else {
					break
				}
			}
			datas[k] = tmp
		}
		fmt.Println(datas)
	}
}

// 生成Sedgewick增量序列数组，最坏时间复杂度为O(N^4/3)；平均时间复杂度约为O(N^7/6)
// 序列中是 9 * 4^i - 9 * 2^i + 1 或 4^i - 3 * 2^i + 1
// 变形为：D=9*(2^(2i)-2^i)+1 或 2^(2i+4)-3*2^(i+2)+1 , i>=0；第二个公式前两项为-1，所以加二
func getSedgewickStepArr(n int) (arr []int) {
	for i := 0; i <= n; i++ {
		tmp := 9*((1<<(2*i))-(1<<i)) + 1
		if tmp < n {
			arr = append(arr, tmp)
		} else {
			break
		}
		tmp = (1 << (2*i + 4)) - 3*(1<<(i+2)) + 1
		if tmp < n {
			arr = append(arr, tmp)
		} else {
			break
		}
	}
	return
}
