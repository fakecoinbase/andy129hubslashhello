package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 生成公钥，私钥 并创建签名
func main() {

	fmt.Println("----------------------生成公钥，私钥-----------------------------")

	// 调用函数生成 公钥和私钥
	private,publicKey := newKeyPair()
	// 打印私钥 (曲线上的 x 点 (私钥的 x 点， 是由 随机种子在 椭圆曲线上生成的 私钥的 点))
	fmt.Printf("%x\n", private.D.Bytes())  // 25f9285a09cb34777c4b7cacee6e0df046e70ee91c19e5e0d83cc59219dce8dd
	// 打印公钥 (曲线上的 x 和 y 这两个点 (公钥对应的 x,y 点， 是由 私钥进行 椭圆曲线算法 得出的公钥的点))
	fmt.Printf("%x\n", publicKey)  // 9f0400042ab3cbef5c9100bc4ae902941d3f4a2e03d2d9a24d7914c35c9c6c789ec9cee341f547507f965f057e2be204506d176910e5bb5fa3de8e828fc24f77


	fmt.Println("----------------------创建签名----------------------------------")

	// 生成签名
	// 1, 将 加密内容进行 hash 运算得出  256位(32字节)的大整数
	hash := sha256.Sum256([]byte("yang to zhang 0.2BTC"))
	// fmt.Printf("hash : %x\n", hash)
	// 通过 私钥将 内容 (经过hash 处理的内容) 经过 ecdsa 签名算法 计算得出  r,s
	r,s,err := ecdsa.Sign(rand.Reader, &private, hash[:])
	if err != nil {
		fmt.Println("Sign failed : ",err)
		return
	}
	// r,s 拼接成 字节流
	signature := append(r.Bytes(), s.Bytes()...)
	fmt.Printf("%X\n", signature)


	fmt.Println("----------------------验证签名----------------------------------")
	curve := elliptic.P256()
	keylen := len(publicKey)

	x := big.Int{}
	y := big.Int{}
	x.SetBytes(publicKey[:(keylen/2)])
	y.SetBytes(publicKey[(keylen/2):])

	// rawPublic 为 ecdsa.PublicKey 结构体对象
	/*
		// PublicKey represents an ECDSA public key.
		type PublicKey struct {
			elliptic.Curve
			X, Y *big.Int
		}
	*/
	rawPublic := ecdsa.PublicKey{curve, &x,&y}
	// 通过公钥，验证 r,s 签名是否 对 hash (经过hash 运算后的内容) 有效
	ok := ecdsa.Verify(&rawPublic, hash[:], r,s )
	if ok {
		fmt.Println("验证成功!")
	}else {
		fmt.Println("验证失败！")
	}

}


// ECDSA 椭圆曲线 ： y^2 = (x^3 + a * x + b) mod p
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	// 生成椭圆曲线, secp256r1 曲线。     比特币当中的曲线是  secp256k1
	// elliptic : 椭圆
	// curve : 曲线
	curve := elliptic.P256()

	/*	crypto/ecdsa/ecdsa.go

			func GenerateKey(c elliptic.Curve, rand io.Reader) (*PrivateKey, error) {
	 */

	// private 为 PrivateKey 结构体
	/*
			// PrivateKey represents an ECDSA private key.
			type PrivateKey struct {
				PublicKey
				D *big.Int
			}
	 */
	// 其中 PublicKey 可以获取到 公钥， D 代表私钥
	private, err := ecdsa.GenerateKey(curve, rand.Reader)   // 通过 rand.Reader 随机数的种子，生成一个 椭圆曲线上的点

	if err != nil {
		fmt.Println("GenerateKey failed : ", err)
	}

	// 将 公钥 x 与 y 坐标 拼接起来 得到 公钥
	pubkey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)

	return *private, pubkey

}

