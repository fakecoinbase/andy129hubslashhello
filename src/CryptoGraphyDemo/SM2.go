// @Title  
// @Description  
// @Author  yang  2020/7/12 15:31
// @Update  yang  2020/7/12 15:31
package main

import (
	"crypto/rand"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
	"io/ioutil"
)

// SM2 国密算法
/*
	SM2: 非对称加密，基于 椭圆加密，签名速度与秘钥生成速度都快于  RSA。
	SM3：消息摘要算法，散列值为 256位。
	SM4: 分组对称加密算法，秘钥长度和分组长度均为 128 位。
 */


func main() {

	// sm2Test1()   // sm2 加解密 示例

	// sm2Test2()   // 将密钥对写入文件 示例

	sm2Test3()      // 从文件中读取秘钥对  示例
}
// sm2 加解密 示例
func sm2Test1() {
	privateKey, err := sm2.GenerateKey()
	if err != nil {
		fmt.Println("秘钥对生成失败！")
		return
	}
	// 从私钥中获取公钥
	publicKey := &privateKey.PublicKey

	// 消息
	msg := []byte("hello sm2")
	// 公钥加密
	encryptMsg, err := publicKey.Encrypt(msg)
	if err != nil {
		fmt.Println("加密失败")
		return
	}else {
		fmt.Printf("encryptMsg : %x\n", encryptMsg)
		/*
				30710220305fd745c70c40a0ff7b5ab9b3ad81c390fb43a14085201b0907bb271e835d6b02200900178595bd12aa5949f7f9e1db48edb39b721eccf1bebaf8766281acf850f1042007ed8c70df273fdd87f7f5a38676a0aa988eec4fbad554c4d6814e55b651e1da0
			40956d425f51232e542ad
		*/
	}

	// 私钥解密
	decryptMsg, err := privateKey.Decrypt(encryptMsg)
	if err != nil {
		fmt.Println("解密失败！")
	}else {
		fmt.Printf("decryptMsg : %s\n", decryptMsg)  // hello sm2
	}
}

// 将密钥对写入文件 示例
func sm2Test2() {

	err := WriteKeyPairToFile("private.pem","public.pem", []byte("12345678"))
	if err != nil {
		fmt.Println("密钥对保存文件失败！")
	}else {
		fmt.Println("密钥对保存文件成功！")
	}
}
// 从文件中读取秘钥对  示例
func sm2Test3() {
	privateKey, publicKey, err := ReadKeyPairFromFile("private.pem", "public.pem", []byte("12345678"))
	if err != nil {
		fmt.Println("密钥对读取失败！")
		return
	}

	// 读取待加密的文件
	info, err := ioutil.ReadFile("G:/node.txt")
	if err != nil {
		fmt.Println("ReadFile failed, err : ", err)
		return
	}
	// 签名
	sigMsg, err := privateKey.Sign(rand.Reader, info, nil)
	if err != nil {
		fmt.Println("签名失败！")
		return
	}

	flag := publicKey.Verify(info,sigMsg)
	if flag {
		fmt.Println("验证成功!")
	}else {
		fmt.Println("验证失败！")
	}
}


/*
	生成公钥私钥并写入文件
	privateKeyPath: 私钥路径
	publicKeyPath: 公钥路径
	password: 用于加密私钥
 */
func WriteKeyPairToFile(privateKeyPath, publicKeyPath string, password []byte) error {

	// 生成秘钥对
	privateKey, err := sm2.GenerateKey()
	if err != nil {
		fmt.Println("秘钥对生成失败！")
		return err
	}

	// 私钥写入文件
	flag,err2 := sm2.WritePrivateKeytoPem(privateKeyPath, privateKey, password)
	if !flag {
		return err2
	}

	// 获取公钥
	publicKey := privateKey.Public().(*sm2.PublicKey) // 类型转换
	flag2, err3 := sm2.WritePublicKeytoPem(publicKeyPath, publicKey, nil)
	if !flag2 {
		return err3
	}

	return nil
}

func ReadKeyPairFromFile(privateKeyPath, publicKeyPath string, password []byte) (*sm2.PrivateKey, *sm2.PublicKey, error ){
	privateKey, err := sm2.ReadPrivateKeyFromPem(privateKeyPath, password)
	if err != nil {
		return nil,nil, err
	}

	publicKey, err2 := sm2.ReadPublicKeyFromPem(publicKeyPath, password)
	if err2 != nil {
		return nil,nil, err2
	}

	return privateKey,publicKey, nil
}
