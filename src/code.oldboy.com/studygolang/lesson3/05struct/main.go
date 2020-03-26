package main

import "fmt"

// struct
/*
 1, struct 是值类型，没初始化，那么每个属性的值都是 默认值
 2, 键值对初始化，与 值初始化 对比

*/

// 创建新的类型要使用 type 关键字
type student struct {
	name   string
	age    int
	gender string
	hobby  []string
}

type school struct {
	name     string
	history  int
	students int
	courses  []string
}

func main() {
	fmt.Println("struct")

	// test2()
	// test3()
	// test4()

	// init1()
	// init2()
	// init3()
	// init4()
}

// 实例化方法3 (声明)
// 创建 结构体指针
func test4() {
	/*
		var zhang = &student{
			name: "张飞",
		}
		fmt.Printf("%T\n", zhang)  // "*main.student"
		fmt.Println((*zhang).name) // "张飞"
	*/
	fmt.Println("---------------------------------------")

	var liu = &student{}
	fmt.Printf("%T\n", liu)  // "*main.student"
	fmt.Println((*liu).name) // ""
}

// 实例化方法2 (声明)
// 创建 结构体指针
func test3() {
	var yang = new(student)
	fmt.Println(yang)        // "&{ 0  []}"
	fmt.Printf("%T\n", yang) // "*main.student"

	//(*yang).name = "阳浩"
	//fmt.Println(*yang) // "{阳浩 0  []}"
}

// 实例化方法1 (声明)
// 创建具体的结构体
func test2() {
	var wangzhan = student{}
	// struct 是值类型，没初始化，那么每个属性的值都是 默认值
	fmt.Println(wangzhan)        // "{ 0  []}"
	fmt.Println(wangzhan.age)    // "0"
	fmt.Println(wangzhan.gender) // ""
	fmt.Println(wangzhan.name)   // ""
	fmt.Println(wangzhan.hobby)  // "[]"
}

// 基本初始化 (赋值)
func init1() {
	var haojie = student{
		name:   "豪杰",
		age:    19,
		gender: "男",
		hobby:  []string{"篮球", "足球", "双色球"},
	}
	fmt.Println(haojie) // "{豪杰 19 男 [篮球 足球 双色球]}"

	// 结构体支持 . 来访问属性
	fmt.Println("name : ", haojie.name)
	fmt.Println("age : ", haojie.age)
	fmt.Println("gender : ", haojie.gender)
	fmt.Println("hobby : ", haojie.hobby)

	fmt.Println("-------------------------------------------------")

	var bfaSchool = school{
		name:     "北京电影学院",
		history:  125,
		students: 300,
		courses:  []string{"导演系", "摄影系", "美术系", "录音系", "制片"},
	}
	fmt.Println(bfaSchool) // "{北京电影学院 125 300 [导演系 摄影系 美术系 录音系 制片]}"

	fmt.Println("name : ", bfaSchool.name)
	fmt.Println("history : ", bfaSchool.history)
	fmt.Println("students : ", bfaSchool.students)
	fmt.Println("courses : ", bfaSchool.courses)
}

// 结构体初始化 (赋值)
// 省略 key ,直接初始化 value
// value 初始化时，需要将结构体里面所有的 属性都初始化
func init2() {
	var stu1 = student{
		"豪杰",
		18,
		"男",
		[]string{"篮球", "足球"},
	}

	fmt.Println(stu1)      // "{豪杰 18 男 [篮球 足球]}"
	fmt.Println(stu1.name) // "豪杰"
}

// 指针初始化
// 键值对初始化时，不需要将全部属性都初始化
func init3() {
	var stu1 = &student{
		name: "豪杰人",
		age:  18,
	}

	fmt.Println(*stu1)        // "{豪杰人 18  []}"
	fmt.Println((*stu1).name) // "豪杰人"
}

// 指针实例化，然后单独对每个属性进行初始化
func init4() {
	var stu1 = new(student) // 采用 new 方式实例化对象之后，不能在后面加 {}
	(*stu1).name = "南京人艺"

	fmt.Println(*stu1)        // "{南京人艺 0  []}"
	fmt.Println((*stu1).name) // "南京人艺"
}
