package main

import (
	"fmt"
	"math"
)

//  希尔排序
func main() {
	var Array = []int{4, 65, 46, 24, 7, 54, 13, 78, 45, 65, 33, 66, 24}
	fmt.Println(Array)
	//for gap := math.Floor(float64(len(Array))); gap > 0; gap = math.Floor(gap / 2) {
	//
	//	for i := gap; i < float64(len(Array)); i++ {
	//		j := i
	//		current := Array[int(i)]
	//		for j-gap >= 0 && current < Array[int(j-gap)] {
	//			Array[int(j)] = Array[int(j-gap)]
	//			j = j - gap
	//		}
	//
	//		Array[int(j)] = current
	//	}
	//}
	for i := math.Floor(float64(len(Array) / 2)); i > 0; i = math.Floor(i / 2) {

		for j := 1 ; j < len(Array) ; j++ {

			for k := j ; k < len(Array) ; k = k + int(i){
				if Array[k] < Array[k-1]{
					tmp := Array[k]
					Array[k] = Array[k-1]
					Array[k-1] = tmp
				}
			}

		}
	}

	fmt.Println(Array)

}
