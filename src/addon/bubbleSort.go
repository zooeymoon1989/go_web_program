package main

import (
	"fmt"
)

// 冒泡排序
func main() {
	var Array = []int{4, 65, 46, 24, 7, 54, 13, 78, 45, 65, 33, 66, 24}
	var tmp int
	for i := 0; i < len(Array)-1; i++ {
		for j := 0; j < len(Array)-1-i; j++ {
			if Array[j] > Array[j+1] {
				tmp = Array[j+1]
				Array[j+1] = Array[j]
				Array[j] = tmp
				fmt.Println(Array)
			}
		}
	}
}
