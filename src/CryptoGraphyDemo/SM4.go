// @Title  
// @Description  
// @Author  yang  2020/7/12 16:37
// @Update  yang  2020/7/12 16:37
package main

import (
	"CryptoGraphyDemo/utils"
	"crypto/cipher"
	"fmt"
	"github.com/tjfoc/gmsm/sm4"
)

// SM4
/*
	SM2: 非对称加密，基于 椭圆加密，签名速度与秘钥生成速度都快于  RSA。
	SM3：消息摘要算法，散列值为 256位。
	SM4: 分组对称加密算法，秘钥长度和分组长度均为 128 位。
*/

// SM4 加密
// src：待加密明文， key : 秘钥
func EncryptSM4(src, key []byte) []byte {
	block, err := sm4.NewCipher(key)
	if err != nil {
		panic(err)
	}

	fmt.Println("sm4 blockSize : ", block.BlockSize())   // 16

	src = utils.PaddingText(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst,src)
	return dst
}
// SM4 解密
// src：待解密的密文， key : 秘钥
func DecryptSM4(src, key []byte) []byte{
	block, err := sm4.NewCipher(key)
	if err != nil {
		panic(err)
	}

	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])   // 将key 中取出 blockSize 个长度的元素作为 初始化向量
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst,src)
	return utils.UnPaddingText(dst)
}


func main() {
	src := []byte("这是一条需要被 SM4 加解密的明文")

	key := []byte("87654321abcdefgh")   // key 必须为 16个字节的长度

	encryptMsg := EncryptSM4(src, key)

	decryptMsg := DecryptSM4(encryptMsg, key)
	fmt.Println("解密后：", string(decryptMsg))


}
