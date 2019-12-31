package main

import (
	"fmt"
	"time"
)

func main() {
	//var s sync.Mutex
	//var wg sync.WaitGroup
	////defer s.Unlock()
	//var value1 , value2 , value3 int32
	//
	//numAdd := func(n *int32) {
	//	*n = *n + 1
	//	//fmt.Println("numAdd:",*n)
	//}
	//
	//syncAdd := func(n *int32) {
	//	s.Lock()
	//	*n++
	//	//fmt.Println("syncAdd:",*n)
	//	s.Unlock()
	//}
	//
	//atomicAdd := func(n *int32) {
	//	atomic.AddInt32(n , 1)
	//	//fmt.Println("atomicAdd:",*n)
	//}
	//
	//wg.Add(1000000)
	//for i := 0 ; i < 1000000 ; i ++ {
	//	go numAdd(&value1)
	//	go syncAdd(&value2)
	//	go atomicAdd(&value3)
	//	wg.Done()
	//}
	//wg.Wait()
	//
	//fmt.Println("value1 is",value1)
	//fmt.Println("value2 is",value2)
	//fmt.Println("value3 is",value3)

	fmt.Println(time.Now().Format("2006-01-02"))

}