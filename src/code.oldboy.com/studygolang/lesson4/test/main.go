package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

// 练习
// 1, 编写一个函数，接收Time 类型的参数，函数内部将传进来的时间格式化输出为 "2017/06/19 20:30:05" 格式
// 2, 编写程序统计一段代码的执行耗时时间，单位精确到 微妙。
func main() {

	//fmt.Println("test")

	// test1()
	// test2()
	// test3()
	test4()
}

// 练习2：
func test2() {

	// codeRunningTime()

	codeRunningTime2()
}

// 计算一段代码执行时间
func codeRunningTime() {
	// 程序执行前：
	before := time.Now()
	beforeNanoSec := before.UnixNano()

	for _, v := range "fso中国fdsf3fsdfjlsdf" {
		fmt.Println(v)
	}

	// 程序执行后
	after := time.Now()
	afterNanoSec := after.UnixNano()

	nanoSec := afterNanoSec - beforeNanoSec
	microSec := nanoSec / 1000
	milliSec := microSec / 1000
	second := milliSec / 1000

	fmt.Println("程序执行时间(纳秒数)：", nanoSec)
	fmt.Println("程序执行时间(微妙数)：", microSec)
	fmt.Println("程序执行时间(毫秒数)：", milliSec)
	fmt.Println("程序执行时间(秒数)：", second)
}

// 采用 time.Since() 计算执行时间
func codeRunningTime2() {

	start := time.Now()

	time.Sleep(time.Second * 2) // 一般建议这种写法,  调用 time 包里的常量时间
	// time.Sleep(20000)   // 虽然不会报错，但不建议直接传入 int 数字

	// since 用法， 传入一个起始时间，得到一个 起始 至 现在时间的差值。
	duration := time.Since(start)
	fmt.Println(duration) // 2.0008534s

}

// 练习1：
func test1() {

	now := time.Now()
	layout := "2006/01/02 15:04:05"
	formatStr := formatTimeToString(now, layout)
	fmt.Println(formatStr) // "2020/03/20 12:46:05"
}

// 将 time.Time 结构体类型 转换为指定格式的 字符串
func formatTimeToString(t time.Time, layout string) string {

	return t.Format(layout)

}

// 实现文件拷贝功能
func test3() {
	writtenCount, err := CopyFile("dst.txt", "xx.txt")
	if err != nil {
		fmt.Printf("copy file failed, written : %d , err : %v\n", writtenCount, err)
		return
	}
	fmt.Println("拷贝成功, 拷贝数 : ", writtenCount)
}

// CopyFile 是一个实现文件拷贝的方法
func CopyFile(dstName, srcName string) (written int64, err error) {
	// 以读方式打开源文件
	src, err := os.Open(srcName)

	if err != nil {
		fmt.Printf("open %s failed, err : %v. \n", srcName, err)
		return
	}
	// 注意 defer 语句块声明的 位置， 不能在 if err != nil  前面声明，
	// 因为 只有在确保 err == nil ，文件正常打开的情况下，我们才会执行 defer 语句，处理 文件关闭操作，所以文件一旦打开失败，我们无需进行 关闭操作
	// 经过测试，我们把 defer 语句块放到  if err != nil 前面，也能正常执行关闭操作， 但是我们一般不这样写。

	// defer 语句1
	defer func() {
		src.Close()
		fmt.Println("源文件关闭")
	}()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		fmt.Printf("open %s failed, err : %v. \n", dstName, err)
		return
	}
	// defer 语句2
	defer func() {
		dst.Close()
		fmt.Println("目标文件关闭")
	}()
	return io.Copy(dst, src) // 调用 io包中的 Copy() 方法

	// 回顾一下原来的内容，  两个 defer 语句， 先执行的是 defer 语句2， 最后执行 defer 语句1
}

// 使用文件操作相关知识，模拟实现 linux 平台的 cat 命令的功能
// 示例： go run main.go cat dst.txt
// test4() 一被执行， flag.Parse() 就开始在获取指令参数，所以我们在执行时，直接在后面带上 命令，否则它无法获取指令
func test4() {
	flag.Parse() // 解析命令行参数
	if flag.NArg() == 0 {
		// 如果没有参数默认从标准输入读取内容
		cat(bufio.NewReader(os.Stdin))
	}
	// 依次读取每个指定文件的内容并打印到终端
	for i := 0; i < flag.NArg(); i++ {
		fmt.Fprintf(os.Stdout, "reading args %s\n", flag.Arg(i))
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stdout, "reading from %s failed, err : %v\n", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(f))
	}
}

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		// 注意这里的参数是 字符， 还需要注意的是，如果你的文本里最后的内容没有敲一个换行符的话，
		// 它这里是不会读取你最后一行的内容的，因为它的规则是以 \n 为一行的结尾来 读取
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout, "%s", buf)
	}
}
