package main

// 导入日志包
import (
	"os"
	"time"

	"code.oldboy.com/studygolang/lesson6/testLogbyChannel/logutils"
)

// 定义一个 Logger 类型的接口
// ConsoleLogger 和 FileLogger 都实现了Logger 接口里的方法
// logger 这个变量就可以接受 ConsoleLogger 和 FileLogger 的实例
var logger logutils.Logger

// 项目入口
func main() {

	// testConsoleLog()   // 测试终端打印日志的方法
	testFileLog() // 测试用日志文件 记录日志的方法

	// testLogger() // logger 接受 ConsoleLogger 和 FileLogger的实例
}

func testLogger() {
	// ConsoleLogger 实现了 Logger 里定义的所有方法
	logger = logutils.NewConsoleLogger("debug")
	// FileLogger 实现了 Logger 里定义的所有方法
	logger = logutils.NewFileLogger("error", "./", "test.log")
}

func testFileLog() {

	// 指定的文件名如果不存在可以创建，但是 filePath 一定要存在
	fileLogger := logutils.NewFileLogger("Error", "./log/", "test.log")

	defer fileLogger.Close() // 程序执行完记得关闭

	var i int64 = 0
	for i = 0; i < 20000; i++ {
		// i++
		fileLogger.Info("===这是一条 %d Info 信息", i)
		fileLogger.Error("---这是一条Error信息")
		fileLogger.Fatal("====这是一条 %d Fatal信息", i)
		time.Sleep(time.Millisecond * 10)
		// 日志信息太多，太频繁，会导致 通道容量快速爆满，就导致后面的log 信息被丢掉处理
		// 目前有两种办法：
		/*
			1, 降低写日志的频率
			2, 增加通道的容量
			// 还可以有其他补救方案 (虽然关系不大，但可以有一定的优化)
			3, 增大日志文件的 默认存储大小。  mylogger.go --> LogFileMaxSize
		*/
	}

	/*  设置一个定时器，一直向日志里面写入数据，便于测试
	ticker := time.Tick(time.Millisecond * 500) // 设置间隔时间为：半秒钟
	for i := range ticker {

		fileLogger.Debug("这是一条DEBUG信息 : %d", i.Nanosecond)
		fileLogger.Info("这是一条Info信息 : %d", i.Nanosecond)
		fileLogger.Warn("这是一条Warn信息 : %d", i.Nanosecond)
		fileLogger.Error("这是一条Error信息 : %d", i.Nanosecond)
		fileLogger.Fatal("这是一条Fatal信息 : %d", i.Nanosecond)
	}
	*/
}

func testConsoleLog() {
	var name string = "yang"
	consoleLogger := logutils.NewConsoleLogger("debug")

	defer consoleLogger.Close() // 程序执行完记得关闭

	consoleLogger.Debug("这是一条debug信息")
	consoleLogger.Debug("这是 %s 指定打印的一条debug信息", name)

	consoleLogger = logutils.NewConsoleLogger("Error")
	consoleLogger.Info("这是一条info 信息")
	consoleLogger.Error("这是一条Error 信息")
	consoleLogger.Fatal("这是一条Fatal 信息")
}

func testFileLogxxx() {

	fileLogger := logutils.NewFileLogger("debug", ".", "test.log")

	var isMonkey = true

	// fileLogger.TestCopyFile()
	// Debug 方法里  参数2 是可变参数，所以可以不传，也可以传多个
	// fileLogger.Debug("这是一条测试的日志。")

	if isMonkey {
		// 循环写入日志信息 (测试用。。。)

		netStr := "https://www.golang.org"
		ticker := time.Tick(time.Millisecond * 500)

		for i := range ticker {
			fileLogger.Debug("id是 %d 的用户一直在尝试登陆服务器: %s", i, netStr) // 往文件里面写入日志信息
		}
	} else {

		// 单词写入日志信息 (测试用。。。)
		userID := 10
		netStr := "https://www.golang.org"
		fileLogger.Debug("id是 %d 的用户一直在尝试登陆服务器: %s", userID, netStr) // 往文件里面写入日志信息
	}

	// fileLogger.Debug("id是 %d 的用户一直在尝试登陆服务器: %s", mylog.STDOUT, userID, netStr)  // 往终端上打印日志信息

	// fileLogger.Info("这是一条Info级别的日志")

	/*
		errorFileLogger := mylog.NewFileLogger(mylog.ERROR, "./", "error.log")
		errorFileLogger.Error("出现了一个错误", mylog.LOGFILE)
		errorFileLogger.Error("出现了一个错误", mylog.STDOUT)

		fatalFileLogger := mylog.NewFileLogger(mylog.FATAL, "./", "fatal.log")
		fatalFileLogger.FataL("超级严重，崩溃了", mylog.LOGFILE)
		fatalFileLogger.FataL("超级严重，崩溃了", mylog.STDOUT)
		fatalFileLogger.Error("error........", mylog.LOGFILE) // FATAL 级别的 logger，调用 Error() ，不会记录log信息
	*/

}

// testRemoveFile("xx.txt")  , 当前文件(main.go) 所在目录下的 xx.txt
func testRemoveFile(name string) {
	os.Remove(name)
}
