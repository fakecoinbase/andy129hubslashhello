package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DB 是一个全局对象
var DB *sql.DB

// sql 事务
// mysql 中只有使用了 Innodb 数据库引擎的数据库或表才支持事务。
func main() {

	dsn := "root:12345687@tcp(127.0.0.1:3306)/go_test"
	err := initDB(dsn)
	if err != nil {
		fmt.Println("初始化数据库失败, err : ", err)
		return
	}

	transDemo()
}

// 将 id 为1 的 age 加两岁， 将 id 为 3 的 age 减两岁， 要么都执行成功，要么都执行失败
func transDemo() {
	// 返回一个事务对象  tx, 后续的操作都必须通过 tx 来处理
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("begin trans failed, err : ", err)
		if tx != nil {
			tx.Rollback()
		}
		return
	}

	// 开始执行事务操作
	sql1 := "update user set age=age+? where id=?"
	_, err = tx.Exec(sql1, 2, 1)
	if err != nil {
		tx.Rollback() // 出错了就回滚
		fmt.Println("update1 failed , err : ", err)
		return
	}

	sql2 := "update user set age=age-? where id=?"
	_, err = tx.Exec(sql2, 2, 3)
	if err != nil {
		tx.Rollback()
		fmt.Println("update2 failed , err : ", err)
		return
	}

	// 两条语句执行完毕
	err = tx.Commit()
	if err != nil {
		tx.Rollback() // 出错了就回滚， 如果第二条语句执行错误，调用 Rollback() 则也会将 第一条语句回滚
		return
	}
	// 全部执行完毕，没有出错
	fmt.Println("两条数据更新成功！")

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
