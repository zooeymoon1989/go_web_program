package main

import "fmt"

func main() {
	arr := [5]int{1, 3, 5, 6, 7}
	fmt.Printf("addr:%p\n", &arr)
	fmt.Printf("addr:%p\n", &arr[1])
	s1 := arr[:]
	fmt.Printf("addr:%p\n", &s1)

	changeSlice(s1)
}

func changeSlice(s []int) {
	fmt.Printf("addr:%p\n", &s)
	fmt.Printf("addr:%p\n", &s[1])
}
