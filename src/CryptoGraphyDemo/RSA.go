// @Title  
// @Description  
// @Author  yang  2020/7/11 17:38
// @Update  yang  2020/7/11 17:38
package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// 非对称加密体系
/*
	1, 简介
		非对称加密也叫做 公钥密码，在公钥秘钥中， 秘钥分为 加密秘钥 和 解密秘钥， 发送至使用加密秘钥对消息进行 加密。
			接收者使用解密秘钥对消息进行解密， 加密秘钥也称之为公钥， 解密秘钥也称之为 私钥。
			使用某一个人的公钥进行加密，那么必须使用这个的私钥才能进行解密，公钥和私钥是配对的。
			a, 发送者只需要公钥
			b, 接收者只需要解密秘钥。
			c, 解密秘钥不可以公开
			d, 加密秘钥可以任意公开。

	2， 通信流程图 见文件夹
	3， 公钥密码存在的问题
		a, 公钥密码解决了秘钥配送问题，但是我们还要判断所得到的公钥是否正确，这个问题被称之为 公钥认证问题。
		b, 公钥密码处理速度很慢，只有对称密码的几百分之一。
	4， 数据完整性
		1，简介
			目前数据的完整性的解决方案主要采用单向散列函数和加密算法。 单向散列函数能够将一段信息 映射成一段小的信息码，
				并且不同文件散列成相同信息的概率很低。通常我们将原始信息使用单向散列函数处理成一段信息码，然后对其加密，
				和文件一起保存。
		2，单相散列函数
			单向散列函数有一个输入和一个输出，其中输入称为消息，输出称为散列值，单向散列函数也称之为 摘要函数， 哈希函数，杂凑函数。

		3，单向散列函数的性质
			a, 根据任意长度的消息计算出固定长度的散列值。
			b, 能够快速计算出散列值
 			c, 具备单向性

		4, 单向散列函数的实际应用
			a, 检测软件是否被篡改。
					图解，详见 非对称加密
			b, 数字签名
				详见示例：MD5.go，SHA256.go
		5，单向散列函数的缺点
			单向散列函数能够辨别出篡改，但是无法辨别出 真伪。(需要与 数字签名 一起使用，能达到 数据的完整性，真伪性)

	5， 数字签名
		1, 简介 （详见非对称加密-- 数字签名介绍1）
			数字签名可以识别篡改 和 伪装， 防止否认，  在数字签名中有两种行为：
				a, 生成消息签名
				b, 验证消息签名
		2，公钥密码和数字签名
			详见 非对称加密 -- 数字签名流程图
			用私钥进行加密 (数字签名)， 用公钥验证签名的真伪。

		3, 数字签名示例 (两种方法)
			a, 直接对消息进行签名 (详见  非对称加密--数字签名示例1)
				详见 DSAsign.go
			b, 对消息的散列值进行签名  (详见  非对称加密--数字签名示例2)
				方法a 直接对消息签名的方法， 这种方法需要对整个消息进行加密，非常耗时，所以我们可以求出消息的散列码在对散列值加密，
					无论消息有多长，散列值都是固定的，因此对散列值加密要快很多。

				详见 RSAsign.go

				使用 曲线算法 进行签名
				详见 ECCsign.go
 */

// bits 代表生成的私钥 有多少位
func RsaGenKey(bits int) error {
	// GenerateKey 函数: 使用随机数生成器生成一对指定长度的公钥和私钥
	// rand.Reader 是一个全局，共享的密码随机生成器
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic(err)
	}

	// x509 是通用的证书格式: 序列号，签名算法，颁发者，有效时间，持有者，公钥 等信息
	// PKCS: RSA 实验室与其他安全系统开发商为 促进公钥密码的发展而指定的一系列标准
	priStream := x509.MarshalPKCS1PrivateKey(privateKey)
	// 将私钥字符串设置 pem 格式的块儿中
	// pem 是一种整数或私钥的格式：
	/*
			-----------BEGIN RSA Private Key-----------


			-----------END RSA Private Key------------
	 */
	block := pem.Block{
		Type:    "RSA Private Key",
		Bytes:   priStream,
	}

	privFile,err := os.Create("private.pem")
	if err != nil {
		panic(err)
	}

	defer privFile.Close()
	// 将块编码到文件
	err = pem.Encode(privFile, &block)
	if err != nil {
		panic(err)
	}


	// 从私钥中获取公钥
	pubKey := privateKey.PublicKey
	// 将公钥序列化
	pubStream := x509.MarshalPKCS1PublicKey(&pubKey)
	// 将公钥字符串设置 pem 格式的块儿中
	block = pem.Block{
		Type:    "RSA Public Key",
		Bytes:   pubStream,
	}

	pubFile,err := os.Create("public.pem")
	if err != nil {
		panic(err)
	}
	defer pubFile.Close()

	err = pem.Encode(pubFile, &block)
	if err != nil {
		panic(err)
	}

	return nil
}

// RSA 使用公钥加密
// src : 待加密的明文，  pathName: 存储公钥的文件名
func EncryptRsaPublic(src, pathName []byte) ([]byte,error) {
	file,err := os.Open(string(pathName))
	msg := []byte("")
	if err != nil {
		return msg, err
	}
	defer  file.Close()
	info,err := file.Stat()
	if err != nil {
		return msg,err
	}
	// 创建切片，用于存储公钥
	recvBuf := make([]byte, info.Size())
	// 读取公钥
	file.Read(recvBuf)
	// 将得到的公钥反序列化
	// 返回值：参数1: 存储公钥的切片， 参数2：剩余未解密的数据
	block, _ := pem.Decode(recvBuf)
	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return msg, err
	}

	msg, err = rsa.EncryptPKCS1v15(rand.Reader, pubKey, src)
	if err != nil {
		return msg, err
	}
	return msg, nil
}
// RSA 使用私钥 解密
// src : 待解密的密文，  pathName: 存储私钥的文件名
func DecryptRsaPrivate(src, pathName []byte) ([]byte,error) {
	file,err := os.Open(string(pathName))
	msg := []byte("")
	if err != nil {
		return msg, err
	}
	defer  file.Close()
	info,err := file.Stat()
	if err != nil {
		return msg,err
	}
	// 创建切片，用于存储私钥
	recvBuf := make([]byte, info.Size())
	// 读取私钥
	file.Read(recvBuf)

	// 将得到的私钥反序列化
	// 返回值：参数1: 存储私钥的切片， 参数2：剩余未解密的数据
	block, _ := pem.Decode(recvBuf)
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return msg, err
	}

	msg, err = rsa.DecryptPKCS1v15(rand.Reader, privKey,src)
	if err != nil {
		return msg, err
	}
	return msg, nil
}

func main() {

	// 生成公钥私钥  (保存在本地：private.pem, public.pem)
	/*
	err := RsaGenKey(1024)
	if err != nil {
		fmt.Println("RsaGenKey err : ", err)
		return
	}
	fmt.Println("私钥公钥已生成！")
	*/

	src := []byte("这是一条需要被RSA 加解密的明文")
	pubKeyPath := "public.pem"
	privKeyPath := "private.pem"

	// 公钥加密
	encryptMsg,_ := EncryptRsaPublic(src, []byte(pubKeyPath))
	fmt.Printf("encryptMsg : %x\n", encryptMsg)

	// 私钥解密
	decryptMsg, _:= DecryptRsaPrivate(encryptMsg, []byte(privKeyPath))
	fmt.Println("decryptMsg : ", string(decryptMsg))

}
