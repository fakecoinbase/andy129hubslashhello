package main

import (
	"fmt"
)

// map(映射) 类型
func main() {

	fmt.Println("map")

	// test()
	// testLenCap()
	testOther()
	// testOther2()
}

// 测试元素类型为 map 的切片 （即: 切片里面保存的是一个个的 map）
func testOther() {

	// 元素类型为 map 的切片
	// 切片里面保存的是一个个的 map
	var mapSlice = make([]map[string]int, 8, 8) // 完成了切片的初始化
	fmt.Println(mapSlice)                       // "[map[] map[] map[] map[] map[] map[] map[] map[]]"
	// 所以，make ,只针对切片进行了初始化，并没有对里面的 map 进行初始化，里面的 map 都是 nil
	fmt.Println(mapSlice[0] == nil) // "true"
	// mapSlice[0]["沙河小王子"] = 100   // 运行报错：panic: assignment to entry in nil map

	// 所以我们还需要完成 内部 map 元素的初始化 (注意make map 时，里面只能有两个参数， (map, 长度）, 否则编译报错:too many arguments to make(map[string]int)
	mapSlice[0] = make(map[string]int, 8) // 完成了 slice 里面第一个元素 map 的初始化
	fmt.Println(mapSlice[0])              // "map[]"
	mapSlice[0]["沙河小王子1"] = 100
	mapSlice[0]["沙河小王子2"] = 100
	mapSlice[0]["沙河小王子3"] = 100
	mapSlice[0]["沙河小王子4"] = 100
	mapSlice[0]["沙河小王子5"] = 100
	mapSlice[0]["沙河小王子6"] = 100
	mapSlice[0]["沙河小王子7"] = 100
	mapSlice[0]["沙河小王子8"] = 100
	mapSlice[0]["沙河小王子9"] = 100
	fmt.Println(mapSlice[0])
	fmt.Println(len(mapSlice[0])) // 9
	// fmt.Println(cap(mapSlice[0])) // 无此方法
	// 以上打印信息：map[沙河小王子1:100 沙河小王子2:100 沙河小王子3:100 沙河小王子4:100 沙河小王子5:100 沙河小王子6:100 沙河小王子7:100 沙河小王子8:100 沙河小王子9:100]

	/*  数组越界：mapSlice[8] : panic: runtime error: index out of range [8] with length 8
	mapSlice[8] = make(map[string]int, 8)
	mapSlice[8]["沙河大魔王"] = 1000
	fmt.Println(mapSlice[8])
	*/

}

// 测试 元素类型为 slice 的 map (即：map 里面保存的是一个个的 切片)
func testOther2() {

	var sliceMap = make(map[string][]int, 8) // 只完成了 map 的初始化，并没有完成 里面切片的初始化
	fmt.Println(sliceMap)                    // "map[]"

	v, ok := sliceMap["中国"]
	if ok { // 找不到元素，不打印下面
		fmt.Println(v)
	} else {
		// 没有完成 map 里面切片的初始化
		// sliceMap["中国"][0] = 100 // 运行错误：panic: runtime error: index out of range [0] with length 0

		// 所以我们要先初始化 切片
		sliceMap["中国"] = make([]int, 8) // 初始化map key 对应的 value(这里是切片)
		sliceMap["中国"][0] = 100
		sliceMap["中国"][1] = 200
		sliceMap["中国"][2] = 300
		sliceMap["中国"][3] = 400

		// 如果还要再创建一个 键值对 ，则需要将对应的 切片初始化
		sliceMap["美国"] = make([]int, 8)
		sliceMap["美国"][0] = 178
		sliceMap["美国"][1] = 158
		sliceMap["美国"][2] = 198
	}

	// 遍历 sliceMap
	for k, v := range sliceMap {
		fmt.Println(k, v)
	}
	/*  打印信息如下：
	美国 [178 158 198 0 0 0 0 0]
	中国 [100 200 300 400 0 0 0 0
	*/

	// 注意全部打印和 上面分开键值对打印的 区别和规律
	fmt.Println(sliceMap) //  "map[中国:[100 200 300 400 0 0 0 0] 美国:[178 158 198 0 0 0 0 0]]"
}

func test() {
	// 仅声明，并没有初始化，此时 a 的值就是 nil
	var a map[string]int  // map[keyType]valueType
	fmt.Println(a == nil) // "true"
	// map 的初始化
	a = make(map[string]int, 8)
	fmt.Println(a == nil) // "false"

	a["沙河娜扎"] = 100
	a["沙河小王子"] = 200
	fmt.Printf("%T\n", a) // "map[string]int"
	fmt.Println(a)        // "map[沙河娜扎:100 沙河小王子:200]"
	// %v 默认格式打印
	fmt.Printf("%v\n", a) // "map[沙河娜扎:100 沙河小王子:200]"
	// %#v 会把map 里面 key,value 相对应的类型也打印出来
	fmt.Printf("%#v\n", a) // "map[string]int{"沙河娜扎":100, "沙河小王子":200}"

	// 声明map 的同时完成初始化
	b := map[int]bool{
		1: true,
		2: false,
	}
	fmt.Printf("%T\n", b) // "map[int]bool"
	fmt.Println(b)        // "map[1:true 2:false]"

	// 错误操作：未初始化就赋值
	// var c map[int]int
	// c 这个map 没有初始化，所以不能直接操作，下面这行代码运行错误
	// c[100] = 200 // 运行报错： panic: assignment to entry in nil map、

	// 判断某个键存不存在
	var scoreMap = make(map[string]int, 8) //指定 map 的长度为 8，   另外一种参数设置： make(map[string]int, 0, 8)   // 长度为0， 容量为8
	fmt.Println(scoreMap)                  // "map[]"
	scoreMap["沙河娜扎"] = 100
	scoreMap["沙河小王子"] = 200
	scoreMap["沙河小魔王"] = 300
	scoreMap["沙河大鳄"] = 400
	scoreMap["沙河剑圣"] = 500
	scoreMap["沙河刘德华"] = 600
	scoreMap["沙河梁朝伟"] = 700
	scoreMap["沙河周星驰"] = 800
	scoreMap["沙河越界"] = 1000000
	fmt.Println(len(scoreMap)) // "9"

	// 判断 张二狗子 在不在 scoreMap 中
	value, ok := scoreMap["张二狗子"]
	fmt.Println(value, ok) // "0 false"
	if ok {
		fmt.Println("张二狗子存在")
	} else {
		fmt.Println("查无此人！")
	}

	v, ok := scoreMap["沙河娜扎"]
	fmt.Println(v, ok) // "100 true"

	fmt.Println("------------------------map的遍历-------------------------")
	// map 的遍历 （map是无序的，遍历除键值对的结果 和 添加的顺序无关）
	// 获取 key 和 value
	for k, v := range scoreMap {
		fmt.Println(k, v)
	}
	/*  遍历结果：  由此可看出 遍历 map 的结果与 加入的顺序不同， map 是无序的。
	沙河刘德华 600
	沙河梁朝伟 700
	沙河周星驰 800
	沙河娜扎 100
	沙河小王子 200
	沙河小魔王 300
	沙河大鳄 400
	沙河剑圣 500
	*/

	// 只获取 key
	for k := range scoreMap {
		fmt.Println(k)
	}

	// 只获取 value
	for _, v := range scoreMap {
		fmt.Println(v)
	}

	// 当for range 时， 只接收一个返回值时，它是 数组的下标或者 map 的key , 如果要获得value, 则要  for _,v:= range c
	var c = []int{1, 2, 3, 5}
	for i := range c {
		fmt.Println(i)
	}

	fmt.Println("------------------------map的删除-------------------------")
	// 删除 map 中的一个 键值对
	// delete 无返回值
	delete(scoreMap, "沙河刘德华")
	delete(scoreMap, "沙河马云") // 当scoreMap 里面没有 "沙河马云" 这个键时，不会报错， 也不会影响 scoreMap 的长度
	delete(scoreMap, "沙河马化腾")
	fmt.Println(scoreMap)
	fmt.Println(len(scoreMap)) // "7"
}

func testLenCap() {
	var a = make([]int, 8)      // 只有两个参数时，后面这个数字，代表 len , 省略了 cap 的指定，则代表 cap == len
	fmt.Println(a)              // "[0 0 0 0 0 0 0 0]"
	fmt.Println(len(a), cap(a)) // "8 8"

	//a[8] = 100
	//fmt.Println(a) // 运行报错： panic: runtime error: index out of range [8] with length 8
}
