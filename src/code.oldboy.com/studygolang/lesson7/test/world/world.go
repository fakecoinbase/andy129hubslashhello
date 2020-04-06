package world

import "unicode"

// IsPalindrome 是一个回文判断函数， 返回 布尔值
/*
func IsPalindrome(s string) bool {

	s2 := []rune(s)   // 将字符串强制转换为 rune 类型 (解决中文字符通过下标取值时的不连续性)
	length := len(s2)

	for i:=0;i<len(s2)/2;i++ {
		if s2[i] != s2[length-i-1]{
			return false
		}
	}
	return true
}
*/

// IsPalindrome 优化2 ： 针对 "Madam,I’mAdam"  带标点符号的 回文判断
func IsPalindrome(s string) bool {

	var letters []rune 
	// 通过 range s 遍历，可以将普通英文按照一个字节导出，中文等其他字符按照 3个字节或者4个字节导出
	for _, l := range s {
		// 判断是否为字符 (除去标点符号)
		if unicode.IsLetter(l) {
			letters = append(letters, unicode.ToLower(l))   // 回文判断 不区分大小写, 所以将其统一小写处理
		}
	}
 
	// 通过去除标点符号，统一小写，并且存放在 支持中文的  []rune里面
	length := len(letters)

	for i:=0;i<length/2;i++ {
		if letters[i] != letters[length-i-1]{
			return false
		}
	}
	return true 
}