package main

import "fmt"

// 函数
// go 语言的函数中没有默认的参数
func main() {

	a, b, s, ok := testReturn()
	fmt.Println(a, b, s, ok) // "0 0 r false"

	//testJoin(a, b, s, ok)
	// 函数参数只有可变参数时，可以省略不传入
	fmt.Println(intSum())       // "0"
	fmt.Println(intSum(10, 20)) // "30"

	// 函数参数包括 固定参数和可变参数时，必须至少要指定一个 固定参数
	// fmt.Println(intSum2())   // 编译不通过
	fmt.Println(intSum2(10))         // 参数：a = 10,  b=[]
	fmt.Println(intSum2(10, 20))     // 参数：a=10, b=[20]
	fmt.Println(intSum2(10, 20, 30)) // 参数：a=10, b=[20,30]

	fmt.Println("------------------------------------------------------")
	sum, sub := testReturn3(10, 20)
	fmt.Println(sum, sub) // "30 -10"

	// testReturn3， 简写了return , 所以按照 函数指定的返回值去返回
	// testReturn4, 函数虽然指定了返回值的名称，但是 return 后面的变量也没有省略，最终会选择哪个返回呢？
	// 看 testReturn4 的结果就明白了，优先级: return sum,sub  >   func ...  (sum,sub)
	sum2, sub2 := testReturn4(10, 20)
	fmt.Println(sum2, sub2) // "-10 30"
}

// 多个返回值，必须要用 ()
func testReturn() (int, int, string, bool) {

	return 0, 0, "r", false
}

// 不带返回值的 return
func testReturn2(a int, b int) (ret int) { // 除了指定返回值的类型为 int 外，还指定了返回值的变量名为 ： ret ，所以它会从函数里面找这个变量名并返回出去

	ret = a + b
	return //  return 不能省略， 只能省略后面的变量名
}

// 返回值类型的简写
func testReturn3(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

// 既在返回值括号里面声明返回变量 ， return 也没有省略变量名称，是否会造成冲突？ 看上面的调用结果
func testReturn4(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return sub, sum
}

// 传入多个参数
func testJoin(a int, b int, s string, ok bool) {

	fmt.Println(a, b, s, ok) // "0 0 r false"
}

// 参数类型的简写
func testJoin2(a, b int, s string, ok bool) {
	fmt.Println(a, b, s, ok) // "0 0 r false
}

// 接收可变参数的函数， 在参数名后面加 ... 表示可变参数
// 可变参数在函数中是 切片类型
func intSum(a ...int) int {
	//fmt.Println(a)        // "[10 20]"
	//fmt.Printf("%T\n", a) // "[]int"

	ret := 0
	for _, v := range a {
		ret = ret + v
	}
	return ret
}

// 固定参数和可变参数同时出现时，可变参数要放在最后
func intSum2(a int, b ...int) int {
	for _, v := range b {
		a = a + v
	}
	return a
}
