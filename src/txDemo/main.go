package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"strings"
)

// 交易  (包含：交易ID，多个输入交易，多个输出交易)
type Transaction struct {
	 ID []byte
	 Vin []TXInput
	 Vout []TXOutput
}

// 交易输入
type TXInput struct {
	TXid []byte
	VoutIndex int   // 对应的 交易输出的 哪一笔 ( index)
	Signature []byte
}
// 交易输出
type TXOutput struct {
	value int
	PubkeyHash []byte
}

const SubSidy = 50

// 比特币交易 实战demo
func main() {

	newTX := NewCoinbaseTX("18rNG8sdeyLDoa88896vMnbXkEkLCsVnYj")

	fmt.Println(newTX)

}
/*
		95651c897012ff80f7e8400e9c0904ad20a894fe2f9e2c2fac7703ba231b1ed4
		3138724e4738736465794c446f613838383936764d6e62586b456b4c4373566e596a
*/
func (tx Transaction) String() string {
	var lines []string
	lines = append(lines, fmt.Sprintf(" Transaction : %x", tx.ID))
	for i,input := range tx.Vin {
		lines = append(lines, fmt.Sprintf("Input : %d", i))
		lines = append(lines, fmt.Sprintf("  TXID : %x", input.TXid))
		lines = append(lines, fmt.Sprintf("  Out : %d", input.VoutIndex))
		lines = append(lines, fmt.Sprintf("  Signature : %x", input.Signature))
	}

	for i,output := range tx.Vout {
		lines = append(lines,fmt.Sprintf("Ouput : %d", i))
		lines = append(lines,fmt.Sprintf("  Value : %d", output.value))
		lines = append(lines,fmt.Sprintf("  Script : %s",string(output.PubkeyHash)))
	}

	return strings.Join(lines, "\n")
}

// 序列化交易
func (tx Transaction)Serialize() []byte {
	var encoded bytes.Buffer
	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		fmt.Println("encode err : ", err)
		return nil
	}
	return encoded.Bytes()
}
// 计算 交易的 hash
func (tx *Transaction)Hash() []byte {
	txcopy := *tx
	txcopy.ID = []byte{}

	hash := sha256.Sum256(txcopy.Serialize())

	return hash[:]
}
// 根据 金额和地址 新建一个交易输出
func NewTXOutput(value int, address string) *TXOutput {
	txo := &TXOutput{value, []byte(address)}
	return txo
}

// 新建一个 coinbase 交易
func NewCoinbaseTX(to string) *Transaction{
	// coinbase 交易没有  输入交易，所以txin 的字段属性都设置为 空
	txin := TXInput{[]byte{}, -1, nil}
	txout := NewTXOutput(SubSidy,to)

	tx := Transaction{nil,[]TXInput{txin}, []TXOutput{*txout}}
	tx.ID = tx.Hash()
	return &tx
}































