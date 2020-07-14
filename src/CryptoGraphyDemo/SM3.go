// @Title  
// @Description  
// @Author  yang  2020/7/12 16:31
// @Update  yang  2020/7/12 16:31
package main

import (
	"fmt"
	"github.com/tjfoc/gmsm/sm3"
)

// SM3
/*
	SM2: 非对称加密，基于 椭圆加密，签名速度与秘钥生成速度都快于  RSA。
	SM3：消息摘要算法，散列值为 256位。
	SM4: 分组对称加密算法，秘钥长度和分组长度均为 128 位。
*/

func main() {
	sm3Test1()

}

func sm3Test1() {
	hash := sm3.New()
	hash.Write([]byte("hello sm3"))
	result := hash.Sum(nil)
	fmt.Printf("length : %d[bit], result : %x\n", len(result)*8, result)
	// length : 256[bit], result : ce2512f4a1487ff23eab376950a7d525cba630696ababadc1136531372e6cce3


	result2 := sm3.Sm3Sum([]byte("hello sm3"))
	fmt.Printf("length : %d[bit], result : %x\n", len(result2)*8, result2)
	// length : 256[bit], result : ce2512f4a1487ff23eab376950a7d525cba630696ababadc1136531372e6cce3
}
