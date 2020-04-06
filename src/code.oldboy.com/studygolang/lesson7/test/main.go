package main

import (
	"fmt"
	"strconv"
	"time"
)

// 回文判断 （函数实现与单元测试，基准测试）
func main() {

}

// 检查 传入的时间是否在 当前时间半年之内的范围 （功能测试）
func test2() {
	dateInt := 20191201
	fmt.Println(checkDateAsHalfYear(dateInt))
}

// go 后面跟 fmt 打印
func test1() {
	fmt.Println("main start")
	go fmt.Println("go fmt ")
	fmt.Println("main end")

	time.Sleep(time.Second)

}

func checkDateAsHalfYear(dateInt int) bool {
	now := time.Now()
	half := now.Add(-(time.Hour * 24 * 180))

	halfInt, err := dateToInt(half)

	if err != nil {
		fmt.Println(err)
	}
	// 传入的时间 是 当前时间半年之内的时间
	if dateInt >= halfInt {
		return true
	}
	return false
}

func dateToInt(date time.Time) (int, error) {
	dateStr := date.Format("20060102")
	dateInt, err := strconv.Atoi(dateStr)
	if err != nil {
		// HandleError(err, "dateToInt --> ParseInt")
	}
	return dateInt, err
}
