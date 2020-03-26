package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"unsafe"
)

// 参考网上方法，貌似不能得到 map 的容量
type hmap struct {
	count      int
	flags      uint8
	B          uint8
	hash0      uint32
	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
}

func main() {
	fmt.Println("第二天练习")

	arr := []int{1, 3, 5, 7, 8}
	num := test1(arr)
	fmt.Println(num)

	result := test2(arr)
	fmt.Println(result) // "(0,3)(1,2)"

	//test3()

	//test4()

	// test5()
	// test6()

	// testStudentInfo()

	// testMap()
	testCoin()

}

/*
	50枚金币，分配给以下几个人： Matthew, Sarah, Augustus, Heidi, Emilie, Peter, Giana, Adriano, Aaron, Elizabeth
	分配规则如下：
	a. 名字中包含e 或者 E： 1枚金币
	b. 名字中包含i 或者 I： 2枚金币
	c. 名字中包含o 或者 O:  3枚金币
	d. 名字中包含u 或者 U:  4枚金币
	写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
*/

// 第一步，根据需求定义如下变量
var (
	coins        = 50
	users        = []string{"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth"}
	distribution = make(map[string]int, len(users))
)

func testCoin() {

	// 采用 if 判断条件来写：
	left := dispatchCoin()
	fmt.Println("剩余金币：", left)

	fmt.Println("--------------------每个人分配金币数量如下：----------------------")
	for k, v := range distribution {
		fmt.Printf("姓名：%s\t金币：%d\n", k, v)
	}

	/*
		姓名：Augustus  金币：4
		姓名：Heidi     金币：3
		姓名：Peter     金币：1
		姓名：Aaron     金币：3
		姓名：Elizabeth 金币：3
		姓名：Matthew   金币：1
		姓名：Emilie    金币：3
		姓名：Giana     金币：2
		姓名：Adriano   金币：5
	*/

	// 采用 switch 流程控制来实现：  无法实现 功能需求
	/*
		left := dispatchCoinBySwitch()
		fmt.Println("剩余金币：", left)

		fmt.Println("--------------------每个人分配金币数量如下：----------------------")
		for k, v := range distribution {
			fmt.Printf("姓名：%s\t金币：%d\n", k, v)
		}
	*/
}

// 第二步，定义一个分发金币的函数，返回 剩余金币数
func dispatchCoin() int {

	// 遍历名字的切片
	for _, v := range users {
		//fmt.Println(v)
		// 判断名字是否满足条件
		if strings.Contains(v, "e") || strings.Contains(v, "E") {

			/*  优化：
			由于distribution 的value 值是int 类型，所以默认值为0， 所以 以下代码块可以优化简写成：
				distribution[v] = distribution[v] + 1
			// 无需再判断  coin, ok := distribution[v]
			*/
			coin, ok := distribution[v]
			//fmt.Println("1, ", coin, ok)
			if ok {
				distribution[v] = coin + 1
			} else {
				distribution[v] = 1
			}
			//fmt.Println("---1", distribution[v])
			coins = coins - 1
		}
		if strings.Contains(v, "i") || strings.Contains(v, "I") {
			coin, ok := distribution[v]
			//fmt.Println("2, ", coin, ok)
			if ok {
				distribution[v] = coin + 2
			} else {
				distribution[v] = 2
			}
			//fmt.Println("---2", distribution[v])
			coins = coins - 2
		}
		if strings.Contains(v, "o") || strings.Contains(v, "O") {
			coin, ok := distribution[v]
			//fmt.Println("3, ", coin, ok)
			if ok {
				distribution[v] = coin + 3
			} else {
				distribution[v] = 3
			}
			//fmt.Println("---3", distribution[v])
			coins = coins - 3
		}
		if strings.Contains(v, "u") || strings.Contains(v, "U") {
			coin, ok := distribution[v]
			//fmt.Println("4, ", coin, ok)
			if ok {
				distribution[v] = coin + 4
			} else {
				distribution[v] = 4
			}
			//fmt.Println("---4", distribution[v])
			coins = coins - 4
		}
	}
	return coins
}

// 我们尝试使用 switch 来写这个需求
// 结果发现并不能实现 题目要求， 求出的结果会是：  e, E 都会计算一遍
func dispatchCoinBySwitch() int {

	for _, v := range users {

		for _, c := range v { // 将名字遍历出 一个一个的字符进行判断，会导致 e 与 E 重复计算
			switch c {
			case 'e', 'E':
				distribution[v] = distribution[v] + 1
				coins = coins - 1
			case 'i', 'I':
				distribution[v] = distribution[v] + 2
				coins = coins - 2
			case 'o', 'O':
				distribution[v] = distribution[v] + 3
				coins = coins - 3
			case 'u', 'U':
				distribution[v] = distribution[v] + 4
				coins = coins - 4
			}
		}
	}
	return coins
}

// 注意 下面 switch 流程控制代码逻辑 与 上面 switch 的区别：

/*   分金币的练习不能用 如下switch 流程控制语句的原因是：
	1, go 语言中 switch case 默认自带 break,  满足一个条件之后就会自动退出， 这跟 java 不一样，
		java 中的case 语句只要不带 break 则会继续向下执行 其它 case

	// 所以我们采用 多个 if 条件并列的判断语句，来多次判断，注意 多个if 条件并列，每个if 条件都要另起一行
	// 在这里能不能用 if{} else if{}  呢， 虽然语法可以，但是效果与  switch case 一样，不会并列判断，只执行一次。

	// 注意 go 中 switch ，  default 可以省略，  Java中也不是必须的，但建议加上。

switch {
		case strings.Contains(v, "e") || strings.Contains(v, "E"):
			coin, ok := distribution[v]
			fmt.Println("1, ", coin, ok)
			if ok {
				distribution[v] = coin + 1
			} else {
				distribution[v] = 1
			}
			coins = coins - 1
		case strings.Contains(v, "i") || strings.Contains(v, "I"):
			coin, ok := distribution[v]
			fmt.Println("2, ", coin, ok)
			if ok {
				distribution[v] = coin + 2
			} else {
				distribution[v] = 2
			}
			coins = coins - 2
		case strings.Contains(v, "o") || strings.Contains(v, "O"):
			coin, ok := distribution[v]
			fmt.Println("3, ", coin, ok)
			if ok {
				distribution[v] = coin + 3
			} else {
				distribution[v] = 3
			}
			coins = coins - 3
		case strings.Contains(v, "u") || strings.Contains(v, "U"):
			coin, ok := distribution[v]
			fmt.Println("4, ", coin, ok)
			if ok {
				distribution[v] = coin + 4
			} else {
				distribution[v] = 4
			}
			coins = coins - 4
		default:
*/

func testMap() {

	// 创建map : key类型为 int , value类型为 string 切片 的
	studentMap := make(map[int][]string, 10)
	studentMap[1] = make([]string, 3, 3)
	studentMap[1][0] = "沙河小王子"
	studentMap[1][1] = "28"
	studentMap[1][2] = "88"

	studentMap[2] = make([]string, 3, 3)
	studentMap[2][0] = "沙河大魔王"
	studentMap[2][1] = "18"
	studentMap[2][2] = "78"

	studentMap[3] = make([]string, 3, 3)
	studentMap[3][0] = "沙河蛮王"
	studentMap[3][1] = "45"
	studentMap[3][2] = "100"

	fmt.Println(studentMap) // "map[1:[沙河小王子 28 88] 2:[沙河大魔王 18 78] 3:[沙河蛮王 45 100]]"
}

// 设计一个程序，存储学员信息： id  姓名 年龄 分数
// 能根据 id 获取学员信息
func testStudentInfo() {
	studentMap := make(map[string][]string, 10)

	studentMap["0001"] = make([]string, 3, 3)
	studentMap["0001"][0] = "沙河小王子"
	studentMap["0001"][1] = "28"
	studentMap["0001"][2] = "88"

	studentMap["0002"] = make([]string, 3, 3)
	studentMap["0002"][0] = "沙河大魔王"
	studentMap["0002"][1] = "18"
	studentMap["0002"][2] = "78"

	studentMap["0003"] = make([]string, 3, 3)
	studentMap["0003"][0] = "沙河蛮王"
	studentMap["0003"][1] = "45"
	studentMap["0003"][2] = "100"

	fmt.Println(studentMap) // map[0001:[沙河小王子 28 88] 0002:[沙河大魔王 18 78] 0003:[沙河蛮王 45 100]]

	fmt.Println("-------------------------------------------------------------------")

	// 查询 studentId 为 "0002" 的学员信息
	studentId := "0002"
	for k, v := range studentMap {
		// fmt.Println("---id : ", k)
		if k == studentId {
			_, ok := studentMap[studentId] // map[key] 返回值为 value  和 ok (是否查到)
			// fmt.Println(v2)  // "[沙河大魔王 18 78]"
			if ok && len(v) == 3 {
				fmt.Printf("name : %s, age : %s, score : %s\n", v[0], v[1], v[2])
			} else {
				fmt.Println("查无此人！")
			}
		}
	}

}

// 求 int 数组里元素之和
func test1(arr []int) int {
	var num int
	for _, v := range arr {
		num += v
	}
	return num
}

// 求数组里面元素和为 8 的 两个索引值
func test2(arr []int) string {
	var result string
	// 让第一个数 和 其他几个数进行相加， 然后让 第二个数 和 除去第一个数之外的元素相加，然后让 第三个数 和除去第一，第二个数之外的元素相加。。。。
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == 8 {
				result += fmt.Sprintf("(%d,%d)", i, j)
			}
		}
	}
	return result
}

// 求 a  输出的值是多少
func test3() {
	var a = make([]string, 5, 10)
	fmt.Println(a) // "[    ]"  // a 默认值是，有 5个空字符串，总容量为10，还剩下 5个空间余量
	for i := 0; i < 10; i++ {
		a = append(a, fmt.Sprintf("%v", i)) // 向a 里面添加元素
	}
	fmt.Println(a) // "[     0 1 2 3 4 5 6 7 8 9]"     // 一定不能忽略默认的 5个空字符串
}

// 使用内置的 sort 包对数组 进行排序
func test4() {
	var a = [...]int{3, 7, 8, 9, 1}
	// a[:] 得到一个切片，指向底层数组 a
	sort.Ints(a[:]) // Ints() 参数里定义的是 切片类型 , 所以要把 a 进行切片操作
	fmt.Println(a)  // "[1 3 7 8 9]"
}

// 统计一个字符串中每个单词出现的次数
func test5() {
	var s = "how do you do"
	words := strings.Split(s, " ")
	var wordCount = make(map[string]int, len(words)) // 定义一个map 用来存储每个单词以及对应的次数， 初始化完成 (10个空字符串)

	fmt.Println("未赋值前：", wordCount)

	// 先将字符串s 按照空格分隔，得到一个 slice
	for _, word := range words { // 遍历 slice, 取出每一个单词
		v, ok := wordCount[word] // 判断单词 是否存在 map 中，(map 默认为10个空字符串)
		if ok {                  // 如果存在，也就是执行了 else 之后，里面有单词了，又发现一个单词与已添加到 map 中的单词一样，则表示存在
			wordCount[word] = v + 1 // 然后拿到已存在的那个单词对应的 value , 也就是次数，然后再 加 1，依次类推，再依次出现就再 加 1
		} else { // 代码的正确执行逻辑，应该会先到 这里来，因为 !ok
			wordCount[word] = 1 // 不存在的话，会将这个单词 作为 key, 加入到 map 中，然后记录单词第一次出现， value 设置为 1
		}
	}
	fmt.Println(wordCount)      // "map[do:2 how:1 you:1]"
	fmt.Println(len(wordCount)) // "3"

	/* // 参考网上方法，貌似不能得到 map 的容量
	l, c := getInfo(wordCount)
	fmt.Printf("len: %d,  cap:%d", l, c)
	*/
}

// 参考网上方法，貌似不能得到 map 的容量
func getInfo(m map[string]int) (int, int) {
	point := (**hmap)(unsafe.Pointer(&m))
	value := *point
	return value.count, int(value.B)
}

// 按照某个固定顺序遍历 map
func test6() {

	var scoreMap2 = make(map[string]int, 50)
	for i := 0; i < 25; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100) // 0~99 的随机整数
		scoreMap2[key] = value
	}

	for k, v := range scoreMap2 {
		fmt.Println(k, v)
	}
	/*  打印结果：
	stu04 81
	stu22 28
	stu24 47
	stu00 81
	stu14 28
	stu15 74
	stu23 58
	stu01 87
	stu06 25
	stu10 94
	stu20 95
	stu07 40
	stu13 89
	stu19 6
	stu09 0
	stu16 11
	stu18 37
	stu08 56
	stu17 45
	stu21 66
	stu12 62
	stu02 47
	stu03 59
	stu05 18
	stu11 11
	*/
	// 问题：我们如何将上面 无序的 map ，进行有序的打印呢?

	fmt.Println("---------------------按照指定的顺序去打印------------------")

	// 1,按照key 从小到大的顺序去遍历 scoreMap2
	keys := make([]string, 0, 50)
	for k := range scoreMap2 {
		keys = append(keys, k)
	}
	// 2,对 key 做排序
	sort.Strings(keys) // 参数为 一个 slice
	for _, key := range keys {
		fmt.Println(key, scoreMap2[key])
	}

	/*  排序后打印信息：
	stu00 81
	stu01 87
	stu02 47
	stu03 59
	stu04 81
	stu05 18
	stu06 25
	stu07 40
	stu08 56
	stu09 0
	stu10 94
	stu11 11
	stu12 62
	stu13 89
	stu14 28
	stu15 74
	stu16 11
	stu17 45
	stu18 37
	stu19 6
	stu20 95
	stu21 66
	stu22 28
	stu23 58
	stu24 47
	*/

}
