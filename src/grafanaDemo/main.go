package main

import (
	"fmt"
	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/shirou/gopsutil/cpu"
	"log"
	"time"
)

var (
	cli client.Client
)

// grafana 是一个 从数据库中查询数据，并将数据以 柱状图，饼状图等图形展示的平台
// grafana 插件下载 https://grafana.com/grafana/plugins?utm_source=grafana_getting_started
// 例如： 下载饼状图 grafana-cli plugins install grafana-piechart-panel
/* 并不涉及 grafana 的go 语言的操作，我们需要做的是 获取我们想展示的数据(监控的数据) 并写入 influxDB, grafana 一直从 influxDB 获取数据 并展示
	0, 下载 grafana 安装包 (windows版 或 mac版或 linux 版), 以 windows版举例， 执行 grafana-server.exe 启动 grafana 服务

		注意(拷贝 sample.ini一份并命名为 custom.ini， 实际上要将 defaults.ini 拷贝一份 命名为 custom.ini)： 以下为官方说明：
				Note: The default Grafana port is 3000. This port might require extra permissions on Windows.
					If it does not appear in the default port, you can try changing to a different port.

			Go into the conf directory and copy sample.ini to custom.ini. Note: You should edit custom.ini, never defaults.ini.
			Edit custom.ini and uncomment the http_port configuration option
			(; is the comment character in ini files) and change it to something like 8080 or similar.
				That port should not require extra Windows privileges. Read more about the configuration options.

	1, 获取需要展示的数据 ---> 使用 gopsutil 获取系统的 CPU,内存，磁盘，IO 等信息
	2, 将监控的信息写入 influxDB
 */
func main(){

	// 连接 influxDB
	err := connectInfluxDB()
	if err != nil {
		fmt.Println("connectInfluxDB err : ", err)
		return
	}

	// 创建一个每一秒中触发的定时器
	ticker := time.Tick(time.Second)
	for _ = range ticker {
		// 获取CPU 使用率信息
		percent, err := getCpuInfo()
		if err == nil {
			// 将信息写入 influxDB
			err := writesPoints(percent)
			if err != nil {
				fmt.Println("writesPoints err : ", err)
			}
		}else {
			fmt.Println("getCupInfo err : ", err)
		}
	}

}

func getCpuInfo() (percent float64, err error){
	infos,err := cpu.Info()
	if err != nil {
		fmt.Println("cpu.Info err : ", err)
		return
	}

	for _,infoState := range infos {
		fmt.Println("infoState : ", infoState)
	}

	percentArr, _ := cpu.Percent(time.Second, false)   // 每一秒更新CPU 使用率
	percent = percentArr[0]
	fmt.Printf("cpu percent:%v\n", percent)
	return
}

// 连接influxDB 服务
func connectInfluxDB() (err error){
	cli, err = client.NewHTTPClient(client.HTTPConfig{
		Addr: "http://127.0.0.1:8086",
		Username : "admin",
		Password : "",
	})
	if err != nil {
		return
	}
	return
}

// insert 插入
func writesPoints(percent float64) (err error){
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  "grafana_test",
		Precision: "s", //精度，默认ns
	})
	if err != nil {
		return
	}
	tags := map[string]string{"cpu": "ih-cpu"}
	fields := map[string]interface{}{
		"percent":   percent,
	}

	pt, err := client.NewPoint("cpu_percent", tags, fields, time.Now())
	if err != nil {
		return
	}
	bp.AddPoint(pt)
	err = cli.Write(bp)
	if err != nil {
		return
	}
	log.Println("insert success")
	return
}
