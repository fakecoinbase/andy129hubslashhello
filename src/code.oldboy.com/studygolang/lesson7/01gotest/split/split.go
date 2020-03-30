package split

import "strings"

// Split 按照sep 分割 str (优化前)
/*
func Split(str, sep string) (result []string) {
	index := strings.Index(str, sep)
	for index >= 0 {
		result = append(result, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}

	// 例如："fsdfafdfa" , 以 "fa" 分隔， 刚好分隔符在结尾，这种情况下，上面的代码 str 最后的结果可能会是 一个 "", 所以要舍弃掉
	if str != "" {
		result = append(result, str)
	}

	return
}
*/

// Split 按照sep 分割 str  （优化版本）
func Split(str, sep string) []string {
	count := strings.Count(str, sep)     // 计算一下字符串 str 中包含多少个 sep
	result := make([]string, 0, count+1) // 根据 sep 的数量初始化切片
	index := strings.Index(str, sep)
	for index >= 0 {
		result = append(result, str[:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}

	// 例如："fsdfafdfa" , 以 "fa" 分隔， 刚好分隔符在结尾，这种情况下，上面的代码 str 最后的结果可能会是 一个 "", 所以要舍弃掉
	if str != "" {
		result = append(result, str)
	}

	return result
}
