package main

import "fmt"

func main() {
	scanDemo()
}

// 示例参考：转载于:https://blog.51cto.com/dreamlinux/2354784
func scanDemo(){
	//要求： 可以从控制台接受用户的信息 【姓名，年龄，薪水，是否通过考试】
	//方式1 fmt.Scanln()ya
	var  name string
	var  age  byte
	var  sal  float32
	var ispass bool
	fmt.Println("请输入姓名：")
	//当程序只是到fmt.Scanln(&name)程序会停止执行等待用户输入
	fmt.Scanln(&name)
	fmt.Println("请输入年龄：")
	//当程序只是到fmt.Scanln(&age)程序会停止执行等待用户输入
	fmt.Scanln(&age)
	fmt.Println("请输入薪水：")
	//当程序只是到fmt.Scanln(&sal)程序会停止执行等待用户输入
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过：")
	//当程序只是到fmt.Scanln(&name)程序会停止执行等待用户输入
	fmt.Scanln(&ispass)
	fmt.Printf(" The name is:%s,age:%d,sal:%f, ispass :%t",name,age,sal,ispass)
}
