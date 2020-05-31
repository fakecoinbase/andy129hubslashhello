package main

import (
	"fmt"
	"math"
)

// RSA 加密算法
// 非对称加密  (公钥 公开， 私钥加密)


func main() {

	// 第一步：加密算法过程
	/*
		RSA 加密算法过程
		1, 随机选取两个质数 p 和 q      (示例：53,59)
		2, 计算 n = pq                (示例：n = 53*59 == 3127)
		3, 选取一个与 φ（n）互质的小奇数e, φ（n）= (p-1)(q-1)      (示例：φ（n）= (53-1)*(59-1) == 3016,  与 φ（n）互质的小奇数 e ==  3  (不会被3 整除) )
		4, 对模 φ（n），计算 e 的乘法逆元 d, 即满足 (e*d) mod φ（n）= 1    (示例：(3*d)%3016 == 1,  求得 d == 2011)
		5, 公钥 (e,n),  私钥 (d,n)     (示例：公钥 (e,n) == (3,3127),  私钥 (d,n): (2011, 3127))
	*/

	// 第二步：加密解密公式
	/*
		加密解密的公式： m 代表 文本信息, m^e 代表 m的e 次方  (在go语言中， ^ 不再用于次方，而是表示“按位异或的运算”， 所以我们使用 math.Pow(m,e))
		加密过程： c = (m^e) % n
		解密过程： m = (c^d) % n
	 */


	// fmt.Println(math.Pow(2,3))   // math.Pow(2,3)  计算 2 的3次方  (注意：传入参数以及 返回值 都是 float64 类型)

	//// 示例：
	//e := 3
	//d := 2011
	//n := 3127

	// 失败，超过了 类型存储的最大值
	m := 35
	c := int64(math.Pow(float64(m),3)) % 3127
	msg := int64(math.Pow(float64(c),2011)) % 3127
	fmt.Println(msg)


	//自己用 go 语言实现有困难，超过了 实际存储大小 (可参考 go 内置包 "crypto/rsa")
	// 参考资料： https://blog.csdn.net/chenxing1230/article/details/83757638
	xx := int64(math.Pow(8,2011))   // -9223372036854775808
	yy := math.Pow(8,2011)  //  +Inf
	fmt.Println(xx)   //
	fmt.Println(yy)


}
