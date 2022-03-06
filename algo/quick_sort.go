package main

import "fmt"

/*
快速排序算法：
- 将数据根据一个值按照大小分成左右两边，左边小于此值，右边大于
- 将两边数据进行递归调用步骤1
- 将所有数据合并
*/

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	splitData := arr[0]
	low := make([]int, 0, 0)
	hight := make([]int, 0, 0)
	mid := make([]int, 0, 0)
	mid = append(mid, splitData)

	for i := 1; i < len(arr); i++ {
		if arr[i] < splitData {
			low = append(low, arr[i])
		} else if arr[i] > splitData {
			hight = append(hight, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}

	low, hight = quickSort(low), quickSort(hight)
	ret := append(append(low, mid...), hight...)
	return ret
}

func main() {
	arr := []int{1, 5, 3, 4, 8}
	ret := quickSort(arr)
	fmt.Println(ret)
}
