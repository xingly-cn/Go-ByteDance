package main

/**
基准测试
*/
import (
	"github.com/awnumar/fastrand"
	"testing"
)

var ServerIndex [10]int

func InitServerIndex() {
	for i := 0; i < 10; i++ {
		ServerIndex[i] = i + 100
	}
}

func Select() int {
	return ServerIndex[fastrand.Intn(10)]
}

// 串行
func BenchmarkSelect(b *testing.B) {
	InitServerIndex()
	b.ResetTimer() // 初始化数组的时间, 刨去
	for i := 0; i > b.N; i++ {
		Select()
	}
}

// 并行
func BenchmarkSelectParallel(b *testing.B) {
	InitServerIndex()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Select()
		}
	})
}
