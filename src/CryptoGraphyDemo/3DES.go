// @Title  
// @Description  
// @Author  yang  2020/7/11 11:55
// @Update  yang  2020/7/11 11:55
package main

import (
	"CryptoGraphyDemo/utils"
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
	"io/ioutil"
	"strings"
)

const KEY_SIZE = 24

// 3DES
/*
	1, 三重 des 为了增强 DES 的强度，将 des 重复 3次得到的一种加密算法
	2, 3DES 加密机制:
			a, 明文经过三次 DES 处理才能变成最后的密文，由于 DES 秘钥的长度实质上是  56比特，因此三重 DES 秘钥长度是 56*3 = 168， 表面上看是 64*3 = 192
			b, 三重DES 并不是进行三次加密，而是 加密 -> 解密 -> 加密 的过程，这种设计是为了 3DES 能够兼容普通的 DES 。
				当三重 DES 所有的秘钥都相同时，三重 DES 也就是等于 普通的 DES，因此 DES加密的密文也就可以使用 三重DES 来进行解密。
			c, 密码算法不能依靠 算法的不公开性来保证 密码算法的安全性， 反而需要公开算法思想，让大家都去破解，如果大家此时都破解不了，这才是安全的密码算法。

	3, 3DES 解密机制

 */

// 3DES 加密
// src: 待加密的明文， key : 秘钥 (key 必须为 24 个字节长度)
func Encrypt3DES(src, key []byte) []byte {
	block, err := des.NewTripleDESCipher(key)   // 注意与 单重 DES 的不同
	if err != nil {
		panic(err)
	}

	src = utils.PaddingText(src, block.BlockSize())
	fmt.Println("--blockSize : ", block.BlockSize())

	// 使用 key 中 block.BlockSize() 个元素作为 初始化向量, 注意 key 的长度万一不够 block.BlockSize() 怎么办？？
	blockMode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(src))

	blockMode.CryptBlocks(dst,src)
	return dst
}

// 3DES 解密
// src: 待解密的密文，  key: 秘钥
func Decrypt3DES(src, key []byte) []byte{
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		panic(err)
	}
	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	dst := make([]byte, len(src))
	blockMode.CryptBlocks(dst, src)
	dst = utils.UnPaddingText(dst)
	return dst
}

// 将用户输入的key, 转换为 24 个字节长度的 key
// key : 用户指定的密钥
func genKey3DES(key []byte) []byte {
	// 用于存储最终密钥
	kkey := make([]byte, 0, KEY_SIZE)
	// 获取原始密钥的长度
	length := len(key)
	if length > KEY_SIZE {
		kkey = append(kkey, key[:KEY_SIZE]...)
	}else {
		// 用指定的长度对实际密钥长度进行求商
		div := KEY_SIZE / length  // 代表要填充 几个  length 长度
		// 用指定的长度对实际密钥长度进行求余
		mod := KEY_SIZE % length  // 代表最后要填充 length 里面的几个元素

		for i:=0;i<div;i++ {
			kkey = append(kkey, key...)
		}
		kkey = append(kkey, key[:mod]...)
		// 最终将 kkey 里面填充至 KEY_SIZE 长度的数据
	}
	return kkey
}


func main() {
	src := []byte("这是一条需要被3DES加解密的明文")

	key := []byte("12345678abcdefgh87654321")   // 由于 3DES 是三重 DES 加密，需要指定 三倍于 DES 秘钥的长度，我们这里指定 key 为 8*3 = 24 个字节

	encryptMsg := Encrypt3DES(src, key)
	fmt.Printf("加密后[16进制]：%x\n", encryptMsg)

	decryptMsg := Decrypt3DES(encryptMsg, key)
	fmt.Printf("解密后：%s\n", decryptMsg)



	fmt.Println("------------------3DES 加密解密文件---------------------------")

	testEncryptFile()

}
// 通过 3DES 对文件进行 加密，解密
func testEncryptFile() {
	var command string
	var filename string
	fmt.Print("请输入命令(加密 or 解密) : ")
	fmt.Scanln(&command)
	if command == "加密" {
		fmt.Print("请输入被加密的文件路径：")
		fmt.Scanln(&filename)
		fmt.Print("请输入密钥：")
		var password string
		fmt.Scanln(&password)
		fmt.Print("请再次输入密钥：")
		var confirmPwd string
		fmt.Scanln(&confirmPwd)
		if !bytes.Equal([]byte(password), []byte(confirmPwd)) {
			fmt.Println("两次输入的密码不一致，请重新输入!")
		}else {
			key := genKey3DES([]byte(password))
			info, err := ioutil.ReadFile(filename)
			if err != nil {
				panic(err)
			}
			fmt.Println("正在进行加密......")
			encryptMsg := Encrypt3DES(info,key)
			index := strings.LastIndex(filename,".")
			newFileName := filename[:index] + "_encrypted" + filename[index:]
			ioutil.WriteFile(newFileName,encryptMsg, 0777)
			fmt.Println("已生成加密文件 : "+newFileName+", 请妥善保管您的密钥！")
		}
	}else if command == "解密" {
		fmt.Print("请输入需要解密的文件路径：")
		fmt.Scanln(&filename)
		fmt.Print("请输入密钥：")
		var password string
		fmt.Scanln(&password)
		key := genKey3DES([]byte(password))
		info,err := ioutil.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		fmt.Println("正在进行解密......")
		decryptMsg := Decrypt3DES(info, key)
		if len(decryptMsg) == 0{
			fmt.Println("密钥不对，请重新输入！")
		}else {
			index := strings.LastIndex(filename, ".")
			newFileName := filename[:index]+"_decrypted"+filename[index:]
			ioutil.WriteFile(newFileName, decryptMsg, 0777)
			fmt.Println("已生成解密文件"+newFileName+", 请妥善保管您的密钥！")
		}
	}
}
