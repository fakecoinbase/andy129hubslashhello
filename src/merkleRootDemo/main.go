package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)


// 计算默克尔根 示例
func main() {
	calcMerkleRoot()
}

/*   // 以比特币 98901 区块信息为示例：https://www.blockchain.com/btc/block/98901


### 区块头部信息

	BTC
Hash
000000000001741120135274584b2a0da45b39c8cc78322a14f9004ae766a8e0
Confirmations
536,365
Timestamp
2010-12-22 20:22
Height
98901
Miner
Unknown
Number of Transactions
2
Difficulty
14,484.16
Merkle root
c7a040cddb20243fb85cdbdee65cb315d678a9e25886ffc5361b89d0dd945852
Version
0x1
Bits
453,281,356
Weight
1,896 WU
Size
474 bytes
Nonce
1,919,102,746
Transaction Volume
1.33000000 BTC
Block Reward
50.00000000 BTC
Fee Reward
0.01000000 BTC



####  区块交易信息：

Block Transactions
Hash
16f0eb42cb4d9c2374b2cb1de4008162c06fdd8f1c18357f0c849eb423672f5f
2010-12-22 20:22
COINBASE (Newly Generated Coins)
15G2rKBCufAMekXEe1D87TvAGeQ15JFCRr
50.01000000 BTC
Fee
0.00000000 BTC
(0.000 sat/B - 0.000 sat/WU - 134 bytes)
50.01000000 BTC
Hash
cce2f95fc282b3f2bc956f61d6924f73d658a1fdbc71027dd40b06c15822e061
2010-12-22 20:22
1P83a5bNdzDpYPxFgxf9XJJ5ET9e1k7m58
1.34000000 BTC
1GYbShmKLQApcP6PFGKLygvitTrSaW2bfJ
1.00000000 BTC
15WabzPcwrHbE7TsrfyMK4VzNJFL4PxAgH
0.33000000 BTC
Fee
0.01000000 BTC
(3861.004 sat/B - 965.251 sat/WU - 259 bytes)
1.33000000 BTC

 */

// 通过 区块中所有交易的 hash ，计算出 默克尔树根
func calcMerkleRoot() {

	/*
			data1,_ := hex.DecodeString("6b6a4236fb06fead0f1bd7fc4f4de123796eb51675fb55dc18c33fe12e33169d")
			data2,_ := hex.DecodeString("2af6b6f6bc6e613049637e32b1809dd767c72f912fef2b978992c6408483d77e")

					--->  33feaa29de38aedfd3767f495a61deb32ffb1d967540636bad9bb72188794038

			data3,_ := hex.DecodeString("6d76d15213c11fcbf4cc7e880f34c35dae43f8081ef30c6901f513ce41374583")
			data4,_ := hex.DecodeString("08c3b50053b010542dca85594af182f8fcf2f0d2bfe8a806e9494e4792222ad2")

					--->  013ae4dd116fd4456b35486298f18569eebebfc2dfbb7b029e14bac4370d476f

			data5,_ := hex.DecodeString("612d035670b7b9dad50f987dfa000a5324ecb3e08745cfefa10a4cefc5544553")

					---> double  f8155014448fbe5e37b545682b0e2755a321623ca4c167b9c89c345eeeb5f11e


							33feaa29de38aedfd3767f495a61deb32ffb1d967540636bad9bb72188794038
							013ae4dd116fd4456b35486298f18569eebebfc2dfbb7b029e14bac4370d476f

								--->  50eef8c1e84c383b2c0cf407c46590da9b31606ee2a1d3f89988662af859d2c4

							f8155014448fbe5e37b545682b0e2755a321623ca4c167b9c89c345eeeb5f11e
								double -->  31e4de07a73d91dee8d3b59be1914d3f24a36b25bd96fc37092fc62afcf4fe49


										50eef8c1e84c383b2c0cf407c46590da9b31606ee2a1d3f89988662af859d2c4
										31e4de07a73d91dee8d3b59be1914d3f24a36b25bd96fc37092fc62afcf4fe49
											--->   c66ee6e01c2332b92e71e17b6c6c3d4e926f6330a06acbb0e203bf7d97d12249
	 */

	hash1,_:= hex.DecodeString("50eef8c1e84c383b2c0cf407c46590da9b31606ee2a1d3f89988662af859d2c4")
	hash2,_:= hex.DecodeString("31e4de07a73d91dee8d3b59be1914d3f24a36b25bd96fc37092fc62afcf4fe49")

	// 大小端转换, 真正的 transaction id
	ReverseBytes(hash1)
	ReverseBytes(hash2)

	rawData := append(hash1,hash2...)
	firsthash := sha256.Sum256(rawData)
	secondhash := sha256.Sum256(firsthash[:])

	result := secondhash[:]

	ReverseBytes(result)
	fmt.Printf("%x", result)   // c7a040cddb20243fb85cdbdee65cb315d678a9e25886ffc5361b89d0dd945852
	// 结果 与 区块信息中  Merkle root 一致
	//c7a040cddb20243fb85cdbdee65cb315d678a9e25886ffc5361b89d0dd945852
}

// 反转  []byte
func ReverseBytes(data []byte) {
	for i,j := 0, len(data)-1; i<j; i,j = i+1,j-1 {
		data[i],data[j] = data[j],data[i]
	}
}


