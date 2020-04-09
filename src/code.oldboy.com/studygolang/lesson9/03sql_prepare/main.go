package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DB 是一个全局对象
var DB *sql.DB

// sql 预处理
/*  为什么要预处理?
1, 优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，节省后续编译的成本。
2, 避免SQL注入问题。
*/
func main() {

	dsn := "root:12345687@tcp(127.0.0.1:3306)/go_test"
	err := initDB(dsn)
	if err != nil {
		fmt.Println("初始化数据库失败, err : ", err)
		return
	}

	// 预处理 插入操作
	// prepareInsertDemo()
	prepareQueryDemo()

}

// 预处理查询操作
func prepareQueryDemo() {
	sqlStr := "select * from user where id=?"
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed , err : ", err)
		return
	}

	defer stmt.Close()

	for i := 0; i < 10; i++ {
		rows, err := stmt.Query(i) // 传入预处理 sql 语句的值

		defer rows.Close()

		if err != nil {
			fmt.Println("query failed , err : ", err)
			return
		}

		var id int
		var name string
		var age int
		for rows.Next() {
			err := rows.Scan(&id, &name, &age)
			if err != nil {
				fmt.Println("scan failed , err : ", err)
				return
			}
			fmt.Printf("查询数据：id = %d, name = %s, age = %d\n", id, name, age)
		}
	}

}

// 预处理 插入操作
func prepareInsertDemo() {
	sqlStr := "insert into user (name,age) values (?,?)"
	stmt, err := DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare failed , err : ", err)
		return
	}

	defer stmt.Close()

	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("stu%2d", i)
		stmt.Exec(name, i+20) // stmt.Exec()  这里面传入参数的个数 与 上面 sqlStr 语句中的占位符数量一致
	}

}

func initDB(dsn string) (err error) {

	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	// 数据库连接成功

	// 设置最大连接数
	DB.SetMaxOpenConns(50)
	// 设置最大空闲连接数
	DB.SetMaxIdleConns(20)
	return nil
}
