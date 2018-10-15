package main

import "testing"

func BenchmarkDecode(b *testing.B)  {
	for i := 0 ; i < b.N ; i ++ {
		decode("something.json")
	}
}

func BenchmarkUnmarshal(b *testing.B)  {
	for i := 0 ; i < b.N ; i ++ {
		unmarshal("something.json")
	}
}