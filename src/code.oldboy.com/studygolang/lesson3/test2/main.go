package main

import (
	"fmt"
	"os"
)

/*   插曲，新学一个技能：
1，如果 Student 这个结构体 单独写在一个文件中 student.go
2, main.go 包里 又调用了 Student 相关方法和字段
3, 那么我运行程序的时候：  go run main.go 是会报错的，提示： Student 没定义，找不到

那么我们可以执行这样的命令：
go run main.go student.go

*/

// 学员管理系统

// Admin 是一个 管理员结构体
type Admin struct {
	name     string
	password string
	gender   string
	age      int
}

// Student 是一个学员结构体
type Student struct {
	id     string
	name   string
	gender string
	age    int
	class  string
}

// 定义一个切片，用于保存学员的信息
var allStudents = make([]*Student, 0, 100)
var admin = &Admin{}

// 实例化一个 Admin 结构体
func newAdmin(name, password, gender string, age int) *Admin {
	return &Admin{
		name:     name,
		password: password,
		gender:   gender,
		age:      age,
	}
}

// 实例化一个 Student 结构体
func newStudent(id, name, gender string, age int, class string) *Student {
	return &Student{
		id:     id,
		name:   name,
		gender: gender,
		age:    age,
		class:  class,
	}
}

func (a *Admin) addStudent(s *Student) {
	allStudents = append(allStudents, s)
}

func (a *Admin) updateStudent(s *Student) bool {
	exsit := false
	for index, v := range allStudents {
		exsit = false

		if s.id == v.id { // 判断要求修改的学员ID 名称 是否存在
			allStudents[index] = s // 整个对象直接赋值
			exsit = true
			break
		}
	}
	/*
		for _, v := range allStudents {
			fmt.Println("修改学员信息如下：")
			fmt.Println(*v)
		}
	*/
	return exsit
}

// 注意 切片删除元素要怎么操作
func (a *Admin) deleteStudent(id string) bool {

	exsit := false
	for index, v := range allStudents {
		exsit = false

		if v.id == id { // 判断要求修改的学员ID 名称 是否存在

			// 进行切片 删除工作

			allStudents = append(allStudents[:index], allStudents[index+1:]...) // 注意参数 2 的写法 ...

			exsit = true
			break
		}
	}
	return exsit
}

func (a *Admin) showAllStudents() {
	if len(allStudents) > 0 {
		fmt.Println("--学员信息如下: ")

		for _, v := range allStudents {
			fmt.Println()
			fmt.Printf("--id:%s,name:%s,gender:%s,age:%d,class:%s\n", v.id, v.name, v.gender, v.age, v.class)
			fmt.Println()
		}
	} else {
		fmt.Println("--没有任何学员，请先添加！")
		fmt.Println()
	}
}

func main() {

	admin = newAdmin("刘老师", "12345678", "女", 33)
	showMenu()

	// testSlice()
}

func testSlice() {
	a := []int{1, 2, 3, 4, 5}

	a = append(a[:4], a[5:]...)

	fmt.Println(a)
}

func showMenu() {

	for {

		// 将 index 声明到 for 循环内，就能保证每次回到菜单界面，index 默认都为0，而不会保存上一次操作的指令值。
		// 指令
		var index int

		fmt.Println("-----------学员管理系统---------")
		fmt.Println()
		fmt.Println("1,添加学员")
		fmt.Println("2,修改学员信息")
		fmt.Println("3,删除学员信息")
		fmt.Println("4,展示所有学员")
		fmt.Println("5,退出")
		fmt.Scanln(&index) // Scanln 的好处在于，当用户按下回车代表 该行输入结束，如果用 scan 的话，用户按下回车，程序不结束，终端会显示空行，界面不美观
		fmt.Println("打印信息：", index)

		switch index {
		case 1:
			addStu()
		case 2:
			updateStu()
		case 3:
			deleteStu()
		case 4:
			showAllStus()
		case 5:
			fmt.Println("正常退出")
			os.Exit(0)
		default:
			fmt.Println("指令不匹配，退出")
			os.Exit(0)
		}
	}

}

// 学员管理系统
func addStu() {

	stu := userInput()

	fmt.Println(*stu)

	admin.addStudent(stu) // 管理员才能调用的方法

	fmt.Println()
}

func updateStu() {

	stu := userInput()

	exsit := admin.updateStudent(stu)

	if exsit {
		fmt.Println("--修改成功！")
	} else {
		fmt.Println("--该学员不存在, 请先添加！")
	}
	fmt.Println()
}

func deleteStu() {

	var id string

	fmt.Println()
	fmt.Println("###请依次根据提示输入相关信息.")
	fmt.Print("id: ")
	fmt.Scanln(&id)

	exsit := admin.deleteStudent(id)

	if exsit {
		fmt.Println("--删除成功！")
	} else {
		fmt.Println("--该学员不存在, 请先添加！")
	}
	fmt.Println()

}

func showAllStus() {
	admin.showAllStudents()
}

func userInput() *Student {
	var (
		id     string
		name   string
		gender string
		age    int
		class  string
	)
	fmt.Println()
	fmt.Println("###请依次根据提示输入相关信息.")
	fmt.Print("id: ")
	fmt.Scanln(&id) // 注意  Scan() 里面传入的都是 指针类型
	fmt.Print("name: ")
	fmt.Scanln(&name)
	fmt.Print("gender: ")
	fmt.Scanln(&gender)
	fmt.Print("age: ")
	fmt.Scanln(&age)
	fmt.Print("class: ")
	fmt.Scanln(&class)

	stu := newStudent(id, name, gender, age, class)
	return stu
}
