package main

import (
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"movie/src/share/config"
)

var schema = `
	CREATE TABLE IF NOT EXISTS user (
		id INT UNSIGNED AUTO_INCREMENT,
		name VARCHAR(20),
		address VARCHAR(20),
		phone VARCHAR(15),
		PRIMARY KEY (id)
	)`

// 对应表的结构体
type User struct {
	Id      int32  `db:"id"`
	Name    string `db:"name"`
	Address string `db:"address"`
	Phone   string `db:"phone"`
}

func main() {
	// 打开并连接数据库
	db, err := sqlx.Connect("mysql", config.MysqlDNS)
	if err != nil {
		logrus.Error("sqlx.Connect err : ", err)
		return
	}
	// 执行建表语句
	db.MustExec(schema)

	// 添加假数据
	// 创建事务
	tx := db.MustBegin()
	// 设置随机种子
	rand.Seed(time.Now().UnixNano())
	// 事务执行 SQL插入
	tx.MustExec("INSERT INTO user(id,name,address,phone) VALUES (?,?,?,?)",
		nil, GetRandomString(10), "beijing"+GetRandomString(10), "1591"+GetRandomString(7))

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return
	}

	fmt.Println("插入完毕")
}

// 按指定个数生成随机数
func GetRandomString(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
