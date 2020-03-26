package main

import "fmt"

func main() {

	a := [...]int{1, 3, 5, 7, 9, 11, 13}

	fmt.Printf("修改前a:%p\n", &a) // 从

	b := a[:]

	b[0] = 100

	fmt.Println(a[0])           // "100"
	fmt.Printf("修改后a:%p\n", &a) // "修改后a:0xc000010100"
	fmt.Printf("b:%p\n", b)     // "b:0xc000010100",
	// b 从 数组a 的第一个元素开始切，所以 b 与 a 的地址相同
	// 虽然修改了第一个元素的值，但是地址没有影响， 根据地址修改值，改动的是值，不会修改到地址

	c := a[2:5]
	fmt.Println(c)      // "[5 7 9]"
	fmt.Println(len(c)) // "3"
	// 长度，从切片a 的第一个元素开始算，一直到它指定切片的位置结束，这之间的元素的个数是它的长度
	// 容量，从切片a 的第一个元素开始算，一直到底层数组最后的元素，这之间元素的个数是它的容量
	fmt.Println(cap(c))     // "5"
	fmt.Printf("c:%p\n", c) // "d:0xc000010110"

	d := c[:5]
	fmt.Printf("d:%p\n", d) // "d:0xc000010110"

	e := c[1:5]
	fmt.Printf("e:%p\n", e) // "e:0xc000010118"

	// c 和 d 起始元素都是同一个，所以 地址相同， e 从 c 的第二个元素开始切，所以地址跟 c,d 都不同

	// 切片的三要素：
	/*
		-- 第一个元素的地址 (从哪一个元素开始切)
		-- 长度
		-- 容量
	*/

	e = append(e, 100, 200, 300)
	fmt.Println(e)              // "[7 9 11 13 100 200 300]"
	fmt.Println(len(e), cap(e)) // "7 8"

	// e 在 append 的时候进行了扩容，所以 e 不再与 c 共用同一个底层数组，所以当修改 e 时，c 不会受到影响
	e[1] = 4
	fmt.Println(e) // "[7 4 11 13 100 200 300]"
	fmt.Println(c) // "[5 7 9]"

}