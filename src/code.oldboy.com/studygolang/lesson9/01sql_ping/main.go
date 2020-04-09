package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// database/sql 使用
func main() {

	// dsn:"user:password@tcp(ip:port)/databasename"
	dsn := "root:12345687@tcp(127.0.0.1:3306)/go_test"
	// 调用标准库中的方法
	// 为什么是 "mysql" ， 不仅仅是因为这次用的是 mysql 数据库，而是 我们导入 "github.com/go-sql-driver/mysql" 第三方包时，它的内部实现注册的名称就是 mysql
	/*  G:\Goworkspace\src\github.com\go-sql-driver\mysql\driver.go 内部实现：
	func init() {
		sql.Register("mysql", &MySQLDriver{})
	}
	*/
	// 这里的 err 检查的是 "mysql" 这个驱动是否已经注册，它不会检查 dsn 中用户名和密码的正确性
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库驱动错误, err : ", err)
		return
	}

	// 尝试连接数据库，校验用户名密码是否正确
	err = db.Ping()
	if err != nil {
		fmt.Println("数据库连接失败，err : ", err)
		return
	}
	fmt.Println("数据库连接成功！")

}
