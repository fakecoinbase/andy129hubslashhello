package main

import (
	"fmt"
	"log"
	"net"
)

// ip 获取

func main() {
	// GetLocalIP()
	 getLocalIP2()
	// getLocalIP3()

	// GetOutboundIP()
}

// 获取本机IP的两种方式 GetLocalIP() 与  GetOutboundIP()

// GetLocalIP() 通过 net.InterfaceAddrs() 接口获取
// GetLocalIP() 与 getLocalIP2(), getLocalIP3()   逐步优化的方案
func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}

		fmt.Println(ipAddr)
		return ipAddr.IP.String(), nil
	}
	return
}

func getLocalIP2(){
	addrs, err := net.InterfaceAddrs()
	if err != nil{
		fmt.Println(err)
		return
	}
	for _, value := range addrs{
		if ipnet, ok := value.(*net.IPNet); ok && !ipnet.IP.IsLoopback(){
			if ipnet.IP.To4() != nil{
				fmt.Println(ipnet.IP.String())
				/*
				169.254.234.80
				192.168.1.2
				192.168.32.1
				192.168.76.1
				169.254.97.162

				*/
			}
		}
	}
}

//  getLocalIP2() 这种实现方式忽略了一个网卡可用性的问题，导致获取出来的IP可能不一定是想要的。
//  需要通过判断net.FlagUp标志进行确认，排除掉无用的网卡。优化后的实现方式：
func getLocalIP3(){
	netInterfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println("net.Interfaces failed, err:", err.Error())
		return
	}

	for i := 0; i < len(netInterfaces); i++ {
		if (netInterfaces[i].Flags & net.FlagUp) != 0 {
			addrs, _ := netInterfaces[i].Addrs()

			for _, address := range addrs {
				if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					if ipnet.IP.To4() != nil {
						fmt.Println(ipnet.IP.String())

						/*
						192.168.1.2
						192.168.32.1
						192.168.76.1
						*/
					}
				}
			}
		}
	}
}

// 通过取巧的方式获取 本地连接外网的IP
// Get preferred outbound ip of this machine
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())   // 在本机中打印： 192.168.1.2:63095
	return localAddr.IP.String()
}
