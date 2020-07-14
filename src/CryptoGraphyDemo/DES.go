// @Title  
// @Description  
// @Author  yang  2020/7/10 18:19
// @Update  yang  2020/7/10 18:19
package main

import (
	"CryptoGraphyDemo/utils"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

// DES 介绍
/*
	DES 是一种将 64比特的明文加密成 64比特的密文的对称加密算法，它的秘钥长度是 56比特 DES 的私钥长度，实际上是 64比特，
	这是由于每个每隔7 位会设置一个错误检查的比特，因此实质上秘钥长度是 56比特。
	DES 以64比特的明文为一个单位进行加密，每64比特为一个分组。
	DES 每次只能加密64比特的明文，如果加密的明文长度比较长，就需要对 DES 进行迭代，而迭代的具体方案就称为 模式(加密模式)
		* 模式：
			CBC 模式：
				CBC 模式的全称是 Cipher Block Chaining (密文分组链接模式),
					之所以叫 密文分组链接模式，是因为 密文分组像链条一样互相链接在一起。
				CBC 模式加密原理：
					1，初始化向量 与 明文分组1 进行异或操作，然后将结果进行加密，得到 密文分组1.
					2，密文分组1 与 明文分组2 进行异或操作，然后将结果进行加密，得到 密文分组2.
					3，密文分组2 与 明文分组3 进行异或操作，然后将结果进行加密，得到 密文分组3.

							。。。 以此类推

                CBC 模式解密原理：
					1，密文分组3 进行解密，然后与 密文分组2 进行异或操作，得到 明文分组3.
					2，密文分组2 进行解密，然后与 密文分组1 进行异或操作，得到 明文分组2.
					3，密文分组1 进行解密，然后与 初始化向量 进行异或操作，得到 明文分组1.

							。。。 以此类推，最后一个密文与 初始化向量进行异或。


				初始化向量 详解(也必须为 8个字节长度, 准确的来说 长度要等于 block.BlockSize())：
					当加密第一个明文分组时， 由于不存在前一个密文分组，因此需要事先准备一个长度为 一个分组的比特序列来代替 前一个密文分组
						这个比特序列就称为 初始化向量，即 iv

					分析：
						1，假设 CBC模式加密的密文分组中有一个分组损坏了(由于硬盘故障导致密文分组的值发生了变化)
							只要密文分组的长度没有发生变化，则解密是最多只会有2个分组收到数据损坏的影响。
						2，假设 CBC模式的密文分组中长度发生了变化，会导致每一个分组会向后一个 分组借位，
						（保证每一个分组的长度是固定的，所以一旦某个密文长度发生了变化，则会从其他密文分组中借位，就导致后续的密文分组数据都发生了变化），
							这样就导致了后面的分组全部发生变化，最终导致后面的数据全部解密失败。


			ECB 模式：
				ECB 模式全称：Electronic CodeBook , 也称之为 电子密码本模式。

				ECB 模式加密原理：
					1, 明文分组1 进行加密 得到密文分组1。
					2，明文分组2 进行加密 得到密文分组2.
					3，明文分组3 进行加密 得到密文分组3.

			CFB 模式：
				CFB 模式全称：Cipher FeedBlock (密文反馈模式), 前一个密文分组会被送回到密码算法的输入端，这就是反馈。

				CFB 模式加密原理:
					1 ,初始化向量IV进行加密,然后与明文分组1进行异或操作,得到密文分组1.
					2 ,密文分组1进行加密,然后与明文分组2进行异或操作,得到密文分组2.
					3 ,密文分组2进行加密,然后与明文分组3进行异或操作,得到密文分组3.
								...... 以此类推

			OFB 模式：
				OFB 模式的全称：Output FeedBack (输出反馈模式)， 密码算法的输出会反馈到 密码算法的输入中。

				OFB 模式加密原理：
					1 ,初始化向量IV进行加密,将结果与明文分组1进行异或,得到密文分组1.
					2 ,将初始化向量V加密后的结果再进行加密,将结果与文分组2进行异或,得到密文分组2.
					3 ,在将加密后的初始化向量再进行加密,将结果与文分组3进行异或,得到文分组3.
								...... 以此类推

				OFB 模式解密原理：
					1,先將初始化向星IV加密,然后结果与密文分组1进行异或，得到明文分组1.
					2,将加密后的初始化向量再加密,然后结果与密文分组2进行异或,得到明文分组2.
					3,将加密后的初始化向量再加密,然后结果与密文分组3进行异或,得到明文分组3.
								...... 以此类推

	DES 每个加密过程称为一轮， 总共加密过程有 16轮。
 */

// 使用 DES 算法进行加密(CBC模式)
// src: 待加密的明文， key : 秘钥  （key 必须为 8个字节的切片）
func EncryptDES_CBC(src, key []byte) []byte{
	// 创建 cipher.Block 接口，其对应的就是一个加密块儿 (也可以称为一个 分组)
	block, err := des.NewCipher(key)

	if err != nil {
		panic(err)
	}
	// 获取每个块的大小 (默认length 为 8，代表  64位)
	length := block.BlockSize()

	// fmt.Printf("block length : %d\n", length)   // 8

	// 对最后一组明文进行填充, 将src 填充为 每组都为 length 长度的标准。
	src = utils.PaddingText(src, length)
	// 初始化向量
	iv := []byte("12345678")   // 必须为 8个字节长度 , 准确的来说 长度要等于 block.BlockSize()
	// 创建CBC 加密模式
	blockMode := cipher.NewCBCEncrypter(block, iv)
	// 创建切片用于存储加密之后的 数据
	dst := make([]byte,len(src))
	blockMode.CryptBlocks(dst, src)
	return dst
}

// 使用 DES 算法进行解密 (CBC模式)
// src: 待解密的密文，  key: 秘钥
func DecryptDES_CBC(src, key[]byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}
	// 初始化向量
	iv := []byte("12345678")
	// 创建CBC 解密模式
	blockMode := cipher.NewCBCDecrypter(block,iv)
	dst := make([]byte,len(src))
	blockMode.CryptBlocks(dst, src)

	// 最后需要将解密的 数据，进行去除尾部填充的操作，从而还原为 原始数据
	return utils.UnPaddingText(dst)
}

// 使用 DES 算法进行加密(ECB模式)
// src: 待加密的明文，  key: 秘钥
func EncryptDES_ECB(src, key[]byte) []byte {
	// 创建 cipher.Block 接口，其对应的就是一个加密块儿 (也可以称为一个 分组)
	block, err := des.NewCipher(key)

	if err != nil {
		panic(err)
	}
	// 获取每个块的大小 (默认length 为 8，代表  64位)
	length := block.BlockSize()
	// fmt.Printf("block length : %d\n", length)   // 8

	// 对最后一组明文进行填充, 将src 填充为 每组都为 length 长度的标准。
	src = utils.ZeroPadding(src, length)

	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		// 每次加密 8 字节数据，将结果放置在 dst 切片中
		block.Encrypt(dst, src[:length])
		// 去除已被加密的数据, 获取未被加密的数据
		src = src[length:]
		// 获取未被放置 加密数据的 切片空间
		dst = dst[length:]

		/*
			可以理解为：src 与 dst 的空间长度是一样的(src为元素数据， dst 为空的切片，用于存放结果)，  每次加密 src 的 8 个字节，将结果放置在 dst 中，一次暂用 8个字节
			当 第二次轮循的时候， src 往后移 8个字节长度，进行下一轮数据加密，然后将结果 放置在 dst 中 (dst 在之前需要移动 8个字节的空间位置 用于放置结果)

			最后，由于 dst := out , 所以 out 与 dst 公用同一块空间，所以 最后返回的 out 就代表了 所有加密后的数据。
		 */
	}

	return out
}

// 使用 DES 算法进行解密(ECB模式)
// src: 待解密的密文，  key: 秘钥
func DecryptDES_ECB(src, key[]byte) []byte {
	block, err := des.NewCipher(key)
	if err != nil {
		panic(err)
	}

	length := block.BlockSize()

	out := make([]byte, len(src))
	dst := out

	for len(src) > 0 {
		block.Decrypt(dst, src[:length])
		src = src[length:]
		dst = dst[length:]
	}

	out = utils.ZeroUnPadding(out)
	return out
}

func main() {


	fmt.Println("--------------------DES CBC模式加密------------------------------------")
	src := []byte("这是一条需要被DES加解密的明文")

	key := []byte("87654321")   // key 必须为 8个字节的长度
	// DES 加密
	encryptMsg := EncryptDES_CBC(src, key)
	fmt.Println("加密后[二进制]：",encryptMsg)                 // [151 171 79 56 20 38 85 174 136 3 149 195 250 245 172 78]
	fmt.Printf("加密后[16进制]：%x\n", encryptMsg)   // 以16进行展示数据：97ab4f38142655ae880395c3faf5ac4e

	// DES 解密
	decryptMsg := DecryptDES_CBC(encryptMsg, key)
	fmt.Printf("解密后：%s\n", decryptMsg)   // 解密后：这是一条需要被DES加解密的明文

	fmt.Println("--------------------DES ECB模式加密------------------------------------")
	src2 := []byte("这是一条需要被DES加解密的明文-ECB模式0")
	encryptMsg2 := EncryptDES_ECB(src2, key)
	fmt.Println("加密后[二进制]：",encryptMsg2)

	decryptMsg2 := DecryptDES_ECB(encryptMsg2, key)
	fmt.Printf("解密后：%s\n", decryptMsg2)

}
