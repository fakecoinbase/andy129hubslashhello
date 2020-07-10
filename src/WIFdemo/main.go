package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

var base58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")


// WIF 私钥生成 (压缩或被压缩)  以及 判断是否为 WIF 私钥
func main() {

	// 8c30f4080de5cad95748c25902ae10f5f83ffe32da8b70eb679b6ccfa75c3845
	hexPrivateKey := "6edfea35d8177af24bf86289d1e052598f0c5a720fa44eb8dd432d77d80eeb47"
	wifPrivateKey := generateWIFPrivatekey(hexPrivateKey,true)
	fmt.Printf("%s\n", wifPrivateKey)

	//第1个 私钥原始信息：8c30f4080de5cad95748c25902ae10f5f83ffe32da8b70eb679b6ccfa75c3845   (用字符串保存的 16进行原始数据)
	//第2个 私钥原始信息：6edfea35d8177af24bf86289d1e052598f0c5a720fa44eb8dd432d77d80eeb47
	/*	WIF 压缩形式的私钥
		1 -->	L1vDxoQ1iZtY6EgueGmxf4muWivbuHCaEZ6E2yr4RM4M2MTpHBQK
		2 -->	KzwEgfjNVagoHjdjEUSYYomNs5q1C76NyA1bySMhygnMy9SqkxZ8

		WIF 未压缩形式的私钥
		1 -->	5Jt2aVBemm9xXe5xt3DTCoAiocKwabUzpvweqw7arUCdiAXHCmn
		2 -->   5Jf7j4nrGzUbWAfECD7EwnVBPgA4fmRnJi1bCt4LFf1mB58vtHN
	*/

	// 可以看出，WIF 压缩形式的私钥，是以  K 或 L 开头，  WIF 未压缩形式的私钥 是以  5  开头


	// 通过 WIF 私钥 获取到 私钥的原始信息
	privateKey := getPrivateKeyFromWIF(string(wifPrivateKey))
	fmt.Printf("%x\n", privateKey)   // 6edfea35d8177af24bf86289d1e052598f0c5a720fa44eb8dd432d77d80eeb47
}

// WIF : Wallet Import Format  (钱包导入格式)
// WIF 私钥 ： 钱包导入格式的 私钥，将原始私钥信息以 Base58 校验和编码显示
// WIF 私钥 又分： WIF 未压缩私钥  和  WIF 压缩私钥
/*
		压缩与未压缩：主要区别在于 ： WIF 未压缩，在原始私钥信息前面加 0x80
								   WIF 压缩，在原始私钥信息前面加 0x80,在末尾加入 0x01 (0x01 代表是 生成压缩形式公钥的私钥)

		最后，压缩与未压缩 都要经过 两次 hash, 计算出校验和，然后再进行 Base58 编码。
*/

// 生成 WIF(钱包导入格式) 私钥  (压缩或未压缩)
func generateWIFPrivatekey(hexPrivateKey string, compressed bool) []byte{
	versionStr := ""
	// 是否生成 WIF 压缩形式的私钥 (0x80 代表是 WIF 形式的私钥，0x01 代表是否为 压缩形式的 私钥)
	if compressed {
		versionStr = "80" + hexPrivateKey + "01"
	}else {
		versionStr = "80" + hexPrivateKey
	}
	// 将字符串转换为 16进制 []byte
	privateKey,_ := hex.DecodeString(versionStr)
	// 两次 hash
	firstHash := sha256.Sum256(privateKey)
	secondHash := sha256.Sum256(firstHash[:])
	// 头部4位 为校验和
	checksum := secondHash[:4]
	// 将校验和 追加到 privatekey 的尾部
	result := append(privateKey, checksum...)
	// 经过 Base58 编码得出最后的 WIF 私钥 (压缩或未压缩形式)
	wifPrivateKey := Base58Encode(result)
	return wifPrivateKey
}

// 通过 WIF 私钥获取 私钥的原始信息
func getPrivateKeyFromWIF(wifPrivateKey string) []byte{
	// 判断是否为 WIF 私钥 (无论是 压缩还是未压缩，只判断是否为 WIF )
	if checkWIF(wifPrivateKey) {
		rawData := []byte(wifPrivateKey)
		// 包含了 80，私钥，checksum
		base58DecodeData := Base58Decode(rawData)
		// 去掉头部，去掉尾部，取中间 32个字节为 私钥原始信息
		return base58DecodeData[1:33]
	}

	return nil
}

// 检查 wifPrivateKey 是否为 真正的 WIF 私钥
func checkWIF(wifPrivateKey string) bool{

	rawData := []byte(wifPrivateKey)
	// 包含了 80，私钥，checksum
	base58DecodeData := Base58Decode(rawData)
	length := len(base58DecodeData)
	if length < 37 {
		return false
	}

	// 去掉尾部4位校验和，取前面所有 (版本号，私钥(包含01或者没有))
	versionPrivateKey := base58DecodeData[:(length-4)]
	firstHash := sha256.Sum256(versionPrivateKey)
	secondHash := sha256.Sum256(firstHash[:])

	// 自己计算出来 校验和
	checksum := secondHash[:4]
	// 取出传入的 wifPrivateKey 的校验和
	originChecksum := base58DecodeData[(length-4):]
	// 两者比较，如果相同，则代表 wifPrivateKey 是有效的 WIF 私钥
	if bytes.Compare(checksum, originChecksum) == 0{
		return true
	}
	return false
}

// base58 加密
func Base58Encode(input []byte) []byte {

	// 用于存储编码后的结果
	var result []byte

	// 把字节数组 input 转化为一个 大整数
	x := big.NewInt(0).SetBytes(input)
	// 长度 58 的大整数
	base := big.NewInt(int64(len(base58Alphabet)))
	// 0 的大整数
	zero := big.NewInt(0)
	// 大整数的 指针  (用于接收 取模后的 余数)
	mod := &big.Int{}

	/*	math/big/ini.go

		// Cmp compares x and y and returns:
		//
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		//
	*/
	// x 与  zero 进行比较  (x, zero 都是 *Int 类型)
	// x
	for x.Cmp(zero) !=0 {
		x.DivMod(x,base,mod)  // x 对 base(也就是 58) 取余, x 保存 整数，mod 保存余数， 直到x 被除尽为 0
		result = append(result, base58Alphabet[mod.Int64()])   // 根据余数到 base58 编码中找到对应的 编码字符
	}

	// 举例: (二进制转换， base58 其实是 4 % 58 的操作，下面以 二进制举例)
	/*	10进制 4 , 转换为 二进制

		4  % 2  == 2  ...    0      append 0  -->  result 0
		2  % 2  == 1  ...    0      append 0  -->  result 00
		1  % 2  == 0  ...    1      append 1  -->  result 001

		so  result ==  001
		但是 4 的二进制应该是  100,  所以我们需要进行一次反转。
	*/

	// 反转 []byte
	ReverseBytes(result)

	// 假如 []byte 里面 前面几个 都为 0 , 则替换为 1  (比特币中特殊的写法)
	for _, b := range input {
		// fmt.Println("b : ",b)
		// 前面 0x00 代表 比特币主网 版本 , input 是[]byte 类型，里面存放的是  256位数据，
		// b 为一个字节(8位) 在 256位长度的数据中 单个长度代表4位，两个长度就代表8位一个字节，所以 可以和 0x00 比较。
		if b == 0x00 {   // 则将其追加到 result 前面，用 base58Alphabet[0] 也就是 '1' 代替
			// 如果 []byte
			result = append([]byte{base58Alphabet[0]}, result...)
		}else {
			break
		}
	}

	return result
}

// 反转  []byte
func ReverseBytes(data []byte) {
	for i,j := 0, len(data)-1; i<j; i,j = i+1,j-1 {
		data[i],data[j] = data[j],data[i]
	}
}

// base58 解码
func Base58Decode(input []byte) []byte{
	result := big.NewInt(0)
	zeroBytes := 0
	for _, b := range input {
		// 只判断前方有多少个 1
		if b == base58Alphabet[0] {
			zeroBytes++
		}else{
			break
		}
	}

	payload := input[zeroBytes:]
	for _,b := range payload {
		charIndex := bytes.IndexByte(base58Alphabet, b)   // 反推出 余数
		result.Mul(result, big.NewInt(58))    // 之前的结果乘以  58，  结果用 result 保存
		result.Add(result, big.NewInt(int64(charIndex)))   // 在加上这个余数

		/*	示例，以二进制示例


		 */

		// fmt.Printf("result : %v\n", result.Int64())
	}

	// 将 result 转换为 字节数组
	decoded := result.Bytes()

	// fmt.Println("decoded1 : ", string(decoded))

	decoded = append(bytes.Repeat([]byte{0x00}, zeroBytes), decoded...)

	// fmt.Println("decoded2 : ", string(decoded))

	return decoded
}

