package logutils

import (
	"fmt"
	"os"
	"path"
	"time"
)

// FileLogger 往文件中记录日志的结构体
type FileLogger struct {
	level      Level    // 日志等级 (详见：mylogger.go 里面定义的Level类型)
	filePath   string   // 文件路径
	fileName   string   // 文件名称
	logFile    *os.File // 日志文件的指针类型 (可以记录各种等级的日志信息)
	errLogFile *os.File // 错误日志文件的指针类型 (记录日志等级为 ErrorLevel 和 FatalLevel等级的日志)
	maxSize    int64
}

// NewFileLogger 是一个生成文件日志结构体示例的构造函数
func NewFileLogger(levelStr, logFilePath, logFileName string) *FileLogger {
	logLevel := parseLogLevel(levelStr)

	flObj := &FileLogger{
		level:    logLevel,
		filePath: logFilePath,
		fileName: logFileName,
		maxSize:  LogFileMaxSize, // 设定日志文件的最大存储容量 (mylogger.go 里面定义)
	}
	flObj.initFileLogger() // 根据上面的文件路径和文件名打开日志文件，把文件句柄赋值给结构体对应的字段
	return flObj
}

// 专门用来初始化文件日志的文件句柄
func (f *FileLogger) initFileLogger() {
	// 通过 path 包 将文件路径和文件名称 拼接起来，方便下面打开文件
	logName := path.Join(f.filePath, f.fileName)
	// 打开文件，如果文件不存在，则创建
	fileObj, err := os.OpenFile(logName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件 %s 失败, %v", logName, err))
	}
	f.logFile = fileObj

	// 打开错误日志文件  (例如：test.log  --> test.log.err)
	errLogName := fmt.Sprintf("%s.err", logName)
	errFileObj, err := os.OpenFile(errLogName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件 %s 失败, %v", errLogName, err))
	}
	f.errLogFile = errFileObj
}

// 检查 log 文件容量是否超过上限，超过则进行拆分 (对源文件进行备份，然后再创建一个新的log文件)
func (f *FileLogger) checkSplit(file *os.File) bool {
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	return fileSize >= f.maxSize
}

// 备份超过容量的日志文件，然后返回一个新的日志文件的句柄
func (f *FileLogger) splitLogFile(file *os.File) *os.File {

	fileName := file.Name() // 拿到需要备份文件的 完整路径
	// 将需要备份的文件名后 加上时间戳
	backupName := fmt.Sprintf("%s_%v.back", fileName, time.Now().Unix())
	// 1, 把原来的文件关闭
	file.Close()
	// 2, 备份原来的文件 (将原来的文件重命名)
	os.Rename(fileName, backupName)
	// 3, 新建一个文件
	fileObj, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		panic(fmt.Errorf("打开日志文件失败"))
	}
	return fileObj
}

func (f *FileLogger) log(level Level, format string, args ...interface{}) {
	if f.level > level {
		return
	}
	//
	msg := fmt.Sprintf(format, args...) // 得到用户要记录的日志信息
	// 日志格式：[时间][文件][函数名()][行号][日志级别] 日志信息
	timeStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileName, funcName, line := getCallerInfo(3)
	logLevelStr := getLoggerLevelStr(level)
	logMsg := fmt.Sprintf("[%s][%s]===>[%s()][%d][%s] %s",
		timeStr, fileName, funcName, line, logLevelStr, msg)

	// 向 log 文件里写入日志之前要判断 这个文件是否已经超过了容量
	if f.checkSplit(f.logFile) {
		f.logFile = f.splitLogFile(f.logFile) // 返回一个新的日志文件的句柄
	}

	// 通过 fmt包里的 Fprintln() 方法 将日志信息写入到 指定的 logFile 里
	fmt.Fprintln(f.logFile, logMsg)

	// 如果是 ErrorLevel 或者 FatalLevel 级别的日志还要记录到 f.errLogFile 里
	if level >= ErrorLevel {
		if f.checkSplit(f.errLogFile) {
			f.errLogFile = f.splitLogFile(f.errLogFile)
		}
		fmt.Fprintln(f.errLogFile, logMsg)
	}
}

// Debug 是一个保存 DebugLevel 级别的日志信息
func (f *FileLogger) Debug(format string, args ...interface{}) {
	f.log(DebugLevel, format, args...)
}

// Info 是一个保存 InfoLevel 级别日志信息的方法
func (f *FileLogger) Info(format string, args ...interface{}) {
	f.log(InfoLevel, format, args...)
}

// Warn 是一个保存 WarningLevel 级别日志信息的方法
func (f *FileLogger) Warn(format string, args ...interface{}) {
	f.log(WarningLevel, format, args...)
}

// Error 是一个保存 ErrorLevel 级别日志信息的方法
func (f *FileLogger) Error(format string, args ...interface{}) {
	f.log(ErrorLevel, format, args...)
}

// Fatal 是一个保存 FatalLevel 级别日志信息的方法
func (f *FileLogger) Fatal(format string, args ...interface{}) {
	f.log(FatalLevel, format, args...)
}

// Close 是一个关闭日志文件的方法
func (f *FileLogger) Close() {
	f.logFile.Close()
	f.errLogFile.Close()

	fmt.Println("-------------FileLogger Close()")
}

// 学习一下 file.Stat() 方法的功能
func (f *FileLogger) isFileStorageEnough() bool {
	file := f.logFile
	fmt.Println("file .... ", file)
	fileInfo, _ := file.Stat() // 注意这里，如果 file 已经执行了 close() 方法，file.Stat() 会返回一个 <nil> 的 fileInfo
	fmt.Println("fileInfo .... ", fileInfo)
	fileSize := fileInfo.Size()
	if fileSize > LogFileMaxSize {
		return true
	}

	return false

	// file.Stat() --->  FileInfo ---> 下面各种信息：
	/*
		fmt.Println("FileInfo --> fileName : ", fileInfo.Name())   // "test.log"
		fmt.Println("FileInfo --> fileSize : ", fileInfo.Size())   // "1801"   // 返回的是字节数
		fmt.Println("FileInfo --> fileMode : ", fileInfo.Mode())   // "-rw-rw-rw-"
		fmt.Println("FileInfo --> IsDir : ", fileInfo.IsDir())     // "false"
		fmt.Println("FileInfo --> ModTime : ", fileInfo.ModTime()) // "2020-03-22 22:35:31.4765782 +0800 CST"
		fmt.Println("FileInfo --> Sys() : ", fileInfo.Sys())       // "&{32 {2849075252 30801987} {2849075252 30801987} {598602710 30802007} 0 1801}"
	*/
}
