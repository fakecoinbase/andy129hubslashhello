package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

// logrus 使用
/*
	Logrus是Go（golang）的结构化logger，与标准库logger完全API兼容。

	它有以下特点：

	完全兼容标准日志库，
		拥有七种日志级别：Trace, Debug, Info, Warning, Error, Fataland Panic。 

	可扩展的Hook机制，
		允许使用者通过Hook的方式将日志分发到任意地方，如本地文件系统，logstash，elasticsearch或者mq等，或者通过Hook定义日志内容和格式等

	可选的日志输出格式，
		内置了两种日志格式JSONFormater和TextFormatter，还可以自定义日志格式

	Field机制，
		通过Filed机制进行结构化的日志记录

	线程安全
		默认的logger在并发写的时候是被mutex保护的，比如当同时调用hook和写log时mutex就会被请求，
	有另外一种情况，文件是以appending mode打开的， 此时的并发操作就是安全的，可以用logger.SetNoLock()来关闭它。
*/
func main() {
	
	/*
	logrus.WithFields(logrus.Fields{
		"name": "yang", 
		"age": 18,
	}).Warn("这是一条warn级别的日志")
	/*  打印信息：
		time="2020-04-13T14:42:33+08:00" level=warning msg="这是一条warn级别的日志" age=18 name=yang
	*/

	// 直接以字符串的形式打印日志 (默认为终端输出)
	// Logrus有七个日志级别：Trace, Debug, Info, Warning, Error, Fataland Panic。
	/*
	logrus.Info("这是一条info级别的日志")
	logrus.Error("这是一条error级别的日志")
	logrus.Fatal("这是一条fatal级别的日志")
	logrus.Warn("这是一条warn级别的日志")
	*/

	// 设置日志级别
	// logrus.SetLevel(logrus.ErrorLevel)


	// 可以在全局定义一个 log 实例
	var log = logrus.New()
	log.Out = os.Stdout    // 设置输出方式 (os.Stdout 在终端输出)

	/*  还可以设置 输出为 file 

		 可以设置像文件等任意`io.Writer`类型作为日志输出
		 file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		 if err == nil {
		  log.Out = file
		 } else {
		 log.Info("Failed to log to file, using default stderr")

	*/

	// 设置日志存储格式 (json 格式，还是 text 格式)
	// 传入指针类型
	log.SetFormatter(&logrus.JSONFormatter{})

	log.WithFields(logrus.Fields{
		"name":"zhang", 
		"age":18,
	}).Warn("这是一条json格式的warn信息")

	/*  JSON 格式的日志信息：
		{"age":18,"level":"warning","msg":"这是一条json格式的warn信息","name":"zhang","time":"2020-04-13T14:59:45+08:00"}
	*/


	log.Warn("这是一条普通的warn信息")
	// JSON 格式的日志信息:
	// {"level":"warning","msg":"这是一条普通的warn信息","time":"2020-04-13T14:59:45+08:00"}

	// 日志信息记录 函数名  
	// 作者不建议开启此项， 注意：，开启这个模式会增加性能开销。
	log.SetReportCaller(true)

	log.Error("这是一条error信息，记录哪里报错了")
	/*  记录函数名的 json 格式的 error 信息
		{"file":"G:/Goworkspace/src/code.oldboy.com/studygolang/lesson11/01gin_logrus/main.go:85","func":"main.main","level":"error","msg":"这是一条error信息，记录哪里报错了","time":"2020-04-13T15:02:32+08:00"}
	*/


	// logrus 的其它操作：

	// Hook 
	/*
		Hooks
		你可以添加日志级别的钩子（Hook）。例如，向异常跟踪服务发送Error、Fatal和Panic、信息到StatsD或同时将日志发送到多个位置，例如syslog。
	*/

	/*
		默认字段
		通常，将一些字段始终附加到应用程序的全部或部分的日志语句中会很有帮助。例如，你可能希望始终在请求的上下文中记录request_id和user_ip。

		区别于在每一行日志中写上log.WithFields(log.Fields{"request_id": request_id, "user_ip": user_ip})，你可以向下面的示例代码一样创建一个logrus.Entry去传递这些字段。

		requestLogger := log.WithFields(log.Fields{"request_id": request_id, "user_ip": user_ip})
		requestLogger.Info("something happened on that request") # will log request_id and user_ip
		requestLogger.Warn("something not great happened")
	*/

}