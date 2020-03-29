package logutils

import (
	"fmt"
	"os"
	"time"
)

// ConsoleLogger 是一个终端日志结构体
type ConsoleLogger struct {
	level Level
}

// NewConsoleLogger 终端日志结构体的构造函数
func NewConsoleLogger(levelStr string) *ConsoleLogger {
	logLevel := parseLogLevel(levelStr)
	cl := &ConsoleLogger{
		level: logLevel,
	}
	return cl
}

// 将公用的记录日志的功能 封装成一个单独的方法
func (c *ConsoleLogger) log(level Level, format string, args ...interface{}) {
	if c.level > level {
		return
	}

	msg := fmt.Sprintf(format, args...) // 得到用户要记录的日志信息
	// 日志格式：[时间][文件][函数名()][行号][日志级别] 日志信息
	timeStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, funcName, line := getCallerInfo(3)
	logLevelStr := getLoggerLevelStr(level)
	logMsg := fmt.Sprintf("[%s][%s]===>[%s()][%d][%s] %s",
		timeStr, fileName, funcName, line, logLevelStr, msg)
	fmt.Fprintln(os.Stdout, logMsg) // 利用 fmt 包将日志信息输出在终端上.
}

// Debug 打印 DebugLevel 级别的日志
func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	c.log(DebugLevel, format, args...)
}

// Info 打印 InfoLevel 级别的日志
func (c *ConsoleLogger) Info(format string, args ...interface{}) {
	c.log(InfoLevel, format, args...)
}

// Warn 打印 WarnLevel 级别的日志
func (c *ConsoleLogger) Warn(format string, args ...interface{}) {
	c.log(WarningLevel, format, args...)
}

// Error 打印 ErrorLevel 级别的日志
func (c *ConsoleLogger) Error(format string, args ...interface{}) {
	c.log(ErrorLevel, format, args...)
}

// Fatal 打印 FatalLevel 级别的日志
func (c *ConsoleLogger) Fatal(format string, args ...interface{}) {
	c.log(FatalLevel, format, args...)
}

// Close 无须操作 (操作系统的标准输出不需要关闭)
func (c *ConsoleLogger) Close() {

	fmt.Println("-------------ConsoleLogger Close()")
}
