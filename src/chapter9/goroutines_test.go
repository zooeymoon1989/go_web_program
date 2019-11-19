package main

import (
	"testing"
)

func TestPrint1(t *testing.T)  {
	print1()
}

func TestGoPrint1(t *testing.T)  {
	goPrint1()

}

//测试方法：
//go test -run x -bench . -cpu 1


//基准测试结果显示：
//goos: darwin
//goarch: amd64
//pkg: go_web_program/src/chapter9
//BenchmarkPrint1         100000000               14.7 ns/op
//BenchmarkGoPrint1        1000000              1296 ns/op
//PASS
//ok      go_web_program/src/chapter9     2.855s
//

//说明开启一个goroutine会有一定的损耗


func BenchmarkPrint1(b *testing.B)  {
	for i := 0 ; i < b.N ; i ++{
		print1()
	}
}

func BenchmarkGoPrint1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint1()
	}
}


func BenchmarkPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print2()
	}
}
func BenchmarkGoPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint2()
	}
}