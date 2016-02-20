package bench

import (
	"testing"
)

func BenchmarkNoSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myFuncA(i, i)
	}
}

func BenchmarkFuncPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myFuncPointer(&i, &i)
	}
}

func Benchmark_Switch文と関数呼び出し(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myFuncSwitch(9, i, i)
	}
}

func Benchmark_if文と直接演算し(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myFuncIfB(9, i, i)
	}
}

func Benchmark_Switch文と直接演算(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myFuncSwitchB(9, i, i)
	}
}
