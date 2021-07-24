package test

import (
	"../demo"
	"testing"
)


func BenchmarkLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		demo.LearnLock()
	}
}