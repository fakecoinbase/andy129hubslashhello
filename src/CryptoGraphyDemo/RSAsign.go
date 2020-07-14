// @Title  
// @Description  
// @Author  yang  2020/7/12 15:08
// @Update  yang  2020/7/12 15:08
package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

// RSA sign
/*

 */
func main() {

	// 生成秘钥对
	priv, _:= rsa.GenerateKey(rand.Reader, 1024)
	// 消息
	msg := []byte("hello world!")
	// 对消息进行 散列处理
	h := md5.New()
	h.Write(msg)
	hashed := h.Sum(nil)

	opts := &rsa.PSSOptions{
		SaltLength: rsa.PSSSaltLengthAuto,
		Hash:       crypto.MD5,
	}
	// 签名
	sig, _ := rsa.SignPSS(rand.Reader, priv, crypto.MD5, hashed, opts)

	// 获取公钥
	pub := priv.PublicKey
	err := rsa.VerifyPSS(&pub, crypto.MD5, hashed, sig, opts)
	if err == nil {
		fmt.Println("--验证通过！")
	}else {
		fmt.Println("验证失败！")
	}


}
