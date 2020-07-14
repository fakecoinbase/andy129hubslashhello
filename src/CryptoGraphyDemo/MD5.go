// @Title  
// @Description  
// @Author  yang  2020/7/12 12:10
// @Update  yang  2020/7/12 12:10
package main

import (
	"crypto/md5"
	"fmt"
)

// MD5


func main() {
	hash := md5.New()
	hash.Write([]byte("hello md5"))

	// Sum 求hash 函数
	result := hash.Sum(nil)
	fmt.Printf("%x\n", result)  // 741fc6b1878e208346359af502dd11c5

	result2 := hash.Sum([]byte("123"))
	fmt.Printf("%x\n", result2)   // 313233741fc6b1878e208346359af502dd11c5

	/*   313233741fc6b1878e208346359af502dd11c5
		拆分为：313233 和  741fc6b1878e208346359af502dd11c5
		这样就明白了 hash.Sum(nil) 与 hash.Sum([]byte("123")) 的区别了

	 */

	// 重置 hash, 重新计算下一条信息
	hash.Reset()
	hash.Write([]byte("hello reset"))
	result = hash.Sum(nil)
	fmt.Printf("%x\n", result)   // a417bf3c48954a61d88458839e3e49b5

	// 代码简洁版：
	res := md5.Sum([]byte("hello reset"))
	fmt.Printf("%x\n",res[:])   // a417bf3c48954a61d88458839e3e49b5
}
