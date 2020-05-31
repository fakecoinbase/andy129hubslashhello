package main

import "fmt"

/*	学习总结：
	1, 使用 树结构来模拟 文件系统目录结构
	2, 树的每一个节点 (父节点，子节点) 相互关联，这种关系 基于 链式结构 (双向)
	3, go 语法问题
		a,  go中结构中的无效递归类型(invalid recursive type in a struct in go)
			--- 什么是无效递归类型呢？
            --- 举例如下: (定义一个结构体 Node)
				type Node struct {
					name string
					fileType string   // "dir", "file"
					children []*Node
					parent *Node
				}
               children 和 parent 也需要使用 node ,但是 如果 定义为 []Node,  Node, 则就会报 上面那个问题： 无效递归类型
				所以我们使用 指针类型 来存储 node ,这样它就会认为是 指针类型，而不是  Node 类型

		b,  当我想打印 Node 对象的时候， fmt.Println(v) 打印出来的是一堆地址，如果我想让 Println() 能打印我指定的内容怎么办呢

			定义一个 String() 方法，如下： (貌似必须是返回 string 类型)

                // String() 返回 需要被打印的内容
				func (n Node) String() string{
					return n.name
				}

 */

// 文件节点 (实际上是个 链式存储)
type Node struct {
	name string
	fileType string   // "dir", "file"
	children []*Node
	parent *Node
}
// 定义一个 String() 方法，当打印 node 对象时，会自动去 调用 String() 方法，打印 自定义的内容
func (n Node) String() string{
	return n.name
}

// 文件系统树
type FileSystemTree struct {
	root *Node  // 根目录
	now *Node   // 当前目录
}

func (f *FileSystemTree) init(n *Node){
	f.root = n
	f.now = f.root
}

func (f *FileSystemTree) mkdir(dir string){  // dir 以
	if dir[len(dir)-1] != '/' {
		dir = dir+"/"
	}
	node := Node{
		name : dir,
	}
	f.now.children = append(f.now.children, &node)   // 将新创建的目录 添加到当前目录的 children 中，作为子节点
	node.parent = f.now   // 将当前目录 设置为 新创建目录的父节点

}

func (f *FileSystemTree) ls() []*Node{
	return f.now.children
}

func (f *FileSystemTree) cd(dir string){
	// 进入根目录
	if dir == "/" {
		f.now = f.root
		return
	}
	// 进入上一层目录
	if dir == "../" {  // 返回上一层目录
		f.now = f.now.parent    // 当前目录指定为 当前目录的父目录 (实现 返回上一层目录的 功能)
		return
	}
	// 如果dir 缺少 "/", 则追加 "/"
	if dir[len(dir)-1] != '/' {
		dir = dir+"/"
	}

	/*	遗留功能：
	    /abc/abc的子目录/        目前 cd() 功能只能进入一层, 进入多层目录怎么办?   (分割 dir 字符串，得到目录名字，然后进入)

		../abc/abc的子目录/     先返回上一层，再进入其他目录  (分割 dir 字符串，得到目录名字，然后进入)

	 */
	for _,child := range f.now.children {
		if dir == child.name {
			f.now = child
			return
		}
	}
	fmt.Println("没有找到这个目录")
}


// 树 的应用：  文件系统目录结构
func main() {
	f := new(FileSystemTree)
	n := Node{
		name:     "/",
		fileType: "",
		children: nil,
		parent:   nil,
	}
	f.init(&n)

	f.mkdir("新建文件夹")   // 创建文件夹
	nodes := f.ls()   // [新建文件夹/]
	fmt.Println(nodes)   // nodes 是 []*Node 类型，当打印时，会去调用 Node 对应的 String() 方法，我们已经定义了这个方法，并且返回我们想要打印的值
	/*
		没有定义 String() 方法前的打印：  [0xc000016100]
		定义了 String() 方法之后的打印：  [新建文件夹/]
	 */

	f.mkdir("abc")
	f.mkdir("go语言学习资料")
	fmt.Println(f.ls())   // [新建文件夹/ abc/ go语言学习资料/]

	f.cd("abc")
	fmt.Println(f.ls())   // []
	f.mkdir("abc的子目录")
	fmt.Println(f.ls())   // [abc的子目录/]

	f.cd("../")   // 返回上一层目录
	fmt.Println(f.ls())  // [新建文件夹/ abc/ go语言学习资料/]

	f.cd("/")
	fmt.Println(f.ls())
}
