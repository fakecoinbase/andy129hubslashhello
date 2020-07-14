// @Title  
// @Description  
// @Author  yang  2020/7/12 12:44
// @Update  yang  2020/7/12 12:44
package main

import (
	"crypto/sha256"
	"fmt"
)

/*
	go 语言 sha256 包中实现了两种 哈希函数，分别是 sha256和 sha224
 */
func main() {
	hash := sha256.New()
	hash.Write([]byte("hello sha256"))
	result := hash.Sum(nil)
	fmt.Printf("%x\n", result)   // 433855b7d2b96c23a6f60e70c655eb4305e8806b682a9596a200642f947259b1

	res := sha256.Sum256([]byte("hello sha256"))
	fmt.Printf("%x\n", res)      // 433855b7d2b96c23a6f60e70c655eb4305e8806b682a9596a200642f947259b1

	res2 := sha256.Sum224([]byte("hello sha224"))
	fmt.Printf("%x\n", res2)     // b9603bd24ce66692e2462c451aaa0e4fd91b8b2a4a345058164f1b54
}
