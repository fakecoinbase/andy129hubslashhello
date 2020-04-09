package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// DB 是一个全局对象
var DB *sql.DB

// 数据库连接池
func main() {

	dsn:= "root:12345687@tcp(127.0.0.1:3306)/go_test"
	err := initDB(dsn)
	if err != nil {
		fmt.Println("初始化数据库失败, err : ",err )
		return 
	}

	// 查询单行
	// queryOne()

	// 查询多行
	// queryMuli()

	// 验证 连接被占用的情况
	/*
	for i:=0;i<100;i++ {
		fmt.Println("--------------i : ", i)
		queryOne()
	}
	*/

	// 插入一条数据
	// insertToDb()

	// 更新数据
	// updateToDb()

	// 删除数据
	deleteFromDb()
	
	fmt.Println("数据库查询结束！")
}

func deleteFromDb(){
	sqlStr := "delete from user where id=?"
	id := 2
	ret, err := DB.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("delete failed, err : ", err)
		return 
	}
	num, err := ret.RowsAffected()
	if err != nil {
		fmt.Println("get Affected Row failed, err : ", err)
		return 
	}
	fmt.Println("num : ", num)
}

func updateToDb(){
	sqlStr := "update user set age=? where id=?"
	id := 2
	age := 18
	ret, err := DB.Exec(sqlStr, age, id)
	if err != nil {
		fmt.Printf("update failed, err : %v\n", err)
		return 
	}
	// 拿到受影响的行数
	lastID, err := ret.LastInsertId()
	num, err := ret.RowsAffected()
	fmt.Printf("lastID : %d, num : %d\n", lastID, num)
}	

func insertToDb(){
	sqlStr := "insert into user(name,age) value(?,?)"
	name := "zhao"
	age := 55
	ret, err := DB.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("insert failed, err : %v\n", err)
		return 
	}
	// 拿到刚插入的数据ID 值 (不同的数据库有不同的实现)
	lastID, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get lastId failed, err : %v\n", err)
	}
	fmt.Println("lastID : ", lastID)
}

// 查询多条数据
func queryMuli(){

	var id int 
	var name string 
	var age int 

	queryStr := "select id,name,age from user where id>?"
	rows, err := DB.Query(queryStr, 0)  // 查询ID 大于 0 的数据
	if err != nil {
		fmt.Println("查询失败, err : ", err)
		return 
	}
	
	// 避免 rows.Next() 取值时出错，导致连接占用的问题。
	defer func(){
		rows.Close()    
	}()

	// 循环读取数据
	// 必须调用 rows.Next() 从里面取值
	for rows.Next() {
		err := rows.Scan(&id, &name, &age)
		if err != nil {
			fmt.Println("读取数据失败, err ", err)
			return 
		}
		fmt.Printf("id : %d, name : %s, age : %d\n", id, name, age)
	}

}

// 查询单条数据
func queryOne(){
	var id int 
	var name string 
	var age int 
	// 查询语句
	queryStr := "select id,name,age from user where id=1"
	// 查询一行结果
	row := DB.QueryRow(queryStr)

	// 注意  Row 是没有 Close() 的方法， 只能使用 Scan() 取值，取完值底层默认就会关闭释放这个连接
	// defer row.Close() 

	// 必须调用 Scan 将值取出来并且 如果不调用 scan 则会一直占用这个 连接。
	row.Scan(&id, &name, &age)
	fmt.Printf("id : %d, name : %s, age : %d\n", id,name, age)
}

func initDB(dsn string)(err error){
	
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