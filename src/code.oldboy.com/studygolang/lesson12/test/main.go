package main

import (
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"
)

// MyInterface 是一个接口
// 定义接口的方法时，参数与返回值可以定义名称，例如：Get()，也可以省略名称 直接指定类型，例如：Set()
type MyInterface interface {
	Get(key string) (value string)
	Set(string, string)
}

// 定义一个全局接口类型
var myInterface MyInterface

// User 是一个用户信息结构体
type User struct {
	name     string
	password string
}

// Get 获取 User 用户名
func (u *User) Get(key string) (value string) {
	if key == "name" {
		return u.name
	}
	return
}

// Set 设置用户字段的值
func (u *User) Set(key string, value string) {
	if key == "name" {
		u.name = value
	}
}

//
func main() {

	// testMapDelete()
	// testMap2()
	// testSlice()
	// testUUID()
	// testMyInterface()
}

// time.Second * 100 , 语法正确
// time.Second * varTmp  ，将100 保存在 int 类型的变量里， 这样就不行
func testTimeCalc() {
	dur := time.Second * 100
	fmt.Println(dur)

	/*
		var double = 2
		// dur2 := time.Second * double   // 编译报错：invalid operation: mismatched types time.Duration and int
	*/

	// 对比 dur 和 dur2 的计算
}

// 测试可变长参数使用
func testParams() {

	// loadServer()   // 至少指定一个参数
	loadServer("127.0.0.1")
	loadServer("127.0.0.1", "pwd:12345")

	// 可变长参数，不传入也不会报错
	loadInfo()
	loadInfo("1234")
	loadInfo("1234", "xxxx")
}

// 一个固定参数 和 一个变长参数
func loadServer(addr string, params ...string) {

}

// 只有一个可变长参数
func loadInfo(params ...string) {

}

// 可变长参数在前， 固定参数在后，  ---》 编译报错：can only use ... with final parameter in list    (可变参数必须放在 参数列表的最后)
/*
func loadMsg(params ...string, maxSize int) {

}
*/

// 测试接口定义方法
func testMyInterface() {

	// myInterface 可以接收的类型，必须是 实现了这个接口定义所有方法的 类型，否则无法类型转换
	myInterface = &User{}
	myInterface.Set("name", "yang")

	fmt.Println(myInterface.Get("name"))

}

func testHttpCookie() {
	// http.SetCookie()
}

func testUUID() (err error) {
	// uuid, err := uuid.NewV4()
	uuid, err := uuid.NewV1()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(uuid)

	return nil
}

// 测试修改 Map1
func testMapDelete() {
	m := make(map[string]string, 8)

	m["name"] = "yang"
	m["age"] = "18"
	m["class"] = "三年二班"

	delete(m, "xxx") // 删除一个不存在的 key ,不会报错
	delete(m, "age") // 删除一个存在的key， 则会将 key,value 这一对都删除
	// map[string]string{"class":"三年二班", "name":"yang"}

	fmt.Printf("%#v\n", m)

	modifyMap(m)
	fmt.Printf("%#v\n", m)
}

// 这里的 m 是对 外面的 map 进行的值拷贝，让它们都指向同一块内存区域，所以 这里对 m 的修改可以同步到 外面的 map
func modifyMap(m map[string]string) {
	m["class"] = "六年一班"
}

// 测试修改Map2
func testMap2() {
	var m map[string]string

	fmt.Printf("%#v\n", m) // "map[string]string(nil)"

	modifyMap2(m)

	fmt.Printf("%#v\n", m) // "map[string]string(nil)"
}

// 由于外层 map 是一个 nil ，并没有申请内存， 所以当做参数传入下面这个方法时，
// 期初 m 也是nil ,只是 在里面进行了单独的内存申请，就变成了一个执行内存的类型，所以针对这块内存进行的修改，不会同步到外面 nil
func modifyMap2(m map[string]string) {
	m = make(map[string]string)
	m["name"] = "阳浩"
	fmt.Printf("---%#v\n", m) // "---map[string]string{"name":"阳浩"}"
}

// 测试 slice
func testSlice() {
	s := make([]int, 4, 8)
	fmt.Printf("%p\n", s)
	s[0] = 1
	fmt.Println(s) // "[1 0 0 0]"

	modifySlice(s)
	fmt.Println(s) // "[1 1 0 0]"
}

// 匪夷所思，为什么 append 之后的没有反馈到 外面的 s
// 详细分析请看： https://www.cnblogs.com/snowInPluto/p/7477365.html
// 大致意思是， 外面的 slice 有4 个元素，当做参数传入下面函数之后， 进行地址的值拷贝，那么 s 就与 外面的 slice 共同指向同一块空间
// s[1] 针对共同的空间进行修改，那么会同步到外面， 但是进行 append 之后, s 的长度变了，就指向了 6个元素，外面的 slice 指向的空间依旧是 4个
func modifySlice(s []int) {
	s[1] = 1
	//fmt.Printf("%p\n", s)
	s = append(s, 2)
	s = append(s, 3)

	s[3] = 100     // s 只要是修改 外面的slice 包含长度的值，都是可以同步的
	fmt.Println(s) // "[1 1 0 0 2 3]"
	//fmt.Printf("%p\n", s)
}
