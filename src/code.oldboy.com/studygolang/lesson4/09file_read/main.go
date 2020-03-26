package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// file read 操作
func main() {
	// test1()
	// test2()
	// test3()
	// test4()
	test5()
}

// 文件的打开与关闭
func test1() {
	file, err := os.Open("./xx.txt")
	if err != nil {
		fmt.Println("open file failed, err: ", err) // " open ./xx.txt: The system cannot find the file specified."
	} else {
		fmt.Println("open file success ")
		file.Close()
	}
}

// 文件读取
func test2() {
	file, err := os.Open("./xx.txt")
	if err != nil {
		fmt.Println("open file failed, err: ", err) // " open ./xx.txt: The system cannot find the file specified."
		return
	}
	// 文件能打开
	defer file.Close()

	var b = make([]byte, 128, 128)
	n, err := file.Read(b)

	if err == io.EOF { // EOF : End of File
		fmt.Println("文件已经读完")
		return
	}

	if err != nil {
		fmt.Println("read from file failed , err: ", err)
	}
	fmt.Printf("读取了 %d 个字节\n", n)
	fmt.Println("读取文件内容：")
	fmt.Println(string(b))
}

// 文件读取2 (循环读取,直到读到文件末尾)
func test3() {
	file, err := os.Open("./xx.txt")
	if err != nil {
		fmt.Println("open file failed, err: ", err) // " open ./xx.txt: The system cannot find the file specified."
		return
	}
	// 文件能打开
	defer file.Close()

	var fileContent string // 循环读完之后，采用字符串保存读取的内容
	var count int          // 统计总共读取的字节数 (一个中文占  3 个字节)

	for { // 每次读 128 个字节，直到读到文件末尾，就 return
		var b = make([]byte, 128, 128)
		n, err := file.Read(b)

		count += n
		fileContent += string(b)

		if err == io.EOF { // EOF : End of File
			fmt.Println("文件已经读完")

			fmt.Printf("读取了 %d 个字节\n", count)
			fmt.Println("读取文件内容：")
			fmt.Println(fileContent)

			return
		}

		if err != nil {
			fmt.Println("read from file failed , err: ", err)
		}

	}
}

// bufio 读数据
// 为什么会有缓冲区
//1, reader.ReadString('\n') 可以设定为 一行一行的读，这样就避免了 test1()中每次只规定读128个字节 导致的乱码问题
//2, 缓冲区不会像 test1() 中每次读取文件的时候都操作 磁盘，效率不高，所以就在 程序 和 磁盘之间加入了整个 缓冲区，提高效率
//3, 缓冲区缺点： 它是一次性将文件读取到 缓存中，相比之下增加了内存使用，并且一旦程序崩溃，内存泄漏，缓冲区里的数据就会丢失。
func test4() {
	file, err := os.Open("./xx.txt")
	if err != nil {
		fmt.Println("打开文件失败")
		return
	}
	defer file.Close()

	// 利用缓冲区从文件读取数据
	reader := bufio.NewReader(file)

	for {
		str, err := reader.ReadString('\n') // 换行符 表示，一行一行的读
		if err == io.EOF {                  // 文件读完，不代表 str 没有数据了， 例如：读到最后一行，读完了，但是 str 保存了最后一行的数据
			fmt.Print(str)
			// fmt.Println("文件已读完")
			return
		}
		if err != nil {
			fmt.Println("读取文件失败")
			return
		}
		fmt.Print(str)
	}
}

// ioutil 读取文件
func test5() {
	content := readFile("./xx.txt")
	fmt.Println(content)
}

// ioutil.ReadFile()
// 此方法直接在内部包含了 打开文件，读取文件的操作 (查看源码发现，还包含了 defer 文件关闭的操作)，无须我们再手动添加
func readFile(fileName string) string {
	content, err := ioutil.ReadFile(fileName) // content 是 []byte 类型
	if err != nil {
		fmt.Println("文件读取失败, err : ", err)
	}
	return string(content) //  将 []byte 强制转换为 string
}
