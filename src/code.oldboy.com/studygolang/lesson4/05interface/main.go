package main

import "fmt"

// WashingMachine 是一个接口类型
type WashingMachine interface {
	wash()
	dry()
}

// dryer 是一个结构体
type dryer struct {
}

// 内部嵌套的结构体 实现 接口中的 dry() 方法
func (d dryer) dry() {
	fmt.Println("甩干")
}

// haier 是一个结构体, 内部嵌套了 一个 dryer 结构体
type haier struct {
	dryer
}

// haier 结构体实现 接口中的 wash() 方法
func (h haier) wash() {
	fmt.Println("洗衣")
}

// 嵌套结构体 实现接口中的方法
// 1, 接口里面的方法，不一定需要由一个类型完全实现，可以通过在类型中 嵌入其他类型或者结构体 来实现
// 如本包中的示例，  WashingMachine 这个接口类型中定义了 两个方法 wash() 和 dry()
// 分别由  haier 实现 wash(),  dryer 实现 dry()
// 但是前提是， haier 和 dryer 这两个结构体类型，必须嵌套在一起，归属于一个大的结构体。
func main() {

	var w WashingMachine
	h := haier{}

	w = h // 实例化 接口

	w.wash()
	w.dry()

	fmt.Println(w) // "{{}}"

	/*  如果单独把 dryer 拿出来赋值给 接口类型，会发生什么？
	// 它会编译报错， 说 d 这个结构体没有实现 wash() 这个方法，  这也验证了 在给接口赋值的时候， 会自动判断你是否 实现了我接口中所有定义的 方法
	d := dryer{}
	// w = d    // 编译报错： cannot use d (variable of type dryer) as WashingMachine value in assignment: missing method wash
	w.wash()

	fmt.Println(w)
	*/
}
