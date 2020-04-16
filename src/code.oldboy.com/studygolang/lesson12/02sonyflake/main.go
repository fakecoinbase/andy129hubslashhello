package main

import (
	"fmt"

	"github.com/sony/sonyflake"
)

// sonyflake 是基于 twitter 的snowflake (雪花算法) 编写的
// 关于雪花算法的原理及 优缺点，请查看 https://developer.51cto.com/art/201909/602525.htm
/*
	3. 特点(自增、有序、适合分布式场景)

	时间位：可以根据时间进行排序，有助于提高查询速度。
	机器id位：适用于分布式环境下对多节点的各个节点进行标识，可以具体根据节点数和部署情况设计划分机器位10位长度，如划分5位表示进程位等。
	序列号位：是一系列的自增id，可以支持同一节点同一毫秒生成多个ID序号，12位的计数序列号支持每个节点每毫秒产生4096个ID序号
	snowflake算法可以根据项目情况以及自身需要进行一定的修改。

	三、雪花算法的缺点

	雪花算法在单机系统上ID是递增的，但是在分布式系统多节点的情况下，所有节点的时钟并不能保证不完全同步，所以有可能会出现不是全局递增的情况。
*/

var (
	sonyFlake *sonyflake.Sonyflake
	machineID uint16 // 真正的分布式环境下必须从 zookeeper 或者 etcd 中获取
)

// 获取机器ID 的回调函数
func getMachineID() (uint16, error) {
	/*   真正的分布式环境下必须从 zookeeper 或者 etcd 中获取

	 */
	return machineID, nil
}

// Init 初始化 sonyflake
func Init(mID uint16) (err error) {
	machineID = mID
	st := sonyflake.Settings{}
	// st.MachineID 是一个函数类型
	st.MachineID = getMachineID
	if err != nil {
		return
	}
	sonyFlake = sonyflake.NewSonyflake(st)

	return

}

// GetID 获取全局唯一ID
func GetID() (id uint64, err error) {
	if sonyFlake == nil {
		fmt.Println("GetID err : ", err)
		return
	}
	return sonyFlake.NextID()
}

// twitter snowflake --> go 语言实现 sonyflake
func main() {

	Init(42445)

	for i := 0; i < 20; i++ {
		id, _ := GetID()
		fmt.Println(id)
	}

	/*  看的出来 雪花算法，ID 能保持趋势递增的效果

	297778451425830349
	297778451425895885
	297778451425961421
	297778451426026957
	297778451426092493
	297778451426158029
	297778451426223565
	297778451426289101
	297778451426354637
	297778451426420173
	297778451426485709
	297778451426551245
	297778451426616781
	297778451426682317
	297778451426747853
	297778451426813389
	297778451426878925
	297778451426944461
	297778451427009997
	297778451427075533

	*/

}
