package pool_test

import (
	"testing"

	pool "github.com/midorigreen/go-sandbox/goroutine-pool"
)

func BenchmarkNotConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.NotConcurrent()
	}
}

func BenchmarkGoroutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.Goroutine()
	}
}

func BenchmarkGoroutinePool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.GoroutinePool()
	}
}
