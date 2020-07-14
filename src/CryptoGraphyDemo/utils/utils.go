// @Title  
// @Description  
// @Author  yang  2020/7/10 18:38
// @Update  yang  2020/7/10 18:38
package utils

import (
	"bytes"
)

// src: 待填充的数据， blockSize: 分组大小
// PaddingText 填充最后一个分组
func PaddingText(src []byte, blockSize int) []byte {
	// 求出最有一个分组需要填充的字节数
	padding := blockSize- len(src)%blockSize
	// 创建新的切片，
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	// 将新创建的切片和待填充的数据进行拼接
	nextText := append(src,padText...)
	return nextText

	/*	示例，加入 src 长度为 10,  blockSize 为 8 (也就是 64位)为一组，则需要将 10 长度填充至 8+2+6 = 16 字节长度

		padding := blockSize- len(src)%blockSize    --->  8 - 10%8 == 8 - 2 == 6
		padText := bytes.Repeat([]byte{byte(padding)}, padding)  ---> 此时 padText 是 填充了 padding (6) 个 byte(padding)（表示数字 6的字节） 的一个字节切片
		nextText := append(src,padText...)    ---> 最后将 填充完毕的 padText 追加到 src 后面，就完成了 src 的填充。
	 */
}

// 删除尾部填充数据 得到原始数据
func UnPaddingText(src []byte) []byte {
	// 获取待处理的数据的长度
	length := len(src)
	// 取出最后一个字符，即可得到我们填充了多少位个数据
	number := int(src[length-1])
	//fmt.Println("number : ", number)
	newText := src[:length-number]

	return newText
}

// 填充最后一个分组，填充 0
// src: 待填充的数据， blockSize: 分组大小
func ZeroPadding(src []byte, blockSize int) []byte {
	// 求出最有一个分组需要填充的字节数
	padding := blockSize- len(src)%blockSize
	// 创建新的切片，
	padText := bytes.Repeat([]byte{0}, padding)

	// 将新创建的切片和待填充的数据进行拼接
	nextText := append(src,padText...)
	return nextText
}

// 去除尾部填充的 0  （问题，万一 src 原本里面就有 0，怎么办呢）
func ZeroUnPadding(src []byte) []byte {
	return bytes.TrimRightFunc(src, func(r rune) bool {
		return r == rune(0)
	})
}
