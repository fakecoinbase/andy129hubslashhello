package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	// RegQQEmail 匹配QQ邮箱的 正则表达式
	RegQQEmail = `\d+@qq.com`
	RegSSQ     = `red num2">\d+</td>`
	// 匹配所有类型邮箱  \w 代表 匹配 数字，字母，下划线等 ， ()?代表 可能有可能无
	RegEmail = `\w+@\w+\.\w+`
	// https? ,  ? 代表有或者没有 s  （https 安全模式，通过爬虫获取，貌似不行）
	// \s\S 代表各种字符
	// +? 代表贪婪模式，如果不加贪婪模式，则会一直在网页中查找最后一个引号
	// 被() 圈出来的内容，会另外保存在数组中，实现了 href=  与  后面超链接的分离
	RegLink = `href="(https?://[\s\S]+?)"`

	RegPhone = `1[3456789]\d\s?\d{4}\s?\d{4}`
	// {} 代表有几位，[] 代表可能出现的数字， ()|()  代表或者满足前面或者满足后面
	RegIdCard = `[123456789]\d{5}((19\d{2})|(20[012]\d))((0[1-9])|(1[012]))((0[1-9])|([12]\d)|(3[01]))\d{3}[\dXx]`

	RegImage = `"https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(git)|(bmp)))"`
)

// regexp 正则表达式
//
func main() {

	// test1()
	// test2()

	// test3()
	// test4()

	// test5()

	test6()
}

func test1() {

	result := GetEmail("https://tieba.baidu.com/p/1467281933?red_tag=0074155440", RegQQEmail)

	for _, v := range result {
		fmt.Println(v[0])
	}
}

// GetEmail 根据地址获取网址中的 email
func GetEmail(urlStr string, regStr string) [][]string {

	// 根据网址获取 网页信息
	resp, err := http.Get(urlStr)
	HandleError(err, "http.Get()")

	// 读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll()")

	pageStr := string(pageBytes)

	// 过滤网页数据
	reg := regexp.MustCompile(regStr) // 根据指定的正则表达式，返回 正则对象

	result := reg.FindAllStringSubmatch(pageStr, -1) // 正则对象 根据网页数据进行匹配查找 ,  -1 代表返回所有， 如果写1 就只返回一个
	fmt.Println(result)

	return result
}

// HandleError 处理错误
func HandleError(err error, why string) {

	if err != nil {
		fmt.Println(why, err)
	}
}

// 获取所有邮箱信息
func test2() {
	result := GetEmail("https://tieba.baidu.com/p/1467281933?red_tag=0074155440", RegEmail)

	for _, v := range result {
		fmt.Println(v[0])
	}
}

// 获取网页里面的超链接
func test3() {
	result := GetContent("https://tieba.baidu.com/p/1467281933?red_tag=0074155440", RegLink)

	for _, v := range result {
		fmt.Println(v[1])
	}
}

// 获取网页里的手机号
func test4() {
	result := GetContent("https://www.zhaohaowang.com/", RegPhone)

	for _, v := range result {
		fmt.Println(v)
	}
}

// 获取网页里的身份证号
func test5() {
	result := GetContent("https://henan.qq.com/a/20171107/069413.htm", RegIdCard)

	for _, v := range result {
		fmt.Println(v)
	}
}

// 获取网页里的图片链接
func test6() {
	result := GetContent("https://image.baidu.com/search/index?tn=baiduimage&ct=201326592&lm=-1&cl=2&ie=gb18030&word=%C3%C0%C5%AE%CD%BC%C6%AC&fr=ala&ala=1&alatpl=cover&pos=0&hs=2&xthttps=111111", RegImage)

	for _, v := range result {
		fmt.Println(v[0])
	}
}

// GetContent 网络爬虫获取内容
func GetContent(urlStr string, regStr string) [][]string {
	// 根据网址获取 网页信息
	resp, err := http.Get(urlStr)
	HandleError(err, "http.Get()")

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

	return result
}

// 获取双色球信息
func test8() {
	result := GetContent("https://zx.500.com/ssq/mediayc.php", RegSSQ)

	for _, v := range result {
		fmt.Println(v[0])
	}
}
