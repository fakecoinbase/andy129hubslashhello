package main

// 导入 time 包
import (
	"fmt"
	"time"
)

// time 包示例
func main() {

	// test1()
	// test2()
	// test3()
	// test4()
	test5()
	// test6()
	// test7()
	// test8()
	// test9()
	// test10()
}

// 时间比较：After()
// 你就想象准时上班 和 迟到上班的人， 先后到公司的顺序
func test10() {
	now := time.Now()
	later := now.Add(time.Hour) // 一个小时后
	isAfter := now.After(later)
	fmt.Println(isAfter) // "false"
}

// 时间比较：Before()
// 你就想象准时上班 和 迟到上班的人， 先后到公司的顺序
func test9() {
	now := time.Now()
	later := now.Add(time.Hour) // 一个小时后
	isBefore := now.Before(later)
	fmt.Println(isBefore) // "true"
}

// 时间比较：Equal()
// Equal() 还会比较时区和位置信息
func test8() {
	now := time.Now()
	later := now.Add(time.Hour)   // 一个小时后
	fmt.Println(now == later)     // "false"
	fmt.Println(now.Equal(later)) // "false",  Equal() 还会比较时区和位置信息
}

// Sub(u Time)  当前时间减去 u,  计算两个时间中间间隔多长时间： 返回值： d Duration
func test7() {
	now := time.Now()
	fmt.Println(now)            // "2020-03-23 15:14:24.9980126 +0800 CST m=+0.005013501"
	later := now.Add(time.Hour) // 一个小时后
	fmt.Println(later)          // "2020-03-23 16:14:24.9980126 +0800 CST m=+3600.005013501"

	subTimeInt := now.Sub(later) // 当前时间，减去 零一个时间 later, 得到他们之间间隔多长时间：
	fmt.Println(subTimeInt)      // -1h0m0s
}

// Add(d Duration) 时间间隔，求 当前时间 一个小时之后是什么时间？: 返回值： t Time
func test6() {
	now := time.Now()
	fmt.Println(now)            // "2020-03-23 15:18:20.535948 +0800 CST m=+0.004002901"
	later := now.Add(time.Hour) // 一个小时后
	fmt.Println(later)          // "2020-03-23 16:18:20.535948 +0800 CST m=+3600.004002901"

	last := now.Add(-time.Hour) // 可以传入一个 负数，代表当前时间的前一个小时 是什么时间
	fmt.Println(last)           // "2020-03-23 14:18:20.535948 +0800 CST m=-3599.995997099"

}

// 格式化时间 （补充）
func test5() {

	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05.000")) // .000 或 .999  显示后面毫秒数
	fmt.Println(now.Format("2006-01-02 15:04:05.999"))

	/* 打印效果：
	2020-03-19 21:26:27.716
	2020-03-19 21:26:27.716
	*/

	// 补充一下， 2006 年如果只想保留2位年份，也可以写成  06-01-02
	fmt.Println(now.Format("06-01-02 15:04"))

	// 打印效果：
	// 20-03-19 21:31

	// 12小时制，  将15转换为 12小时制为： 03 , PM 代表12小时制，  Mon 显示星期几， Jan 显示几月
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	// 打印效果：
	// (如果只把 15 换成 12 小时制为 03， 那么打印出来的时间你无法分辨是 上午还是下午，
	// 所以在后面还要加上 PM 代表是 12小时制，那么就可以清楚的知道是 上午还是下午了， 例如下：
	//
	/* 我将系统时间 修改为 2020年6月23日，上午9点34分， 星期二
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))

	打印： 2020-06-23 09:34:19.527 AM Tue Jun

	*/
}

// 格式化时间 (将日期转化为 字符串)
func test4() {
	formatTime()

}

// 要严格遵循以下规则：
// 2006-01-02 15:04:05  (年-月-日 时:分:秒)   快速记忆：2006 1 2 3 4 5
func formatTime() {
	now := time.Now()
	// 格式化的模板为 Go语言诞生的时间： 2006年1月2日15点04分
	fmt.Println(now.Format("2006-01-02 15:04"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))

	// 示例1 ：(如上)
	/*  打印如下：
	2020-03-19-21:07
	2020/03/19 21:07
	21:07 2020/03/19
	2020/03/19
	*/

	// 示例2：
	/*
			fmt.Println(now.Format("2006  -01- 02-15:  04"))
			fmt.Println(now.Format("我是2006/01/02 15:04"))
			fmt.Println(now.Format("15:04&&2006/01/02"))
			fmt.Println(now.Format("   2006/01/02   "))

		打印如下 ：（依旧能按照指定的格式 打印准确的时间）
			2020  -03- 19-21:  09
			我是2020/03/19 21:09
			21:09&&2020/03/19
				2020/03/19

	*/

	// 示例3： (不遵循 2006 01 02 15:04 这几个数字的规则，则会格式化错误)
	/*

		fmt.Println(now.Format("2007-01-02 15:04"))
		fmt.Println(now.Format("2006/02/02 15:04"))
		fmt.Println(now.Format("16:04 2006/01/03"))
		fmt.Println(now.Format("2006/01/02 17"))

			打印如下 ：
				19007-03-19 21:12
				2020/19/19 21:12
				36:12 2020/03/09
				2020/03/19 37
	*/
}

// Ticker ， 计时器
func test3() {

	// 创建一个计时器， 时间间隔为 一秒  (time.Second)

	/* time 包中关于 计时器里面的值  设定如下：
	const (
		Nanosecond  Duration = 1
		Microsecond          = 1000 * Nanosecond
		Millisecond          = 1000 * Microsecond
		Second               = 1000 * Millisecond
		Minute               = 60 * Second
		Hour                 = 60 * Minute
	)
	*/
	ticker := time.Tick(time.Second * 5) // 也可修改为 ：time.Second * 5,   5秒

	for i := range ticker {
		fmt.Println(i) //  打印 计时器返回的值
	}

	/* 打印信息如下(每隔一秒打印一次，一直进行下去...)：

	2020-03-19 20:58:45.9168371 +0800 CST m=+1.005458801
	2020-03-19 20:58:46.9177445 +0800 CST m=+2.006366201
	2020-03-19 20:58:47.9170795 +0800 CST m=+3.005701201
	2020-03-19 20:58:48.9177681 +0800 CST m=+4.006389801
	2020-03-19 20:58:49.9176337 +0800 CST m=+5.006255401
	2020-03-19 20:58:50.9168419 +0800 CST m=+6.005463601

					......

					......
	*/
}

// 时间戳
func test2() {
	now := time.Now()        // 返回  time.Time 这个结构体
	fmt.Printf("%#v\n", now) // "time.Time{wall:0xbf94f579eb05adf4, ext:3503001, loc:(*time.Location)(0x587b40)}"

	// 从 1970年1月1日 (GMT +8) 至今的秒数
	fmt.Println(now.Unix()) // "1584620207"    // 该 Unix() 是 time.Time 结构体内部私有的方法
	// 从 1970年1月1日 (GMT +8) 至今的 纳秒数
	fmt.Println(now.UnixNano()) // "1584620207792210200"

	timestampToTimeObj(now.Unix())

}

// 时间戳 转换为一个 Time 结构体
func timestampToTimeObj(timeStamp int64) {
	// time.Unix(arg1,arg2):  arg1 代表一个时间戳 (秒),  arg2 代表一个偏移量 (纳秒) [0,999999999]
	timeObj := time.Unix(timeStamp, 0) // 该 Unix() 这是 time 包里公开的方法
	fmt.Println(timeObj.Year())
	fmt.Println(timeObj.Month())
	fmt.Println(timeObj.Day())
	fmt.Println(timeObj.Hour())
	fmt.Println(timeObj.Minute())
	fmt.Println(timeObj.Second())
}

// 通过调用 time.Time 这个结构体里的函数来 获取时间
func test1() {

	// time.Time struct
	now := time.Now()        // 返回  time.Time 这个结构体
	fmt.Printf("%#v\n", now) // "time.Time{wall:0xbf94f579eb05adf4, ext:3503001, loc:(*time.Location)(0x587b40)}"

	// 时间信息，是根据你系统时间的设定来获取的。
	fmt.Printf("year : %v\n", now.Year()) // 注意是调用函数， 而不是 字段 Year
	fmt.Printf("month : %v\n", now.Month())
	fmt.Printf("day : %v\n", now.Day())
	fmt.Printf("Hour : %v\n", now.Hour())
	fmt.Printf("Minute : %v\n", now.Minute())
	fmt.Printf("Second : %v\n", now.Second())
	fmt.Printf("Nanosecond : %v\n", now.Nanosecond()) //  纳秒 （9位数） ( 指当前时间秒后面的 9位数)

	/* 打印信息如下：
	year : 2020
	month : March
	day : 19
	Hour : 20
	Minute : 4
	Second : 18
	Nanosecond : 106045600
	*/
}
