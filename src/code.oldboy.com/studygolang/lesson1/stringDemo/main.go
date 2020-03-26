package main

import "fmt"

// 字符串反转操作
func main() {

	s1 := "hello"
	byteArray := []byte(s1) // [h e l l o]
	s2 := ""
	for i := len(byteArray) - 1; i >= 0; i-- {

		s2 = s2 + string(byteArray[i]) //  将字符强制转换为 字符串，然后 与 s2 拼接
	}
	fmt.Println(s2) // "olleh"

	// 方法2
	length := len(byteArray)
	for i := 0; i < length/2; i++ {
		byteArray[i], byteArray[length-1-i] = byteArray[length-1-i], byteArray[i]
	}
	fmt.Println(string(byteArray)) // "olleh"
}
