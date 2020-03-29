package logutils

/*  日志包使用说明：
1, LogFileMaxSize  暂定为 5KB，便于测试。使用时注意根据项目不同修改一个 合适的值。
2，使用 logutils.NewFileLogger("Error", "./log/", "test.log")  创建一个日志结构体时，注意里面的路径，必须先在本地创建成功(例如：./log/)
3, 后续优化：
	a, 增加 日志文件按照时间拆分 (每隔一小时拆分一次)的功能
	b, 增加配置文件 (里面配置log 存放路径，文件名，最大容量，模式(终端，还是写文件)， 按照大小还是时间拆分文件 等)
*/

import (
	"strings"
)

// Level 是一个自定义的类型，表示 日志级别
type Level uint16

// 定义日志级别常量
const (
	DebugLevel Level = iota
	InfoLevel
	WarningLevel
	ErrorLevel
	FatalLevel
)

// LogFileMaxSize 定义日志最大存储量  (暂定为 5KB， 便于测试)
const (
	KB             = 1024
	LogFileMaxSize = 1024 * KB
)

// 日志信息太多，太频繁，会导致 通道容量快速爆满，就导致后面的log 信息被丢掉处理
// 目前有两种办法：
/*
	1, 降低写日志的频率
	2, 增加通道的容量
	// 还可以有其他补救方案 (虽然关系不大，但可以有一定的优化)
	3, 增大日志文件的 默认存储大小。  mylogger.go --> LogFileMaxSize
*/

// Logger 定义一个 logger 接口
type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Close()
}

// 根据传进来的 Level 类型的值获取对应的 字符串
// (用于结构体中的 Level类型的字段 转换为字符串 保存在日志信息中)
func getLoggerLevelStr(level Level) string {
	switch level {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarningLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case FatalLevel:
		return "FATAL"
	default:
		return "DEBUG"
	}
}

// 根据用户传入的字符串类型的日志级别，解析出对应的Level
func parseLogLevel(levelStr string) Level {
	levelStr = strings.ToLower(levelStr) // 用使用者传入的 levelStr 转换为小写统一处理
	switch levelStr {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarningLevel
	case "error":
		return ErrorLevel
	case "fatal":
		return FatalLevel
	default:
		return DebugLevel
	}
}
