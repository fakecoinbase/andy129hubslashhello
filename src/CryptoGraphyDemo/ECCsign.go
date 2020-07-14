// @Title  
// @Description  
// @Author  yang  2020/7/12 15:20
// @Update  yang  2020/7/12 15:20
package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

// ECC sign

func main() {
	message := []byte("hello world!")
	// 参数1 ： 曲线类型
	privateKey, _ := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)


	// 散列明文
	digest := sha256.Sum256(message)
	r,s, _ := ecdsa.Sign(rand.Reader, privateKey, digest[:])


	// 模拟接收者，操作

	// 获取公钥
	pub := privateKey.PublicKey
	// 接收的数据
	recvMsg := []byte("hello world!")
	digest2 := sha256.Sum256(recvMsg)

	// 使用公钥验证 签名，得到的 散列值是否为 digest2[:]
	flag := ecdsa.Verify(&pub,digest2[:],r,s)
	if flag {
		fmt.Println("==验证通过！")
	}else {
		fmt.Println("==验证失败！")
	}

}
