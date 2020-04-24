package main

import "sync"

// 求 斐波那契数列, 强制性让CPU 使用率达到顶峰 用于验证 grafana 监控CPU 使用率是否起作用
func main() {

	var wg sync.WaitGroup
	for i:=0;i<100;i++ {
		wg.Add(1)
		go fib(10000)
	}
	wg.Wait()
}

func fib(n int64) int64 {
	if n<=2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}
