package logutils

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

// 定义日志最大存储量  (暂定为 5KB， 便于测试)
const (
	KB             = 1024
	LogFileMaxSize = 5 * KB
)

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
