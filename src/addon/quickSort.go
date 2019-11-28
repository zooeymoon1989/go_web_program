package main

import "fmt"

func main() {

	var Array = []int{4, 65, 46, 24, 7, 54, 13, 78, 45, 65, 33, 66, 24}
	fmt.Println(Array)
	fmt.Println(quickSort(Array))

}

// 快速排序
// 这里我是找了最后一个元素当作pivot
// 之后设置了两个数组
// 小于等于pivot的元素放入left中
// 大于pivot的元素放入right中
// 同时对左右两个数组进行递归
// 整理返回的数据
func quickSort(array []int) []int {

	if len(array) <= 1 {
		return array
	}

	var pivot = array[len(array)-1]
	var left []int
	var right []int

	for i := 0; i < len(array)-1; i++ {
		if array[i] <= pivot {
			left = append(left, array[i])
		} else {
			right = append(right, array[i])
		}
	}

	afterQuickSortLeft := quickSort(left)
	afterQuickSortRight := quickSort(right)

	foo := append(afterQuickSortLeft, pivot)
	foo = append(foo, afterQuickSortRight...)

	return foo

}
