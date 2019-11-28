package main

import (
	"fmt"
	"math"
)

//两个数组合并排序
func mergeArray(left []int, right []int) []int {

	var i, j, k = 0, 0, 0
	var done1, done2 = false, false
	length := len(left) + len(right)
	merged := make([]int, length)
	for i < len(left) || j < len(right) {

		if done1 && done2 {
			break
		}

		//如果第一个数组循环完毕
		if done1 {
			merged[k] = right[j]
			j++
			if len(right) == j {
				done2 = true
			}
			k++
			continue
		}

		//如果第二个数组循环完毕
		if done2 {
			merged[k] = left[i]
			i++
			if len(left) == i {
				done1 = true
			}
			k++
			continue
		}

		if left[i] < right[j] {
			merged[k] = left[i]
			i++
			if len(left) == i {
				done1 = true
			}
		} else {
			merged[k] = right[j]
			j++
			if len(right) == j {
				done2 = true
			}
		}
		k++
	}

	return merged
}

func mergeSort(array []int) []int {

	if len(array) < 2 {
		return array
	}

	mid := int(math.Floor(float64(len(array) / 2)))

	left := append(array[0:mid])
	right := append(array[mid:])
	mergedLeft := mergeSort(left)
	mergedRight := mergeSort(right)

	return mergeArray(mergedLeft, mergedRight)

}

func main() {

	var Array = []int{4, 65, 46, 24, 7, 54, 13, 78, 45, 65, 33, 66, 24}
	fmt.Println(Array)
	fmt.Println(mergeSort(Array))

}
