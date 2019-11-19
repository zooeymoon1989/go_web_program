package main

import (
	"fmt"
	"time"
)

func callerA(c chan int)  {
	c <- 5
	close(c)
}

func callerB(c chan string)  {
	c <- "foo bar"
	close(c)   //close关闭一个channel，但是只能是发送方或者是双向的channel。
}

func main() {
	a , b := make(chan int) , make(chan string)
	go callerA(a)
	go callerB(b)

	var msg string
	var integer int

	ok1 , ok2 := true , true

	for ok1 || ok2 {
		for i := 0 ; i < 5 ; i ++ {
			time.Sleep(1*time.Microsecond)
			select {
			case integer , ok1 = <-a:	//如果channel被关闭了，再从channel中获取值为初始变量,同时ok1为false
				if ok1{
					fmt.Printf("%d from A \n" , integer)
				}
			case msg , ok2 = <-b:
				if ok2 {
					fmt.Printf("%s from B \n" , msg)
				}
			}
		}
	}


}
