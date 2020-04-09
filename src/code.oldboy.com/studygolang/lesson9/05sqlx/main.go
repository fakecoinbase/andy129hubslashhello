package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DB 是一个全局对象 (内置连接池)
var db *sqlx.DB     // 注意与 sql.DB 的不同

// User 是一个用户结构体  (通过定义 tag , 方便sqlx 知道从数据库中的查询出来的值 与 User 结构体中哪个字段对应 )
type User struct {
	ID int `db:"id"`
	Name string  `db:"name"`
	Age int `db:"age"`
}

// sqlx 第三方数据库包
// 获取： go get github.com/jmoiron/sqlx
// sqlx 是兼容 go语言内置 sql 包的
func main() {
	err := initDB()
	if err != nil {
		fmt.Println("初始化数据库失败, err : ", err)
		return 
	}

	// queryRowDemo()
	// queryRowsDemo()

	// transDemo()

	transPrepareDemo()
}

// 事务与 预处理一起操作
func transPrepareDemo(){
	tx, err := db.Beginx()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Println("beginx failed , err : ", err)
		return 
	}

	sqlStr1 := "update user set age=age+? where id=?"
	stmt, err := tx.Preparex(sqlStr1)
	if err != nil {
		tx.Rollback()
		fmt.Println("preparex failed, err : ", err)
		return 
	}

	defer stmt.Close()

	for i:=0;i<10;i++ {
		stmt.MustExec(2, i)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		fmt.Println("更新失败, err : ", err)
		return 
	}
	fmt.Println("所有数据更新完成！")
	
}

// sqlx 事务操作
func transDemo(){
	tx, err := db.Beginx()
	if err != nil {
		if tx != nil {
			tx.Rollback()
		}
		fmt.Println("beginx failed , err : ", err)
		return 
	}
	sqlStr1 := "update user set age=age-? where id=?"
	tx.MustExec(sqlStr1, 2,1)
	// err = tx.Commit()     // 定义位置错误，不能写在两条 sql 语句之间

	sqlStr2 := "update hhh user set age=age+? where id=?"
	tx.MustExec(sqlStr2, 2,3)

	err = tx.Commit()  // 关于事务的提交，必须定义在所有 sql 语句执行代码块之后，保持一致性。
	if err != nil {
		tx.Rollback()
		fmt.Println("commit failed , err : ",err)
		return 
	}
	fmt.Println("两条数据更新成功！")
}

// sqlx 查询多条数据
func queryRowsDemo(){
	sqlStr := "select * from user"

	var users []User  // 由于 select 中要传入结构体指针，所以结构体中字段的名字首字母必须要大写，才能被 sqlx 底层操作进行反射赋值
	// 第一个参数必须传入指针类型的结构体，用于接收 数据库返回的值
	err := db.Select(&users, sqlStr)   // 查询多条数据
	if err != nil {
		fmt.Println("query failed, err : ", err)
		return 
	}

	for _, user := range users {
		fmt.Println("user : ", user)
	}
	

}
		   
// sqlx 查询单条数据
func queryRowDemo(){
	sqlStr := "select id,name,age from user where id=?"

	var user User
	// 第一个参数必须传入指针类型的结构体，用于接收 数据库返回的值
	err := db.Get(&user, sqlStr, 1)
	if err != nil {
		fmt.Println("query failed, err : ", err)
		return 
	}
	fmt.Println("user : ", user)

}

func initDB() (err error) {

	dsn := "root:12345687@tcp(127.0.0.1:3306)/go_test"
	// 由于定义了全局 db, 所以这里直接 = 赋值即可，不能 := , 否则db 会被当做局部变量，而全局 db 没有被初始化。
	db, err = sqlx.Connect("mysql", dsn)   // 包含了sqlx 的connnect() 包含了 open() 和 ping() 两个操作
	
	if err != nil {
		fmt.Println("connect failed , err : ", err)
		return err
	}
	// fmt.Println(db == nil )

	// 数据库连接成功

	// 设置最大连接数
	db.SetMaxOpenConns(50)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(20)
	return nil
}