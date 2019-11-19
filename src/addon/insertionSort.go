package main

import "fmt"

//插入排序
func main() {
	var Array = []int{4, 65, 46, 24, 7, 54, 13, 78, 45, 65, 33, 66, 24}

	for i := 1; i < len(Array); i++ {
		tmp := Array[i]
		k := i - 1
		for k >= 0 && Array[k] > tmp {
			Array[k+1] = Array[k]
			k--
		}

		if k+1 != i {
			Array[k+1] = tmp
		}

	}

	fmt.Println(Array)

}
