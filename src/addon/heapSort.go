package main

import "math"

func main() {

	var Array = []int{4, 65, 46, 24, 7, 54, 13, 78, 45, 65, 33, 66, 24}

}

func makeHeap(array []int) {
	length := len(array)
	for i := math.Floor(float64(length / 2)); i > 0; i-- {
		modifyHeap()
	}
}

func modifyHeap() {

}

func heapSort() {

}

func swapHeap() {

}
