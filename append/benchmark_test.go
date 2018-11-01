package main_test

import (
	"testing"
)

const size = 1000000

func BenchmarkLenAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sa := make([]int, size)
		for i := 0; i < size; i++ {
			sa = append(sa, i)
		}
	}
}

func BenchmarkCapAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sa := make([]int, 0, size)
		for i := 0; i < size; i++ {
			sa = append(sa, i)
		}
	}
}

func BenchmarkNotInitAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var sa []int
		for i := 0; i < size; i++ {
			sa = append(sa, i)
		}
	}
}

func BenchmarkLenIdx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sa := make([]int, size)
		for i := 0; i < size; i++ {
			sa[i] = i
		}
	}
}
