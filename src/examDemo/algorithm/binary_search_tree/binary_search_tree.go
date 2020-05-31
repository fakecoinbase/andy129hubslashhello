package main

import "strconv"
import "fmt"

/*  学习总结：

	1, 节点的删除，其实就是 子节点与节点 关系的解除过程，或者 子节点位置的上升，与其他节点 关系建立的过程
	2，go 语法：
			注意 if() if() ， 与 if() else if() 的调用规则

 */

// 二叉搜索树
type BST struct {
	root *BiTreeNode
}

func NewBST(d []int) *BST{

	b := &BST{root:nil,}
	// 初始化二叉搜索树
	for _,v := range d {
		// fmt.Println(v)
		b.root = b.insert(b.root, v)  // 从根节点开始插入数据
	}
	return b
}
// 递归插入法
func (b *BST) insert(node *BiTreeNode, val int) *BiTreeNode{
	if node == nil {
		node = NewBiTreeNode(val)
	}else {
		if val < node.data {
			node.lChild = b.insert(node.lChild, val)
			node.lChild.parent = node
		}
		if val > node.data {
			node.rChild = b.insert(node.rChild, val)
			node.rChild.parent = node
		}
	}
	return node
}
// 非递归插入法
func (b *BST) insertNoRec(val int){
	p := b.root
	if p == nil {
		p = NewBiTreeNode(val)
		return
	}
	for {
		if val < p.data {
			if p.lChild != nil {  // 如果 p.lChild 不为 nil ,则将 当前节点赋值为 当前节点的 左孩子节点，继续往下面找
				p = p.lChild      // 这里赋值的目的就是 继续往下面找合适的位置
			}else {   // 如果 p.lChild 为nil , 则就根据 val 创建一个 BiTreeNode, 插入到 p.lChild的位置上，并指定关系
				p.lChild = NewBiTreeNode(val)
				p.lChild.parent = p  // 指定关系
				return
			}
		}else if val > p.data {
			if p.rChild != nil {   // 这里是为了继续往下查找合适的位置
				p = p.rChild
			}else {   // 如上，这里是真正的插入操作
				p.rChild = NewBiTreeNode(val)
				p.rChild.parent = p
				return
			}
		}else {
			return
		}
	}
}
// 递归查询法
func (b *BST)query(node *BiTreeNode, val int) *BiTreeNode{
	if node == nil {
		return nil
	}
	if node != nil {
		// 从 lChild 继续找
		if val < node.data {
			return b.query(node.lChild,val)   // 注意，递归调用，有返回值则要返回
		}else if val > node.data {  // 从 rChild 继续找
			return b.query(node.rChild,val)
		}else { // 相等，则返回该 node
			return node
		}
	}
	return nil
}

// 非递归查询法
func (b *BST)queryNoRec(val int) *BiTreeNode{
	p := b.root
	for p != nil {
		if val < p.data {
			p = p.lChild
		}else if val > p.data {
			p = p.rChild
		}else {
			return p
		}
	}
	return nil
}
// Delete 是二叉搜索树的  删除方法
func Delete(b *BST, val int) {
	if b.root != nil {
		node := b.queryNoRec(val)
		if node == nil {
			fmt.Println(val, "不存在")
			return
		}
		if node.lChild == nil && node.rChild == nil {  // 情况1
			b.delete1(node)
		}else if node.rChild == nil {     // 情况2.1  只有一个左孩子节点
			b.delete21(node)
		}else if node.lChild == nil {     // 情况2.2  只有一个右孩子节点
			b.delete22(node)
		}else {
			// 找到右节点最小节点 (将值赋值给当前节点，然后删除最小节点，当前节点的关系依旧, 只更改值，不更改关系，达到删除的目的)
			minNode := node.rChild
			for minNode.lChild != nil {
				minNode = minNode.lChild
			}
			node.data = minNode.data  // 将右节点最小节点的值赋值给当前节点
			// 删除 min_node
			if minNode.rChild != nil {
				b.delete22(minNode)
			}else {
				b.delete1(minNode)
			}
		}
	}
}
// 删除： 情况1：node 是叶子节点
func (b *BST)delete1(node *BiTreeNode) {
	// 先判断 叶子节点是不是 根节点
	if node.parent == nil {
		b.root = nil
	}
	// 左孩子节点
	if node == node.parent.lChild {
		node.parent.lChild = nil
		node.parent = nil
	}else if node == node.parent.rChild {
		node.parent.rChild = nil
		node.parent = nil
	}
}
// 删除： 情况2.1 ： node 只有一个左孩子节点
func (b *BST)delete21(node *BiTreeNode) {
	// 先判断 node 的父节点是不是nil （则代表 node 是不是根节点）
	if node.parent == nil {
		b.root = node.lChild  // 将 node 的左孩子节点置为 根节点
		node.lChild.parent = nil
	}else if node == node.parent.lChild {
		node.parent.lChild = node.lChild
		node.lChild.parent= node.parent
	}else {
		node.parent.rChild = node.lChild
		node.lChild.parent = node.parent
	}
}
// 删除： 情况2.2 ： node 只有一个右孩子节点
func (b *BST)delete22(node *BiTreeNode) {
	// 先判断 node 的父节点是不是nil （则代表 node 是不是根节点）
	if node.parent == nil {
		b.root = node.rChild  // 将 node 的左孩子节点置为 根节点
		node.rChild.parent = nil
	}else if node == node.parent.lChild {
		node.parent.lChild = node.rChild
		node.rChild.parent= node.parent
	}else {
		node.parent.rChild = node.rChild
		node.rChild.parent = node.parent
	}
}

type BiTreeNode struct {
	data int           // 当前节点数据
	lChild *BiTreeNode    // 左孩子节点
	rChild *BiTreeNode    // 右孩子节点
	parent *BiTreeNode    // 当前节点的 父节点
}

func NewBiTreeNode(val int) *BiTreeNode{

	return &BiTreeNode{
		data:   val,
		lChild: nil,
		rChild: nil,
		parent: nil,
	}
}

func (b *BiTreeNode) String() string{
	if b != nil {
		return strconv.Itoa(b.data)
	}
	return ""
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

// 二叉搜索树
func main() {

	bst := NewBST([]int{1,4,3,6,7,2,0,9,8,5})

	preRange(bst.root)
	fmt.Println()
	inOrderRange(bst.root)
	fmt.Println()
	postOrderRange(bst.root)
	// fmt.Println(bst.root)

	fmt.Println()

	/*
	node := bst.query(bst.root, 6)
	if node != nil {
		fmt.Println("查到值：")
		fmt.Println("当前节点data : ", node)
		fmt.Println("左孩子节点data : ", node.lChild)
		fmt.Println("右孩子节点data : ", node.rChild)
		fmt.Println("父节点data : ", node.parent)
	}else {
		fmt.Println("没找到值")
	}

	 */


	Delete(bst, 3)
	inOrderRange(bst.root)
	fmt.Println()
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++")

	/*
	node2 := bst.queryNoRec( 1)
	if node2 != nil {
		fmt.Println("查到值：")
		fmt.Println("当前节点data : ", node2)
		fmt.Println("左孩子节点data : ", node2.lChild)
		fmt.Println("右孩子节点data : ", node2.rChild)
		fmt.Println("父节点data : ", node2.parent)
	}else {
		fmt.Println("没找到值")
	}
	 */

	Delete(bst, 0)
	//inOrderRange(bst.root)

	Delete(bst, 1)
	preRange(bst.root)
	fmt.Println()
	fmt.Println("----------------------------------")
	inOrderRange(bst.root)

	Delete(bst, 7)
	fmt.Println()
	fmt.Println("----------------------------------")
	preRange(bst.root)
	fmt.Println()
	fmt.Println("----------------------------------")
	inOrderRange(bst.root)

	fmt.Println()
	Delete(bst, 3)   // 3 不存在

}
