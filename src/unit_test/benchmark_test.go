package service

import "testing"

func Benchmark_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(i, 1)
	}
}


// go test -v -bench=. benchmark_test.go

/*
goos: darwin
goarch: arm64
Benchmark_Add
Benchmark_Add-8         1000000000               0.3181 ns/op
PASS
ok      command-line-arguments  0.824s

“0.31 ns/op”表示每一个操作耗费多少时间（纳秒）
*/
