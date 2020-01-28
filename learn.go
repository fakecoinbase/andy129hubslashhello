package main

import "fmt"

const boilingF = 212.0  // 沸水的温度
// 不同于java 常量定义，不需要在前面加入 int 或者 float，直接赋值浮点数就代表是 float 类型

func main() {
	var f = boilingF
	var c = (f - 32)*5 / 9  // 华氏度与摄氏度的 转换
	fmt.Printf("boiling point = %g℉ or %g℃\n",f,c)

	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("freezing: %g℉ = %g℃\n", freezingF, fToC(freezingF))
	fmt.Printf("boiling: %g℉ = %g℃\n", boilingF,fToC(boilingF))
}

// 将 华氏度与摄氏度的转换 包装成一个函数里，其他地方可以任意调用。
func fToC(f float64) float64 {
	return (f - 32)*5 / 9
}
