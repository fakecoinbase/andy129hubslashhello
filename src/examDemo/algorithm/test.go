package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 算法面试题
func main() {
	// test1()
	//test1Other()

	// test2()

	// test3()

	// test3Other()
	// test4()   // 合并两个有序列表

	//test5()    // 栈 的应用 (判断 括号的有效性)

	// test6()    // 模拟队列的操作

	// test6Other()  // 队列的应用：  读取文件的最后 5 行数据

	// test7()   // 栈的应用: 迷宫的问题 (深度搜索)

	// test7Other()  // 队列的应用： 迷宫的问题 (广度搜索)

	// test8()     // 贪心算法
	// test9()     // 背包算法

	// test10Other()  // 字符串拼接 比较大小
	// test10()     // 给定一些数字，拼接成最大值

	// test11()        // 活动最有解

	// test12()        // // 动态规划 (DP)的思想：求 斐波那契数列

	// test13()   // 钢条切割的问题

	// test14()   // 最长公共子序列  (字符串模糊匹配)

	// test15()    // 求最大公约数，欧几里得算法, 以及运用最大公约数 进行 分数的 四则运算(加减乘除)

	test16()   // 数组元素依次移位
}
// 数组移位
//  示例：给定一个数组 data := [10]int{1,4,5,7,0,2,3,8,6,9} , 输入n , 数组元素依次移动 n 个位置
func test16(){
	data := [10]int{1,4,5,7,0,2,3,8,6,9}
	moveArr(data[:], 4)
}
func moveArr(data []int, n int){
	data = append(data[len(data)-n:len(data)],data[0:len(data)-n]... )
	fmt.Println(data)  // [3 8 6 9 1 4 5 7 0 2]
}

// 创建一个代表 分数的的结构体
type Fraction struct {
	a int      // 分子
	b int      // 分母
}

func (f *Fraction)String() string {
	return fmt.Sprintf("%d/%d", f.a,f.b)
}
// 构造 Fraction 对象 (针对 分子a, 分母b ,自动进行约分)
func NewFraction(a,b int) *Fraction{
	x := gcd(a,b)   // 求分子和分母的最大公约数，进行约分

	return &Fraction {
		a : a/x,
		b : b/x,
	}
}
// 两分数 相加
func (f *Fraction) Add(fra *Fraction) *Fraction {

	fenmu := zgs(f.b, fra.b)   // 求两分数中 分母的最小公倍数, 也就是通分之后的 分母

	fenzi := f.a * (fenmu /f.b) + fra.a * (fenmu /fra.b)   // 分子也根据分母扩大相同的倍数 ，然后两数的分子数 相加

	return &Fraction{
		a: fenzi,
		b: fenmu,
	}
}

// 两数相减...
func (f *Fraction) Sub(fra *Fraction) *Fraction {

	return  nil
}

// 两数相乘...
func (f *Fraction) Mul(fra *Fraction) *Fraction {

	return  nil
}

// 两数相除...
func (f *Fraction) Div(fra *Fraction) *Fraction {

	return  nil
}

func test15(){

	fmt.Println(gcd(12,18))  // 6

	fmt.Println(gcdNoRec(12,18))  // 6

	// 欧几里得的应用：  分数 (分子与分母的约分)
	fra := NewFraction(12,18)
	fmt.Println(fra)   // 2/3


	fmt.Println(zgs(12,18))   // 36

	fra2 := NewFraction(20,50)
	sum := fra.Add(fra2)
	fmt.Println(sum)   // 16/15


}
// 求两数的最小公倍数
func zgs(a,b int) int {
	x := gcd(a,b)
	m := a/x * b/x * x    // 算出最小公倍数 (a除于最大公约数，b 除以最大公约数，然后再乘以 最大公约数，就是 最小公倍数)
	return m
}


// 求 x ,y 两数的最大公约数
/*  欧几里得算法 gcd(x,y) == gcd(y, x%y)
	示例： gcd(60,21) == gcd(21,18) == gcd(18,3) == gcd(3,0) == 3
 */
func gcd(x,y int) int{
	if y == 0 {
		return x
	}else {
		return gcd(y, x%y)
	}
}
// 求最大公约数，欧几里得算法，非递归方式
func gcdNoRec(x,y int)int{
	for y > 0 {
		r := x%y

		x = y
		y = r
	}

	return x
}

// 最长公共子序列
func test14(){

}

func lcs(){

}

// 钢条切割的问题
func test13(){

}


//
func test12(){
	fmt.Println(fib(50))

	fmt.Println(fibNoRec(100))
}

// 动态规划 (DP)的思想：递推式 + 重复子问题


// 递归求 斐波那契数列 (子问题的重复计算)
/*  举例： fib(6),  可以看出 f(4), f(3) 重复计算了
	f(6) = f(5) + f(4)
    f(5) = f(4) + f(3)
    f(4) = f(3) + f(2)
    f(4) = f(3) + f(2)
    f(3) = f(2) + f(1)
    f(3) = f(2) + f(1)
    f(3) = f(2) + f(1)

 */
func fib(n int) int64{
	if n == 1 || n ==2 {
		return 1
	}
	return fib(n-1)+fib(n-2)
}
// 非递归 求斐波那契数列 （与递归法不同的是， 非递归方式会创建一个 切片保存每一位斐波那契数列的值）
func fibNoRec(n int ) int {
	// 创建一个切片用于存储 斐波那契数值 (将 n 对应 f 中下标指向的值)
	// 0 作为补充，目的是为了 n 对应下标的值
	f := []int{0,1,1}
	if n >2 {
		for i:=0;i<n-2;i++ {
			// 从 f 中保存的前几个斐波那契数的值
			num := f[len(f)-1] + f[len(f)-2]
			f = append(f, num)   // 将结果添加到  f 中
		}
	}
	return f[n]   // 从 f 中获取 值 (也就是 n 的斐波那契数列值)
}

type Activity struct {
	start int
	end int
}

type Activitys []Activity

// fmt.Println(acts)  时，系统会自动调用 String() 方法
// 但是 fmt.Printf("%#v", acts)  ， 这样就不会调用 String() 方法
func (a Activitys) String() string{
	str := ""
	for i:=0;i<len(a);i++ {
		str = str + fmt.Sprintf("(%d,%d)", a[i].start,a[i].end)
	}
	return str
}

func (a Activitys)Less(i,j int) bool{
	return a[i].end < a[j].end
}

func (a Activitys)Len() int {
	return len(a)
}

func (a Activitys)Swap(i,j int){
	a[i],a[j] = a[j],a[i]
}


// 测试11： 活动最优解
func test11(){
	acts := Activitys {
		{5,9},
		{3,5},
		{0,6},
		{1,4},
		{8,11},
		{5,7},
		{12,16},
		{2,14},
		{8,12},
		{3,9},
		{6,10},
	}

	// 按照 end 时间排序
	sort.Sort(acts)
	fmt.Println(acts)  // (1,4)(3,5)(0,6)(5,7)(5,9)(3,9)(6,10)(8,11)(8,12)(2,14)(12,16)

	// 计算活动选择最优解
	fmt.Println(activitySelection(acts)) // (1,4)(5,7)(8,11)(12,16)

}

func activitySelection(n []Activity) Activitys{

	// 创建一个 slice 用于保存有效的活动 (时间不冲突的活动)
	res := make([]Activity, 0, len(n))

	res = append(res, n[0])   // 排序后的第一个活动（活动结束的时间最早），肯定是排在第一个

	for i:=1;i<len(n);i++ {
		if n[i].start >= res[len(res) -1].end {  // 如果下一个活动的开始时间，大于 res 中末尾活动的结束时间，则可以连接上 (时间不冲突)
			res = append(res, n[i])
		}
	}
	return res
}

// 字符串拼接 比较大小
func test10Other(){
	x := "12"
	y := "45"

	str1 := x+y    // "1245"
	str2 := y+x    // "4512"

	fmt.Println(str1 > str2)   // false
}


type Numbers []string

func (n Numbers)Less(i,j int) bool{
	return n[i]+n[j] > n[j]+n[i]
}

func (n Numbers)Len() int {
	return len(n)
}

func (n Numbers)Swap(i,j int){
	n[i],n[j] = n[j],n[i]
}

// 测试10， 给定一些数字，能拼接出的最大数是 多少。
func test10(){
	num := []int{32,94,128,1286,6,71}
	numJoin := numberJoin(num)
	fmt.Println("拼接的最大数：", numJoin)
}

func numberJoin(num []int) string{
	numSlice := make(Numbers, len(num))
	for i:=0;i<len(numSlice);i++ {
		numSlice[i] = strconv.Itoa(num[i])
	}
	fmt.Printf("%#v\n", numSlice)

	//sort.Sort(sort.Reverse(sort.StringSlice(numSlice)))  // 效果与下面  重写的一样
	sort.Sort(numSlice)  // 重写 Less(), Swap(), Len() 进行自定义比较排序

	fmt.Printf("%#v\n", numSlice)

	return strings.Join(numSlice,"")   // 将字符串切片转换为 字符串，使用 Join()

	/*  使用 sort 自带方法排序： sort.Sort(sort.Reverse(sort.StringSlice(numSlice)))
	[]string{"32", "94", "128", "1286", "6", "71"}
	[]string{"94", "71", "6", "32", "1286", "128"}
	 */
}



// 定义 Len(), Less(), Swap() 这几个方法，就可以使用 sort 排序了,  Less() 里面自定义排序规则
type Good struct {
	price float64
	weight float64
}

type Goods []Good

func (g Goods)Len() int {
	return len(g)
}
func (g Goods)Less(i,j int) bool {
	return (g[i].price)/(g[i].weight) > (g[j].price)/(g[j].weight)
}

func (g Goods)Swap(i,j int){
	g[i],g[j] = g[j],g[i]
}

// 测试9：分数背包
/*
	举例：
		-- 商品1： v1=60元, w1=10斤
		-- 商品2： v1=100元, w1=20斤
		-- 商品3： v1=120元, w1=30斤

      背包容量：50斤
	  求，背包塞满，拿走的最大价值 (尽量每种商品都能拿到)
 */
func test9(){
	// 商品价格以及重量  (假设已经按照价格降序排列好)
	n := Goods {
		{60,10},    // 单价为 6
		{120,30},   // 单价为 4
		{100,20},   // 单价为 5
	}

	sort.Sort(n)   // 排序
	// fmt.Printf("%#v\n",n)
	/*  排序后：
		n := Goods {
		{60,10},    // 单价为 6
		{100,20},   // 单价为 5
		{120,30},   // 单价为 4
	}
	 */

	fractionalBackpack(n, 50)
}

// 参数1，各种商品价格以及重量;  参数2： 背包的最大容量
func fractionalBackpack(n Goods, w float64){

	fmt.Printf("背包最大容量: %f 斤\n", w)

	// 创建一个切片用于存储，每个商品可拿走的最大容量
	m := make([]float64,len(n))
	totalPrice := 0.0  // 拿走的最大价值
	totalWeight := w
	for i:=0;i<len(n);i++ {
		price := n[i].price   // 商品价格
		weight := n[i].weight  // 商品重量
		if w >= weight {
			m[i] = 1
			w = w- weight
			totalPrice = totalPrice + price
		}else {
			m[i] = w/weight  // 最后的重量除于 商品重量,可能只能拿走一部分 (小数，非1 情况)
			w = 0   // 贪心算法，最后可定要把背包的容量塞满，所以最后 背包的容量为 0
			totalPrice = totalPrice + m[i]*price
			break
		}
	}
	fmt.Printf("总共拿走：%f 斤\n", totalWeight-w)
	fmt.Printf("总共拿走价值：%f 元\n", totalPrice)
	for i:=0;i<len(m);i++ {
		fmt.Printf("商品价格 %f, 商品重量 %f,  拿走 %f 个\n",n[i].price,n[i].weight,m[i])
	}
}

// 测试8：贪心算法:
/*
	假设商店老板需要找零 m 元钱，钱币的面额有：100元，50元，20元，10元，5元，1元，
	如何找零使得所需钱币的数量最少？
 */
func test8(){

	// 人名币面额
	n := []int{100,50,20,10,5,1}
	change(n, 435)
}

// 参数1：所有面额,  参数2：消费金额
func change(n []int, m int) {

	fmt.Printf("消费金额: %d\n", m)

	// 创建一个切片 用于存储 每张面额找出多少张
	k := make([]int, len(n))
	for i,v := range n {
		k[i] = m / v
		m = m % v
	}

	for i:=0;i<len(k);i++ {
		fmt.Printf("面额 %d : %d 张\n", n[i], k[i])
	}
	fmt.Printf("余额：%d 元\n", m)
}



type QueuePath struct {
	data [][]int
	front int
	rear int
	size int
}
func (q *QueuePath) init(length int){
	q.size = length
	q.rear = 0
	q.front = 0
	q.data = make([][]int,length, length)   // 必须要初始化 length 个长度的默认值，否则下面 push 时，根据下标修改值，会找不到 元素
}

// 向队列里添加元素
func (q *QueuePath) append(elem []int){
	//fmt.Println("push : ", elem)
	if q.isFilled() {
		// fmt.Println("queue is filled!")
		q.popLeft() // 版本改进： 如果满了，则弹出一个元素 留出空间给新元素
	}
	q.rear = (q.rear +1) % q.size // rear 初始位置与 front 一致都指向 下标为0 的位置，往队列里面添加值，从 +1 的位置开始
	q.data[q.rear] = elem
	//  fmt.Println(q.data)
}
// 从队列中弹出元素
func (q *QueuePath) popLeft() ([]int, error){
	if q.isEmpty() {
		fmt.Println("queue is empty!")
		return []int{-1,-1}, fmt.Errorf("queue is empty!")
	}
	q.front = (q.front +1) % q.size  // front 始终指向的是没有元素的那一位 (也就是 数组类型的默认值，它的下一个位置才是第一个元素)
	return q.data[q.front], nil
}
// 判断队列里是否为空
func (q *QueuePath) isEmpty() bool{
	return q.rear == q.front   // rear == front 代表队列里没元素，为空
}
// 判断队列里是否已满
func (q *QueuePath) isFilled() bool{
	return (q.rear+1) % q.size == q.front  // rear +1 的位置 等于 front 则代表 队列已满
}


// 队列的应用： 迷宫问题
func test7Other(){
	// 迷宫数组 (1 代表围墙，0代表可走路线)
	maze := [][]int {
		{1,1,1,1,1,1,1,1,1,1},
		{1,0,0,1,0,0,0,1,0,1},
		{1,0,0,1,0,0,0,1,0,1},
		{1,0,0,0,0,1,1,0,0,1},
		{1,0,1,1,1,0,0,0,0,1},
		{1,0,0,0,1,0,0,0,0,1},
		{1,0,1,0,0,0,1,0,0,1},
		{1,0,1,1,1,0,1,1,0,1},
		{1,1,0,1,0,0,0,0,0,1},
		{1,1,1,1,1,1,1,1,1,1},
	}

	mazePathByQueue(maze, 1,1, 8,8)
}

func pathReverse(path [][]int){
	for i:=0;i<len(path)/2;i++ {
		path[i],path[len(path)-i-1] = path[len(path)-i-1], path[i]
	}
}

// 通过队列实现 迷宫的广度搜索
/*	原理：
	1, 多个方向一同寻找， 不像 栈的实现， 一条道走到黑，走不通再回退，继续找方向
	2, 通过队列的形式保存 每条路线的上一个可走的节点，通过循环实现 不同路线的寻找 (与栈不同的是，这里除了要保存坐标还要保存 上一个节点的 index ,具体关系查看 path 的定义)
	3, 为了实现最后 终点回溯至起点的操作，需要创建一个 数组用来保存 每条路线可走节点的坐标以及上一个节点的index
	4, 队列与 path 都是保存坐标与 上一个节点的 index, 那有什么区别呢?
		1, 队列里的坐标是 动态变化的，每找到一条路线 或者 可走的节点  都会 pop 出上一条路线的 可走的节点
		2, path 里保存的是历史可走的节点 (每条路线可走的节点都会保存下来)
 */
func mazePathByQueue(maze [][]int, x1,y1,x2,y2 int){

	// 创建一个 path 用于保存可走的节点，以及 这个节点的 上一个可走的节点的 index (用于最后 反向寻找 终点至起点之间的所有节点)
	// (index , 当前节点，上一个节点的 关系是) :  上一个节点的位置是 index,  当前节点 是由上一个节点引出来的节点
	path := make([][]int, 0,20)

	q := new(QueuePath)
	q.init(20)
	q.append([]int{1,1, -1}) // []int{x,y,index}   // 当前节点x，y的坐标以及 上一个节点的 index

	for !q.isEmpty() {
		curNode, _ := q.popLeft()  // 前一个点出栈，根据前一个点寻找 下面可走的节点 加入到 队列里
		// 注意：将下一个可走的节点加入到队列中，一旦寻找到 下一个可走的节点，则弹出上一个节点，
		// 注意：一旦出现分叉路：则将不同的路线的节点 加入到 队列里
		path = append(path, curNode)   // 保存可走的路径节点 (广度搜索，保存 不同路线的 可走的节点), 节点：[]int{x,y,index}
		if curNode[0] == x2 && curNode[1] == y2 {
			// 到达终点
			realPath := make([][]int, 0,20)
			i := len(path) -1    // 最后的一个路径位置  (到达终点的位置)
			for i>=0 {   // 通过上一个节点的index 倒序查找路线至 起点, （当i 等于0 时，代表是起点）
				realPath = append(realPath, path[i][0:2])   // 取路径的 x,y坐标
				i = path[i][2]   // 取上一个节点的 index
			}
			pathReverse(realPath)  // 再将路径反转 得到 起点至终点的位置
			fmt.Println("路线图：", realPath)
			return
		}

		for i:=0;i<4;i++ {
			nextNode := dir(curNode[0], curNode[1], i)
			// 如果下一个节点等于 0， 则代表 可以走
			if maze[nextNode[0]][nextNode[1]] == 0 {
				q.append([]int{nextNode[0],nextNode[1],len(path)-1}) // []int{x,y,index}   // 当前节点x，y的坐标以及 上一个节点的 index
				maze[nextNode[0]][nextNode[1]] = 2 // 并且标记为 2，代表已经走过，则不会再走
			}
		}

	}
}

// 栈的 应用： 迷宫问题
func test7() {

	// 迷宫数组 (1 代表围墙，0代表可走路线)
	maze := [][]int {
		{1,1,1,1,1,1,1,1,1,1},
		{1,0,0,1,0,0,0,1,0,1},
		{1,0,0,1,0,0,0,1,0,1},
		{1,0,0,0,0,1,1,0,0,1},
		{1,0,1,1,1,0,0,0,0,1},
		{1,0,0,0,1,0,0,0,0,1},
		{1,0,1,0,0,0,1,0,0,1},
		{1,0,1,1,1,0,1,1,0,1},
		{1,1,0,1,0,0,0,0,0,1},
		{1,1,1,1,1,1,1,1,1,1},
	}

	mazePathByStack(maze, 1,1, 8,8)
}


// 参数1：迷宫,  参数2,3 代表起点坐标， 参数4,5 代表 终点坐标
// 通过 栈 实现迷宫的深度搜索
func mazePathByStack(maze [][]int, x1,y1,x2,y2 int){

	// 创建一个 [][]int 模拟栈的操作，存储 路线 (寻找迷宫可走的每一步)
	stack := make([][]int,0,20)
	stack = append(stack, []int{x1,y1})  // 起点的 (x,y) 坐标

	// 如果 stack 为空，则代表没有路
	for len(stack) > 0 {
		curNode := stack[len(stack)-1]   // 栈顶元素 (当前可以走的坐标) 其实是 上一次可以走的坐标 (请仔细理解以下 curNode 与 nextNode)
		// fmt.Println(curNode)
		// 如果栈顶元素等于 终点坐标，则代表 已走到终点
		if curNode[0] == x2 && curNode[1] == y2 {

			fmt.Println("找到路了!")
			// 循环打印，栈里存储的 路线
			for _,v := range stack {
				fmt.Println(v)
				/*
						找到路了!
						[1 1]
						[2 1]
						[3 1]
						[4 1]
						[5 1]
						[5 2]
						[5 3]
						[6 3]
						[6 4]
						[6 5]
						[7 5]
						[8 5]
						[8 6]
						[8 7]
						[8 8]
				*/
			}
			return
		}
		// x,y 的四个方向:  上: [x-1,y] ; 下：[x+1,y] ; 左：[x,y-1] ; 右: [x,y+1]
		nextNode := []int{}
		isFindPath := false
		for i:=0;i<4;i++ {
			nextNode = dir(curNode[0], curNode[1], i)
			// 如果下一个节点等于 0， 则代表 可以走
			if maze[nextNode[0]][nextNode[1]] == 0 {
				stack = append(stack, nextNode)   // 将下一个可走的节点 加入到 stack
				maze[nextNode[0]][nextNode[1]] = 2 // 并且标记为 2，代表已经走过，则不会再走
				isFindPath = true  // 找到可以走的下一步路
				break
			}
		}
		// 如果下一个节点 四个方向都走不通
		if !isFindPath {
			// 如果下一个节点走不通，则也标记为2， 不会再走
			maze[nextNode[0]][nextNode[1]] = 2
			stack = stack[:len(stack)-1]   // 弹出栈顶的元素, 回退路线, 回退到上一个可以走的路线，然后再继续四个方向寻找
		}
	}

	if len(stack) == 0{
		fmt.Println("没有路")
	}
}
// 当前节点x, y 坐标，和 指定方向
// 寻找方向 (上右下左，或者 下左上右，等等等都可以，只是你给的方向不一样，它找出来的路径就会不一样)
func dir(curNodeX,curNodeY int, ori int) []int{
	if ori == 1 {  // 上
		return []int{curNodeX-1, curNodeY}
	}
	if ori == 3 {  // 右
		return []int{curNodeX, curNodeY+1}
	}
	if ori == 0 {  // 下
		return []int{curNodeX+1, curNodeY}
	}
	if ori == 2 {  // 左
		return []int{curNodeX, curNodeY-1}
	}
	return []int{0,0}
}

type Queue struct {
	data []string
	front int
	rear int
	size int
}
// 模拟 队列的先入先出， 环形结构
func (q *Queue) init(length int){
	q.size = length
	q.rear = 0
	q.front = 0
	q.data = make([]string,length, length)   // 必须要初始化 length 个长度的默认值，否则下面 push 时，根据下标修改值，会找不到 元素
}
/*
	注意： (以下解释说明针对 Queue 的第一版  (data []int,  并且 push 的时候当满的时候 则弹出错误，不能添加新元素， 第二版改进版中这里做了处理))
	1, % 运算是为了 实现环形结构，有效的利用里面的空间
		示例：队列里容量为 size == 5， push了 4个元素, rear == 4, front 为 0(起始位置)
			如果还要向 队列里添加值，则 (4+1) % 5 == 0,   (rear+1)/size == front, 则代表已满，不能添加元素

			此时 pop 了 3个元素出去，此时 front == 3,  继续 push 元素， push 了 3个元素,
			rear == (rear+3) % size == (4+3) % 5 == 2, 此时 rear 指向下标为 2 的位置，如果不用 % ,则会出现 rear == rear+3 == 7,下标越界的问题
	2, pop 函数 是真的从 数组中删除了 那些元素吗？  其实不是的， 它只是 移动了一个位置的下标，
			队列里 元素的个数与实际的值的范围是 rear 和 front 之间的范围决定，(注意由于是环形结构，所以 rear 可能大于 front ,有可能小于 front)

		并且之前那些pop 的元素，到后面新添加的元素 会将 pop 掉的元素的位置霸占，重新赋值，有效的利用了空间
		所以也就是 为什么 push() 里面的实现是用的 根据下标修改值，而不是简单的 append()
 */

// 向队列里添加元素
func (q *Queue) push(elem string){
	//fmt.Println("push : ", elem)
	if q.isFilled() {
		// fmt.Println("queue is filled!")
		q.pop() // 版本改进： 如果满了，则弹出一个元素 留出空间给新元素
	}
	q.rear = (q.rear +1) % q.size // rear 初始位置与 front 一致都指向 下标为0 的位置，往队列里面添加值，从 +1 的位置开始
	q.data[q.rear] = elem
	fmt.Println(q.data)
}
// 从队列中弹出元素
func (q *Queue) pop() (string, error){
	if q.isEmpty() {
		fmt.Println("queue is empty!")
		return "", fmt.Errorf("queue is empty!")
	}
	q.front = (q.front +1) % q.size  // front 始终指向的是没有元素的那一位 (也就是 数组类型的默认值，它的下一个位置才是第一个元素)
	return q.data[q.front], nil
}
// 判断队列里是否为空
func (q *Queue) isEmpty() bool{
	return q.rear == q.front   // rear == front 代表队列里没元素，为空
}
// 判断队列里是否已满
func (q *Queue) isFilled() bool{
	return (q.rear+1) % q.size == q.front  // rear +1 的位置 等于 front 则代表 队列已满
}

// 环形结构 的应用 :  读取文件内容的后5行
func test6Other(){

	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Println("file open failed : ", err)
		return
	}
	defer file.Close()

	queue := new(Queue)
	queue.init(5)

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			if len(strings.TrimSpace(string(line))) != 0 {
				queue.push(string(line))
			}
			break
		}
		if err != nil {
			fmt.Println("readline err : ", err)
			break
		}
		queue.push(string(line))
	}

	fmt.Printf("队列里的内容：%#v\n",queue.data)

	/*  打印内容：
		队列里的内容：[]string{"10ddddddddddd", "11fggbbbbbbbb", "12fdffssssssssddddd", "13vvvvvfffffff", "14gggghhhhhh"}
	 */
}


// 测试6，队列：  使用 slice 实现队列的先进先出 模式
/*	环形队列介绍：
	环形队列：当队尾指针 rear/front == MaxSize -1 时，再前进一个位置就自动到 0， (实现 环形结构)

		-- 队首指针前进1：front = (front+1) % MaxSize
        -- 队尾指针前进1: rear = (rear+1) % MaxSize
		-- 队空条件：     rear == front
        -- 队满条件：     (rear+1) % MaxSize == front
 */
func test6(){

	queue := new(Queue)
	queue.init(5)
	fmt.Println("isEmpty : ", queue.isEmpty())

	queue.push("12")
	queue.push("13")
	queue.push("14")
	queue.push("15")
	fmt.Println("isFilled : ", queue.isFilled())

	// fmt.Println(queue.pop())

	queue.push("16")
	queue.push("17")
	queue.push("18")

	fmt.Println(queue.data)


	/*

		打印信息如下：
		第一版与第二版的区别： push() 里面 当判断已满的时候，第一个版本不让添加新元素，
					第二个版本弹出旧元素，让新元素进来, 但第二个版本弹的是 front 那个0 下标还是 1 下标的值，这里貌似还有点而小问题

		第二版：
		isEmpty :  true
		[ 12   ]
		[ 12 13  ]
		[ 12 13 14 ]
		[ 12 13 14 15]
		isFilled :  true
		[16 12 13 14 15]   //然后再 push ,它会把 之前 12 的那个位置给利用上， 这就是环形结构， 有效的利用了空间
		[16 17 13 14 15]
		[16 17 18 14 15]
		[16 17 18 14 15]



		// 第一版：
		isEmpty :  true

		[0 12 0 0 0]       // push 第一个元素，给的是 数组的第 1 个元素，而不是 第 0 个
		[0 12 13 0 0]
		[0 12 13 14 0]
		[0 12 13 14 15]    // 12，13，14，15，依次被加入

		isFilled :  true   // 5 个长度的数组，实际上只能添加 4个元素  (N-1)
		12 <nil>           // 弹出了 第一个push 的元素  12
		[17 12 13 14 15]

	*/
}

// 测试5：判断 括号的有效性
/*
	给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

	有效字符串需满足：

	左括号必须用相同类型的右括号闭合。
	左括号必须以正确的顺序闭合。
	注意空字符串可被认为是有效字符串。
 */
func test5(){
	//s := "{[({[ ()] [( )]})] } "
	// s := "(])"

	s := "  [ "
	fmt.Println(isValid(s))    // 第一版

	fmt.Println(isValid2(s))   // 第二版 （代码复杂，但是 执行效率以及内存占用 都达到了很大的优化）
}


// 使用自定义类型，模拟栈的操作 (先入后出)
type MyByte []byte
// (不能值传递) 传入指针，方便修改值
func (b *MyByte) push(by byte){
	*b = append(*b, by)
}
func (b *MyByte) pop(){
	*b = (*b)[:len(*b)-1]
}
func (b *MyByte) isEmpty() bool{
	return len(*b) == 0
}
func (b *MyByte) getTop() byte {
	return (*b)[len(*b)-1]
}

// 优化第二版： 主要优化内存占用这一块儿，  将 第一版isValid() 中的 data 和 map 都做了最小优化，降低内存占用
func isValid2(s string) bool{
	// 如果s 里面只有空字符串则直接返回 true ，不需要执行下面的操作
	if len(strings.TrimSpace(s)) == 0 {
		return true
	}

	// 创建一个 map, 用于保存 有效括号的 对应关系
	m := map[byte]byte{')':'(','}':'{',']':'['}   // 优化前：m := map[int32]string{')':"(", '}':"{",']':"["}
	// 创建一个 []byte 模拟 栈的先进后出操作
	b := new(MyByte)    // 优化前：data := make([]string,0,len(s))
	for i,v := range s {
		if v == '(' || v == '{' || v == '['{
			b.push(s[i])
		} else {
			// 如果是 ' ' 则不做处理
			if v == ' ' {  // // 题目要求 空字符串也是有效字符，所以
				continue
			}
			if b.isEmpty() {  // 如果右括号进来，但是 data 已经为空，则代表没有 左括号与之匹配，则返回 false
				return false
			}
			// 如果 v 等于 data 末尾的元素 (想象成 栈顶), 那么就删除末尾的元素 (类似于 出栈)
			if b.getTop() == m[s[i]] {  // v 是 ')', 则 m[v]对应 "("
				b.pop()
			}else {   // 示例："(])",  '( 被加入到 data 中, 然后] 进来 通过 map 查询对应结果为 '[', '('与'[' 不匹配，  if data[len(data)-1] == m[v]  不满足，则代表不匹配，也要返回 false
				return false
			}
		}
	}
	return b.isEmpty()
}

// 第一版 （内存还需要优化）
func isValid(s string) bool {

	// 如果题目要求空字符串不算有效字符，那么可以如下判断
	/*
	// 如果字符串s 里面不包含任何一个 如下字符，则返回 false
	if !strings.ContainsAny(s, "(){}[]") {
		return false
	}
	*/

	// 如果s 里面只有空字符串则直接返回 true ，不需要执行下面的操作
	if len(strings.TrimSpace(s)) == 0 {
		return true
	}

	// 创建一个 map, 用于保存 有效括号的 对应关系
	m := map[int32]string{')':"(", '}':"{",']':"["}
	// 创建一个 slice 模拟 栈的先进后出操作
	data := make([]string,0,len(s))
	for _,v := range s {
		if v == '(' || v == '{' || v == '['{
			data = append(data, string(v))
		} else {
			// 如果是 ' ' 则不做处理
			if v == ' ' {  // // 题目要求 空字符串也是有效字符，所以
				continue
			}
			if len(data) == 0 {  // 如果右括号进来，但是 data 已经为空，则代表没有 左括号与之匹配，则返回 false
				return false
			}
			// 如果 v 等于 data 末尾的元素 (想象成 栈顶), 那么就删除末尾的元素 (类似于 出栈)
			if data[len(data)-1] == m[v] {  // v 是 ')', 则 m[v]对应 "("
				data = data[:len(data)-1]
			}else {   // 示例："(])",  '( 被加入到 data 中, 然后] 进来 通过 map 查询对应结果为 '[', '('与'[' 不匹配，  if data[len(data)-1] == m[v]  不满足，则代表不匹配，也要返回 false
				return false
			}
		}
	}
	return len(data) == 0
}

// 占用内存最少 (此方法要保证 numbers 是有序数组(升序))
func twoSum4(numbers []int, target int) []int {
	if len(numbers) <= 1 {
		return []int{0,0}
	}
	// i, j 初始值为 数组开始和末尾位置
	for i,j:=0, len(numbers)-1; i<j; {
		// 两数相加等于 target ，就返回坐标
		if numbers[i]+numbers[j] == target {
			return []int{i+1, j+1}
		}
		// 两数相加大于 target, 则j 往左移 (寻找小一点儿的值)
		if numbers[i]+numbers[j] > target {
			j--
		}else {  // 两数相加小于 target, 则i 往右移 (寻找大一点儿的值)
			i++
		}
	}

	return []int{0,0}
}


// 速度最快 (此方法不需要用到二分查找，就避免了数组一定要有序的前提，无序状态，此方法依旧可以找到下标值)
// 同样是使用 map ，为什么如下实现没有出现 同一个值返回相同的下标呢，

// 关键点：
// 0, map 不是提前遍历数组保存值，而是一边遍历，一边查询
// 1，map中依次保存元素的值，通过取 差值去确定是否有 差值存在map中， 找到就直接返回，
// 2, 返回的时候 j 从map 中查找下标，i 从 数组中取出，所以不会冲突

/*
	示例：[2,3,6], 求5

		第一次 2进来, map[5-2]， map[3] 找不到值，则将 2 对应的下标保存起来, m[2] = 0
		进行迭代， 3进来， map[5-3], map[2]就存在了，就返回 2 的坐标， 3 的坐标直接取 i 就行了

    示例：[0,0,2,3,6], 求0
		第一次 0 进来，map[0-0]， map[0]找不到值，则将 0 对应的下标保存起来， m[0] = 0
		进行迭代，0 进来, map[0-0], map[0]存在, 就返回 0 的坐标，第二次 0 的坐标直接取 i 就行了，所以是 [0,1]
*/

func twoSum3(numbers []int, target int) []int {
	// index 负责保存map[整数]整数的序列号
	index := make(map[int]int, len(numbers))

	for i, b := range numbers {
		// 通过查询map，获取a = target - b的序列号
		if j, ok := index[target-b]; ok {
			return []int{j+1, i+1}
			// 注意，顺序是j，i, j放前面的目的是 保证小数在前，(因为前提是 numbers 里面有升序排列)
		}
		// 把b和i的值，存入map
		index[b] = i   // 如果
	}

	return nil
}

// 二维数组的排序
func test3Other(){

	data := [][]int{{1,3,4,5},{0,0,45},{8,9,6,15}}

	newData := make([]int, 0,10)
	for i:=0;i<len(data);i++ {
		newData = append(newData, data[i]...)   // 将二维数组中的元素 全部追加到 newData 中
	}

	// 对 newData 进行排序
	sort.Ints(newData)   // [0 0 1 3 4 5 6 8 9 15 45]
	fmt.Println(newData)

}


// 测试4：给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
/*
	说明:

	初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
	你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
	 

	示例:

	输入:
	nums1 = [1,2,3,0,0,0], m = 3
	nums2 = [2,5,6],       n = 3

	输出: [1,2,2,3,5,6]
 */

func test4(){
	nums1 := []int{1,2,3,0,0,0,0}   // , m = 3
	nums2 := []int{2,5,6}       //   n = 3

	mergeArr(nums1,3,nums2,3)
	fmt.Println(nums1)
}

func mergeArr(nums1 []int, m int, nums2 []int, n int)  {
	if len(nums1) - m >= n {
		copy(nums1[m:], nums2)    // 先拷贝
		sort.Ints(nums1)          // 再使用 sort 排序
	}
}


// 测试3 ： 给定一个列表和一个整数，设计算法找到两个数的下表，是的两个数之和为给定的整数。保证肯定仅有一个结果。
/*
	要求：返回的下标值（index1 和 index2）不是从零开始的。
		你可以假设每个输入只对应唯一的答案，而且你不可以重复使用相同的元素。

	示例：
		输入: numbers = [2, 7, 11, 15], target = 9
		输出: [1,2]
		解释: 2 与 7 之和等于目标数 9 。因此 index1 = 1, index2 = 2 。
 */
func test3(){
	data := []int{1,3,4,5,0,0,45,8,9,6,15}
	// fmt.Println(twoSum(data, 0))

	fmt.Println(twoSum4(data, 15))
}
// 自己实现的方法，第一种map 的方式有漏洞， 第二种 [][]int 的方式 效率并不高，并且内存占用不少 (推荐使用 twoSum3() (更通用，速度更快) 或 twoSum4()(内存占用更少))
func twoSum(nums []int, target int) []int {

	/*  map 存储的方式不适用于所有情况，例如：{1,0,0,5,23,34,45}, 求两数和值为0 的下表， map 方式 只会返回一个 相同的下标[2,2] ,而不是 [1,2]
	// 创建一个 map 用于存储nums 里面的元素 和 下标
	m := make(map[int]int, len(nums))
	for i:=0;i<len(nums);i++ {
		m[nums[i]] = i     // map 中key 为nums 元素值， value 为 元素对应的下标
	}

	sort.Sort(sort.IntSlice(nums))   // 将 nums 进行排序, 便于 下面的二分查找

	 */

	s := make([][]int, len(nums))
	for i:=0;i<len(nums);i++ {
		s[i] = append(s[i], nums[i], i)   //
	}

	// sort.Sort(sort.IntSlice(s[]))

	for i:=0;i<len(s);i++ {
		a := s[i][0]
		b := target - a
		j := -1
		// 如果 b>=a 代表 b 在 a 的右侧区间,  从 右侧区间查找 b,
		if b >= a {
			// 如果找到 返回 a 值和 b 值对应的 下标
			j = binarySearch(s, i+1, len(s)-1, b)
		}
		// 如果 b<a 代表 b 在 a 的左侧区间,  从 左侧区间查找 b,
		if b < a {
			// 如果找到 则从 map 中找到 a 值和 b 值对应的 下标
			j = binarySearch(s, 0, i-1,b)
		}
		if j != -1 {
			return []int{s[i][1], s[j][1]}    // s[][], 后一个[], [0]代表值，[1]代表值对应的下标
		}
	}
	// 如果没找到，则返回 [-1,-1]
	return []int{-1,-1}
}

// 二分查找 target 的下标 (前提是数组中的元素是有序的)
func binarySearch(nums [][]int, left, right, val int) int{
	// 给定一个 left 与 right ,固定 区间
	// 在区间内再进行二分查找，直到找到 val
	for left <= right {
		mid := (left+right)/2
		if nums[mid][0] == val {  // 如果找到则返回 true
			return mid
		}
		if nums[mid][0] > val {
			right = mid -1
		}
		if nums[mid][0] < val {
			left = mid +1
		}
	}
	return -1  // 如果找不到则返回 false
}

// 测试2：给定一个 m*n 的二维列表，查找一个数是否存在。 列表有下列特性：
/*
	每一行的列表从左到右已经排序好。
	每一行第一个数比上一行最后一个数大。
    [
		[1,3,5,7],
		[10,11,16,20],
		[23,30,34,50]
	]
 */
func test2(){

	 data := [][]int {{1,3,5,7},{10,11,16,20},{23,30,34,50},{54,55,56,57},{64,65,66,67},{74,75,76,77},{84,85,86,87}}

	// data := [][]int {{}}
	fmt.Println(searchMatrix(data, 88))
}

// 第三版：最后优化版 （递归 二分查找，并且加入 二维数组长度的判断 避免取值时数组越界）
func searchMatrix(matrix [][]int, target int) bool {
	dLen := len(matrix)
	// 避免出现 matrix = [][]int {{}}
	if dLen == 0 {
		return false
	}

	i:= dLen/2
	for i >=0 && i<dLen {
		// 避免出现 matrix = [][]int {{}}
		if len(matrix[i]) >0 && target <matrix[i][0] {  // 如果 target 小于该数组的第一个元素，则代表 target 有可能出现在 matrix[0:i][:] 范围的数组中
			return searchMatrix(matrix[0:i][:], target)
		}
		if len(matrix[i]) >0 && target > matrix[i][len(matrix[i])-1] {  // 如果 target 大于该数组的最后一个元素，则代表 target 有可能出现在 matrix[i+1:][:] 范围的数组中
			return searchMatrix(matrix[i+1:][:], target)
		}
		// 这里还可以优化（缩小区间）
		// 如果既不小于 第一个元素，也不大于第二个元素，则 target 可能在该数组中
		for j:=0;j<len(matrix[i]);j++ {
			if target == matrix[i][j]{
				return true
			}
		}
		// 如果能执行到这里，则代表 在 matrix[i] 中没有找到 target 这个数字
		return false
	}
	return false
}


// 第二版：递归 二分查找 (采用这个方案)
func isExsit(data [][]int, num int) bool{
	dLen := len(data)
	fmt.Println(dLen)
	i:= dLen/2
	for i >=0 && i<dLen {
		if num <data[i][0] {  // 如果 num 小于该数组的第一个元素，则代表 num 有可能出现在 data[0:i][:] 范围的数组中
			fmt.Println("less")
			return isExsit(data[0:i][:], num)
		}
		if num > data[i][len(data[i])-1] {  // 如果 num 大于该数组的最后一个元素，则代表 num 有可能出现在 data[i+1:][:] 范围的数组中
			fmt.Println("more")
			return isExsit(data[i+1:][:], num)
		}
		// 这里还可以优化（缩小区间）
		// 如果既不小于 第一个元素，也不大于第二个元素，则 num 可能在该数组中
		for j:=0;j<len(data[i]);j++ {
			if num == data[i][j]{
				return true
			}
		}
		// 如果能执行到这里，则代表 在 data[i] 中没有找到 num 这个数字
		return false
	}
	return false
}

// 第一版：从中间查找，小于则从前面依次倒叙查找， 大于则从后面顺序查找
func isExsit2(data [][]int, num int) bool{
	dLen := len(data)
	fmt.Println(dLen)
	i:= dLen/2
	for i >=0 && i<dLen {
		if num <data[i][0] {  // 如果 num 小于该数组的第一个元素，则代表 num 有可能出现在前一个数组中
			i--    // 移动到前面一个数组 进行比较
			fmt.Println("i--")
			continue
		}
		if num > data[i][len(data[i])-1] {  // 如果 num 大于该数组的最后一个元素，则代表 num 有可能出现在后一个数组中
			i++   // 移动到后面一个数组进行比较
			fmt.Println("i++")
			continue
		}
		// 这里还可以优化（缩小区间）
		// 如果既不小于 第一个元素，也不大于第二个元素，则 num 可能在该数组中
		for j:=0;j<len(data[i]);j++ {
			if num == data[i][j]{
				return true
			}
		}
		// 如果能执行到这里，则代表 在 data[i] 中没有找到 num 这个数字
		return false
	}
	return false
}

// string 排序
func test1(){

	s := "anagram"
	t := "nagaram"

	// fmt.Println(isAnagram(s,t))
	start := time.Now().UnixNano()
	isAnagram(s,t)
	end := time.Now().UnixNano()
	fmt.Println(end-start)
}
// 测试1：判断s 与 t 两个字符串 是否为 相同字符异序
func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	//  使用 map (可避免 key 值越界, 如果创建一个 []int slice 则要考虑 下标越界的问题)
	mLen := len(s)
	if len(s) > 52 {   // 包括大写字母
		mLen = 52
	}

	m := make(map[byte]int, mLen)
	for i:=0;i<len(s);i++ {
		m[s[i]] ++   // 统计 s 中每个字符的个数
		m[t[i]] --   // 如果 t 中某个字符等于 s 中的某个字符，则减减 会抵消个数的统计
	}

	fmt.Printf("%#v\n", m)

	for _,v := range m {
		if v != 0 {    // 不等于0 代表，统计到了个数 (只有 不同的字符才会被统计到，所以两个字符串是不符合要求的)
			return false
		}
	}

	return true
}

// slice 排序 (可用系统自带的 sort.Strings() 排序)
func test1Other(){
	ss := []string{"surface", "ipad", "mac pro", "mac air", "think pad", "idea pad"}

	sort.Strings(ss)

	fmt.Printf("%#v\n", ss)
}

/*
		示例： s="anagram", t="nagaram", return true
              s="rat", t="car", return false
*/
