package main

type AVLNode struct {
	data int           // 当前节点数据
	lChild *AVLNode    // 左孩子节点
	rChild *AVLNode    // 右孩子节点
	parent *AVLNode    // 当前节点的 父节点
}

type AVLTree struct {

}

// AVL 树
/*
	AVL 树： AVL树是一棵自平衡的二叉搜索树
	AVL 树具有以下性质：
		-- 根的左右子树的高度之差的绝对值不能超过 1
		-- 根的左右子树都是平衡二叉树
 */
func main() {

}
