package main

import (
	"fmt"
	client "github.com/influxdata/influxdb1-client/v2"
	"log"
	"time"
)

// influxDB 时序数据库
// 参考文章： https://www.liwenzhou.com/posts/Go/go_influxdb/

/* 下载： 参考文章：https://blog.csdn.net/v6543210/article/details/86934941
官网：https://portal.influxdata.com/downloads/

下载链接：https://dl.influxdata.com/influxdb/releases/influxdb-1.7.3_windows_amd64.zip

Mac和Linux用户可以点击https://v2.docs.influxdata.com/v2.0/get-started/下载。
 */

// Go操作influxDB  (官方推荐使用 v1.8 版本)
/*
influxDB 1.x版本  (官方文档：https://docs.influxdata.com/influxdb/v1.7/introduction/getting-started/   )
go get github.com/influxdata/influxdb1-client/v2
influxDB 2.x版本 （2020-04-22目前还是测试版本，还未正式发布）
go get github.com/influxdata/influxdb-client-go
 */


// windows 下启动 influxDB 服务
/*
第一步： 修改 influxdb.conf
	dir = "G:/Share/influxdb-1.7.3_windows_amd64/meta"
	dir = "G:/Share/influxdb-1.7.3_windows_amd64/data"
    wal-dir = "G:/Share/influxdb-1.7.3_windows_amd64/wal"

第二步：启动服务
G:\Share\influxdb-1.7.3_windows_amd64\influxdb-1.7.3-1\

	influxd.exe    // 启动服务端
	influx.exe     // 启动客户端

 */
// 命令行操作 influxDB
/*
	CREATE DATABASE mydb    // 创建数据库
    SHOW DATABASES          // 显示所有数据库
	USE mydb                // 使用某个数据库
	SELECT "host", "region", "value" FROM "cpu_usage"   // 查询某张表里的某些字段
	show measurements       // 显示某张表
 */
func main() {
	conn := connectInfluxDB()
	fmt.Println(conn)

	// insert
	writesPoints(conn)

	// 获取10条数据并展示
	qs := fmt.Sprintf("SELECT * FROM %s LIMIT %d", "cpu_usage", 10)
	res, err := queryDB(conn, qs)
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range res[0].Series[0].Values {
		for j, value := range row {
			log.Printf("j:%d value:%v\n", j, value)
		}
	}
}

// 连接influxDB 服务
func connectInfluxDB() client.Client{
	cli, err := client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://127.0.0.1:8086",
		Username : "admin",
		Password : "",
	})
	if err != nil {
		fmt.Println(err)
	}
	return cli
}

// query 查询
func queryDB(cli client.Client, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: "test",
	}
	if response, err := cli.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}

// insert 插入
func writesPoints(cli client.Client) {
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "test",
		Precision: "s", //精度，默认ns
	})
	if err != nil {
		log.Fatal(err)
	}
	// 设置关键字， 可以通过 检索 "cpu" 这个 key 对应的 value 所属 的  "idle", "system", "user" 信息
	// 如果是 获取磁盘信息，或者 网络io 信息，则这里要 设置多个 key  (磁盘对应 多个磁盘， 网络io 可能对应多个网卡)
	tags := map[string]string{"cpu": "ih-cpu"}
	fields := map[string]interface{}{
		"idle":   201.1,
		"system": 43.3,
		"user":   86.6,
	}

	pt, err := client.NewPoint("cpu_usage", tags, fields, time.Now())
	if err != nil {
		log.Fatal(err)
	}
	bp.AddPoint(pt)
	err = cli.Write(bp)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("insert success")
}


