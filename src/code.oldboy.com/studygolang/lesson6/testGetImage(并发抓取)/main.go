package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

// RegImage 抓取图片的正则表达式
var RegImage = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(git)|(bmp)))`

// DownloadPath 图片下载地址
var DownloadPath = "./img/"

// 存储图片链接 通道
var urlStrChan chan string

// 任务统计通道
var checkChan chan struct{}

var wg sync.WaitGroup

var imageUrlFormat string = "https://www.umei.cc/p/gaoqing/cn/%d.htm"

// 并发抓取网页上的图片
func main() {

	// test1()

	testDownloadImage()

}

func testDownloadImage() {
	urlStrChan = make(chan string, 500)

	checkChan = make(chan struct{}, 26)

	// 创建26个 goroutine, 每个 goroutine 抓取不同的 网页
	for i := 1; i < 27; i++ {
		wg.Add(1)
		go getImageUrl(imageUrlFormat, i, urlStrChan, checkChan)
	}

	// 创建一个 检查通道 的goroutine
	wg.Add(1)
	go checkTask()

	// 创建下载图片的 goroutine  (注意，下载协程， 不易开启过多，容易被封IP)
	for j := 0; j < 5; j++ {
		wg.Add(1)
		go DownloadImage()
	}

	wg.Wait()
	fmt.Println("抓图程序结束！")

	// 练习总结：

	// 以下是 自己写的， 注意对比上面优化的写法，
	// 1, 何时关闭 图片链接通道？
	// 2, 保存文件名，是否严谨？
	// 3, 是否想到了使用  sync.WaitGroup (何时用它)
	// 4, Download 图片时，要开启多少 goroutine
	// 5, 整个设计模式 (a 还是 b, 为什么呢？)
	//     a, 边解析边下载，
	//     b, 还是先解析，解析完了之后，通道关闭，最后开启几个goroutine 从通道里拿链接并下载)

	/*
		for v := range urlStrChan {
			fileName := fmt.Sprintf("%v.jpg", time.Now().Nanosecond())
			go DownloadFile(v, fileName)
		}
	*/

	/*
		for {
			select {
			case v := <-urlStrChan:
				fileName := fmt.Sprintf("%v.jpg", time.Now().Nanosecond())
				go DownloadFile(v, fileName)
			default:
				fmt.Println("无所事事！")
			}
		}
	*/

}

// checkChan
// 检查26 个goroutine 是否都完成了任务，然后关闭 检查通道和 储存图片链接的通道
func checkTask() {

	var i int
	for {

		<-checkChan // 取不到值，则会在这里一直阻塞
		i++
		if i == 26 { // 26 个goroutine 完成

			// close(checkChan)  // checkChan 里面的值取完了之后，代表 26个goroutine 都结束任务了
			close(urlStrChan) // 关闭 储存 图片链接的管道
			break
		}
	}
	wg.Done()
}

func getImageUrl(format string, i int, urlStrChan chan<- string, checkChan chan struct{}) {
	urlStr := fmt.Sprintf(format, i)
	result := GetContent(urlStr, RegImage)

	// checkChan
	defer func() {
		checkChan <- struct{}{} // 获取网页的 图片链接，无论成功失败，都加一个标识代表这个 goroutine 结束了任务
		wg.Done()
	}()

	// 网址访问不成功，返回 nil [][]string
	if result == nil {
		return
	}

	for _, v := range result {
		urlStrChan <- v[0] // 从网页里解析出来的 url 发送到通道中
	}
}

// DownloadImage 下载图片
func DownloadImage() {
	for url := range urlStrChan {
		fileName := GetFileNameFromURL(url)
		ok := DownloadFile(url, fileName)
		if ok {
			//fmt.Printf("%s 下载成功！", fileName)
		} else {
			//fmt.Printf("%s 下载失败！", fileName)
		}
	}
	wg.Done()
}

// GetFileNameFromURL 通过 URL 获取图片名
func GetFileNameFromURL(url string) string {
	lastIndex := strings.LastIndex(url, "/")
	fileName := url[lastIndex+1:]
	fileName = fmt.Sprintf("%v_%s", time.Now().Nanosecond(), fileName)
	return fileName
}

// 示例： 抓取和下载图片
func test1() {
	result := GetContent("https://www.umei.cc/p/gaoqing/cn/26.htm", RegImage)

	for _, v := range result {
		fmt.Println(v[0])
	}

	ok := DownloadFile("http://i1.shaodiyejin.com/uploads/tu/201910/10333/7d52844e60_444.jpg", "1.jpg")
	if ok {
		fmt.Println("图片下载完成！")
	}
}

// GetContent 网络爬虫获取内容
func GetContent(urlStr string, regStr string) [][]string {
	// 根据网址获取 网页信息
	resp, err := http.Get(urlStr)
	HandleError(err, "http.Get()")

	if resp == nil {
		fmt.Printf("%s 访问失败！\n", urlStr)
		return nil
	}

	defer func() {
		if resp != nil {
			//fmt.Printf("resp.Body 类型：%T\n", resp.Body)
			resp.Body.Close()
		}
	}()

	// 读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll()")

	pageStr := string(pageBytes)

	// 过滤网页数据
	reg := regexp.MustCompile(regStr) // 根据指定的正则表达式，返回 正则对象

	result := reg.FindAllStringSubmatch(pageStr, -1) // 返回的是 [][]string , 正则对象 根据网页数据进行匹配查找 ,  -1 代表返回所有， 如果写1 就只返回一个
	// fmt.Println(result)

	// 备用，注意
	// result := reg.FindAllString(pageStr, -1)   // 此返回的是 []string

	if result != nil {
		fmt.Printf("抓取 %d 条数据\n", len(result))
	}

	return result
}

// HandleError 处理错误
func HandleError(err error, why string) {

	if err != nil {
		fmt.Println(why, err)
	}
}

// DownloadFile 下载
func DownloadFile(url string, fileName string) (ok bool) {

	resp, err := http.Get(url)
	HandleError(err, "DownloadFile http.get.url")

	if resp == nil {
		fmt.Printf("%s 下载失败！\n", fileName)
		return false
	}

	defer func() {
		if resp != nil {
			// fmt.Printf("resp.Body 类型：%T\n", resp.Body)
			resp.Body.Close()
		}
	}()

	// 读取页面内容
	bytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "DownloadFile ioutil.ReadAll()")

	fileName = DownloadPath + fileName

	// 写入数据
	ioutil.WriteFile(fileName, bytes, 0666)

	if err != nil {
		return false
	}

	return true

}
