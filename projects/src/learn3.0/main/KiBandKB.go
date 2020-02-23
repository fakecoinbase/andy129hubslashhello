package main

import "fmt"

// iota 机制存在局限。比如，因为不存在指数运算符，所以无从生成更为人熟知的 1000的幂 (KB,MB 等)。
// 对比 learn3.6.1 里面 对  KiB, MiB 等的定义
// 定义 KB,MB,GB 等常量
const (
	KB = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

/*
	KiB 与 KB 的区别
 */
func main() {

	fmt.Println("KiBandKB")
}

/*
		KiB、MiB与KB、MB的区别
		原创starshine 最后发布于2012-11-26 16:16:42 阅读数 80836  收藏
		展开
		    原来没太注意MB与MiB的区别，甚至没太关注还有MiB这等单位，今天认真了一下，发现两者还是有区别的，具体的差别是MB等单位以10为底数的指数，MiB是以2为底数的指数，如：1KB=10^3=1000, 1MB=10^6=1000000=1000KB,1GB=10^9=1000000000=1000MB,而 1KiB=2^10=1024,1MiB=2^20=1048576=1024KiB。与我们密切相关的是我们在买硬盘的时候，操作系统报的数量要比产品标出或商家号称的小一些，主要原因是标出的是以MB、GB为单位的，1GB就是1,000,000,000 Byte，而操作系统是以2进制为处理单位的，因此检查硬盘容量时是以MiB、GiB为单位，1GB=2^30=1,073,741,824，相比较而言，1GiB要比1GB多出1,073,741,824-1,000,000,000=73,741,824，所以检测实际结果要比标出的少一些。

		具体的对比关系如下：


		    十进制单位                              二进制单位
		------------------------------------------------------	 
		名字	缩写	次方	 名字	缩写	次方
		kilobyte	KB	10^3	kibibyte	KiB	2^10
		megabyte	MB	10^6	mebibyte	MiB	2^20
		gigabyte	GB	10^9	gibibyte	GiB	2^30
		terabyte	TB	10^12	tebibyte	TiB	2^40
		petabyte	PB	10^15	pebibyte	PiB	2^50
		exabyte	    EB	10^18	exbibyte	EiB	2^60
		zettabyte	ZB	10^21	zebibyte	ZiB	2^70
		yottabyte	YB	10^24	yobibyte	YiB	2^80
		————————————————
		版权声明：本文为CSDN博主「starshine」的原创文章，遵循 CC 4.0 BY-SA 版权协议，转载请附上原文出处链接及本声明。
		原文链接：https://blog.csdn.net/starshine/article/details/8226320
 */


/*
		Kib Kb KB KIB 区别
		今天和同事聊了一下Kib Kb KB KIB这几个单位的含义及其区别,自己在网上也查了查资料,总结如下:

		Ki 和 K 只是数学单位

		Ki = 1024

		K  = 1000

		这二者之间没有任何联系



		B 和 b 表示物理存储单位

		B = Byte 即一个字节

		b = bit    即一个二进制位

		了解计算机基本原理的都知道,二者的关系是 : 1 Byte = 8 bit



		所以有如下结果:

		1Kib = 1024 bit

		1Kb  = 1000 bit

		1KiB = 1024 Byte

		1KB  = 1000 Byte



		至于MB MiB Mb Mib 以此类推
 */