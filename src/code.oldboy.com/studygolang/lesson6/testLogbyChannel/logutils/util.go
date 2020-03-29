package logutils

import (
	"path"
	"runtime"
)

/*  Go 语言 runtime 包中 Caller 源码
// Caller reports file and line number information about function invocations on
// the calling goroutine's stack. The argument skip is the number of stack frames
// to ascend, with 0 identifying the caller of Caller.  (For historical reasons the
// meaning of skip differs between Caller and Callers.) The return values report the
// program counter, file name, and line number within the file of the corresponding
// call. The boolean ok is false if it was not possible to recover the information.
func Caller(skip int) (pc uintptr, file string, line int, ok bool) {
	rpc := make([]uintptr, 1)
	n := callers(skip+1, rpc[:])
	if n < 1 {
		return
	}
	frame, _ := CallersFrames(rpc).Next()
	return frame.PC, frame.File, frame.Line, frame.PC != 0
}
*/

// Caller(skip int) 详解

/*  假如：runtime.Caller(0) 这行代码写在 a.go 文件里面的 func a() 里面
b.go 文件里面的 func b() 调用了 a.go 的 func a()
c.go 文件里面的 func c() 调用了 b.go 里面的 func b()

那么：
当 skip = 0 时，将返回 a.go 文件名, 可通过 pc 指针类型 获取 a() 方法名，以及 哪一行
当 skip = 1 时，将返回 b.go 文件名, 可通过 pc 指针类型 获取 b() 方法名，以及 哪一行
当 skip = 2 时，将返回 c.go 文件名, 可通过 pc 指针类型 获取 c() 方法名，以及 哪一行

依此类推。。。。。
*/

func getCallerInfo(skip int) (fileName, funcName string, line int) {

	// pc 指针类型， fileName (文件名), line(第几行), ok (返回是否正确)
	pc, fileName, line, ok := runtime.Caller(skip) // Caller(skip int)
	if !ok {
		return
	}

	// 根据 pc 拿到当前执行的函数名
	funcName = runtime.FuncForPC(pc).Name()

	// fmt.Println(fileName) // G:/Goworkspace/src/code.oldboy.com/studygolang/lesson4/testLog/myproject/main.go
	// fmt.Println(funcName) // main.main

	// 只获取文件名
	fileName = path.Base(fileName)
	// 只获取方法名
	funcName = path.Base(funcName)

	// fmt.Println(fileName) // main.go
	// fmt.Println(funcName) // main.main

	return
}
