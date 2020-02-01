package tempconv


func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32)*5 / 9)
}

// 函数首字母没有大写，所以在外面无法直接通过 包名直接调用
func tempFunc(){

}
