package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"
)

// 区块头信息
type BlockHeader struct {
	version int32
	prevHash []byte
	merkleRoot []byte
	time int32
	bits int32
	nonce int32
}

var (
	MAX_NONCE int32= math.MaxInt32
)


// 计算区块头 hash
func main() {

	fmt.Println("---------------------------------------------------------------")


	// 计算区块头的 hash
	// calcBlockHeaderHash()

	// 通过 bits 数据 计算出 区块目标困难值
	// 第一种方法：
	var bits int32 = 404454260
	// int 转换为 字符串形式的 16进制
	str := strconv.FormatInt(int64(bits), 16)
	// fmt.Println(str)   // 181b7b74
	data,_:= hex.DecodeString(str)
	result := calcTarget(data)
	fmt.Printf("result : %x\n", result)   // 00000000000000001b7b74000000000000000000000000000000000000000000

	// 第二种方法：结果与方法1 一致
	result2 := calcTarget2(data)
	fmt.Printf("result2 : %s\n", result2) // 00000000000000001b7b74000000000000000000000000000000000000000000

	// 测试 big.Int 写的求 x 的 n次方
	// testBigIntDemo()

	// 计算与创世区块难度 的差值， 可以形象的了解到 当前区块的难度 (数值与难度成正比，一个越来越大的数字)
	targetBitsStr := string(result2)   // 必须是可转为 字符串形式的 16进制  (calcTarget2() 这个方法返回的 data 可转为 字符串形式的 16进制)
	difficulty := CalculateDifficulty(targetBitsStr)
	fmt.Println("difficulty : ",difficulty)

	// 40,007,470,271.27

	fmt.Println("---------------------------------------------------------------")
	fmt.Println("----------------------------挖矿开始-----------------------------------")

	testMinerWork()

}

func (header *BlockHeader)serial() []byte{

	result := bytes.Join([][]byte{IntToHex(header.version),
								header.prevHash,
								header.merkleRoot,
								IntToHex(header.time),
								IntToHex(header.bits),
								IntToHex(header.nonce)},
								[]byte{}) // 拼接各个 []byte

	firstHash := sha256.Sum256(result)
	blockHeaderHash := sha256.Sum256(firstHash[:])

	ReverseBytes(blockHeaderHash[:])
	return blockHeaderHash[:]
}

// 挖矿
func minerWork(version int32,prevHash,merkleRoot []byte, time,bits int32) *BlockHeader{

	// 初始化 区块头信息  (版本，前一区块hash, 默克尔根, 时间戳, 目标难度值，随机数)
	header := BlockHeader{
		version:    version,
		prevHash:   prevHash,
		merkleRoot: merkleRoot,
		time:       time,
		bits:       bits,
		nonce:      0,
	}

	// 通过 bits (目标难度值) 计算出 目标难度 hash 值
	str := strconv.FormatInt(int64(header.bits), 16)
	// fmt.Println(str)   // 181b7b74
	data,_:= hex.DecodeString(str)
	target := calcTarget2(data)

	t,_ := hex.DecodeString(string(target))    // 将 target 字符串形式的 16进制 转换为  []byte 的 16进制

	var targetHash = big.NewInt(0)
	targetHash.SetBytes(t)   // 为了和 serial() 返回的 序列化 data 比较，targetHash.setBytes() 必须传入 16进制的 []byte
	// fmt.Printf("%x\n", targetHash)

	var headerHash = big.NewInt(0)

	// header.nonce = 1865996585
	// 18 6595
	for header.nonce < MAX_NONCE {
		data := header.serial()
		headerHash.SetBytes(data)
		fmt.Printf("nonce : %d , hash : %x\n", header.nonce, data)
		if headerHash.Cmp(targetHash) == -1 {

			fmt.Printf("找到随机数：%d\n", header.nonce)
			fmt.Printf("区块难度hash：%064x\n", targetHash)
			fmt.Printf("区块头 hash ：%064x\n", headerHash)
			break
		}
		header.nonce++
	}
	return &header
}

func testMinerWork(){
	var version int32 = 2
	// 区块的上一个区块的 hash  ---> 链上前一个区块的散列值的参考值
	prevHash, _ := hex.DecodeString("000000000000000016145aa12fa7e81a304c38aec3d7c5208f1d33b587f966a6")
	ReverseBytes(prevHash)
	// 区块的默克尔根  ---> 当前区块中所有交易产生的默克尔树根节点的散列值
	merkleRootHash, _ := hex.DecodeString("3a4f410269fcc4c7885770bc8841ce6781f15dd304ae5d2770fc93a21dbd70d7")
	ReverseBytes(merkleRootHash)

	var timeInt int32 = 1418755780
	//fmt.Println(timeInt)

	//fmt.Printf("time : %x\n", IntToHex(timeInt))
	// 目标难度值  ---> 当前区块 POW 算法的目标难度值
	var bits int32 = 404454260
	//fmt.Printf("bits : %x\n", IntToHex(bits))
	// 随机数  ---> 用于 POW 算法的计数器
	//var nonce int32 = 1865996595
	//fmt.Printf("nonce : %x\n", IntToHex(nonce))

	blockHeader := minerWork(version, prevHash,merkleRootHash,timeInt,bits)
	fmt.Println("--------------------------挖矿成功，出块儿！------------------------")
	fmt.Printf("区块信息： %#v\n", blockHeader)
}

/// 以 334599 区块为例： https://www.blockchain.com/btc/block/334599
// 计算区块头 hash
func calcBlockHeaderHash() {

	/*		区块头结构
		版本
		上一区块的hash
		默克尔根 hash
		时间戳
		目标难度值
		随机数
	 */

	// 版本   ---> 版本信息用于跟踪软件和协议的更新
	var version int32 = 2
	// 区块的上一个区块的 hash  ---> 链上前一个区块的散列值的参考值
	prevHash, _ := hex.DecodeString("000000000000000016145aa12fa7e81a304c38aec3d7c5208f1d33b587f966a6")
	ReverseBytes(prevHash)
	// 区块的默克尔根  ---> 当前区块中所有交易产生的默克尔树根节点的散列值
	merkleRootHash, _ := hex.DecodeString("3a4f410269fcc4c7885770bc8841ce6781f15dd304ae5d2770fc93a21dbd70d7")
	ReverseBytes(merkleRootHash)

	// 时间戳  ---> 当前区块的大致生成时间 (从 Unix 纪元时间开始的秒数)
	/*  注意 ： 比特币浏览器中 区块的时间戳显示 是做了 时区转换
		浏览器显示： 2014-12-17 02:49
		实际计算如下：
			timeStr := "2014-12-16 18:49:40"
			timeInt := stringToDate(timeStr).Unix()

			timeInt ---> 1418755780
	 */
	var timeInt int32 = 1418755780
	fmt.Println(timeInt)

	fmt.Printf("time : %x\n", IntToHex(timeInt))
	// 目标难度值  ---> 当前区块 POW 算法的目标难度值
	var bits int32 = 404454260
	fmt.Printf("bits : %x\n", IntToHex(bits))
	// 随机数  ---> 用于 POW 算法的计数器
	var nonce int32 = 1865996595
	fmt.Printf("nonce : %x\n", IntToHex(nonce))

	result := bytes.Join([][]byte{IntToHex(version), prevHash, merkleRootHash, IntToHex(timeInt), IntToHex(bits), IntToHex(nonce)}, []byte{}) // 拼接各个 []byte

	firstHash := sha256.Sum256(result)
	resultHash := sha256.Sum256(firstHash[:])

	ReverseBytes(resultHash[:])
	fmt.Printf("%x\n", resultHash)

	// block header hash
	// 00000000000000000a1f57cd656e5522b7bac263aa33fc98c583ad68de309603

	// target :  block header hash  要小于  target
	// 00000000000000001b7b74000000000000000000000000000000000000000000
}

// 计算目标难度值
func calcTarget(bits []byte) []byte{

	// 第一个字节代表 指数
	exponent := bits[:1]
	fmt.Printf("%x\n", exponent)
	// 后面的字节代表系数
	coefficient := bits[1:]
	fmt.Printf("%x\n",coefficient)

	// 将 exponent 转换为 字符串形式的 16进制
	str := hex.EncodeToString(exponent)
	// base:16 代表 str 为16进制, bitSize:8, 代表转换为 int8 类型
	exp, _ := strconv.ParseInt(str, 16, 8)

	// 在前面补0
	result := append(bytes.Repeat([]byte{0x00}, 32-int(exp)),coefficient...)
	// 在后面补0，使其达到 32个字节的长度
	result = append(result, bytes.Repeat([]byte{0x00}, 32-len(result))...)

	return result

}
// 使用 公式求 target
// 目标位 = 系数* 2 ^ (8 * (指数-3))
// result = coefficient * 2 ^(8 * (exponent-3))
func calcTarget2(bits []byte) []byte{

	// 第一个字节代表 指数
	exponent := bits[:1]
	//fmt.Printf("%x\n", exponent)
	// 后面的字节代表系数
	coefficient := bits[1:]
	//fmt.Printf("%x\n",coefficient)

	// 将 exponent 转换为 字符串形式的 16进制
	str := hex.EncodeToString(exponent)
	// base:16 代表 str 为16进制, bitSize:8, 代表转换为 int8 类型
	exp, _ := strconv.ParseInt(str, 16, 8)
	//fmt.Println("exp : ", exp)

	str2 := hex.EncodeToString(coefficient)
	coe,_ := strconv.ParseInt(str2, 16, 32)

	//fmt.Println("coe : ", coe)

	num1 := 8 * (exp-3)
	//fmt.Println("num1 : ",num1)

	var a = big.NewInt(2)
	var b = big.NewInt(num1)
	power := Powerf(a,b)  // 计算 a 的 b 次方 （big.Int 类型）

	var coeInt = big.NewInt(coe)
	power.Mul(power,coeInt)

	fmt.Printf("%x\n", power.Bytes())  // 1b7b74000000000000000000000000000000000000000000
	fmt.Printf("%d\n", len(power.Bytes()))  // 24

	// 目前还剩下的问题是， 要将 24位 前面补零 至 32位(还要补 16个 0)，有没有什么简单的办法。 (不使用 bytes.repeat)
	// 解决方案：
	/*
	    通过 fmt.Sprintf("%064x", power)   格式化拼接字符串, 将 power 的10进制值 以 %x(16进制)格式化输出，
	  	输出 64个长度 (没2个长度为 一个字节，共32个字节)，并且前面必须补 0
	*/

	target := fmt.Sprintf("%064x", power)
	fmt.Printf("格式化： %s\n", target)

	// targetHash,_ := hex.DecodeString(target)
	return []byte(target)


	// big.Int .Text(base int) 可以将10进制的结果 以字符串的形式显示出来
	/*
	fmt.Printf("%d\n", power)      // 673862533877092685902494685124943911912916060357898797056
	fmt.Println("power : ",power)  // 673862533877092685902494685124943911912916060357898797056
	fmt.Println(power.Text(10))   // 673862533877092685902494685124943911912916060357898797056
	fmt.Println(power.Text(16))   //  1b7b74000000000000000000000000000000000000000000
	*/

}

// 与创世区块 对比，计算难度值 (计算与创世区块难度 的差值， 可以形象的了解到 当前区块的难度 (数值与难度成正比，一个越来越大的数字))
func CalculateDifficulty(strTargetHash string) string {
	strGeniusBlockHash := "00000000ffff0000000000000000000000000000000000000000000000000000" // 创世块编号

	var biGeniusHash big.Int
	var biTargetHash big.Int
	biGeniusHash.SetString(strGeniusBlockHash, 16)
	biTargetHash.SetString(strTargetHash, 16)

	difficulty := big.NewInt(0)
	difficulty.Div(&biGeniusHash, &biTargetHash)
	//fmt.Printf("%T \n" , difficulty)
	return fmt.Sprintf("%s", difficulty)
}

// OK
func testBigIntDemo(){
	var a = big.NewInt(2)
	var b = big.NewInt(10)

	res := Powerf(a,b)
	fmt.Println(res.Int64())
}

/*  go 语言 自己实现 x 的 n 次方
func powerf(x float64, n int) float64 {
    ans := 1.0

    for n != 0 {
        if n%2 == 1 {
            ans *= x
        }
        x *= x
        n /= 2
    }
    return ans
}

 */

//  使用 big.Int 实现 x 的 n 次方 （对比如上）
func Powerf(x *big.Int, n *big.Int) *big.Int {
	var ans = big.NewInt(1)
	var zero = big.NewInt(0)
	var one = big.NewInt(1)
	var second = big.NewInt(2)

	var modRes = big.NewInt(0)

	for n.Cmp(zero) !=0 {
		if modRes.Mod(n,second).Cmp(one) == 0 {
			ans.Mul(ans,x)
		}
		x.Mul(x,x)
		n.Div(n,second)
	}
	return ans
}

// 将 int 类型转换为 16进制小端模式
func IntToHex(num int32) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff,binary.LittleEndian,num)

	if err != nil {
		panic("IntToHex failed")
	}
	return buff.Bytes()
}

func Int64ToHex(num uint64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff,binary.LittleEndian,num)

	if err != nil {
		panic("float64ToHex failed")
	}
	return buff.Bytes()
}


// 将日期字符串转换为 time.Time
func stringToDate(dateStr string) time.Time {
	date, err := time.Parse("2006-01-02 15:04:05", dateStr)
	if err != nil {
		panic("stringToDate failed")
	}
	// fmt.Println(date)
	return date
}

// 反转  []byte (针对 tx hash, merkle root hash 进行大小端转换)
func ReverseBytes(data []byte) {
	for i,j := 0, len(data)-1; i<j; i,j = i+1,j-1 {
		data[i],data[j] = data[j],data[i]
	}
}
