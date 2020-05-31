package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
)

type People interface {
	Speak(string) string
}

type Student struct{
	Name string
}

// Student 的指针类型实现了接口
func (stu *Student) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

type User struct {
	Name string
	Age int
}

type StudentFunc func(*Student)

// -----------------
//定义Integer类型
type Integer int

type LessAddInf interface{
	Less(n Integer) bool
	Add(n Integer) Integer
}

func (this Integer) Less(n Integer) bool{
	return this < n
}

func (this *Integer) Add(n Integer) Integer{
	*this += n
	return *this
}

type Computer struct{
	CPU string "计算器"
	Memory string "内存"
}

type Thing interface{
	Name() string
	Attribute() string
}

func (this Computer) Name() string  {
	return "Computer"
}

func (this *Computer) Attribute()string  {
	return fmt.Sprintf("CPU=%v Memory=%v", this.CPU, this.Memory)
}

// -------------------------


// -----------------------------

type S1 struct {

}

func (s1 S1) f(){
	fmt.Println("S1.f()")
}

func (s1 S1) g() {
	fmt.Println("S1.g()")
}

type S2 struct {
	S1
}

func (s2 S2)f() {
	fmt.Println("S2.f()")
}

type I interface {
	f()
}

func printType(i I){
	if s1,ok := i.(S1);ok {
		s1.f()
		s1.g()
	}
	if s2, ok:= i.(S2);ok {
		s2.f()
		s2.g()   // 由于 s2 没有实现自己的 g() 方法，但 又由于 s2 嵌套的 s1, 则会调用 s1 的g() 方法
	}
}

// ----------------------------

// A是B的子集
type A interface {
	data1(int)
}

type B interface {
	data1(int)
	data2(int)
}

type Impl_B struct {

}

func (i Impl_B)data1(a int)  {
	fmt.Println("Impl_B data1 : ", a)
	return
}

func (i Impl_B)data2(b int)  {
	fmt.Println("Impl_B data2 : ", b)
	return
}

type Impl_A struct {

}

func (i Impl_A)data1(a int)  {
	fmt.Println("Impl_A data1 : ", a)
	return
}

// ----------------------------

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}

func (s Slice) Add(elem int) *Slice {
	s = append(s, elem)
	fmt.Print(elem)
	return &s
}

// ----------------------------

// 面试题
func main() {
	// 考核接口的用法
	// testInterface()

	// 考核 for range 的用法
	// testForRange()

	// testPtr()
	// testPtr2()

	// testNewAndMake()

	//testCalcTag()

	// testCalcTag2()

	// testString()

	//testFor()

	//testChan()

	// testVar()

	// testArgs()

	// testSwitch()

	// testInterface2()

	// testInterface3()

	// testInterface4()

	// testInterface5()

	// testInterface6()

	// testInterface7()

	// testMap()

	// testMap2()

	// testPrint()

	// testForMap()

	// testFunc()

	// testFunc2()

	// testFunc3ForMap()

	// testFunc4ForMap()

	// testSlice()

	// testSlice2()

	testDefer()     // 考核点： return 之后 defer 的执行顺序

	// testSlice3()     // 考核点：defer 语句的执行顺序
	// testSlice4()   // 考核点：defer 语句的执行顺序

	// testJson()

	// testBool()

	// testChan2()

	// testPanic()
}

// 考核点： panic 与 defer
// 首先：defer 的设计原则就是 在函数退出(包括 panic)之后，依次运行 (进行一些关闭连接，释放资源的操作)
// 但是 panic 之后的 defer 就不会执行，这也说明了 panic 的优先级高于 defer
func testPanic(){
	defer fmt.Println(1)
	defer fmt.Println(2)
	panic("abc")
	defer fmt.Println(3)

	/*  运行结果：
		2
		1
		panic: abc

		... (异常信息)

	 */
}

// 考核点：通道的赋值操作
func testChan2(){
	var c chan int
	fmt.Println(c)   // nil

	// c <- 2   // 给一个 nil channel 发送数据，造成永远阻塞并报错
	// 报错信息： fatal error: all goroutines are asleep - deadlock!
}

// 考核点：bool 类型的编码规范
func testBool(){
	flag := false

	// if flag == 1 {}  // 1 是 int 类型，不能与 bool 类型比较
	if flag == false {   // 虽然正常执行，但不符合编码规范，要采用 !flag 的方式
		fmt.Println("flag == false")
	}

	if !flag {
		fmt.Println("!flag")
	}
}

// 考核点：go 语言中可以被json 序列化的类型
// golang中大多数数据类型都可以转化为有效的JSON文本，channel, 函数 除外
func testJson(){

	// 指针可以被 json 序列化
	var num int = 10
	var ptr *int
	ptr = &num
	data, _ := json.Marshal(*ptr)
	fmt.Println(data)
	fmt.Println(string(data))   // 10

	// 通道不可以通过 json 序列化
	c := make(chan int ,3)
	c <- 2
	data2, _ := json.Marshal(c)
	fmt.Println(data2)
}

// 考核点：互斥锁
func testMutex(){
	/*  正确答案：ABC
		A. 当一个goroutine获得了Mutex后，其他goroutine就只能乖乖的等待，除非该goroutine释放这个Mutex
		B. RWMutex在读锁占用的情况下，会阻止写，但不阻止读
		C. RWMutex在写锁占用情况下，会阻止任何其他goroutine（无论读和写）进来，整个锁相当于由该goroutine独占
		D. Lock()操作需要保证有Unlock()或RUnlock()调用与之对应，否则编译错误
	 */


}

// 考核点： defer
func testDefer() {
	fmt.Println(f(3))
}

func f(n int) (r int) {

	defer  func() {
		r += n
		fmt.Println("+++++++r : ",r)
		recover()
		fmt.Println("@@@@@@@@r : ",r)
	}()

	var f func()

	fmt.Println(f)   // <nil>
	defer f()    // f 只是声明了一个 func() 类型的变量，并没有初始化，所以 defer f() 会报错 : panic: runtime error: invalid memory address or nil pointer dereference
	f = func() {  // 由于上一行代码报错，所以这个语句块不会被执行。
		r+=2
		fmt.Println("-----r : ",r)
	}

	return n+1
}


// 考核点：defer 语句的执行顺序
func testSlice4(){

	s := NewSlice()
	defer func(){    // 最后处理 一整段 func()
		s.Add(1).Add(2).Add(3).Add(4)
	}()

	s.Add(3)

	/*  执行顺序
		s.Add(3)

		s.Add(1).Add(2).Add(3).Add(4)

	打印结果：
		31234

	 */
}

// 考核点：defer 的压栈处理 原则  (对比 testSlice4() 查看)
func testSlice3(){
	s := NewSlice()
	defer s.Add(1).Add(2).Add(4)   // defer 最后执行 .Add(4),  但是前面依旧提前执行 (最后一个. 被压栈处理)

	s.Add(3)

	/* 所以它的执行顺序是：
		s.Add(1)   // 执行
		s.Add(2)   // 执行
		-- 暂不执行，s.Add(4) 被压栈到最后执行
		s.Add(3)   // 执行

		打印结果：
		1234

	 */
}

// 考核点：切片的 range 遍历
func testSlice2(){
	x := []string{"a", "b","c"}
	// for k,v := range x {  // range slice 返回两个参数，一个是 slice 的 index ,一个是 value
	for v := range x { // 如果只接受一个参数，则是 index
		fmt.Println(v)
		/*
			0
			1
			2

		 */
	}

	fmt.Println()

	m := make(map[int]int, 5)
	m[0]= 3
	m[2] = 4
	m[-1] = 7
	m[100] = 8
	// for k, v := range m {   // range map 返回两个参数，一个是 map 的key , 一个是 value
	for v := range m {   // 如果只接受一个参数，则是 key
		fmt.Println(v)
		/*
			0
			2
			-1
			100

		 */
	}
}

// 考核点：切片的初始化语法规范以及 比较
func testSlice(){

    s := []int{1,2,3,4,5}
	fmt.Println(s)
	fmt.Printf("len : %d, cap : %d\n", len(s), cap(s))

	s2 := []int{1,2,3,4,5,}    // 格式不报错，与 s 一样
	fmt.Println(s2)
	fmt.Printf("len : %d, cap : %d\n", len(s2), cap(s2))

	// fmt.Println(s == s2)   // 注意，slice 之间不能比较


	a := [...]int{1,2,3,4,5}

	b := [5]int{1,2,3,4,5}
	fmt.Println(a == b)    // true ,    a 和 b 是有指定长度的数组，所以可以比较


	// 对比x 与 x2 的初始化语法， 花括号内每一行结尾都要有 逗号， (例如：x),  除非 花括号在最后一行结尾 (例如：x2)
	x := []int{
		1,2,3,
		4,5,6,
	}

	x2 := []int{
		1,2,3,
		4,5,6}

	fmt.Println(x)
	fmt.Println(x2)
}

// 考核点：go 语言中 栈与堆
func testStackAndHeap(){

	// 描述 golang 中的 stack 和 heap 的区别，分别在什么情况下会分配 stack , 又在何时会分配到 heap 中
	// 资料参考：https://blog.csdn.net/u010853261/article/details/102846449

	// 通过命令查看 变量逃逸情况
	// 第一： 编译器命令
	/*
		可以看到详细的逃逸分析过程。而指令集 -gcflags 用于将标识参数传递给 Go 编译器，涉及如下：

		-m 会打印出逃逸分析的优化策略，实际上最多总共可以用 4 个 -m，但是信息量较大，一般用 1 个就可以了
		-l 会禁用函数内联，在这里禁用掉 inline 能更好的观察逃逸情况，减少干扰

		go build -gcflags "-m -l" main.go

	 */

	// 第二：反编译命令查看
	// go tool compile -S main.go



	/*
		区别：
			stack (栈): 一般存放变量名，比较小，  ---- 系统自动处理内存分配和释放 (栈中的分配的变量名其实也是有 地址的，方便于 系统管理， 变量名又保存着指向堆空间的地址)

			heap (堆): 一般存放具体数据，比较大， ---- 一般由程序员处理 (但是现在的高级语言都有 垃圾回收机制，所以golang 也会自动回收)
	 */


	/*  其实示例说明：

		func NewRect(x, y, width, height float64) *Rect {

		return &Rect{x, y, width, height}

		}

		注意，这里与C/C++不同的是，返回一个局部变量的地址在Go语言中是绝对没有问题的；变量关联的存储在函数返回之后依然存在。

		更直接的说，在Go语言中，如果一个局部变量在函数返回后仍然被使用，这个变量会从heap，而不是stack中分配内存。
	 */


}



// 考核点：go 程序初始化顺序
func testMainInit(){

	// go 程序执行，程序加载的顺序
	/*
	    main 包
	    import
	        --- 初始化引用包中的 全局 const ，全局 var , init() ...
		全局 const
	    全局 var
	    init()
	    main()
	 */
}


// 结合 testFunc4ForMap(), testFunc3ForMap(), testFunc2(),  testFunc() 中知道，
// 1, 闭包的传值要 显示传入，
// 2, 并且 并发针对 map 的写入操作(同一个key的写入)，必须加锁,
// 3, 并且要加入 等待组 WaitGroup，等待所有值写入完成
// 3, 并发读并没有出现问题 (目前测速)

// 考核点：闭包的传值
func testFunc(){

	m := make(map[int]int)

	wg := sync.WaitGroup{}

	mu := sync.Mutex{}

	wg.Add(10)
	for i:=0;i<10;i++ {
		go func(i int) {   // 必须显示定义传入 i
			defer wg.Done()
			mu.Lock()   // 加锁保证 map 的写入是安全的
			m[i] = i   // i 使用的是 func() 外部的变量，针对 闭包传值时，必须要显示传入 func(i int),   (i)
			mu.Unlock()
		}(i)   // 别忘了，这里传值
	}
	wg.Wait()
	fmt.Println(len(m))

}

// go 中 针对 map 的写入是 不安全的，必须要加锁
// 考核点：map 的并发写入问题
func testFunc2() {
	var wg sync.WaitGroup

	mu := sync.Mutex{}   // 加锁

	var m = make(map[int]int)
	wg.Add(20)
	for i := 0; i < 20; i++ {

		go func(n int) {
			defer wg.Done()
			mu.Lock()   // 加锁
			m[n] = n    // 此处针对 map 的修改，没有加锁，所以运行会报错误: fatal error: concurrent map writes
			mu.Unlock() // 解锁
		}(i)
	}
	wg.Wait()
}

// 考核点：并发针对map 的写入问题 (随机key ,会有隐患)
func testFunc3ForMap(){

	fmt.Println("testFunc3ForMap")
	var wg sync.WaitGroup
	mu := sync.Mutex{}   // 加锁

	var m = make(map[int]int)
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()   // 加锁
			m[rand.Int()] = rand.Int()    // 此处针对 map 的修改，没有加锁，可能会出现修改同一个key 的问题(rand 可能会随机到同一个key )，所以运行会报错误: fatal error: concurrent map writes
			mu.Unlock() // 解锁
		}()
	}
	wg.Wait()

	fmt.Println(len(m))
}

// 考核点：并发 针对 map 同一个key 的写入
func testFunc4ForMap(){
	fmt.Println("testFunc4ForMap")
	var wg sync.WaitGroup
	mu := sync.Mutex{}   // 加锁

	var m = make(map[int]int)
 	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()   // 加锁
			 m[2] = rand.Int()    // 此处针对 map 的修改，如果没有加锁，会出现由于修改同一个key 的问题，所以运行会报错误: fatal error: concurrent map writes
			// _ = m[2]   // map 并发 读没有出现问题
			mu.Unlock() // 解锁
		}()
	}
	wg.Wait()

	fmt.Println(len(m))
}


// 考核点：map中 value 为地址的赋值操作，以及 i++ 的使用
func testForMap(){
	m := make(map[int]*int)

	for i:=0;i<3;i++ {
		m[i] = &i
	}

	for k,v := range m {
		fmt.Println(k, *v)
	}
}

// 考核点： print 与 fmt.Print() 的区别
func testPrint(){
	// print 在golang中 是属于输出到标准错误流中并打印,官方不建议写程序时候用它。可以再debug时候用
	print("hello\n")   // print 是 go 语言中的 日志打印


	fmt.Println("----------")
	// fmt.print 在golang中 是属于标准输出流,一般使用它来进行屏幕输出.
	fmt.Print("hello fmt\n")
}

// 考核点： map 中结构体寻址赋值操作
func testMap2(){

	type S struct {
		name string
	}
	m := map[string]S{"x":S{"one"}}  // 看这里，这里的写法分为两步，第一步：定义 map[string]S ， 第二步：初始化 map  {"x", S{"one"}}

	// m["x"].name = "two"    // 语法错误，map 中寻址操作错误

	// 改为如下：
	// (m["x"]).name = "two"  // 也是语法错误，map 中寻址操作错误

	// 不是针对结构体中的字段赋值，则 取值完全没有问题
	v := m["x"].name
	fmt.Println("-----------v : ", v)


	// map 中的值若是结构体，则2无法直接寻址
	// 必须是 指针类型
	m2 := map[string]*S{"x":&S{"one"}}

	m2["x"].name = "two"    // 操作正确
	v2 := m2["x"].name
	fmt.Println("-----------v2 : ", v2)

}

func testMap(){
	m := make(map[string]string, 8)
	m["name"] = "zhang"

	fmt.Printf("传入前：m 的地址：%x\n", &m)   // 传入前：m 的地址：&map[6e616d65:7a68616e67]
	modify(m)



	fmt.Println(m["name"])   // "yang"
}

func modify(m map[string]string){

	fmt.Printf("传入后：m 的地址：%x\n", &m)   // 传入后：m 的地址：&map[6e616d65:7a68616e67]

	m["name"] = "yang"
}

// 考核点：接口的继承
func testInterface7(){
	var a A = Impl_B{}   // A 接口是 B 接口的子集，所以 A 能接收 B 的实现
	a.data1(10)

	fmt.Println(a)

	/*
	var b B = Impl_A{}   // 编译错误： B 接口却不能接收 A 的实现，因为 A 的实现缺少B 接口中的某些方法 (data2(i int))
	b.data1(20)
	b.data2(30)
	*/

}

// 考核点： interface 之间的比较 之 结构体指针
func testInterface6(){

	type S struct {
		a,b,c string
	}

	x := interface{} (&S{"a", "b", "c"})
	y := interface{} (&S{"a", "b", "c"})

	fmt.Println(x == y)  // false

	// 去掉指针类型，m == n
	m := interface{} (S{"a", "b", "c"})
	n := interface{} (S{"a", "b", "c"})
	fmt.Println(m == n)  // true

}

// 考核点：结构体嵌套与接口
func testInterface5(){
	printType(S1{})
	printType(S2{})

	/* 假如 S2 没有实现 f() 方法
	var i I
	i = S2{}   // I 接口依旧可以接收 S2{} 的值，因为 S2 嵌套了 S1 结构体，  S1 实现了 I 接口
	i.f()
	*/
}

// 考核点： 接口实现以及接口接收值 的规则
func testInterface4(){
	var inf LessAddInf   // 接口
	var n Integer  // 自定义类型
	inf = &n    // 由于 自定义类型的 值类型 和引用类型 都分别实现了 LessAddInf 接口，如果想用 inf 接口类型来接收 自定义类型，那么必须是 接收指针类型的变量，也就是指针
	fmt.Printf("inf.Less(20)=%v\n",inf.Less(20))  // 语法糖， Less 是 值类型实现的方法， inf 这个接口为什么可以调用呢，因为 语法糖
	fmt.Printf("inf.Add(30)=%v\n", inf.Add(30))

	var thing Thing
	var computer = Computer{CPU:"英特尔至强-v3440", Memory:"三星DDR4(8g)"}
	thing = &computer
	fmt.Printf("thing.Name()=%v\n", thing.Name())
	fmt.Printf("thing.Attribute()=%v\n", thing.Attribute())
}

// 考核点：动态类型之间的比较
func testInterface3(){
	var x interface{} = []int{1,2,3}
	fmt.Println(x ==x)   // 运行报错： panic: runtime error: comparing uncomparable type []int

	/*
		在比较两个接口值时，如果两个接口值的动态类型一致，但对应的动态值是不可比较的 (比如 slice)， 那么这个比较会以崩溃的方式失败。
	 */
}

// 考核点：接口是不是引用类型？？？？
func testInterface2(){
	var x = 10
	fmt.Printf("传入前：x 的地址：%x\n", &x)   // 传入前：x 的地址：c0000140b0
	change(x)
}

func change(x interface{}){
	v, ok := x.(int)
	if ok {
		fmt.Println(v)
	}

	fmt.Printf("传入后：x 的地址：%x\n", &x)   // 传入后：x 的地址：c0000401f0
}

// 考核点：switch case 的操作
func testSwitch(){
	// 虽然提示重复的 case ,但是 程序能正常运行，代表 go 语言中支持两个 条件相同的 case
	switch  {
	case true:
		fmt.Println("case 1")
		fallthrough    // fallthrough  表示在满足条件1 时，紧接着执行条件2 的代码 （条件2 是否满足，就算是 条件2 为 false ,也直接执行条件2 下面的代码）
	case true:
		fmt.Println("case 2")
	}
}

// 考核点：可变参数调用
func testArgs(){

	// 正确调用如下：
	add()
	add(1)
	add(1,3,6)

	add([]int{11}...)
	add([]int{11,22,33}...)

	// 错误
	// add([]int{1,2,3})  // 编译不通过

}

func add(args ...int) int {
	sum := 0

	for arg := range args {
		sum += arg
	}
	return sum
}

// 考核点：命名规则
func testVar(){

	// go 语言支持中文命名的 变量名, 不支持数字开头的变量名
	姓名:= "小明"

	/*
	var 2a = "小米"   // 错误命名

	var a&b = "ddd"   // 错误命名 , 不支持 &*$ 等命名

	var a*b = 1234   // 错误命名

	var func int     // func ,int  都是关键字，关键字冲突
	*/

	fmt.Println(姓名)
}

// 考核点：通道的取值
func testChan(){
	m := make(chan int, 10)

	// 直接从通道中取值
	a := <-m
	fmt.Println(a)

	// 从通道中取值，还会返回一个 bool, 来表示是否取到值
	v, ok := <-m
	if ok {   // 能取到☞6
		fmt.Println(v)
	}
}

// 考核点：for 循环遍历通道
func testFor(){

	m := make(chan int, 10)

	/*  for 循环错误写法
	for v,ok := range m {    // 错误
		// fmt.Println(v)
	}

	for _,ok := range m {    // 错误
		// fmt.Println(v)
	}

	for ok := range m;ok {    // 错误
		// fmt.Println(v)
	}
	*/

	// 正确
	for v := range m {
		fmt.Println(v)
	}

	// 正确
	for {
		if v,ok := <-m; ok {
			fmt.Println(v)
		}
	}
}

// 考核点：字符串拼接
func testString(){
	str1 := "abc"+"123"
	fmt.Println(str1)
	fmt.Printf("%#v\n", str1)

	str2 := `abc` + `123`
	fmt.Println(str2)
	fmt.Printf("%#v\n", str2)

	str3 := fmt.Sprintf("abc%d", 123)
	fmt.Println(str3)
	fmt.Printf("%#v\n", str3)
}

// 考核点：go 语言中 变量自增自减的操作
func testCalcTag2(){
	i := 1
	i++            // 变量自增，
	// j := i++    // 错误语法， i++ 就是  i = i+1, 不能再赋值给 j

	// --i   // 错误语法，go 语言中 ++, -- 只支持后置，不支持 前置


}

// 考核点：go 语言中的运算符
func testCalcTag(){
	// 参考： https://www.sojson.com/operation/go.html

	// 问： (1+6)/2*4^2+10%3<<3的值是多少
	/*  分析解答：

	^是异或的意思，是二进制运算
	1是二进制是0001，2的二进制是0010
	异或规则是相同为0，不同为1
	1^2   0001 ^ 0010  0011

	7/2*4^2+10%3<<3
	12^2+10%3<<3
	1100 ^ 0010  1110
	14+10%3<<3
	14+1<<3          1左移1位乘以2的1次方，移3位乘以2的3次方
	14+8=22

	 */

	v := (1+6)/2*4^2+10%3<<3

	fmt.Println(v)   //  22

}


// 考核点: new 与 make 的区别
func testNewAndMake(){
	// make 只针对 slice, map, channel 进行申请内存操作 并且初始化
	// new 可以针对任意类型进行 申请内存的操作，但不进行初始化

	m := new([]int)   // 只申请内存
	fmt.Println(m)   // &[]


	m2 := make([]int, 4,8)   // 申请内存并初始化
	fmt.Println(m2)  // [0 0 0 0]
}

// 考核点：指针的操作
func testPtr(){

	// 取地址
	var a int = 10
	fmt.Printf("%p\n", &a)   // 0xc00000e0c0

	fmt.Printf("%x\n", &a)   // c00000e0c0


	// 定义指针类型的变量
	var b int = 20
	var ip *int

	ip= &b
	fmt.Printf("b 的地址：%p, b 的值：%v\n", &b, b)       // b 的地址：0xc00000e0c8, b 的值：20
	fmt.Printf("ip 的地址：%v, ip 的值：%v\n", ip, *ip)  // ip 的地址：0xc00000e0c8, ip 的值：20

	// 空指针
	var p *int
	fmt.Println(p)  // <nil>
	fmt.Printf("p (x)的值 %x\n", p)    // p (x)的值 0

	fmt.Printf("p (p)的值 %p\n", p)    // p (p)的值 0x0

	fmt.Printf("p (v)的值 %v\n", p)    // p (v)的值 <nil>

	// 值传递与引用传递
	a,b = 3,4
	swap(a,b)
	fmt.Println(a,b)   // 3,4

	swapPtr(&a,&b)
	fmt.Println(a,b)   // 4,3


}

// 考核点： 指针的连锁反应
func testPtr2(){

	a,b := 3,4

	c,d := swapReturn(&a, &b)
	fmt.Println(*c, *d)   // 4,3

	a = *c
	b = *d
	fmt.Println(a, b)  // 4,4   // b 为什么会是 4 呢？

	/* 分析
		第一步： c,d := swapReturn(&a, &b)   , 将a,b 的地址交换 并且赋值给 c, d , 那就说明 c 指向 b 的值, d 指向 a 的值 , 所以 *c,*d = 4,3
		第二步： a = *c,  那就是把 c 指向的 b 的值 (4) 赋值给 a, 此时 a 等于 4,  然后呢，因为 d 指向 a 的值，所以 *d = 4
		第三步： b = *d,  经过 a = *c 的操作 (连锁反应), 此时 *d = 4, 所以 b 的值为 4
	 */

}

func swapReturn(a,b *int) (*int,*int) {
	a,b = b,a
	return a, b
}

// 引用传递 (传递的是地址，要达到值的交换目的，那就必须在 里面进行 * 取值并交换)
func swapPtr(a,b *int){
	*a, *b = *b, *a
}

// 值传递
func swap(a,b int) {
	// 交换的语法
	a, b = b, a
}


// 考核 for range 的用法
func testForRange(){
	m1 := make(map[string]*User, 8)

	userSlice := []User{
		{Name:"yang",Age:12},
		{Name:"zhagn",Age:18},
		{Name:"liu",Age:34},
	}

	// 问题1：问题出在这里， for range 遍历对象出来之后 赋值给了 stu 这个临时变量，然后每次都将 stu 的地址赋值给 m1,所以 m1里面只放入了一个 stu指向的值
	for _,stu := range userSlice {
		m1[stu.Name] = &stu
	}

	for k,v := range m1 {
		fmt.Println(k,v)
	}

	/*  打印信息：为什么呢？  问题出在 for _,stu := range userSlice
		yang &{liu 34}
		zhagn &{liu 34}
		liu &{liu 34}
	*/


	// 问题2：
	// 对 for range 遍历出来的值进行 修改操作
	// 遍历出来的值赋值给了 stu， 是值拷贝，所以针对 stu 的修改不会影响到  userSlice
	for _, stu := range userSlice{
		stu.Age = stu.Age + 10
	}

}

// 考核接口的用法
func testInterface(){
	// People 是一个接口，所以它只能接收 实现它方法的 类型 (所以必须是 Student 的指针类型)
	// var peo People = Student{}  // 编译失败
	var peo People = &Student{}  // 考核点

	fmt.Println(peo.Speak("sb"))
}


func getStudentInfo(sf StudentFunc){

}

// 考核点： type StudentFunc func(*Student) 的用法
func testTypeFunc(){

	studentFunc := func(stu *Student) {
		stu.Name = "yang"
	}

	getStudentInfo(studentFunc)


}

