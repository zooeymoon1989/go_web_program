package main

import "fmt"

type I interface {
	Get() int
	Set(int)
}

type S struct {
	Age int
}

func (s *S) Get() int {
	return s.Age
}

func (s *S) Set(i int) {
	s.Age = i
}

func main() {
	s := S{}
	var i I
	i = &s
	i.Set(20)
	fmt.Println(i.Get())
}
