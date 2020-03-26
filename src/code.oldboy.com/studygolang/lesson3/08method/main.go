package main

import "fmt"

type people struct {
	name   string
	gender string
	dream  string
}

// method:  方法
// 函数是谁都可以调用的。
// 方法就是某个具体的类型才能调用的函数
func main() {

	fmt.Println("方法")

	eat() // 函数

	var yang = people{
		"阳",
		"男",
		"不上班也有钱拿",
	}

	yang.printDream()        // 方法
	fmt.Println(yang.gender) // "男"

	// 接受者为 指针类型
	// yang.printDreamByPtr()   // 可以使用语法糖 (简写)
	(&yang).printDreamByPtr() // 方法
	fmt.Println(yang.gender)  // "女"

	// 通过 printDream() 与 printDreamByPtr() 修改 gender 的结果对比
	// 再次验证 结构体是值传递，拷贝操作。  如果想实现修改，则可以使用指针。

	/* 什么时候用指针？
	1，需要修改接受者中的值
	2，接受者是 拷贝代价比较大的 大对象
	3, 保证一致性，如果有某个方法使用了指针接受者，那么其他的方法也应该使用指针接受者。
	*/

	// check 里面指定 接受者 (sss *people) ,  虽然可以正常调用，但是 命名规范 会建议你 写成 (p *people) , 使用后面类型(people)的首字母 p 的小写
	yang.check()
	fmt.Println(yang.gender) // "瞎写"
}

// 函数指定接受者之后就是方法
// (p people) 指定接受者
// 在 go 语言中约定俗成 不用 this 也不用 self , 而是使用后面类型(people)的首字母 p 的小写，
// 例如 struct 的名字是  DOG, 那么指定接受者应该为： (d DOG)
func (p people) printDream() { // 另外，结构体的方法名 不能与 结构体里的属性名字 相同。 比如：结构体里面有个 dream 属性，那么它的方法就不能定义为 dream()
	p.gender = "女"
	fmt.Printf("%s的梦想: %s\n", p.name, p.dream)
}

func (p *people) printDreamByPtr() {
	(*p).gender = "女"

}

func (sss *people) check() {
	sss.gender = "瞎写"
}

func eat() {
	fmt.Println("吃....")
}
