package main

import (
	mysort "dataStructAlg/mysort/shellSort"
)

func main() {
	data := []int{34, 37, 2, 58, 67, 34, 567, 67, 88, 54, 1, 4, 8, 34, 0, 67, 3, 6, 434}
	mysort.ShellSort(data)
}
