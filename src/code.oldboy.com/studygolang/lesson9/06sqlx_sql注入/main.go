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

// sqlx -- sql 注入
/*  什么是 sql 注入问题
	1, 过分相信用户
	2, 拿着用户输入的内容直接拼接字符串去执行
*/
func main() {
	// 正常用户输入查询
	// sqlInjectDemo("yang")

	// 不正常输入
	// select id,name,age from user where name='xxx' or 1=1#'   // 1=1 永远成立，  # 代表把后面的语句注释掉
	// sqlInjectDemo("xxx' or 1=1#")

	// 不正常输入
	// sqlInjectDemo("xxx' union select * from user #")   // union 代表合并

	// 不正常输入
	sqlInjectDemo("yang' and (select count(*) from user) < 10 #")  // 后面满足条件， 并且前面用户名 也存在的情况下，返回  yang 的信息
}

func init(){
	err := initDB()
	if err != nil {
		fmt.Println("初始化数据库失败, err : ", err)
		return 
	}
}

func sqlInjectDemo(name string){
	// 字符串拼接 sql 语句 (会带来 sql 注入的问题)
	sqlStr := fmt.Sprintf("select id,name,age from user where name='%s'", name)
	var users []User
	err := db.Select(&users, sqlStr)
	if err != nil {
		fmt.Println("query failed, err : ", err)
		return 
	}

	for _,user := range users {
		fmt.Println(user)
	}
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
