package main

import (
	"github.com/sirupsen/logrus"
	"net"
	"strings"
)

// ConfigChan 通道流通的是 etcd 服务中新的配置信息
var ConfigChan chan []CollectEntry

// CollectEntry 是 从etcd 服务中 获取日志目录以及topic 的配置项的结构体
type CollectEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

// 通过取巧的方式获取 本地连接外网的IP
// Get preferred outbound ip of this machine
func GetOutboundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// fmt.Println(localAddr.String()) // 在本机中打印： 192.168.1.2:63095

	ip = strings.Split(localAddr.IP.String(), ":")[0]
	logrus.Infof("GetOutboundIP(), ip : %s", ip)
	return
}
