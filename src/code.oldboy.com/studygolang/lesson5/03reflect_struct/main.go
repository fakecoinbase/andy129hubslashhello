package main

import (
	"fmt"
	"reflect"
)

// animal 是一个结构体
type animal struct {
	name  string `json:"n" xx:"na"`
	count int    `json:"c" xx:"co"`
}

// Student 是一个结构体
type Student struct {
	name   string
	age    int
	gender string
}

func (a *animal) Move() {
	fmt.Println("move")
}

func (a *animal) Loud(str string) {
	fmt.Println("lound : ", str)
}

// Walk 是一个方法
func (s Student) Walk() {
	fmt.Println("walk")
}

// Speak 是一个方法
func (s Student) Speak(str string) string {
	fmt.Println("speak : ", str)
	return "人类说话"
}

/* 待研究， go 语言里面是否有 java 中的 toString() 方法
func (a *animal) String() string {
	// fmt.Println("animal toString方法")
	return "animal toString方法"
}
*/

// 结构体反射
// 总结：
// 1, 取字段的名称，不管结构体是否大小写,结构体内的 字段是否大小写 都能正常获得 (go 版本： go version go1.13.6 windows/amd64)
// 注意注意：设置字段的值， ValueOf(),   里面的字段的首字母必须要大写
//  目前来看，结构体的名字 大小写无所谓。
// 2, 取方法名，结构体中的方法名首字母一定要大写。否则获取的方法名数量为 0 ，   t.NumMethod() == 0 , v.NumMethod()== 0
// 3, 通过反射操作时，需要进行 nil 判断，如果是反射结构体则 还需判断是否为结构体类型，
//    传入指针类型的结构体或者 接口时，注意 t.Elem(),  v.Elem()
//    获取它们包含的值或者指向的值。否则后续操作 (例如获取 字段数目，获取字段名称等操作)容易 panic
// 4, 通过反射 TypeOf() 获得字段的名字，字段的类型，通过判断类型，我们可以 ValueOf() 获得这个字段的值
// 5, 通过反射不仅可以获得字段的值，还能修改字段的值， 例如: t.Elem().SetInt(23)
// 6, 通过反射结构体 可以调用结构体里的方法，  注意传参和接收方法返回值。

// 反射的优缺点：
// 反射是一个强大并富有表现力的工具，能让我们写出更灵活的代码。反射常用于各种框架模块。 但是反射不应该被滥用，原因有以下三个：
// 1, 基于反射的代码是极其脆弱的，反射中的类型错误会在真正运行的时候才会引发 panic, 那很可能是在代码写完的很长时间之后。
// 2, 大量使用反射的代码通常难以理解
// 3, 反射的性能低下，基于反射实现的代码通常比正常代码运行速度慢 一 到两个数量级。

func main() {

	// test1()
	// test2()
	// test3()
	test4()

}

func test1() {

	a := animal{
		name:  "猫",
		count: 15,
	}

	h := Student{
		name:   "学生",
		gender: "男",
		age:    18,
	}

	printField(a)
	printField(h)

	// 传入空的接口，看看会怎么样
	/*
		var xx interface{}
		printFunc(xx) // "TypeOf 返回值为 : nil"
	*/
}

// animal , Human 两个结构体里面的 字段都是小写， 依然可以使用反射拿到它们
func printField(x interface{}) {

	t := reflect.TypeOf(x)

	if t == nil {
		// 空类型，返回
		fmt.Println("TypeOf 返回值为 : nil")
		return
	}
	fmt.Printf("类型名：%s , 类型种类：%s\n", t.Name(), t.Kind())

	/*  方法介绍：
	//NumField() int
	t.NumField() // 返回字段数量

	// Field(i int) reflect.StructField
	t.Field()
		// reflect.StructField 里面的字段：
		structField.Name   // 字段的名字
		structField.Index  // 字段位于字段列表中第几个 (Index 是[]int 类型，取值时 是以 0为开始)
		structField.Type   // 字段类型
		structField.Tag    // 字段的 tag  (可获取字段的 tag 信息)
			// StructTag   ---->  示例：json:"n" xx:"na"
				structField.Tag.Get("json")  --->"n"
	*/

	// 通过索引获取 字段信息
	for i := 0; i < t.NumField(); i++ {
		structField := t.Field(i)
		fmt.Printf("name : %s, index : %d, type : %v, tag : %v\n", structField.Name, structField.Index, structField.Type, structField.Tag)
	}

	/* 打印如下：

	类型名：animal , 类型种类：struct
	name : name, index : [0], type : string, tag : json:"n" xx:"na"
	name : count, index : [1], type : int, tag : json:"c" xx:"co"
	类型名：Student , 类型种类：struct
	name : name, index : [0], type : string, tag :
	name : age, index : [1], type : int, tag :
	name : gender, index : [2], type : string, tag :

	*/

	fmt.Println("--------------通过指定 字段名称 获取字段信息------------")

	// 通过指定 字段名称 获取字段信息：
	st, ok := t.FieldByName("name")
	// 经过测试， animal , Student 两个结构体中的 name 都能分别打印出来，
	// 这说明，不需要特别指定结构体或者 字段名为 大写，难道是 go 1.13版本改进了?
	if !ok {
		fmt.Println("获取不到 gender 字段")
		return
	} else {
		fmt.Printf("name : %s, index : %d, type : %v, tag : %v\n", st.Name, st.Index, st.Type, st.Tag)
	}

}

/*  go 语言中  StructField 的定义：

// A StructField describes a single field in a struct.
type StructField struct {
	// Name is the field name.
	Name string
	// PkgPath is the package path that qualifies a lower case (unexported)
	// field name. It is empty for upper case (exported) field names.
	// See https://golang.org/ref/spec#Uniqueness_of_identifiers
	PkgPath string

	Type      Type      // field type
	Tag       StructTag // field tag string
	Offset    uintptr   // offset within struct, in bytes
	Index     []int     // index sequence for Type.FieldByIndex
	Anonymous bool      // is an embedded field
}

*/

func test2() {

	a := &animal{
		name:  "猫",
		count: 15,
	}

	h := Student{
		name:   "学生",
		gender: "男",
		age:    18,
	}

	printMethod(a)
	printMethod(h)
}

// 获取结构体中方法的名称，前提: 1, 方法名首字母要大写。 2，方法名中指定接收者 是 指针类型还是 值类型，如果想通过反射获取方法名，那就必须传入一个 指针类型的结构体
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	//fmt.Println(t.NumMethod())   // 获取结构体中方法有多少个
	//fmt.Println(v.NumMethod())

	for i := 0; i < t.NumMethod(); i++ {
		methodType := v.Method(i).Type() // 拿到方法的类型
		fmt.Printf("method name : %s, method : %s\n", t.Method(i).Name, methodType)
	}

	// 打印信息如下：
	/*
		method name : Loud, method : func(string)
		method name : Move, method : func()
		method name : Speak, method : func(string)
		method name : Walk, method : func()

	*/
}

// 通过反射调用方法
func test3() {
	a := &animal{
		name:  "猫",
		count: 15,
	}

	h := Student{
		name:   "学生",
		gender: "男",
		age:    18,
	}

	runMethodByReflect(a, "Move")
	runMethodByReflect(h, "Walk")
}

// 通过反射调用方法 （传参与接收返回值）
func runMethodByReflect(x interface{}, methodStr string) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	//fmt.Println(t.NumMethod())   // 获取结构体中方法有多少个
	//fmt.Println(v.NumMethod())

	for i := 0; i < t.NumMethod(); i++ {
		// methodType := v.Method(i).Type() // 拿到方法的类型
		// fmt.Printf("method name : %s, method : %s\n", t.Method(i).Name, methodType)

		methodName := t.Method(i).Name

		// 通过反射 调用有参数及有返回值的方法
		if methodName == "Speak" {

			//var args = []reflect.Value{cv}

			args := make([]reflect.Value, 1) //参数
			args[0] = reflect.ValueOf("哈哈哈")

			rs := v.Method(i).Call(args) // rs 为调用方法之后的返回值,  返回值为 []value

			fmt.Println("返回值：", rs[0].Interface().(string))

			return
		}

		// 通过反射 调用方法 (无需传入参数，也不需要接受返回值)
		if methodName == methodStr {
			var args = []reflect.Value{}
			v.Method(i).Call(args)
		} else {
			fmt.Println("找不到要调用的方法")
		}
		// 通过反射调用方法，  参数必须是 []reflect.Value 类型

	}
}

// 测试获取结构体中的字段的值
func test4() {

	a := &animal{
		name:  "猫",
		count: 15,
	}
	getFieldValue(a)

	h := Student{
		name:   "学生",
		gender: "男",
		age:    18,
	}

	getFieldValue(h)
}

// 获取结构体中的字段，并根据字段的类型 获得字段的值
func getFieldValue(x interface{}) {

	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	if t == nil {
		// 空类型，返回
		fmt.Println("TypeOf 返回值为 : nil")
		return
	}
	// 当传入的是一个指针类型时，我们不能直接操作 t, v ，而是需要 t.Elem(),  v.Elem() 获得一个指针指向的值
	// 翻译： t.Elem() 会返回 t 这个接口包含的值 或者  t 这个指针指向的值， 如果 t 的种类不是 interface 或者 ptr ,则就会 panic, 如果 t is nil, 则返回 zero value
	/*   go 语言源码中对于 Elem() 的说明：

	// Elem returns the value that the interface v contains
	// or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr.
	// It returns the zero Value if v is nil.
	func (v Value) Elem() Value {

	*/
	if v.Kind() == reflect.Ptr {
		fmt.Println("-----------reflect.Ptr")
		t = t.Elem()
		v = v.Elem()
	}
	if t.Kind() != reflect.Struct {
		fmt.Println("不是结构体，不能进行反射字段和方法的操作")
		return
	}
	fmt.Printf("类型名：%s , 类型种类：%s\n", t.Name(), t.Kind())

	// 通过索引获取 字段信息
	for i := 0; i < t.NumField(); i++ {
		structField := t.Field(i)
		// fmt.Printf("name : %s, index : %d, type : %v, tag : %v\n", structField.Name, structField.Index, structField.Type, structField.Tag)

		fmt.Println("---------------------------------------------------")
		// fieldValue := v.FieldByName(structField.Name)   // 通过字段名称拿到字段的值
		fieldValue := v.Field(i) // 通过索引拿到 字段的值
		fieldValueKind := fieldValue.Kind()
		fmt.Printf("fieldValue : %v, fieldValueKind : %v\n", fieldValue, fieldValueKind)

		if fieldValueKind == reflect.Int {
			fmt.Printf("fieldName : %s , fieldValue : %d\n", structField.Name, fieldValue.Int())
		}
		if fieldValueKind == reflect.String {
			fmt.Printf("fieldName : %s , fieldValue : %s\n", structField.Name, fieldValue.String())
		}
		fmt.Println("---------------------------------------------------")
	}

	// 打印信息如下：

	/*  传入一个 animal 的指针类型时，打印如下：：

	-----------reflect.Ptr
	类型名：animal , 类型种类：struct
	---------------------------------------------------
	fieldValue : 猫, fieldValueKind : string
	fieldName : name , fieldValue : 猫
	---------------------------------------------------
	---------------------------------------------------
	fieldValue : 15, fieldValueKind : int
	fieldName : count , fieldValue : 15

	*/

	/*  传入一个 Student 结构体值类型时，打印如下：

	类型名：Student , 类型种类：struct
	---------------------------------------------------
	fieldValue : 学生, fieldValueKind : string
	fieldName : name , fieldValue : 学生
	---------------------------------------------------
	---------------------------------------------------
	fieldValue : 18, fieldValueKind : int
	fieldName : age , fieldValue : 18
	---------------------------------------------------
	---------------------------------------------------
	fieldValue : 男, fieldValueKind : string
	fieldName : gender , fieldValue : 男
	*/

}
