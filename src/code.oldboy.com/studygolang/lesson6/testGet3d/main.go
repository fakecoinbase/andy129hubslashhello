package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime/pprof"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"code.oldboy.com/studygolang/lesson6/testLogbyChannel/logutils"
	"github.com/PuerkitoBio/goquery"
)

// FIRST_BALL 代表第一个球
const (
	FIRSTBALL = iota
	SECONDBALL
	THIRDBALL
)

// D 是一个结构体，储存 红球 和 蓝球，以 int 类型存储
type D struct {
	FirstNum  int
	SecondNum int
	ThirdNum  int
}

// DInfo 是一个结构体，储存双色球的开奖d信息
type DInfo struct {
	date     time.Time // 开奖日期
	qiStr    string    // 第几期
	numInfo  string    // 开奖号码
	sellInfo string    // 销售金额
	d        *D        // 单独定义一个 SSQ 结构体，存放 中奖号码(数字)
}

// NumberCount 是一个结构体，储存 1-33 个红号每个号码在历史数据中出现的次数， 1-16 个蓝号每个号码在历史数据中出现的次数
type NumberCount struct {
	ZeroCount  int
	OneCount   int
	TwoCount   int
	ThreeCount int
	FourCount  int
	FiveCount  int
	SixCount   int
	SevenCount int
	EightCount int
	NineCount  int
}

// 存储双色球信息的通道
var dChan chan *DInfo

// 任务统计通道
var checkChan chan struct{}

var wg sync.WaitGroup

var dUrlFormat string = "http://kaijiang.zhcw.com/zhcw/html/3d/list_%d.html"

var File_Path = "./3d.txt" // 双色球信息存放路径

// 双色球开奖信息会隔几天更新一次，所以网页数量也会有变动，是一个 增加的过程
var total = 275 // 总共抓取 127个网页，总共开启127个goroutine，每个 goroutine 抓取一个网页

var logger logutils.Logger

// 示例： 抓取双色球网页信息
func test1() {
	GetContent("http://kaijiang.zhcw.com/zhcw/html/3d/list_1.html")

	// writeSSQInfoToFile()
}

// 并发抓取网页上的图片
func main() {

	var isCPUPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	// CPU使用情况记录
	if isCPUPprof {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed, err:%v\n", err)
			return
		}
		pprof.StartCPUProfile(file)
		defer file.Close()
		defer pprof.StopCPUProfile()
	}

	logger = logutils.NewFileLogger("Info", "./log/", "3d.log")

	defer logger.Close() // 程序执行完记得关闭

	// 总网络上抓取数据并保存
	// "http://kaijiang.zhcw.com/zhcw/html/3d/list_1.html"  中 list_1 至 list_275 总共275个网页
	// 其中 list_1 是最新一页的双色球数据

	// fetchAll3D(275)

	parseSSQ()
	time.Sleep(time.Millisecond * 500)

	// 内存信息记录
	if isMemPprof {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed, err:%v\n", err)
			return
		}
		pprof.WriteHeapProfile(file)
		file.Close()
	}
}

func parseSSQ() {
	// 从文件中把保存好的双色球信息，再读取出来进行分析
	all3DInfoMap := get3DInfoFromFile()

	// 创建一个 []int 用于保存 map 的 key
	var keyArr = make([]int, 0, len(all3DInfoMap))
	for key := range all3DInfoMap {
		keyArr = append(keyArr, key)
		// fmt.Println("key : ", key)
	}

	// 通过sort 包里的功能将 keyArr 排序
	// sort.Ints(keyArr)

	// IntSlice 升序排序，然后返回排序后的[]int
	// Reverse 将升序进行反转，成为降序
	// (暂时不太明白) sort.Sort 并不保证排序的稳定性。如果有需要, 可以使用 sort.Stable
	// sort.Sort(sort.Reverse(sort.IntSlice(keyArr)))
	sort.Stable(sort.Reverse(sort.IntSlice(keyArr)))
	/*
		for i, v := range keyArr {

			logger.Info("i : %d , v : %d", i, v)
			fmt.Printf("i : %d , v : %d\n", i, v)
		}
	*/
	// fmt.Println("keyArr len : ", len(keyArr))

	// 查询 map 中任意一期彩票数据是否与 历史其它期 开奖结果一样。
	// searchMatch("088", all3DInfoMap)

	// 查询 map 中任意一期彩票数据是否与 历史其它期 开奖结果一样。
	// searchMatchFromMap(keyArr, all3DInfoMap)

	// calcNumPer(keyArr, all3DInfoMap)

	// 000 - 999 这之间的数字在 历史数据中出现的频率
	calcZeroToNine(all3DInfoMap)

	// 无序与有序 遍历map
	// 有序： 根据排序后的 key 进行 遍历 map , 达到有序遍历的目的
	/*
		for _, k := range keyArr {
			ssqObj := allSsqInfoMap[k]
			fmt.Printf("key : %d , numInfo : %s , blueBall : %d\n", k, ssqObj.numInfo, ssqObj.ssq.BlueBall)
		}
	*/
	// fmt.Printf("map---> len : %d\n", len(allSsqInfoMap))

	/*  无序： 默认的 map 是无序的
	for i, v := range allSsqInfoMap {
		fmt.Printf("key : %d , numInfo : %s , blueBall : %d\n", i, v.numInfo, v.ssq.BlueBall)
	}
	*/
}

// 产生 000-999， 这1000 个字符串 （模拟 3D单选）
func calcZeroToNine(all3DInfoMap map[int]*DInfo) {

	// 000 - 999 每个数字都与 历史数据进行匹配
	for i := 0; i < 1000; i++ {
		var numStr string
		if i < 10 {
			numStr = "00" + strconv.Itoa(i)
		} else if i < 100 {
			numStr = "0" + strconv.Itoa(i)
		} else {
			numStr = strconv.Itoa(i)
		}
		calc3DNumberCount(numStr, all3DInfoMap)
	}
}

// 000-999 筛选匹配历史数据，从而预测下一期号码出现的范围
func calc3DNumberCount(numStr string, all3DInfoMap map[int]*DInfo) {

	// logger.Info("############################################################################################\n")
	// logger.Info("正在匹配查询 : %s", numStr)
	// fmt.Printf("正在匹配查询 : %s\n", numStr)

	var count int = 0

	var ok1 bool = true

	for k2, v2 := range all3DInfoMap {

		// 每一期的号码 与历史数据进行对比 (比当前开奖号码 早的历史数据)
		// (废除)筛选条件1：(记录历史中重复的次数, 范围20100101 - 20190331)
		// k - 10000,  例如： 20200331 - 10000 = 20190331 (当前日期一年之前的时间)
		// (更新)筛选条件1：(不加范围，全局搜索)
		if numStr == v2.numInfo {
			// 筛选条件2：如果号码在 半年内出现过，则抛弃
			if checkDateAsHalfYear(k2) {
				ok1 = false
				break
			} else {
				count++
			}
		}
	}
	// 筛选条件3：(去除历史重复率小于3次的号码)
	// 历史数据中 有重复出现3次及以上的，我们打印出来
	if ok1 && count >= 3 { // 半年内没有出现过 并且出现次数至少3次
		// logger.Info("号码：%s", numStr)

		// 筛选条件4：(去除连号)
		if !checkNumberSeq(numStr) {
			// logger.Info("筛掉一年之内开出的号码，在历史中重复出现3次的号码(非连号)：%s", numStr)

			//  筛选条件5：(去除和值小于10的号 或者 大于20的 号码)
			if !checkNumberSum(numStr) {
				// logger.Info("和值大于10 并且小于20的 号：%s", numStr)

				// 筛选条件6：去除 全奇数或 全偶数的号码
				if !isOddOrEvenNumber(numStr) {
					// logger.Info("和值大于10 并且小于20, 并且 非全奇或非全偶的 号：%s", numStr)
					// 筛选条件7： 去除和值的个位数等于 N， 跨度值等于 M 的号码
					n, m := 3, 7
					if !isSumNumberEqualN(numStr, n) && !isValueWidthEqualM(numStr, m) {
						// logger.Info("和值大于10 并且小于20, 并且 非全奇或非全偶的 , 抛去和值 %d , 跨度值 %d 的号：%s", n, m, numStr)
						// 筛选条件8：(筛选每一位不经常出现的数字)
						first, second, third := 3, 0, 3
						if !isLessAppear(numStr, first, second, third) {
							//logger.Info("10<=sum<=20 ,和值非 %d , 跨度值非 %d, filter：%d %d %d %s", n, m, first, second, third, numStr)

							// 筛选条件9： (二次筛选 和值与跨度值)
							n2, m2 := 5, 1
							if !isSumNumberEqualN(numStr, n2) && !isValueWidthEqualM(numStr, m2) {
								// logger.Info("10<=sum<=20 ,和值非 %d,%d , 跨度值非 %d,%d, filter：%d %d %d %s", n, n2, m, m2, first, second, third, numStr)
								// 筛选条件9： (三次筛选 和值与跨度值)
								n3, m3 := 4, 6
								if !isSumNumberEqualN(numStr, n3) && !isValueWidthEqualM(numStr, m3) {
									//logger.Info("10<=sum<=20 ,和值非 %d,%d,%d , 跨度值非 %d,%d,%d, filter：%d %d %d, *%s*", n, n2, n3, m, m2, m3, first, second, third, numStr)

									// 筛选条件10： (筛选首位，中位，末位 范围)
									if isCoverIn(numStr, 2, 9, FIRSTBALL) && isCoverIn(numStr, 1, 8, SECONDBALL) && isCoverIn(numStr, 0, 6, THIRDBALL) {
										//logger.Info("10<=sum<=20 ,和值非 %d,%d,%d , 跨度值非 %d,%d,%d, filter：%d %d %d, 筛选三位范围 *%s*", n, n2, n3, m, m2, m3, first, second, third, numStr)

										// 筛选条件11：（第二次筛选首位，中位，末位 不太可能出现的号码）
										if isFilterIn(numStr, 5, FIRSTBALL) && isFilterIn(numStr, -1, FIRSTBALL) && isFilterIn(numStr, -1, SECONDBALL) && isFilterIn(numStr, 5, SECONDBALL) && isFilterIn(numStr, 3, SECONDBALL) && isFilterIn(numStr, 9, THIRDBALL) && isFilterIn(numStr, -1, THIRDBALL) && isFilterIn(numStr, -1, THIRDBALL) {

											// logger.Info("10<=sum<=20 ,和值非 %d,%d,%d , 跨度值非 %d,%d,%d, filter：%d %d %d, 筛选三位范围, 过滤3位不常出的号 *%s*", n, n2, n3, m, m2, m3, first, second, third, numStr)
											// 筛选条件12 ： 抛弃与上一期开奖号码有 两个数字相同的号码
											if !isEqualTowNum(numStr, "292") {
												logger.Info("最终结果：%s", numStr)
											}

										}
									}
								}
							}
						}
					}

				}
			}

			// 筛选条件 5B （保留和值小于10的号）
			/*
				if checkNumberSumLessTen(numStr) {
					logger.Info("和值小于10号：%s", numStr)
				}
			*/
		}
	} else {

		// logger.Info("重复率不高，只有：%d 次", len(arr))
		fmt.Printf(" %s 重复率不高，或者半年内出现过， 只有：%d 次\n", numStr, count)

	}
	//logger.Info("############################################################################################\n")
}

// 判断 numStr 是否有两个数字与 上一期开奖号码 一样
func isEqualTowNum(numStr string, lastNum string) bool {

	var count int = 0
	var matchStr string
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if numStr[i:i+1] != matchStr && numStr[i:i+1] == lastNum[j:j+1] {
				matchStr = numStr[i : i+1]
				count++
				continue
			}
		}
		// 循环完一遍之后，将matchStr 置为 "", 再进行下一次比较
		// 避免出现 "661" 与 "663" 之间比较失败的问题
		matchStr = ""
	}

	if count >= 2 {
		return true
	}

	return false

}

// 判断numStr 中 flag 球的值是否等于 filterNum
func isFilterIn(numStr string, filterNum, flag int) bool {
	if filterNum == -1 { // 如果等于 -1， 则直接返回 true
		return true
	}
	firstNum, err := strconv.Atoi(numStr[0:1])
	secondNum, err := strconv.Atoi(numStr[1:2])
	thirdNum, err := strconv.Atoi(numStr[2:3])
	if err != nil {
		HandleError(err, "checkNumberSum --> Atoi")
	}

	if flag == FIRSTBALL {
		if firstNum != filterNum {
			return true
		}
	}
	if flag == SECONDBALL {
		if secondNum != filterNum {
			return true
		}
	}
	if flag == THIRDBALL {
		if thirdNum != filterNum {
			return true
		}
	}
	return false
}

// 判断numStr 中 flag 球是否在  [left, right] 范围之内
func isCoverIn(numStr string, left, right int, flag int) bool {
	firstNum, err := strconv.Atoi(numStr[0:1])
	secondNum, err := strconv.Atoi(numStr[1:2])
	thirdNum, err := strconv.Atoi(numStr[2:3])
	if err != nil {
		HandleError(err, "checkNumberSum --> Atoi")
	}

	if flag == FIRSTBALL {
		if firstNum >= left && firstNum <= right {
			return true
		}
	}
	if flag == SECONDBALL {
		if secondNum >= left && secondNum <= right {
			return true
		}
	}
	if flag == THIRDBALL {
		if thirdNum >= left && thirdNum <= right {
			return true
		}
	}
	return false
}

// 判断这个号码的每一位的数字， 是否都很少出现
func isLessAppear(numStr string, first, second, third int) bool {
	firstNum, err := strconv.Atoi(numStr[0:1])
	secondNum, err := strconv.Atoi(numStr[1:2])
	thirdNum, err := strconv.Atoi(numStr[2:3])
	if err != nil {
		HandleError(err, "checkNumberSum --> Atoi")
	}
	// 判断每位数字 与 fist, second, third 比较， 只要有一组成立，则判断为 很少出现的数字
	if firstNum == first || secondNum == second || thirdNum == third {
		return true
	}
	return false
}

// 解析 numStr(三位数) 返回 int 类型的最大值和最小值
func parseMaxMin(numStr string) (int, int) {
	firstNum, err := strconv.Atoi(numStr[0:1])
	secondNum, err := strconv.Atoi(numStr[1:2])
	thirdNum, err := strconv.Atoi(numStr[2:3])
	if err != nil {
		HandleError(err, "checkNumberSum --> Atoi")
	}

	var arrSlice = [3]int{firstNum, secondNum, thirdNum}

	// 总共执行 len -1 趟
	for i := 0; i < 3-1; i++ {
		for j := 0; j < 3-i-1; j++ {
			if arrSlice[j] >= arrSlice[j+1] {
				// 交换数据
				arrSlice[j], arrSlice[j+1] = arrSlice[j+1], arrSlice[j]
			}
		}
	}
	return arrSlice[2], arrSlice[0]
}

// 最大值与最小值之间的差值 跨度是否等于 M
func isValueWidthEqualM(numStr string, m int) bool {

	max, min := parseMaxMin(numStr)

	if (max - min) == m {
		return true
	}
	return false
}

// 和值的个位数是否 等于 N
func isSumNumberEqualN(numStr string, n int) bool {
	firstNum, err := strconv.Atoi(numStr[0:1])
	secondNum, err := strconv.Atoi(numStr[1:2])
	thirdNum, err := strconv.Atoi(numStr[2:3])

	if err != nil {
		HandleError(err, "checkNumberSum --> Atoi")
	}

	sum := (firstNum + secondNum + thirdNum)

	if sum%10 == n {
		return true
	}
	return false
}

// 检查是否为 全奇数或 全偶数
func isOddOrEvenNumber(numStr string) bool {
	firstNum, err := strconv.Atoi(numStr[0:1])
	secondNum, err := strconv.Atoi(numStr[1:2])
	thirdNum, err := strconv.Atoi(numStr[2:3])

	if err != nil {
		HandleError(err, "checkNumberSum --> Atoi")
	}
	if (firstNum&1) == 0 && (secondNum&1) == 0 && (thirdNum&1) == 0 {
		//fmt.Printf(" %s 是偶数\n", numStr)
		return true
	}
	if (firstNum&1) != 0 && (secondNum&1) != 0 && (thirdNum&1) != 0 {
		//fmt.Printf(" %s 是奇数\n", numStr)
		return true
	}
	return false

}

// 检查传入的时间 是否为 当前日期的半年之内
func checkDateAsHalfYear(dateInt int) bool {
	now := time.Now()
	// 当前日期减去 180天（半年）的时间
	half := now.Add(-(time.Hour * 24 * 180))

	halfInt, err := dateToInt(half)

	if err != nil {
		HandleError(err, "checkDateAsHalfYear --> dateToInt")
	}
	// 传入的时间 是 当前时间半年之内的时间
	if dateInt >= halfInt {
		return true
	}
	return false
}

// 获取和值小于 10 的数字
func checkNumberSumLessTen(numStr string) bool {
	firstNum, err := strconv.Atoi(numStr[0:1])
	secondNum, err := strconv.Atoi(numStr[1:2])
	thirdNum, err := strconv.Atoi(numStr[2:3])

	if err != nil {
		HandleError(err, "checkNumberSum --> Atoi")
	}

	sum := (firstNum + secondNum + thirdNum)

	// fmt.Printf("firtNum : %d, secondNum : %d, thirdNum : %d, sum : %d, 号码：%s\n", firstNum, secondNum, thirdNum, sum, numStr)

	if sum <= 10 {
		return true
	}
	return false
}

// 检查和值是否小于 10 或者 大于 20
func checkNumberSum(numStr string) bool {
	firstNum, err := strconv.Atoi(numStr[0:1])
	secondNum, err := strconv.Atoi(numStr[1:2])
	thirdNum, err := strconv.Atoi(numStr[2:3])

	if err != nil {
		HandleError(err, "checkNumberSum --> Atoi")
	}

	sum := (firstNum + secondNum + thirdNum)

	// fmt.Printf("firtNum : %d, secondNum : %d, thirdNum : %d, sum : %d, 号码：%s\n", firstNum, secondNum, thirdNum, sum, numStr)

	if sum < 10 || sum > 20 {
		return true
	}
	return false
}

// 检查是否为 连续的号码 (123 或 321 或 666)
func checkNumberSeq(numStr string) bool {

	firstNum, err := strconv.Atoi(numStr[0:1])
	secondNum, err := strconv.Atoi(numStr[1:2])
	thirdNum, err := strconv.Atoi(numStr[2:3])

	if err != nil {
		HandleError(err, "checkNumberSeq --> Atoi")
	}

	// 例如：222, 三个数相等，则返回 true
	if (firstNum == secondNum) && (secondNum == thirdNum) {
		return true
	}
	// 例如： 789, 则返回 true
	if (firstNum+1 == secondNum) && (secondNum+1 == thirdNum) {
		return true
	}
	// 例如：987, 则返回 true
	if (firstNum-1 == secondNum) && (secondNum-1 == thirdNum) {
		return true
	}

	return false
}

// 与历史数据想比较， 计算每一期彩票，每个数字出现的频率
func calcNumPer(keyArr []int, all3DInfoMap map[int]*DInfo) {
	for _, k := range keyArr {
		dObj := all3DInfoMap[k]
		logger.Info("############################################################################################\n")
		logger.Info("开奖日期 : %d , 开奖号码 : %s", k, dObj.numInfo)
		fmt.Printf("开奖日期 : %d , 开奖号码 : %s\n", k, dObj.numInfo)
		if k == 20100103 { // 只提取2010年1月3日至 2020年3月31 之间的开奖号码 与 历史数据作对比
			logger.Info("统计至 %d", k)
			//fmt.Printf("统计至 %d\n", k)
			break
		}

		d := dObj.d // 拿出需要分析的 中奖号码

		// 统计每个号码出现的频率
		var dFirstCount, dSecondCount, dThirdCount int

		// 统计 0-9 个红号，每个号码出现的频率
		var numCountObj = &NumberCount{}

		// fmt.Printf("key : %d , numInfo : %s , blueBall : %d\n", k, ssqObj.numInfo, ssqObj.ssq.BlueBall)
		for k2, v2 := range all3DInfoMap {
			if k2 < k { // 开奖号码 只与 历史数据相比： 例如：19年的开奖号码 只与 19年之前的数据进行对比，而不能找 2020年的数据

				history := v2.d

				/*  不分位置，只判断数字出现的频率 (类似于  组选)
				if check3dNum(d.FirstNum, history) {
					dFirstCount++
				}
				if check3dNum(d.SecondNum, history) {
					dSecondCount++
				}
				if check3dNum(d.ThirdNum, history) {
					dThirdCount++
				}
				*/

				// 判断位置与数字 是否都相同 （类似于 直选）
				if d.FirstNum == history.FirstNum {
					dFirstCount++
				}
				if d.SecondNum == history.SecondNum {
					dSecondCount++
				}
				if d.ThirdNum == history.ThirdNum {
					dThirdCount++
				}

				// 计算 0-9 每个数字在历史数据中出现的次数
				calcNumberCount(history, numCountObj)

			}
		}

		logger.Info("3D：%d (%d次), %d (%d次), %d (%d次) ", d.FirstNum, dFirstCount, d.SecondNum, dSecondCount, d.ThirdNum, dThirdCount)
		logger.Info("\n-----------------------------------数据分析--------------------------------------\n")
		logger.Info("0(%d), 1(%d), 2(%d), 3(%d), 4(%d), 5(%d), 6(%d), 7(%d), 8(%d), 9(%d)", numCountObj.ZeroCount, numCountObj.OneCount, numCountObj.TwoCount, numCountObj.ThreeCount, numCountObj.FourCount, numCountObj.FiveCount, numCountObj.SixCount, numCountObj.SevenCount, numCountObj.EightCount, numCountObj.NineCount)

	}
}

// 计算 0-9 每个数字在历史数据中出现的次数
func calcNumberCount(history *D, numCount *NumberCount) {
	for i := 0; i < 10; i++ {
		if check3dNum(i, history) {
			switch i {
			case 0:
				numCount.ZeroCount++
			case 1:
				numCount.OneCount++
			case 2:
				numCount.TwoCount++
			case 3:
				numCount.ThreeCount++
			case 4:
				numCount.FourCount++
			case 5:
				numCount.FiveCount++
			case 6:
				numCount.SixCount++
			case 7:
				numCount.SevenCount++
			case 8:
				numCount.EightCount++
			case 9:
				numCount.NineCount++
			}
		}
	}
}

// 检查一注彩票里的 每个红号 是否出现在历史数据中(与历史数据中每一注里面的每个红号进行对比)，如果有 则返回 true
func check3dNum(num int, history *D) bool {

	switch num {
	case history.FirstNum, history.SecondNum, history.ThirdNum:
		return true
	default:
		return false
	}
}

//  查询 map 中任意一期彩票数据是否与 历史其它期 开奖结果一样。
func searchMatchFromMap(keyArr []int, all3DInfoMap map[int]*DInfo) {
	for _, k := range keyArr {
		dObj := all3DInfoMap[k]
		logger.Info("check : %d", k)
		fmt.Printf("check : %d\n", k)
		// fmt.Printf("key : %d , numInfo : %s , blueBall : %d\n", k, ssqObj.numInfo, ssqObj.ssq.BlueBall)
		for k2, v2 := range all3DInfoMap {
			// 这是与所有的历史数据进行匹配（但 17年的会与 2020年的数据对比）, 这个可以查询 宏观下，一个数出现的频率
			/*
				if k2 != k && dObj.numInfo == v2.numInfo {
					logger.Info("开奖日期：%d == %d , 相同号码：%s", k2, k, dObj.numInfo)
					fmt.Printf("开奖日期：%d == %d , 相同号码：%s\n", k2, k, dObj.numInfo)
				}
			*/

			// 每一期的号码 与历史数据进行对比 (比当前开奖号码 早的历史数据)
			// k - 10000,  例如： 20200331 - 10000 = 20190331 (当前日期一年之前的时间)
			if k2 <= (k-10000) && k2 > 20100000 { // 在当前时间一年以前， 并且是在 2010后，这段时间的数据进行对比
				if dObj.numInfo == v2.numInfo {
					logger.Info("开奖日期：%d == %d , 相同号码：%s", k2, k, dObj.numInfo)
					fmt.Printf("开奖日期：%d == %d , 相同号码：%s\n", k2, k, dObj.numInfo)
				}
			}

		}
	}
}

// 指定一组彩票数据，与历史数据进行查询
func searchMatch(numInfo string, all3DInfoMap map[int]*DInfo) {

	var matchArr = make([]int, 0, 10)
	for k2, v2 := range all3DInfoMap {
		if numInfo == v2.numInfo {
			matchArr = append(matchArr, k2) // 有匹配的则加入到 arr 里面
		}
	}
	if len(matchArr) > 0 {
		fmt.Printf("%s 查找到 %d 匹配！\n", numInfo, len(matchArr))
		for i, v := range matchArr {
			fmt.Printf("i : %d, v : %d\n", i, v)
		}
	} else {
		fmt.Println("未找到匹配项！")
	}
}

// 从文件中读取双色球信息
func get3DInfoFromFile() (all3DInfoMap map[int]*DInfo) {
	file, err := os.Open(File_Path)
	if err != nil {
		HandleError(err, "get3DInfoFromFile  --> OpenFile")
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	all3DInfoMap = make(map[int]*DInfo, 10000)

	for {

		lineStr, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Printf("文件已读完 EOF : %s\n", lineStr)
			// fmt.Println("文件已读完")
			return
		}
		if err != nil {
			HandleError(err, "getSsqInfoFromFile --> ReadString")
			return
		}
		lineStr = strings.TrimSpace(lineStr) // 去除空格，等

		dInfo := get3DObjFromLineStr(lineStr) // 读取每一行信息，将其转换为 DInfo 结构体

		key, err := dateToInt(dInfo.date) // 将开奖日期转换为 int 类型作为 map 的 key
		if err != nil {
			return
		}
		all3DInfoMap[key] = dInfo // 将 DInfo 结构体指针作为 value

		// fmt.Printf("----ReadString : %s\n", lineStr)

	}

}

// 将结构体中的日期对象 转换为 int64类型
func dateToInt(date time.Time) (int, error) {
	dateStr := date.Format("20060102")
	dateInt, err := strconv.Atoi(dateStr)
	if err != nil {
		HandleError(err, "dateToInt --> ParseInt")
	}
	return dateInt, err
}

// [2020-01-21] @2020009@ *03 06 08 14 19 26 +12* $405,733,380$ ^16^ !(京 辽 沪 苏..)! ~253~
// 把从文件中读取的每一行数据 解析为 SSQInfo 结构体
func get3DObjFromLineStr(lineStr string) *DInfo {

	var dInfo = &DInfo{}

	dInfo.date = getDateFromLineStr(lineStr)
	// fmt.Printf("date : %v\n", (*ssqInfo).date)

	dInfo.qiStr = getQiStrFromLineStr(lineStr)

	dInfo.numInfo = getNumInfoFromLineStr(lineStr)

	dInfo.sellInfo = getSellInfoFromLineStr(lineStr)

	// fmt.Printf("ssqInfo : %#v\n", *ssqInfo)

	d := getDObjFromNumInfo(dInfo.numInfo)

	dInfo.d = d

	return dInfo
}

// 327
// 将双色球 字符串形式的信息 转换为 D 结构体中的int
func getDObjFromNumInfo(numInfo string) (d *D) {

	firstNum, err := strconv.Atoi(numInfo[0:1])
	secondNum, err := strconv.Atoi(numInfo[1:2])
	thirdNum, err := strconv.Atoi(numInfo[2:3])

	if err != nil {
		HandleError(err, "getDObjFromNumInfo --> Atoi")
	}

	return &D{
		FirstNum:  firstNum,
		SecondNum: secondNum,
		ThirdNum:  thirdNum,
	}
}

func stringToInt(str string) int {
	value, err := strconv.Atoi(str)
	if err != nil {
		HandleError(err, "stringToInt")
	}
	return value
}

/*  除了获取 date, 其他的可以通用该方法
func getItemFromLineStr(lineStr, sub string) string {
	index := strings.Index(lineStr, sub)
	lastIndex := strings.LastIndex(lineStr, sub)
	str := lineStr[index+1 : lastIndex]
	return str
}
*/

// 从文本中读取的每一行数据中，解析出日期字符串，返回 time.Time 对象
func getDateFromLineStr(lineStr string) time.Time {
	dateIndex := strings.Index(lineStr, "[")
	dateLastIndex := strings.Index(lineStr, "]")
	dateStr := lineStr[dateIndex+1 : dateLastIndex]
	// fmt.Println("dateStr : ", dateStr)
	date := stringToDate(dateStr)
	return date
}

// 从文本中读取的每一行数据中，解析出期号字符串，返回 string
func getQiStrFromLineStr(lineStr string) string {
	qiStrIndex := strings.Index(lineStr, "@")
	qiStrLastIndex := strings.LastIndex(lineStr, "@")
	qiStr := lineStr[qiStrIndex+1 : qiStrLastIndex]
	return qiStr
}

// 从文本中读取的每一行数据中，解析出开奖号码字符串，返回 string
func getNumInfoFromLineStr(lineStr string) string {
	numInfoIndex := strings.Index(lineStr, "*")
	numInfoLastIndex := strings.LastIndex(lineStr, "*")
	numInfo := lineStr[numInfoIndex+1 : numInfoLastIndex]
	return numInfo
}

// 从文本中读取的每一行数据中，解析出销售金额字符串，返回 string
func getSellInfoFromLineStr(lineStr string) string {
	sellInfoIndex := strings.Index(lineStr, "$")
	sellInfoLastIndex := strings.LastIndex(lineStr, "$")
	sellInfo := lineStr[sellInfoIndex+1 : sellInfoLastIndex]
	return sellInfo
}

// 从通道中取出双色球信息，然后写入到 一个默认文件中。
func write3DInfoToFile() {
	file, err := os.OpenFile(File_Path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0664)
	if err != nil {
		HandleError(err, "write3DInfoToFile  --> OpenFile")
		return
	}

	defer file.Close() // 写入完毕后 关闭文件

	for dObj := range dChan {
		dMsg := format3DInfo(dObj)
		fmt.Fprintln(file, dMsg) // 写入
	}
}

// 将 双色球结构体里各字段的信息 拼接成一个指定格式的字符串，然后返回该字符串。
func format3DInfo(s *DInfo) string {
	date := s.date
	qiStr := s.qiStr
	numInfo := s.numInfo
	sellInfo := s.sellInfo

	dateStr := dateToString(date)
	ssqMsg := fmt.Sprintf("[%s] @%s@ *%s* $%s$\r", dateStr, qiStr, numInfo, sellInfo)
	// fmt.Println("ssqMsg : ", ssqMsg)
	return ssqMsg
}

// 网络抓取3d信息，入口程序
func fetchAll3D(htmlTotal int) {

	total = htmlTotal
	dChan = make(chan *DInfo, 10000)

	checkChan = make(chan struct{}, total)

	// 创建275个 goroutine, 每个 goroutine 抓取不同的 网页
	for i := 1; i < total+1; i++ {
		wg.Add(1)
		go get3DInfo(dUrlFormat, i, checkChan)

		time.Sleep(time.Millisecond * 200)
	}

	// 创建一个 检查通道 的goroutine
	wg.Add(1)
	go checkTask()
	wg.Wait()

	fmt.Printf("总共抓取 %d 条数据,  通道总容量: %d\n", len(dChan), cap(dChan))
	write3DInfoToFile()

	fmt.Println("抓取程序结束！")

}

func get3DInfo(format string, i int, checkChan chan struct{}) {
	urlStr := fmt.Sprintf(format, i)
	err := GetContent(urlStr)

	if err != nil {
		HandleError(err, "get3DInfo --> GetContent")
	}

	// checkChan
	defer func() {
		checkChan <- struct{}{} // 获取双色球网页的内容，无论成功失败，都加一个标识代表这个 goroutine 结束了任务
		wg.Done()
	}()
}

// GetContent 网络爬虫获取内容
func GetContent(urlStr string) (err error) {
	// 根据网址获取 网页信息
	resp, err := http.Get(urlStr)
	HandleError(err, "http.Get()")

	fmt.Println("抓取网址：", urlStr)

	if resp == nil {
		// fmt.Printf("%s 访问失败！\n", urlStr)
		err = fmt.Errorf("%s 访问失败！\n", urlStr)
		return
	}

	defer func() {
		if resp != nil {
			//fmt.Printf("resp.Body 类型：%T\n", resp.Body)
			resp.Body.Close()
		}
	}()

	doc, err2 := goquery.NewDocumentFromReader(resp.Body)

	if err2 != nil {
		return err2
	}

	// Find the review items
	doc.Find("table[class=wqhgt]").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		//fmt.Printf("第 %d 个 <table>\n", i)
		s.Find("tr").Each(func(i int, se *goquery.Selection) {

			// fmt.Printf("第 %d 个 <tr>\n", i)
			// fmt.Println("tr 个数 : ", s.Find("tr").Length())
			if i == 0 || i == 1 || i == (s.Find("tr").Length()-1) {
				// fmt.Printf("第 %d 个 <tr> 不解析， return\n", i)
				return
			}

			var dInfo = &DInfo{}

			se.Find("td").Each(func(i int, sel *goquery.Selection) {
				// fmt.Printf("第 %d 个 <td>\n", i)

				switch i {
				case 0: // 3d开奖日期
					dateStr := sel.Text()
					dInfo.date = stringToDate(dateStr)
					//fmt.Printf("开奖日期：%v\n", dInfo.date)
				case 1: // 3d开奖期数
					qiStr := sel.Text()
					dInfo.qiStr = qiStr
					//fmt.Printf("开奖期号：%s\n", dInfo.qiStr)
				case 2: // 3d 开奖号码
					ssqNumber := sel.Find("em").Text()
					// ssqNumber = format3dNumInfo(ssqNumber)
					dInfo.numInfo = ssqNumber
					//fmt.Printf("开奖号码：%s\n", dInfo.numInfo)
					/*
						case 3: // 单选  中的注数
						case 4: // 组选3 中的注数
						case 5: // 组选6 中的注数
					*/
				case 6: // 3d 销售额信息
					sellStr := sel.Text()
					sellStr = strings.TrimSpace(sellStr)
					dInfo.sellInfo = sellStr
					//fmt.Printf("销售额：%s 元\n", dInfo.sellInfo)
				}
			})
			dChan <- dInfo // 解析完一条3D信息后， 放入通道中

		})
	})

	fmt.Println("抓取网址：抓取完成")
	return
}

// checkChan
// 检查127 个goroutine 是否都完成了任务，然后关闭 检查通道和 储存图片链接的通道
func checkTask() {

	var i int
	for {

		<-checkChan // 取不到值，则会在这里一直阻塞
		i++
		if i == total { // 275 个goroutine 完成
			fmt.Println("--------------------total : 275")
			close(checkChan) // checkChan 里面的值取完了之后，代表 275个goroutine 都结束任务了
			close(dChan)     // 关闭 储存 图片链接的管道
			break
		}
	}
	wg.Done()
}

// 将日期字符串转换为 time.Time
func stringToDate(dateStr string) time.Time {
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		HandleError(err, "stringToDate")
	}
	// fmt.Println(date)
	return date
}

// 将日期对象转换为 字符串形式
func dateToString(date time.Time) string {
	dateStr := date.Format("2006-01-02")

	return dateStr
}

// 从网页中解析出来的 双色球号码，进行格式化整理
// 例如：02040715202704　　---> 02 04 07 15 20 27 +04
func formatSsqNumInfo(ssqAllInfo string) string {
	// var index = 0
	for i := 2; i < len(ssqAllInfo); i = i + 2 + 1 {
		if i == len(ssqAllInfo)-2 {
			ssqAllInfo = ssqAllInfo[:i] + " +" + ssqAllInfo[i:]
			break
		}
		ssqAllInfo = ssqAllInfo[:i] + " " + ssqAllInfo[i:]
	}

	return ssqAllInfo
}

// HandleError 处理错误
func HandleError(err error, why string) {

	if err != nil {
		fmt.Println(why, err)
	}
}
