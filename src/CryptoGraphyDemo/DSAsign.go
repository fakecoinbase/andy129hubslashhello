// @Title  
// @Description  
// @Author  yang  2020/7/12 14:38
// @Update  yang  2020/7/12 14:38
package main

import (
	"crypto/dsa"
	"crypto/rand"
	"fmt"
)

// DSA sign
/*
	验证签名的作用：
		1, 保证数据的完整性
		2, 确保数据的来源
 */
func main() {

	// Parameters
	var param dsa.Parameters
	// GenerateParameters 函数 随机设置合法的参数到 param
	// dsa.L1024N160 , 根据第三个参数就决定 L 和 N 的长度，长度越长，加密强度越高
	dsa.GenerateParameters(&param, rand.Reader, dsa.L1024N160)

	// 生成私钥
	var priv dsa.PrivateKey
	priv.Parameters = param
	dsa.GenerateKey(&priv, rand.Reader)

	// 利用私钥签名数据
	message := []byte("hello world!")
	r,s, _ := dsa.Sign(rand.Reader, &priv, message)

	// 通过私钥获取公钥
	pub := priv.PublicKey

	recvMsg := []byte("hello world!")
	// 使用公钥验证签名
	if dsa.Verify(&pub, recvMsg, r,s) {
		fmt.Println("验证通过！")
	}else {
		fmt.Println("验证失败！")
	}


}
