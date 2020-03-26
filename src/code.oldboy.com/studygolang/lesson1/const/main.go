package main

import "fmt"

const pi = 3.14

// 批量定义
const (
	a = 100
	b = 1000
	c
	d
)

// iota 枚举
const (
	aa = iota // 0
	bb
	cc
	dd
)

const (
	n1 = iota
	n2 = 100
	n3
	n4
)

const n5 = iota

// const 中每新增一行常量声明 将使 iota 计数一次 (iota 可理解为const 语句块中的行索引)
// k2 虽然没有赋值为 iota , 但是它属于 const 里面新增加的一行，所以
// k2 = 100 (隐式的 iota 了一次)， 所以到 k3 时，值为2
const (
	k1 = iota
	k2 = 100
	k3 = iota
	k4 = iota
)

// const 里面只要有 iota 关键字，则 从第一行开始 iota 已经被隐式赋值了
// m1 = 200 (隐式, 赋值为 iota )
// m5 未赋值，则与 m4 相同 ， m5 也被赋值为 iota ,则 继续加1
const (
	m1 = 200
	m2 = iota
	_
	m3 = 100
	m4 = iota
	m5
)

// const 常量定义里面出现了一行 空格，不影响结果，go format 时会自动将空行去掉
// 空行 不属于 常量声明代码，所以不算加 行
const (
	p1 = 200

	p2 = iota
	p3 = 100
	p4
)

const (
	_  = iota
	KB = 1 << (10 * iota) // 1<<10  ==>  2 的 10次方 =  1024 .
	MB = 1 << (10 * iota) // 1<<20  ==>  2 的 20次方 = 1024*1024
	GB = 1 << (10 * iota) // 1<<30
	TB = 1 << (10 * iota) // 1<<40
	PB = 1 << (10 * iota) // 1<<50
)

// 同一行的 iota 不增加 1
const (
	a1, b1 = iota + 1, iota + 2 // iota = 0;  a1=1, b1=2
	a2, b2                      // iota = 1;  a2=2, b2=3
	a3, b3                      // iota = 2;  a3=3, b3=4
)

func main() {

	fmt.Println(pi)

	// 常量不能被赋值
	// pi = 3.12434543    // cannot assign to pi

	// c, d 默认值与 b 相同
	fmt.Println(a, b, c, d) // "100 1000 1000 1000"

	// 依次加1
	fmt.Println(aa, bb, cc, dd) // "0 1 2 3"

	fmt.Println(n1, n2, n3, n4, n5) // "0 100 100 100 0"

	fmt.Println(k1, k2, k3, k4) // "0 100 2 3"

	fmt.Println(m1, m2, m3, m4, m5) // "200 1 100 4 5"

	fmt.Println(p1, p2, p3, p4) // "200 1 100 100"

	fmt.Println(KB, MB, GB, TB, PB) // "1024 1048576 1073741824 1099511627776 1125899906842624"

	fmt.Println(a1, b1, a2, b2, a3, b3) // "1 2 2 3 3 4"

	// iota 总结：
	/*
		const 声明如果不写， 默认就和上一行一样
		遇到 const iota 就初始化为 零
		const 中每新增加一行变量声明 iota 就递增 1
	*/
}
