package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

// gopsutil  Go语言部署简单、性能好的特点非常适合做一些诸如采集系统信息和监控的服务
/*
	psutil是一个跨平台进程和系统监控的Python库，而gopsutil是其Go语言版本的实现。本文介绍了它的基本使用。
	安装：
	go get github.com/shirou/gopsutil
 */
func main() {

	// getCpuInfo()
	// getCpuLoad()
	// getMemInfo()
	// getHostInfo()
	getDiskInfo()
	//getNetInfo()
}

// 获取 cpu 信息
func getCpuInfo(){
	infos,err := cpu.Info()
	if err != nil {
		fmt.Println("cpu.Info err : ", err)
		return
	}

	for _,infoState := range infos {
		fmt.Println("infoState : ", infoState)
	}

	// CPU 使用率
	for {
		percent, _ := cpu.Percent(time.Second, false)   // 每一秒更新CPU 使用率
		fmt.Printf("cpu percent:%v\n", percent)
	}
}

// 获取CPU负载信息  (在 windows 下不支持)
func getCpuLoad(){
	avgStat, err := load.Avg()
	if err != nil {
		fmt.Println("load.Avg() err : ", err)
		return
	}
	fmt.Printf("%v\n", avgStat)  // load.Avg() err :  not implemented yet
}

// 获取内存信息
func getMemInfo(){
	memStat, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("mem.VirtualMemory() err : ", err)
		return
	}
	fmt.Printf("%v\n", memStat)
	/*
	{"total":8524779520,"available":2686705664,"used":5838073856,"usedPercent":68,"free":0,"active":0,"inactive":0,"wired":0,"laundry":0,"buffers":0,"cached":0,"writeback":0,"dirty":0,"writebacktmp":0,"shared":0,"slab":0,"srec
	laimable":0,"sunreclaim":0,"pagetables":0,"swapcached":0,"commitlimit":0,"committedas":0,"hightotal":0,"highfree":0,"lowtotal":0,"lowfree":0,"swaptotal":0,"swapfree":0,"mapped":0,"vmalloctotal":0,"vmallocused":0,"vmallocch
	unk":0,"hugepagestotal":0,"hugepagesfree":0,"hugepagesize":0}

	*/
}

// 获取主机信息
func getHostInfo(){
	hInfo, err := host.Info()
	if err != nil {
		fmt.Println("host.Info() err : ", err)
		return
	}
	fmt.Printf("host info:%v uptime:%v boottime:%v\n", hInfo, hInfo.Uptime, hInfo.BootTime)
	/*
	host info:{"hostname":"DESKTOP-9QU11GO","uptime":94229,"bootTime":1587391446,"procs":156,"os":"windows","platform":"Microsoft Windows 10 Pro","platformFamily":"Standalone Workstation","platformVersion":"10.0.14393 Build 14
	393","kernelVersion":"","kernelArch":"x86_64","virtualizationSystem":"","virtualizationRole":"","hostid":"d4b26c4a-a676-4973-aaa3-6ebd330d8e08"} uptime:94229 boottime:1587391446
	 */
}

// 获取磁盘信息
func getDiskInfo(){
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return
	}

	for _,part := range parts {
		fmt.Printf("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	}
	/*
	part:{"device":"C:","mountpoint":"C:","fstype":"NTFS","opts":"rw.compress"}
	disk info:used:91.38529171470235 free:4582834176
	part:{"device":"E:","mountpoint":"E:","fstype":"exFAT","opts":"rw"}
	disk info:used:50.08738785395656 free:64045580288
	part:{"device":"G:","mountpoint":"G:","fstype":"NTFS","opts":"rw.compress"}
	disk info:used:92.17543493509316 free:5190569984
	*/
	ioStat,_ := disk.IOCounters()
	for k,v := range ioStat {
		fmt.Printf("%v -- %v\n", k, v)
	}
	/*
	C: -- {"readCount":266148,"mergedReadCount":0,"writeCount":366058,"mergedWriteCount":0,"readBytes":6941647872,"writeBytes":12925803008,"readTime":505,"writeTime":622,"iopsInProgress":0,"ioTime":0,"weightedIO":0,"name":"C
	:","serialNumber":"","label":""}
	G: -- {"readCount":288125,"mergedReadCount":0,"writeCount":138883,"mergedWriteCount":0,"readBytes":10120664064,"writeBytes":8323694080,"readTime":282,"writeTime":235,"iopsInProgress":0,"ioTime":0,"weightedIO":0,"name":"G
	:","serialNumber":"","label":""}

	*/
}

func getNetInfo(){
	infos, err := net.IOCounters(true)
	if err != nil {
		fmt.Println("net.IOCounters err : ", err)
		return
	}
	for index, v := range infos {
		fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	}

	/*
	0:{"name":"SSTAP 1","bytesSent":0,"bytesRecv":0,"packetsSent":0,"packetsRecv":0,"errin":0,"errout":0,"dropin":0,"dropout":0,"fifoin":0,"fifoout":0} send:0 recv:0
	1:{"name":"以太网","bytesSent":324391800,"bytesRecv":3450076728,"packetsSent":1560431,"packetsRecv":2999858,"errin":0,"errout":0,"dropin":7267,"dropout":0,"fifoin":0,"fifoout":0} send:324391800 recv:3450076728
	2:{"name":"VMware Network Adapter VMnet1","bytesSent":3536,"bytesRecv":88,"packetsSent":85,"packetsRecv":88,"errin":0,"errout":0,"dropin":0,"dropout":0,"fifoin":0,"fifoout":0} send:3536 recv:88
	3:{"name":"VMware Network Adapter VMnet8","bytesSent":13518,"bytesRecv":88,"packetsSent":85,"packetsRecv":88,"errin":0,"errout":0,"dropin":0,"dropout":0,"fifoin":0,"fifoout":0} send:13518 recv:88
	4:{"name":"Loopback Pseudo-Interface 1","bytesSent":0,"bytesRecv":0,"packetsSent":0,"packetsRecv":0,"errin":0,"errout":0,"dropin":0,"dropout":0,"fifoin":0,"fifoout":0} send:0 recv:0
	5:{"name":"本地连接* 2","bytesSent":8993,"bytesRecv":2888,"packetsSent":93,"packetsRecv":19,"errin":0,"errout":0,"dropin":0,"dropout":0,"fifoin":0,"fifoout":0} send:8993 recv:2888
	6:{"name":"isatap.{764103F8-8087-4714-982A-652C1FCD9632}","bytesSent":0,"bytesRecv":0,"packetsSent":0,"packetsRecv":0,"errin":0,"errout":0,"dropin":0,"dropout":0,"fifoin":0,"fifoout":0} send:0 recv:0
	7:{"name":"isatap.{FA47EBDD-A3EC-4BC8-8103-B962CBDD0ECD}","bytesSent":0,"bytesRecv":0,"packetsSent":0,"packetsRecv":0,"errin":0,"errout":0,"dropin":0,"dropout":0,"fifoin":0,"fifoout":0} send:0 recv:0
	8:{"name":"isatap.{8EAC5DC2-A08D-42DF-89B9-51F13456770A}","bytesSent":0,"bytesRecv":0,"packetsSent":0,"packetsRecv":0,"errin":0,"errout":0,"dropin":0,"dropout":0,"fifoin":0,"fifoout":0} send:0 recv:0
	9:{"name":"WLAN 2","bytesSent":0,"bytesRecv":0,"packetsSent":0,"packetsRecv":0,"errin":0,"errout":0,"dropin":0,"dropout":0,"fifoin":0,"fifoout":0} send:0 recv:0
	 */
}
