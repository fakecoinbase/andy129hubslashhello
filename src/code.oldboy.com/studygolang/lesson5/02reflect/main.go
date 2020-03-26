package main

import (
	"fmt"
	"reflect"
)

type myinterface interface {
	walk()
}

type human struct {
}

type student struct {
}

// reflect (反射)
// 反射的应用：  各种web框架、配置文件解析库， ORM 框架
/*  总结：
1, map, 切片（[]int, []string等）, 指针类型， 通过 reflect.TypeOf(x) 获取的 type,
	type.Name() 为空 ，其它类型则有相应的名字 (基本数据类型，结构体等)
2, 注意 引用类型与基本数据类型，Name() 与  Kind() 的区别
3, 特别注意 [3]int{} 与 []int{} 的区别
4, reflect.TypeOf() 可能会返回一个 nil , 所以注意 空指针报错
5, 查看 reflect 中 对 kind 的定义 (注意这里 kind 的定义 与 Kind() 返回的值是 不一样的，虽然类型，但是两个值的类型不一致)

*/
func main() {
	// test1()
	test2()
	//test3()
}

//reflect.TypeOf()
func test1() {

	// TypeOf(i interface{}) reflect.Type  可以接收任意类型的变量，返回一个 reflect.Type  类型

	var x = 100
	t := reflect.TypeOf(x)                   // 传入任意类型的变量，可返回 这个类型具体是 什么类型
	fmt.Printf("t : %v , type : %T\n", t, t) // "v : int , type : *reflect.rtype"

	// map, 切片（[]int, []string等）, 指针类型， 通过 reflect.TypeOf(x) 获取的 type,
	// type.Name() 为空 ，其它类型则有相应的名字 (基本数据类型，结构体等)

	reflectType(map[string]int{}) // "type : map[string]int"
	// type : map[string]int
	// typeName :  , typeKind : map

	dd := 100
	ptr := &dd
	reflectType(ptr)
	// type : *int
	// typeName :  , typeKind : ptr

	reflectType([]int{}) // "type : []int"
	// type : []int
	// typeName :  , typeKind : slice

	reflectType([3]int{1, 2, 3})
	// type : [3]int
	// typeName :  , typeKind : array

	reflectType("沙河") // "type : string"
	// type : string
	// typeName : string , typeKind : string

	reflectType(false) // "type : bool"
	// type : bool
	// typeName : bool , typeKind : bool

	reflectType(human{}) // "main.human"
	// type : main.human
	// typeName : human , typeKind : struct

	reflectType(student{}) // "main.student"
	// type : main.student
	// typeName : student , typeKind : struct

	var mi myinterface
	fmt.Println(mi) // "<nil>"
	reflectType(mi) // "type : <nil>"

	var xx interface{}
	reflectType(xx) //  "type : <nil>"

}

//reflect.TypeOf()
// fmt.Prinf("%T\n", x)   // %T 实现原理就是用的 反射
func reflectType(x interface{}) { // 传入进来的类型，必须是实例化完全的
	t := reflect.TypeOf(x) // 当 x 为 nil 时，则返回 t 为 nil
	fmt.Printf("type : %v\n", t)
	if t != nil { // TypeOf() 返回值 必须手动进行  t != nil 判断
		fmt.Printf("typeName : %v , typeKind : %v\n", t.Name(), t.Kind())
	}

}

/*  reflect 包中，关于 kind 的定义：

type Kind uint
const (
    Invalid Kind = iota  // 非法类型
    Bool                 // 布尔型
    Int                  // 有符号整型
    Int8                 // 有符号8位整型
    Int16                // 有符号16位整型
    Int32                // 有符号32位整型
    Int64                // 有符号64位整型
    Uint                 // 无符号整型
    Uint8                // 无符号8位整型
    Uint16               // 无符号16位整型
    Uint32               // 无符号32位整型
    Uint64               // 无符号64位整型
    Uintptr              // 指针
    Float32              // 单精度浮点数
    Float64              // 双精度浮点数
    Complex64            // 64位复数类型
    Complex128           // 128位复数类型
    Array                // 数组
    Chan                 // 通道
    Func                 // 函数
    Interface            // 接口
    Map                  // 映射
    Ptr                  // 指针
    Slice                // 切片
    String               // 字符串
    Struct               // 结构体
    UnsafePointer        // 底层指针
)

*/

// reflect.ValueOf()
func test2() {
	var x interface{}

	reflectValue(x)       // "<invalid reflect.Value>"
	reflectValue(human{}) // "{}"

	reflectValue(100)

	var a int32 = 300
	reflectValue(a)

	var b int64 = 4345235
	reflectValue(b)

	var f float32 = 3.4
	reflectValue(f)

	reflectValue([]int{1, 2, 3})

	fmt.Println("-------------------------------------------")

	// // IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic

	testValueIsNil([]int{}) // "IsNil :  false"
	// testValueIsNil(100)     // 运行报错："panic: reflect: call of reflect.Value.IsNil on int Value"

	// testValueIsNil(human{}) // 运行报错： panic: reflect: call of reflect.Value.IsNil on struct Value
	// testValueIsNil(x) // 传入一个空接口， 运行报错：panic: reflect: call of reflect.Value.IsNil on zero Value

	testValueIsNil(reflectValue) // "IsNil :  false"    // 传入一个 函数名

	testValueIsNil(emptyFunc)   // "IsNil :  false"     // 传入一个函数名 (函数里面的实现为空)
	testValueIsValid(emptyFunc) // "IsValid :  true"

	s := make([]string, 0, 8)
	fmt.Println(s)        // "[]"
	fmt.Println(s == nil) // "false"
	testValueIsNil(s)     // "IsNil :  false"  // 传入一个 slice (slice 里面为空字符)
	testValueIsValid(s)   // "IsValid :  true"  // 这个？？

	var m map[string]int

	fmt.Println(m)        // "map[]"
	fmt.Println(m == nil) // "true"
	testValueIsNil(m)     // "true"   // 传入一个 map 变量 (只声明未初始化)
	testValueIsValid(m)   // "IsValid :  true"   ????  这是为什么呢

	// 就算是传入的 map 没有初始化, 甚至  == nil ， 通过反射依旧能取到 它是什么类型，

	var p *int
	// 空指针，通过反射打印 v 和 k :   <nil> ptr
	testValueIsNil(p)   // "IsNil :  true"
	testValueIsValid(p) // "IsValid :  true"

	// IsNil() 常被用于判断指针是否为空；  IsValid() 常被用于判断返回值是否有效

	// *int 类型空指针
	fmt.Println("var p *int IsNil : ", reflect.ValueOf(p).IsNil())
	// nil 值
	fmt.Println("nil IsValid : ", reflect.ValueOf(nil).IsValid()) // "nil IsValid :  false"
	// 实例化一个匿名结构体
	st := struct{}{}
	fmt.Println("不存在的结构体成员：", reflect.ValueOf(st).FieldByName("abc").IsValid())  // "不存在的结构体成员： false"
	fmt.Println("不存在的结构体方法：", reflect.ValueOf(st).MethodByName("abc").IsValid()) // "不存在的结构体方法： false"

	// map
	ma := map[string]int{}
	ma["阳"] = 12
	fmt.Println("map 中不存在的键: ", reflect.ValueOf(ma).MapIndex(reflect.ValueOf("娜扎")).IsValid()) // map 中不存在的键:  false
	fmt.Println("map 中不存在的键: ", reflect.ValueOf(ma).MapIndex(reflect.ValueOf("阳")).IsValid())  // map 中不存在的键:  true
}

// 测试空函数
func emptyFunc() {}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x) // 获取接口的值信息
	fmt.Println("x 的value 为 : ", v)
	/*  注意这里的 v 是 reflect.Value 类型，不能与 nil 比较
	if v != nil {

	}
	*/
	// v.IsNil() // 使用要注意，具体介绍看下面 testValueIsNil() 测试 nil 函数， 以及下面的注释部分

	k := v.Kind() // 拿到值对应的种类

	switch k {

	case reflect.Int:
		fmt.Printf("(默认输出：)x 的值是 : %v \n", v)           // "x 的值是 : 100" (注意这里不能直接用  %d 输出 v, 因为 v 是 reflect.Value 类型，不是 int 类型)
		fmt.Printf("(强转输出：)x 的值是 ：%d \n", int(v.Int())) // "100"   // int(Type)  ---> int((reflect.Value).Int())
	case reflect.Int32:
		fmt.Printf("(默认输出：)x 的值是 : %v \n", v)             // "x 的值是 : 300" (注意这里不能直接用  %d 输出 v, 因为 v 是 reflect.Value 类型，不是 int 类型)
		fmt.Printf("(强转输出：)x 的值是 ：%d \n", int32(v.Int())) // "300"  // int32(Type)  ---> int32((reflect.Value).Int())
	case reflect.Int64:
		fmt.Printf("(默认输出：)x 的值是 : %v \n", v)             // "x 的值是 : 4345235" (注意这里不能直接用  %d 输出 v, 因为 v 是 reflect.Value 类型，不是 int 类型)
		fmt.Printf("(强转输出：)x 的值是 ：%d \n", int64(v.Int())) // "4345235"   // int64(Type)  ---> int64((reflect.Value).Int())
	case reflect.Float32:
		fmt.Printf("(默认输出：)x 的值是 : %v \n", v)                 // "x 的值是 : 3.4" (注意这里不能直接用  %f 输出 v, 因为 v 是 reflect.Value 类型，不是 float 类型)
		fmt.Printf("(强转输出：)x 的值是 ：%f \n", float32(v.Float())) // 3.400000"   //  float32(Type)  ---> float32((reflect.Value).Float())
	case reflect.Slice:
		fmt.Printf("(默认输出：)x 的值是 : %v \n", v) // "x 的值是 : [1 2 3]"
	default:

	}

}

// IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic
func testValueIsNil(x interface{}) {
	v := reflect.ValueOf(x) // 获取接口的值信息
	k := v.Kind()
	fmt.Println(v, k)
	fmt.Println("IsNil : ", v.IsNil()) // 判断v 里面保存的值是否为 nil
}

func testValueIsValid(x interface{}) {
	v := reflect.ValueOf(x)                // 获取接口的值信息
	fmt.Println("IsValid : ", v.IsValid()) // 判断v 里面保存的是否有 值
}

// 关于  v.IsNil() 源码说明:
// 1, 返回传入类型的参数或者 nil
// 2, 参数必须接收 chan, func, interface, map , pointer, or slice value (只有这些类型 才能使用 .IsNil() ，否则抛出  panic)
/*
	// IsNil reports whether its argument v is nil. The argument must be
	// a chan, func, interface, map, pointer, or slice value; if it is
	// not, IsNil panics. Note that IsNil is not always equivalent to a
	// regular comparison with nil in Go. For example, if v was created
	// by calling ValueOf with an uninitialized interface variable i,
	// i==nil will be true but v.IsNil will panic as v will be the zero
	// Value.
	func (v Value) IsNil() bool {

*/

// 通过反射 修改传入参数的值
func test3() {
	var a int64 = 100
	modifyValue(&a)

	fmt.Println(a) // "3453"
}

func modifyValue(x interface{}) {
	v := reflect.ValueOf(x) // reflect.Value
	k := v.Kind()
	fmt.Println(k) // "ptr"   // 指针类型

	if k == reflect.Ptr {
		// v.SetFloat(34.54) // 反射里面的方法 执行在代码执行阶段，所以编译不会报错
		// 运行报错： panic: reflect: reflect.flag.mustBeAssignable using unaddressable value
		// 说明 v 不能直接进行设置值的操作

		// Elem() 用法
		// v.Elem().SetFloat(12344.3) // 运行又报错了，这是为什么?
		// 运行报错： panic: reflect: call of reflect.Value.SetFloat on int64 Value
		// 发现报错信息与上面不一样， 而是 你设置的值与 之前 传入的值 类型不匹配

		v.Elem().SetInt(3453) // 运行正常，到外面打印结果是 ： 3453 ,  修改成功
	}

}
