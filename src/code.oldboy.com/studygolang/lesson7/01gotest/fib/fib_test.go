package fib

import "testing"

// 性能比较

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(2)
	}
}

// 供内部调用
func benchmarkFib(b *testing.B, n int) {
	for i := 0; i < b.N; i++ {
		Fib(n)
	}
}

// 计算第2 个斐波那契数列
func BenchmarkFib2(b *testing.B) {
	benchmarkFib(b, 2)
}

// 计算第 20个斐波那契数列
func BenchmarkFib20(b *testing.B) {
	benchmarkFib(b, 20)
}

// 输入指令： go test -bench=Fib -benchmem

/*  打印信息：

goos: windows
goarch: amd64
pkg: code.oldboy.com/studygolang/lesson7/01gotest/fib
BenchmarkFib-8          207282060                5.77 ns/op            0 B/op          0 allocs/op
BenchmarkFib2-8         199359391                6.00 ns/op            0 B/op          0 allocs/op
BenchmarkFib20-8           25086             45416 ns/op               0 B/op          0 allocs/op
PASS
ok      code.oldboy.com/studygolang/lesson7/01gotest/fib        5.250s

*/
