package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// sha256 操作
func main(){
	// test1()
	// test2()


	applySha256()
}

// sha256 的第一种形式
func test1(){

	// crypto/sha256/sha256.go
	/*
			// The size of a SHA256 checksum in bytes.
			const Size = 32

			// Sum256 returns the SHA256 checksum of the data.
			func Sum256(data []byte) [Size]byte {
	 */

	// sum 为 [Size]byte ===>  [32]byte
	sum := sha256.Sum256([]byte("hello yang"))
	// fmt.Println(string(sum[:]))   // 可以这样打印，但是打印出来为 乱码
	// fmt.Println(sum)  // [162 250 207 190 173 240 158 104 159 123 107 181 2 24 117 177 218 112 230 164 17 8 182 246 247 169 239 224 145 223 86 150 67]
	// 64 个长度，每个字母代表 4位， 两个字母代表一个字节， 总共是 32个字节， 256位
	fmt.Printf("%X", sum)   // A2FACFBEADF09E689F7B6BB5021875B1DA70E6A4B2B6F6F7A9EFE091DF569643
}

/*
		 %x    每个字节用两字符十六进制数表示（使用a-f）
		 %X    每个字节用两字符十六进制数表示（使用A-F）

 */

// sha256 的第二种形式
func test2() {

	// crypto/sha256/sha256.go
	/*
			// New returns a new hash.Hash computing the SHA256 checksum. The Hash
			// also implements encoding.BinaryMarshaler and
			// encoding.BinaryUnmarshaler to marshal and unmarshal the internal
			// state of the hash.
			func New() hash.Hash {

	 */
	// h 为  hash.Hash 类型
	h := sha256.New()
	h.Write([]byte("hello yang"))

	fmt.Printf("%X", h.Sum(nil))  // A2FACFBEADF09E689F7B6BB5021875B1DA70E6A4B2B6F6F7A9EFE091DF569643a2facfbeadf09e68

	fmt.Printf("%x", h.Sum(nil))  // 9f7b6bb5021875b1da70e6a4b2b6f6f7a9efe091df569643a2facfbeadf09e689f7b6bb5021875b1
}

// 在文件操作中使用  sha256
func applySha256() {

	// h 为  hash.Hash 接口类型, 定义了 io.Writer
	h := sha256.New()
	f,err := os.Open("hash.test")
	if err != nil {
		fmt.Println("open failed : ", err)
		return
	}

	defer f.Close()

	_,err = io.Copy(h, f)   // 将文件内容 写入 h
	if err != nil {
		fmt.Println("copy failed : ", err)
		return
	}
	fmt.Printf("%X", h.Sum(nil))
}

