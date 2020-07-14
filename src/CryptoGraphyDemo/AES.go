// @Title  
// @Description  
// @Author  yang  2020/7/11 16:13
// @Update  yang  2020/7/11 16:13
package main

import (
	"CryptoGraphyDemo/utils"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

// AES  资料参考：https://blog.csdn.net/qq_28205153/article/details/55798628

// AES 加密
// src: 待加密的明文， key : 秘钥
func EncryptAES(src, key []byte) []byte {
	block,err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	src = utils.PaddingText(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

// AES 解密
// src: 待解密的密文， key : 秘钥
func DecryptAES(src, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	dst = utils.UnPaddingText(dst)
	return dst
}

func main() {
	src := []byte("这是一条需要被AES加解密的明文")

	key := []byte("87654321ABCDEFGH")   // key 必须为 16,24,或 32 个字节的长度

	encryptMsg := EncryptAES(src, key)

	fmt.Printf("加密后[16进制]：%x\n", encryptMsg)

	decryptMsg := DecryptAES(encryptMsg, key)
	fmt.Printf("解密后：%s\n", decryptMsg)
}
