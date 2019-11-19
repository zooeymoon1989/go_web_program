package main

import "fmt"

// 选择排序
func main() {
	var Array = []int{4, 65, 46, 24, 7, 54, 13, 78, 45, 65, 33, 66, 24}

	var lowest int
	var tmp int

	for i := 0; i < len(Array) - 1; i++ {
		lowest = i
		for j := i; j < len(Array); j++ {
			if Array[lowest] > Array[j] {
				lowest = j
			}
		}
		tmp = Array[i]
		Array[i] = Array[lowest]
		Array[lowest] = tmp
	}

	fmt.Println(Array)
}
