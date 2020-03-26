package main

import (
	"fmt"
	"strings"
)

// 定义各种常量  (不能全用大写 (CAT_HUNGRY)， 要使用驼峰命名法)
const (
	CatHungry = "喵，喵，..."
	CatThirst = "喵喵喵...."

	DogHungry  = "汪，汪，..."
	DogThirsty = "汪汪汪...."

	HumanHungry  = "饿了"
	HumanThirsty = "渴了"

	HumanUnkown = "无法识别！"
)

// animal 是一个接口类型
// 注意观察 里面方法的命名规范
// 1, 接口命名： 在 Go 语言中的，一般会在单词后面添加 er, 如有写操作的接口叫 Writer
// 2, 接口中方法命名规范，如下：
type animal interface {
	move()
	convertLanguage(string) string
	// convertLanguage(str string) (abc string)    // 也可以指定 变量名
}

// 声明一个切片用来保存 所有的 animal
var animalList []animal

// cat 是一个结构体
type cat struct {
	name  string
	color string
}

// dog 是一个结构体
type dog struct {
	name  string
	color string
}

// cat 内部方法 (实现 animal 接口)
func (c cat) move() {
	fmt.Printf("cat name : %s ---> move()\n", c.name)
}

// cat 内部方法 (实现 animal 接口)
func (c cat) convertLanguage(animalLangStr string) (humanLangStr string) {

	if strings.Compare(animalLangStr, CatHungry) == 0 {
		humanLangStr = HumanHungry
		return
	}
	if strings.Compare(animalLangStr, CatThirst) == 0 {
		humanLangStr = HumanThirsty
		return
	}
	humanLangStr = HumanUnkown
	return
}

// dog 内部方法 (实现 animal 接口)
func (d dog) move() {
	fmt.Printf("dog name : %s ---> move()\n", d.name)
}

// dog 内部方法 (实现 animal 接口)
func (d dog) convertLanguage(animalLangStr string) (humanLangStr string) {

	if strings.Compare(animalLangStr, DogHungry) == 0 {
		humanLangStr = HumanHungry
		return
	}
	if strings.Compare(animalLangStr, DogThirsty) == 0 {
		humanLangStr = HumanThirsty
		return
	}
	humanLangStr = HumanUnkown
	return
}

// interface
/*
// 1，接口是什么?
	接口定义了一个对象的行为规范，只定义规范不实现，由具体的对象来实现规范的细节。
	在 Go 语言中接口 (interface) 是一种类型，一种抽象的类型。它不关心属性(数据)， 只关心行为 (方法)
// 2，接口的命名规范
// 3, 接口的实现过程
// 4, 为什么会用到接口
*/
func main() {
	fmt.Println("interface")

	c := cat{
		name:  "娜娜",
		color: "橘色",
	}

	d := dog{
		name:  "小黑",
		color: "黑色",
	}

	// c.move()
	// d.move()

	// fmt.Println(len(animalList), cap(animalList)) // "0 0"
	// fmt.Println(animalList == nil)                // "true"

	// 注意看这里，c 和 d 是不同的类型，为什么都可以加入到 animalList 这个用来存放 animal 接口类型的 集合呢？
	// 原因在于 cat 和 dog 都实现了 animal 接口类型里面定义的  move()  和  convertLanguage(string)string 方法
	animalList = append(animalList, c, d)
	fmt.Println(len(animalList), cap(animalList)) // "2 2"

	// 接口类型使用的好处，就是可以减少冗余代码， 把各种结构体的可以公用的方法 统一定义在接口类型中。
	for _, v := range animalList {
		v.move()
		fmt.Println(v.convertLanguage(DogHungry))
	}

	/*	打印结果：

		cat name : 娜娜 ---> move()
		无法识别！
		dog name : 小黑 ---> move()
		饿了
	*/

	fmt.Println("------------------------------------------------")

	// 实例化 接口的方法，如下。
	// 声明一个接口变量
	var a animal
	fmt.Println(a)        // "<nil>"
	fmt.Printf("%T\n", a) // "<nil>"
	// 将实现 a 接口的 struct c 赋值给 a, 从而达到 a 的实例化
	a = c
	fmt.Println(a)        // "{娜娜 橘色}"
	fmt.Printf("%T\n", a) // "main.cat"

}
