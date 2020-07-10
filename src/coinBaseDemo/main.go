package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
)

// 创币交易： coinbase
// 区块中的第一笔交易 (奖励给 挖出该区块的矿工)
func main() {

	coinBaseDemo()


	// 创世区块，输入交易中的 Sigscript 中的内容，解密之后的内容如下：
	data,_ := hex.DecodeString("5468652054696d65732030332f4a616e2f32303039204368616e63656c6c6f72206f6e206272696e6b206f66207365636f6e64206261696c6f757420666f722062616e6b73")
	fmt.Println("data : ", string(data))   //  The Times 03/Jan/2009 Chancellor on brink of second bailout for banks

}

func coinBaseDemo() {

	// 解密区块中的 第一笔交易： 创币交易  (下面为 创币交易 交易输入的解锁脚本(scriptSig 的 hex 形式))

	/*  创币交易前几个字节也曾是可以任意填写的，不过在后来的第 34号比特币改进提议(BIP34) 中规定了版本2 的区块(版本字段为2 的区块),
	区块的高度必须填充在创币交易字段的起始处。
	第一个字节 03， 脚本执行引擎执行这个指令将后面 3个字节压入脚本栈，紧接着的 3个字节是：37bb07, 是以小端格式(最低有效字节在先) 编码的区块高度。
		翻转字节序得到 07bb37, 表示十进制是 506679,  代表该区块高度为 506679。
	*/

	// 0337bb07192f5669614254432f4d696e656420627920783636353537362f2cfabe6d6df237cd19122c7690b64b8c9ba5012fddd653fcff7b52cbf1ccc111c9bce04896010000000000000010a78866001e97f4c956bd0435acadaa2a

	// 将 字符串形式的16进制 转换为 []byte, 便于后面进行 小端转换
	data3,_ := hex.DecodeString("37bb07")
	// 进行小端转换
	ReverseBytes(data3)
	fmt.Println("data3 : ", data3)   // [7 187 55]

	// 再将 []byte 形式的 小端16进制  转换为 字符串形式的 小端16进制
	str := hex.EncodeToString(data3)
	fmt.Println("str : ",str)  // 07bb37
	// 通过 ParseInt 将16进制转换为 10进制  -->  int64
	result,_ := strconv.ParseInt(str, 16, 64)
	fmt.Printf("%d\n",result)
}

// 反转  []byte (针对 tx hash, merkle root hash 进行大小端转换)
func ReverseBytes(data []byte) {
	for i,j := 0, len(data)-1; i<j; i,j = i+1,j-1 {
		data[i],data[j] = data[j],data[i]
	}
}
