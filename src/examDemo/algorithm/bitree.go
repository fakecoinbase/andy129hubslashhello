package main

import "fmt"


type QueueLevel struct {
	data []*BiTreeNode
	front int
	rear int
	size int
}
func (q *QueueLevel) init(length int){
	q.size = length
	q.rear = 0
	q.front = 0
	q.data = make([]*BiTreeNode,length, length)   // 必须要初始化 length 个长度的默认值，否则下面 push 时，根据下标修改值，会找不到 元素
}

// 向队列里添加元素
func (q *QueueLevel) append(elem *BiTreeNode){
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
func (q *QueueLevel) popLeft() (*BiTreeNode, error){
	if q.isEmpty() {
		fmt.Println("queue is empty!")
		return &BiTreeNode{}, fmt.Errorf("queue is empty!")
	}
	q.front = (q.front +1) % q.size  // front 始终指向的是没有元素的那一位 (也就是 数组类型的默认值，它的下一个位置才是第一个元素)
	return q.data[q.front], nil
}
// 判断队列里是否为空
func (q *QueueLevel) isEmpty() bool{
	return q.rear == q.front   // rear == front 代表队列里没元素，为空
}
// 判断队列里是否已满
func (q *QueueLevel) isFilled() bool{
	return (q.rear+1) % q.size == q.front  // rear +1 的位置 等于 front 则代表 队列已满
}



type BiTreeNode struct {
	data string           // 当前节点数据
	lChild *BiTreeNode    // 左孩子节点
	rChild *BiTreeNode    // 右孩子节点
}

func testNode(){
	a := &BiTreeNode{
		data:   "A",
	}

	b := &BiTreeNode{
		data:   "B",
	}

	c := &BiTreeNode{
		data:   "C",
	}

	d := &BiTreeNode{
		data:   "D",
	}

	e := &BiTreeNode{
		data:   "E",
	}

	e.lChild = a
	e.rChild = b
	b.lChild = c
	b.rChild = d

	/*    二叉树关系模型：
				E
	         A        B
                   C      D
	 */

	// fmt.Println(e.rChild.data)   // B

	preRange(e)         //  E,A,B,C,D,

	fmt.Println()

	inOrderRange(e)     //  A,E,C,B,D,

	fmt.Println()

	postOrderRange(e)   // A,C,D,B,E,

	fmt.Println()

	levelOrderRange(e)   // E,A,B,C,D,
}

// 前序遍历  (根据 根节点，遍历所有节点的值)
func preRange(node *BiTreeNode){
	if node != nil {   // 判断节点是否为 nil 
		fmt.Print(node.data,",")
		preRange(node.lChild)   // 遍历左节点
		preRange(node.rChild)   // 遍历右节点
	}
}

// 中序遍历 (先遍历左子节点，再打印自己，最后遍历右子节点)
func inOrderRange(node *BiTreeNode){
	if node != nil {   // 判断节点是否为 nil
		inOrderRange(node.lChild)   // 遍历左节点
		fmt.Print(node.data,",")
		inOrderRange(node.rChild)   // 遍历右节点
	}
}

// 后续遍历 (先遍历左子节点，再遍历右子节点，最后打印自己)
func postOrderRange(node *BiTreeNode){
	if node != nil {   // 判断节点是否为 nil
		postOrderRange(node.lChild)   // 遍历左节点
		postOrderRange(node.rChild)   // 遍历右节点
		fmt.Print(node.data,",")
	}
}

// 层次遍历 (使用队列)
func levelOrderRange(node *BiTreeNode){
	q := new(QueueLevel)
	q.init(20)
	q.append(node)    // 添加 根节点入队列
	for !q.isEmpty() {   // 记住要用这个方法，而不是直接 len(q.data)
		node,_ := q.popLeft()
		fmt.Print(node.data,",")
		if node.lChild != nil {
			q.append(node.lChild)
		}
		if node.rChild != nil {
			q.append(node.rChild)
		}
	}
}

// 判断结构体对象是否为空
func testOther(){
	// 1, 结构体为指针类型时:   可直接与 nil 比较
	node := &BiTreeNode{
		data:   "A node",
		lChild: nil,
		rChild: nil,
	}
	fmt.Println(node == nil)   // false

	// 2, 结构体为值类型时：
	node2 := BiTreeNode{}

	fmt.Println(node2 == BiTreeNode{})   // true

	/*  或者定义一个 方法，方法内部使用 reflect.DeepEqual() 方法
		func (b BiTreeNode) IsEmpty() bool {
	    	return reflect.DeepEqual(b, BiTreeNode{})
		}
	*/
}

// 二叉树
func main() {
	testNode()

	// testOther()
}
