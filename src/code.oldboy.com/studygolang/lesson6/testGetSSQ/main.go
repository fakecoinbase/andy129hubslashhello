package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"code.oldboy.com/studygolang/lesson6/testLogbyChannel/logutils"
	"github.com/PuerkitoBio/goquery"
)

// SSQ 是一个结构体，储存 红球 和 蓝球，以 int 类型存储
type SSQ struct {
	RedOne   int
	RedTwo   int
	RedThree int
	RedFour  int
	RedFive  int
	RedSix   int
	BlueBall int
}

// SSQInfo 是一个结构体，储存双色球的开奖信息
type SSQInfo struct {
	date        time.Time // 开奖日期
	qiStr       string    // 第几期
	numInfo     string    // 开奖号码
	sellInfo    string    // 销售金额
	oneNum      string    // 一等奖中奖数
	oneLocation string    // 一等奖中奖地点(哪些省)
	secondNum   string    // 二等奖中奖数
	ssq         *SSQ      // 单独定义一个 SSQ 结构体，存放 中奖号码(数字)
}

// NumberCount 是一个结构体，储存 1-33 个红号每个号码在历史数据中出现的次数， 1-16 个蓝号每个号码在历史数据中出现的次数
type NumberCount struct {
	OneCount         int
	TwoCount         int
	ThreeCount       int
	FourCount        int
	FiveCount        int
	SixCount         int
	SevenCount       int
	EightCount       int
	NineCount        int
	TenCount         int
	ElevenCount      int
	TwelveCount      int
	ThirteenCount    int
	FourteenCount    int
	FifteenCount     int
	SixteenCount     int
	SeventeenCount   int
	EighteenCount    int
	NineteenCount    int
	TwentyCount      int
	TwentyOneCount   int
	TwentyTwoCount   int
	TwentyThreeCount int
	TwentyFourCount  int
	TwentyFiveCount  int
	TwentySixCount   int
	TwentySevenCount int
	TwentyEightCount int
	TwentyNineCount  int
	ThirtyCount      int
	ThirtyOneCount   int
	ThirtyTwoCount   int
	ThirtyThreeCount int

	BlueOneCount      int
	BlueTwoCount      int
	BlueThreeCount    int
	BlueFourCount     int
	BlueFiveCount     int
	BlueSixCount      int
	BlueSevenCount    int
	BlueEightCount    int
	BlueNineCount     int
	BlueTenCount      int
	BlueElevenCount   int
	BlueTwelveCount   int
	BlueThirteenCount int
	BlueFourteenCount int
	BlueFifteenCount  int
	BlueSixteenCount  int
}

// DownloadPath 图片下载地址
var DownloadPath = "./img/"

// 存储双色球信息的通道
var ssqChan chan *SSQInfo

// 任务统计通道
var checkChan chan struct{}

var wg sync.WaitGroup

var ssqUrlFormat string = "http://kaijiang.zhcw.com/zhcw/html/ssq/list_%d.html"

var File_Path = "./ssq.txt" // 双色球信息存放路径

// 双色球开奖信息会隔几天更新一次，所以网页数量也会有变动，是一个 增加的过程
var total = 127 // 总共抓取 127个网页，总共开启127个goroutine，每个 goroutine 抓取一个网页

var logger logutils.Logger

// 示例： 抓取双色球网页信息
func test1() {
	GetContent("http://kaijiang.zhcw.com/zhcw/html/ssq/list_2.html")

	// writeSSQInfoToFile()
}

// 并发抓取网页上的图片
func main() {

	logger = logutils.NewFileLogger("Info", "./log/", "ssq.log")

	defer logger.Close() // 程序执行完记得关闭

	// test1()
	// testStringToDate()
	// testMap()

	// 总网络上抓取数据并保存
	// "http://kaijiang.zhcw.com/zhcw/html/ssq/list_1.html"  中 list_1 至 list_127 总共127个网页
	// 其中 list_1 是最新一页的双色球数据
	// fetchAllSSQ(127)

	parseSSQ()

	time.Sleep(time.Millisecond * 300)

}

func parseSSQ() {
	// 从文件中把保存好的双色球信息，再读取出来进行分析
	allSsqInfoMap := getSsqInfoFromFile()

	// 创建一个 []int 用于保存 map 的 key
	var keyArr = make([]int, 0, len(allSsqInfoMap))
	for key := range allSsqInfoMap {
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

	calcNumPer(keyArr, allSsqInfoMap)

	// 查询 map 中任意一期彩票数据是否与 历史其它期 开奖结果一样。
	// searchMatchFromMap(keyArr, allSsqInfoMap)

	// 查询 map 中任意一期彩票数据是否与 历史其它期 开奖结果一样。
	// searchMatch("07 11 17 18 24 29 +05", allSsqInfoMap)

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

// 与历史数据想比较， 计算每一期彩票，每个数字出现的频率
func calcNumPer(keyArr []int, allSsqInfoMap map[int]*SSQInfo) {
	for _, k := range keyArr {
		ssqObj := allSsqInfoMap[k]
		logger.Info("############################################################################################\n")
		logger.Info("开奖日期 : %d , 开奖号码 : %s", k, ssqObj.numInfo)
		fmt.Printf("开奖日期 : %d , 开奖号码 : %s\n", k, ssqObj.numInfo)
		if k == 20100103 { // 只提取2010年1月3日至今 之间的开奖号码 与 历史数据作对比
			logger.Info("统计至 %d", k)
			//fmt.Printf("统计至 %d\n", k)
			break
		}

		ssq := ssqObj.ssq // 拿出需要分析的 中奖号码

		// 统计每个号码出现的频率
		var redOneCount, redTwoCount, redThreeCount, redFourCount, redFiveCount, redSixCount, blueBallCount int

		// 统计 1-33 个红号，每个号码出现的频率
		var numCountObj = &NumberCount{}

		// fmt.Printf("key : %d , numInfo : %s , blueBall : %d\n", k, ssqObj.numInfo, ssqObj.ssq.BlueBall)
		for k2, v2 := range allSsqInfoMap {
			if k2 < k { // 开奖号码 只与 历史数据相比： 例如：19年的开奖号码 只与 19年之前的数据进行对比，而不能找 2020年的数据

				history := v2.ssq

				if checkRedNum(ssq.RedOne, history) {
					redOneCount++
				}
				if checkRedNum(ssq.RedTwo, history) {
					redTwoCount++
				}
				if checkRedNum(ssq.RedThree, history) {
					redThreeCount++
				}
				if checkRedNum(ssq.RedFour, history) {
					redFourCount++
				}
				if checkRedNum(ssq.RedFive, history) {
					redFiveCount++
				}
				if checkRedNum(ssq.RedSix, history) {
					redSixCount++
				}

				if ssq.BlueBall == history.BlueBall {
					blueBallCount++
				}

				// 计算 1-33， 1-16 每个数字在历史数据中出现的次数
				calcNumberCount(history, numCountObj)

			}
		}

		logger.Info("红号：%d (%d次), %d (%d次), %d (%d次), %d (%d次), %d (%d次), %d (%d次)  +蓝号：%d (%d次)",
			ssq.RedOne, redOneCount, ssq.RedTwo, redTwoCount, ssq.RedThree, redThreeCount, ssq.RedFour,
			redFourCount, ssq.RedFive, redFiveCount, ssq.RedSix, redSixCount, ssq.BlueBall, blueBallCount)
		logger.Info("\n-----------------------------------数据分析--------------------------------------\n")

		logger.Info("##红号区：")
		logger.Info("1(%d), 2(%d), 3(%d), 4(%d), 5(%d), 6(%d), 7(%d), 8(%d), 9(%d),\n 10(%d), 11(%d), 12(%d), 13(%d), 14(%d), 15(%d), 16(%d), 17(%d), 18(%d),\n 19(%d), 20(%d), 21(%d), 22(%d), 23(%d), 24(%d), 25(%d), 26(%d), 27(%d),\n 28(%d), 29(%d), 30(%d), 31(%d), 32(%d), 33(%d)", numCountObj.OneCount, numCountObj.TwoCount, numCountObj.ThreeCount, numCountObj.FourCount, numCountObj.FiveCount, numCountObj.SixCount, numCountObj.SevenCount, numCountObj.EightCount, numCountObj.NineCount, numCountObj.TenCount, numCountObj.ElevenCount, numCountObj.TwelveCount, numCountObj.ThirteenCount, numCountObj.FourteenCount,
			numCountObj.FifteenCount, numCountObj.SixteenCount, numCountObj.SeventeenCount, numCountObj.EighteenCount, numCountObj.NineteenCount, numCountObj.TwentyCount, numCountObj.TwentyOneCount,
			numCountObj.TwentyTwoCount, numCountObj.TwentyThreeCount, numCountObj.TwentyFourCount, numCountObj.TwentyFiveCount, numCountObj.TwentySixCount, numCountObj.TwentySevenCount, numCountObj.TwentyEightCount, numCountObj.TwentyNineCount, numCountObj.ThirtyCount, numCountObj.ThirtyOneCount, numCountObj.ThirtyTwoCount, numCountObj.ThirtyThreeCount)

		logger.Info("##蓝号区：")
		logger.Info("1(%d), 2(%d), 3(%d), 4(%d), 5(%d), 6(%d), 7(%d), 8(%d), 9(%d),\n 10(%d), 11(%d), 12(%d), 13(%d), 14(%d), 15(%d), 16(%d)\n", numCountObj.BlueOneCount, numCountObj.BlueTwoCount, numCountObj.BlueThreeCount, numCountObj.BlueFourCount, numCountObj.BlueFiveCount, numCountObj.BlueSixCount, numCountObj.BlueSevenCount, numCountObj.BlueEightCount, numCountObj.BlueNineCount, numCountObj.BlueTenCount, numCountObj.BlueElevenCount, numCountObj.BlueTwelveCount, numCountObj.BlueThirteenCount, numCountObj.BlueFourteenCount, numCountObj.BlueFifteenCount, numCountObj.BlueSixteenCount)

		/*
			fmt.Printf("%d (%d次), %d (%d次), %d (%d次), %d (%d次), %d (%d次), %d (%d次), %d (%d次)\n",
				ssq.RedOne, redOneCount, ssq.RedTwo, redTwoCount, ssq.RedThree, redThreeCount, ssq.RedFour,
				redFourCount, ssq.RedFive, redFiveCount, ssq.RedSix, redSixCount, ssq.BlueBall, blueBallCount)
		*/
	}
}

// 计算 1-33， 1-16 每个数字在历史数据中出现的次数
func calcNumberCount(history *SSQ, numCount *NumberCount) {
	// 红号区：
	for i := 1; i < 34; i++ {
		if checkRedNum(i, history) {
			switch i {
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
			case 10:
				numCount.TenCount++
			case 11:
				numCount.ElevenCount++
			case 12:
				numCount.TwelveCount++
			case 13:
				numCount.ThirteenCount++
			case 14:
				numCount.FourteenCount++
			case 15:
				numCount.FifteenCount++
			case 16:
				numCount.SixteenCount++
			case 17:
				numCount.SeventeenCount++
			case 18:
				numCount.EighteenCount++
			case 19:
				numCount.NineteenCount++
			case 20:
				numCount.TwentyCount++
			case 21:
				numCount.TwentyOneCount++
			case 22:
				numCount.TwentyTwoCount++
			case 23:
				numCount.TwentyThreeCount++
			case 24:
				numCount.TwentyFourCount++
			case 25:
				numCount.TwentyFiveCount++
			case 26:
				numCount.TwentySixCount++
			case 27:
				numCount.TwentySevenCount++
			case 28:
				numCount.TwentyEightCount++
			case 29:
				numCount.TwentyNineCount++
			case 30:
				numCount.ThirtyCount++
			case 31:
				numCount.ThirtyOneCount++
			case 32:
				numCount.ThirtyTwoCount++
			case 33:
				numCount.ThirtyThreeCount++
			}
		}
	}

	// 蓝号区：
	for j := 1; j < 17; j++ {
		if j == history.BlueBall {
			switch j {
			case 1:
				numCount.BlueOneCount++
			case 2:
				numCount.BlueTwoCount++
			case 3:
				numCount.BlueThreeCount++
			case 4:
				numCount.BlueFourCount++
			case 5:
				numCount.BlueFiveCount++
			case 6:
				numCount.BlueSixCount++
			case 7:
				numCount.BlueSevenCount++
			case 8:
				numCount.BlueEightCount++
			case 9:
				numCount.BlueNineCount++
			case 10:
				numCount.BlueTenCount++
			case 11:
				numCount.BlueElevenCount++
			case 12:
				numCount.BlueTwelveCount++
			case 13:
				numCount.BlueThirteenCount++
			case 14:
				numCount.BlueFourteenCount++
			case 15:
				numCount.BlueFifteenCount++
			case 16:
				numCount.BlueSixteenCount++
			}
		}
	}
}

// 检查一注彩票里的 每个红号 是否出现在历史数据中(与历史数据中每一注里面的每个红号进行对比)，如果有 则返回 true
func checkRedNum(red int, history *SSQ) bool {

	switch red {
	case history.RedOne, history.RedTwo, history.RedThree, history.RedFour, history.RedFive, history.RedSix:
		return true
	default:
		return false
	}
}

//  查询 map 中任意一期彩票数据是否与 历史其它期 开奖结果一样。
func searchMatchFromMap(keyArr []int, allSsqInfoMap map[int]*SSQInfo) {
	for _, k := range keyArr {
		ssqObj := allSsqInfoMap[k]
		logger.Info("check : %d", k)
		fmt.Printf("check : %d\n", k)
		// fmt.Printf("key : %d , numInfo : %s , blueBall : %d\n", k, ssqObj.numInfo, ssqObj.ssq.BlueBall)
		for k2, v2 := range allSsqInfoMap {
			if k2 != k && ssqObj.numInfo == v2.numInfo {
				logger.Info("开奖日期：%d == %d , 相同号码：%s", k2, k, ssqObj.numInfo)
				fmt.Printf("开奖日期：%d == %d , 相同号码：%s\n", k2, k, ssqObj.numInfo)
			}
		}
	}
}

// 指定一组彩票数据，与历史数据进行查询
func searchMatch(numInfo string, allSsqInfoMap map[int]*SSQInfo) {

	match := false
	var dateInt int
	for k2, v2 := range allSsqInfoMap {
		if numInfo == v2.numInfo {
			dateInt = k2
			match = true
			break
		}
	}
	if match {
		fmt.Printf("%d 匹配\n", dateInt)
	} else {
		fmt.Println("没找到匹配的！")
	}
}

// 证实 map 是引用传递
func testMap() {
	var m = make(map[int]string, 8)
	m[0] = "123ee"
	m[1] = "ffff"
	m[6] = "aaaa"

	modifyMap(m)

	fmt.Println(m)
}

// 可修改 map 中的值
func modifyMap(m map[int]string) {
	m[6] = "cccc"
}

// 从文件中读取双色球信息
func getSsqInfoFromFile() (allSsqInfoMap map[int]*SSQInfo) {
	file, err := os.Open(File_Path)
	if err != nil {
		HandleError(err, "writeSSQInfoToFile  --> OpenFile")
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	allSsqInfoMap = make(map[int]*SSQInfo, 5000)

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

		ssqInfo := getSsqObjFromLineStr(lineStr) // 读取每一行信息，将其转换为 SSQInfo 结构体

		key, err := dateToInt(ssqInfo.date) // 将开奖日期转换为 int 类型作为 map 的 key
		if err != nil {
			return
		}
		allSsqInfoMap[key] = ssqInfo // 将 SSQInfo 结构体指针作为 value

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
func getSsqObjFromLineStr(lineStr string) *SSQInfo {

	var ssqInfo = &SSQInfo{}

	ssqInfo.date = getDateFromLineStr(lineStr)
	// fmt.Printf("date : %v\n", (*ssqInfo).date)

	ssqInfo.qiStr = getQiStrFromLineStr(lineStr)

	ssqInfo.numInfo = getNumInfoFromLineStr(lineStr)

	ssqInfo.sellInfo = getSellInfoFromLineStr(lineStr)

	ssqInfo.oneNum = getOneNumFromLineStr(lineStr)

	ssqInfo.oneLocation = getOneLocationFromLineStr(lineStr)

	ssqInfo.secondNum = getSecondNumFromLineStr(lineStr)

	// fmt.Printf("ssqInfo : %#v\n", *ssqInfo)

	ssq := getSsqObjFromNumInfo(ssqInfo.numInfo)

	ssqInfo.ssq = ssq

	return ssqInfo
}

// 11 16 17 22 26 32 +04
// 将双色球 字符串形式的信息 转换为 SSQ 结构体中的 []int , int
func getSsqObjFromNumInfo(numInfo string) (ssq *SSQ) {

	var redOne, redTwo, redThree, redFour, redFive, redSix, blueBall int

	// 空格分隔字符串
	arr := strings.Split(numInfo, " ")
	if len(arr) != 7 {
		fmt.Println("----------------len(arr) != 7")
		return
	}

	redOne = stringToInt(arr[0])
	redTwo = stringToInt(arr[1])
	redThree = stringToInt(arr[2])
	redFour = stringToInt(arr[3])
	redFive = stringToInt(arr[4])
	redSix = stringToInt(arr[5])

	str := arr[6]          // 此时的字符串应该是： +04
	blueBallStr := str[1:] // 去除 + 字符
	blueBall = stringToInt(blueBallStr)

	return &SSQ{
		RedOne:   redOne,
		RedTwo:   redTwo,
		RedThree: redThree,
		RedFour:  redFour,
		RedFive:  redFive,
		RedSix:   redSix,
		BlueBall: blueBall,
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

// 从文本中读取的每一行数据中，解析出一等奖中奖数 字符串，返回 string
func getOneNumFromLineStr(lineStr string) string {
	oneNumIndex := strings.Index(lineStr, "^")
	oneNumLastIndex := strings.LastIndex(lineStr, "^")
	oneNum := lineStr[oneNumIndex+1 : oneNumLastIndex]
	return oneNum
}

// 从文本中读取的每一行数据中，解析出一等奖中奖地点 字符串，返回 string
func getOneLocationFromLineStr(lineStr string) string {
	oneLocationIndex := strings.Index(lineStr, "!")
	oneLocationLastIndex := strings.LastIndex(lineStr, "!")
	oneLocation := lineStr[oneLocationIndex+1 : oneLocationLastIndex]
	return oneLocation
}

// 从文本中读取的每一行数据中，解析出二等奖中奖数 字符串，返回 string
func getSecondNumFromLineStr(lineStr string) string {
	secondNumIndex := strings.Index(lineStr, "~")
	secondNumLastIndex := strings.LastIndex(lineStr, "~")
	secondNum := lineStr[secondNumIndex+1 : secondNumLastIndex]
	return secondNum
}

// 从通道中取出双色球信息，然后写入到 一个默认文件中。
func writeSSQInfoToFile() {
	file, err := os.OpenFile(File_Path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0664)
	if err != nil {
		HandleError(err, "writeSSQInfoToFile  --> OpenFile")
		return
	}

	defer file.Close() // 写入完毕后 关闭文件

	for ssqObj := range ssqChan {
		ssqMsg := formatSsqInfo(ssqObj)
		fmt.Fprintln(file, ssqMsg) // 写入
	}

}

// 将 双色球结构体里各字段的信息 拼接成一个指定格式的字符串，然后返回该字符串。
func formatSsqInfo(s *SSQInfo) string {
	date := s.date
	qiStr := s.qiStr
	numInfo := s.numInfo
	sellInfo := s.sellInfo
	oneNum := s.oneNum
	oneLocation := s.oneLocation
	secondNum := s.secondNum

	dateStr := dateToString(date)
	ssqMsg := fmt.Sprintf("[%s] @%s@ *%s* $%s$ ^%s^ !%s! ~%s~\r", dateStr, qiStr, numInfo, sellInfo, oneNum, oneLocation, secondNum)
	// fmt.Println("ssqMsg : ", ssqMsg)
	return ssqMsg
}

// 网络抓取双色球信息，入口程序
func fetchAllSSQ(htmlTotal int) {

	total = htmlTotal
	ssqChan = make(chan *SSQInfo, 5000)

	checkChan = make(chan struct{}, total)

	// 创建128个 goroutine, 每个 goroutine 抓取不同的 网页
	for i := 1; i < total+1; i++ {
		wg.Add(1)
		go getSSQInfo(ssqUrlFormat, i, ssqChan, checkChan)

		time.Sleep(time.Millisecond * 200)
	}

	// 创建一个 检查通道 的goroutine
	wg.Add(1)
	go checkTask()
	wg.Wait()

	fmt.Printf("总共抓取 %d 条数据,  通道总容量: %d\n", len(ssqChan), cap(ssqChan))
	writeSSQInfoToFile()

	fmt.Println("抓取程序结束！")

}

func getSSQInfo(format string, i int, ssqChan chan<- *SSQInfo, checkChan chan struct{}) {
	urlStr := fmt.Sprintf(format, i)
	err := GetContent(urlStr)

	if err != nil {
		HandleError(err, "getSSQInfo --> GetContent")
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

			var ssqInfo = &SSQInfo{}

			se.Find("td").Each(func(i int, sel *goquery.Selection) {
				// fmt.Printf("第 %d 个 <td>\n", i)

				switch i {
				case 0: // 双色球开奖日期
					dateStr := sel.Text()
					ssqInfo.date = stringToDate(dateStr)
					//fmt.Printf("开奖日期：%v\n", ssqInfo.date)
				case 1: // 双色球开奖期数
					qiStr := sel.Text()
					ssqInfo.qiStr = qiStr
					//fmt.Printf("开奖期号：%s\n", ssqInfo.qiStr)
				case 2: // 双色球销售额信息
					ssqNumber := sel.Find("em").Text()
					ssqNumber = formatSsqNumInfo(ssqNumber)
					ssqInfo.numInfo = ssqNumber
					//fmt.Printf("开奖号码：%s\n", ssqInfo.numInfo)
				case 3: // 双色球销售额信息
					sellStr := sel.Text()
					ssqInfo.sellInfo = sellStr
					//fmt.Printf("销售额：%s 元\n", ssqInfo.sellInfo)
				case 4: // 双色球一等奖 中奖注数，以及 中奖地点

					tdInfo := sel.Text()
					tdInfo = strings.TrimSpace(tdInfo)

					str := strings.Split(tdInfo, "\n")

					for index, line := range str {
						line = strings.TrimSpace(line)
						if index == 0 {
							ssqInfo.oneNum = line
						}
						if index == 1 {
							ssqInfo.oneLocation = line
						}
					}

					//fmt.Printf("一等奖中奖数 : %s, 中奖地点 : %s\n", ssqInfo.oneNum, ssqInfo.oneLocation)

				case 5: // 双色球二等奖 中奖注数
					secondInfo := sel.Text()
					ssqInfo.secondNum = secondInfo
					//fmt.Printf("二等奖中奖数：%s\n", ssqInfo.secondNum)
				}
			})

			ssqChan <- ssqInfo // 解析完一条双色球信息后， 放入通道中

		})
	})
	return
}

// checkChan
// 检查127 个goroutine 是否都完成了任务，然后关闭 检查通道和 储存图片链接的通道
func checkTask() {

	var i int
	for {

		<-checkChan // 取不到值，则会在这里一直阻塞
		i++
		if i == total { // 127 个goroutine 完成

			close(checkChan) // checkChan 里面的值取完了之后，代表 127个goroutine 都结束任务了
			close(ssqChan)   // 关闭 储存 图片链接的管道
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
