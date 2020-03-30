package fib

// Fib 计算第n 个斐波那契数列
func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2) // 递归
}
