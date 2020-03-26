package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// file write 操作
func main() {
	fmt.Println("文件写入")

	// test1()
	// test2()
	test3()
}

// 写入文件
// go 语言内置包 os 中的方法：func OpenFile(name string, flag int, perm FileMode) (*File, error)
// 参数1， name string :  表示文件名字
// 参数2， flag int 如下(下面几个常用)：
// os.O_WRONLY , 只写 (write only)
// os.O_CREATE , 如果要读取的文件不存在， 则会创建，不添加这个属性, 则会报错
// os.O_APPEND , 追加文本
// os.O_TRUNC  , 每次写入文件时，都先清空文件里的内容，然后再追加新的内容
// 参数3, perm FileMode ,   采用 linux 里面权限操作:  0777,0666, 等
func test1() {

	file, err := os.OpenFile("xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	defer file.Close()
	// Write([]byte),  WriteString(string)  两个常用的方法
	file.Write([]byte("字符串转化为字节数组\n"))
	file.WriteString("直接写入一个字符串")

}

// 通过 缓冲区 写入文件
func test2() {
	file, err := os.OpenFile("xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}

	defer func() {
		file.Close()
		fmt.Println("文件关闭")
	}()

	writer := bufio.NewWriter(file)
	writer.WriteByte('c') // 单独写一个字节
	writer.WriteByte('\n')
	writer.WriteString("直接写入一个字符串\n")
	writer.WriteRune('中') // 可以写入一个中文 字符

	writer.Flush() // 最后一定要调用 Flush() 方法， 将缓冲区内容写入到文件中
}

// 通过 ioutil.Write  写入文件  （写入参数为 []byte）
func test3() {
	err := ioutil.WriteFile("xx.txt", []byte("ioutil.Write 写入字节数组"), 0666)
	if err != nil {
		fmt.Println("写入文件失败", err)
		return
	}
}

/*  go 语言 ioutil 包中 WriteFile  源码，发现 默认的 flag 就是  os.O_WRONLY|os.O_CREATE|os.O_TRUNC
	func WriteFile(filename string, data []byte, perm os.FileMode) error {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return err
	}
			.....
	return err
}

*/
