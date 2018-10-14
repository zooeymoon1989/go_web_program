package main

import (
	"fmt"
	"time"
)

type A struct {}


type IA interface {
	doSomething() int
	healSomething() int
	something()
}

func (a *A)something()  {
	fmt.Println("foo")
}

func (a *A) doSomething() int {
	return 1
}

func (a *A) healSomething() int {
	return 2
}

func pinger(c chan string)  {
	for i :=0 ; ; i++ {
		c <- "ping"
	}
}

func printer(c chan string)  {
	for {
		msg := <- c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c = make(chan string)
	go pinger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)


}
