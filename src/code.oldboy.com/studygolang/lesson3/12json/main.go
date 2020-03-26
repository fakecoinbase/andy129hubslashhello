package main

import (
	"encoding/json"
	"fmt"
)

// Student 是一个：注意字段首字母大写，方便 json包访问
type Student struct {
	ID     int
	Gender string
	Name   string
}

// Teacher 是一个：测试小写： json 是否能读取
type Teacher struct {
	id     int
	gender string
	name   string
}

// Doctor 是一个：标注 语言信息 (tag),  指明 当是json 序列化的时候采用 小写, 还用在 xml, db 操作时
type Doctor struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
}

// json 序列化
// 序列化: 将编程语言里面的数据 转换为 json 格式的数据
func main() {

	// test1()
	// test2()
	test3()
}

// 如何让web 前端能使用小写字段解析，而又能让我们后台 json 包 访问到 结构体里的字段呢？
func test3() {

	var d = Doctor{
		ID:     100,
		Gender: "男",
		Name:   "终南山",
	}
	// 序列化
	v, err := json.Marshal(d) // 可正常序列化
	if err != nil {
		fmt.Println("JSON格式化出错！", err)
	} else {
		// 并且也保证了字段首字母的 小写，方便 web 前端解析数据
		fmt.Println(v)                 // "[123 34 105 100 34 58 49 48 48 44 34 103 101 110 100 101 114 34 58 34 231 148 183 34 44 34 110 97 109 101 34 58 34 231 187 136 229 141 151 229 177 177 34 125]"
		fmt.Println(string(v))         // {"id":100,"gender":"男","name":"终南山"},   // 将 []byte 转换为字符串
		fmt.Printf("%#v\n", string(v)) // "{\"id\":100,\"gender\":\"男\",\"name\":\"终南山\"}"
	}

	// 反序列化：
	var docInfo = &Doctor{}
	str := string(v)
	// 反序列化
	json.Unmarshal([]byte(str), docInfo) //
	fmt.Println(*docInfo)                // "{100 男 终南山}"
	fmt.Printf("%T\n", docInfo)          // "*main.Doctor"

	fmt.Println(docInfo.Name) // "终南山"
}

// 结构体里的字段首字母为 小写， 无法让 json 访问
func test2() {
	var teacher = Teacher{
		id:     100,
		gender: "女",
		name:   "徐",
	}

	// teacher 里面的字段首字母声明为 小写，所以导致 json 包无法访问，所以无法正常解析
	v, err := json.Marshal(teacher)
	if err != nil {
		fmt.Println("JSON格式化出错！", err)
	} else {
		fmt.Println(v)                 // "[123 125]"
		fmt.Println(string(v))         // {},   // 将 []byte 转换为字符串
		fmt.Printf("%#v\n", string(v)) // "{}"
	}
}

// 结构体里的字段首字母为 大写， json 包正常访问
func test1() {
	var stu1 = Student{
		ID:     1,
		Gender: "男",
		Name:   "阳",
	}
	// 序列化: 将编程语言里面的数据 转换为 json 格式的数据
	v, err := json.Marshal(stu1)
	if err != nil {
		fmt.Println("JSON格式化出错！", err)
	} else {
		fmt.Println(v)                 // "[123 34 73 68 34 58 49 44 34 71 101 110 100 101 114 34 58 34 231 148 183 34 44 34 78 97 109 101 34 58 34 233 152 179 34 125]"
		fmt.Println(string(v))         // {"ID":1,"Gender":"男","Name":"阳"},   // 将 []byte 转换为字符串
		fmt.Printf("%#v\n", string(v)) // "{\"ID\":1,\"Gender\":\"男\",\"Name\":\"阳\"}"
	}

	// 反序列化：
	var stuInfo = &Student{}
	str := string(v)
	// 反序列化
	json.Unmarshal([]byte(str), stuInfo) //
	fmt.Println(*stuInfo)                // "{1 男 阳}"
	fmt.Printf("%T\n", stuInfo)          // "*main.Student"
}
