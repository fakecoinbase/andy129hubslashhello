package main

import (
	"fmt"
	"os"
)

// 练习需求：
/*
	使用函数实现一个简单的图书管理系统
	每本书有书名、作者、价格、上架信息
	用户可以在控制台添加书籍、修改书籍信息、打印所有的书籍列表。
*/

// 练习总结：
/*
	1, (注意：)要思考项目中的书籍内容，用什么数据类型存储， 多本书籍要使用什么 复合数据类型 来存储
	2，根据需求，一条一条的拆分， 然后针对每个需求 定义一个 函数 实现功能
	3，（严重注意：）涉及到添加，修改，删除，这类的需求，要考虑我是用 值类型还是 指针类型。
	4，检查逻辑代码是否清晰，严谨，结构是否还能再优化。(多次重复调用的代码，可以写一个函数，到处调用。)

*/

// Book 是一个 书籍的结构体
type Book struct {
	title   string
	author  string
	price   float32
	publish bool
}

// 必须将Book 以指针类型传入到 allBooks 这个 切片里，不然后续更新无法同步修改。
var allBooks = make([]*Book, 0, 10) // 注意指针类型的写法，例如： func test(p *People){},   []*Book, :  ( *紧跟类型名)

func main() {

	showMenu()

}

// 添加一本书，创建一个 Book 实例。
func newBook(title, author string, price float32, publish bool) *Book {
	return &Book{
		title:   title,
		author:  author,
		price:   price,
		publish: publish,
	}
}

func showMenu() {

	// 采用 for 循环，当每个操作执行完毕之后回 重新回到这个主菜单，一直运行。 除非执行 os.Exit(0),  程序才会完全退出。
	for {

		// 将 index 声明到 for 循环内，就能保证每次回到菜单界面，index 默认都为0，而不会保存上一次操作的指令值。
		// 指令
		var index int

		fmt.Println("-----------书籍管理系统---------")
		fmt.Println()
		fmt.Println("1,添加书籍")
		fmt.Println("2,修改书籍信息")
		fmt.Println("3,展示所有书籍")
		fmt.Println("4,退出")
		fmt.Scanln(&index) // Scanln 的好处在于，当用户按下回车代表 该行输入结束，如果用 scan 的话，用户按下回车，程序不结束，终端会显示空行，界面不美观
		fmt.Println("打印信息：", index)

		switch index {
		case 1:
			addBook()
		case 2:
			modifyBook()
		case 3:
			showAllBooks()
		case 4:
			fmt.Println("正常退出")
			os.Exit(0)
		default:
			fmt.Println("指令不匹配，退出")
			os.Exit(0)
		}
	}

}

// 书籍管理系统
func addBook() {

	var (
		title   string
		author  string
		price   float32
		publish bool
	)
	fmt.Println()
	fmt.Println("###请依次根据提示输入相关信息.")
	fmt.Print("title: ")
	fmt.Scanln(&title) // 注意  Scan() 里面传入的都是 指针类型
	fmt.Print("author: ")
	fmt.Scanln(&author)
	fmt.Print("price: ")
	fmt.Scanln(&price)
	fmt.Print("publish: ")
	fmt.Scanln(&publish)

	book := newBook(title, author, price, publish)

	fmt.Println(*book)

	allBooks = append(allBooks, book)
	fmt.Println()
}

func modifyBook() {

	var (
		title   string
		author  string
		price   float32
		publish bool
	)
	fmt.Println()
	fmt.Println("###请依次根据提示输入相关信息.")
	fmt.Print("title: ")
	fmt.Scanln(&title) // 注意  Scan() 里面传入的都是 指针类型
	fmt.Print("author: ")
	fmt.Scanln(&author)
	fmt.Print("price: ")
	fmt.Scanln(&price)
	fmt.Print("publish: ")
	fmt.Scanln(&publish)
	fmt.Println()

	exsit := false
	for _, v := range allBooks {
		exsit = false
		// 必须要把 book 定义为指针类型，不然的话，无法将终端输入的书籍信息 更新过去，
		// 也就是说下面的四行赋值语句，v 是迭代出来的，可以理解为值类型, 所以只是一份拷贝，
		// 然后我们将 终端的信息赋值给一份拷贝过来的值类型，那么实际上并没有修改到 allBooks里面的值。
		// var allBooks = make([]*Book, 0, 10)

		if title == v.title { // 判断要求修改的书籍名称 是否存在
			v.author = author
			v.price = price
			v.publish = publish
			exsit = true
			break
		}
	}
	/*
		for _, v := range allBooks {
			fmt.Println("修改书籍信息如下：")
			fmt.Println(*v)
		}
	*/

	if exsit {
		fmt.Println("--修改成功！")
	} else {
		fmt.Println("--该书籍不存在, 请先添加！")
	}
	fmt.Println()
}
func showAllBooks() {

	if len(allBooks) > 0 {
		fmt.Println("--书籍信息如下: ")

		for _, v := range allBooks {
			fmt.Println()
			fmt.Printf("--title:%s,author:%s,price:%.2f,publish:%t\n", v.title, v.author, v.price, v.publish)
			fmt.Println()
		}
	} else {
		fmt.Println("--没有任何书籍，请先添加！")
		fmt.Println()
	}

}

/*
// 在终端输入的时候，直接给 Book 结构体里的字段赋值：
func addBook() {

	var book = Book{}
	fmt.Println()
	fmt.Println("###请依次根据提示输入相关信息.")
	fmt.Print("title: ")
	fmt.Scanln(&(book.title)) // 注意  Scan() 里面传入的都是 指针类型
	fmt.Print("author: ")
	fmt.Scanln(&(book.author))
	fmt.Print("price: ")
	fmt.Scanln(&(book.price))
	fmt.Print("publish: ")
	fmt.Scanln(&(book.publish))

	fmt.Println(book)

	allBooks = append(allBooks, &book)
	fmt.Println()
}
*/
