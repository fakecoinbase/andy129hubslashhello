package main

// 文章参考：https://www.liwenzhou.com/posts/Go/go_dependency/  
// https://zhuanlan.zhihu.com/p/59687626

/*
	go module是Go1.11版本之后官方推出的版本管理工具，并且从Go1.13版本开始，go module将是Go语言默认的依赖管理工具。

GO111MODULE
要启用go module支持首先要设置环境变量GO111MODULE，通过它可以开启或关闭模块支持，它有三个可选值：off、on、auto，默认值是auto。

GO111MODULE=off禁用模块支持，编译时会从GOPATH和vendor文件夹中查找包。
GO111MODULE=on启用模块支持，编译时会忽略GOPATH和vendor文件夹，只根据 go.mod下载依赖。
GO111MODULE=auto，当项目在$GOPATH/src外且项目根目录有go.mod文件时，开启模块支持。
简单来说，设置GO111MODULE=on之后就可以使用go module了，以后就没有必要在GOPATH中创建项目了，并且还能够很好的管理项目依赖的第三方包信息。

使用 go module 管理依赖后会在项目根目录下生成两个文件go.mod和go.sum。

GOPROXY
Go1.11之后设置GOPROXY命令为：

export GOPROXY=https://goproxy.cn
Go1.13之后GOPROXY默认值为https://proxy.golang.org，在国内是无法访问的，所以十分建议大家设置GOPROXY，这里我推荐使用goproxy.cn。

go env -w GOPROXY=https://goproxy.cn,direct
go mod命令
常用的go mod命令如下：

go mod download    下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
go mod edit        编辑go.mod文件
go mod graph       打印模块依赖图
go mod init        初始化当前文件夹, 创建go.mod文件
go mod tidy        增加缺少的module，删除无用的module
go mod vendor      将依赖复制到vendor下
go mod verify      校验依赖
go mod why         解释为什么需要依赖
go.mod
go.mod文件记录了项目所有的依赖信息，其结构大致如下：

module github.com/Q1mi/studygo/blogger

go 1.12

require (
	github.com/DeanThompson/ginpprof v0.0.0-20190408063150-3be636683586
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/satori/go.uuid v1.2.0
	google.golang.org/appengine v1.6.1 // indirect
)
其中，

module用来定义包名
require用来定义依赖包及版本
indirect表示间接引用

*/

func main() {
	
	// go mod 具体使用

	/*	可以在 $GOPATH/src 以外的目录创建 一个新的项目
		1, 创建一个 .go 文件，里面写入正常代码，以及需要导入的第三方包
		2, 设置 set GO111MODULE=on 
		3, go mod init main.go (或者其它.go 文件)
		4, 执行完第  3 条命令， 会在同级目录下 生成一个 go.mod 文件 (文件格式， 类似与下：)

				module main.go

				go 1.13

				require (
					github.com/go-sql-driver/mysql v1.5.0
					github.com/jmoiron/sqlx v1.2.0
				)
		
		   并且还会生成一个 go.sum 文件 (文件格式，类似与下：)

		    github.com/go-sql-driver/mysql v1.4.0/go.mod h1:zAC/RDZ24gD3HViQzih4MyKcchzm+sOG5ZlKdlhCg5w=
			github.com/go-sql-driver/mysql v1.5.0 h1:ozyZYNQW3x3HtqT1jira07DN2PArx2v7/mN66gGcHOs=
			github.com/go-sql-driver/mysql v1.5.0/go.mod h1:DCzpHaOWr8IXmIStZouvnhqoel9Qv2LBy8hT2VhHyBg=
			github.com/jmoiron/sqlx v1.2.0 h1:41Ip0zITnmWNR/vHV+S4m+VoUivnWY5E4OJfLZjCJMA=
			github.com/jmoiron/sqlx v1.2.0/go.mod h1:1FEQNm3xlJgrMD+FBdI9+xvCksHtbpVBBw5dYhBSsks=
			github.com/lib/pq v1.0.0/go.mod h1:5WUZQaWbwv1U+lTReE5YruASi9Al49XbQIvNi/34Woo=
			github.com/mattn/go-sqlite3 v1.9.0/go.mod h1:FPy6KqzDD04eiIsT53CuJW3U88zkxoIYsOqkbpncsNc=

		5, 执行 go mod download 命令下载 依赖包( gomod 不会在 $GOPATH/src 目录下保存 第三方包的源码，而是包源码和链接库保存再 $GOPATH/pkg/mod 目录下) , 下载信息如下：
				
				go: finding github.com/lib/pq v1.0.0
				go: finding github.com/mattn/go-sqlite3 v1.9.0

		6, 下载完毕后，即可正常执行 go build ,  go run 等命令


		注意： go mod 常用命令：

			常用的go mod命令如下：

			go mod download    下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
			go mod edit        编辑go.mod文件
			go mod graph       打印模块依赖图
			go mod init        初始化当前文件夹, 创建go.mod文件
			go mod tidy        增加缺少的module，删除无用的module
			go mod vendor      将依赖复制到vendor下
			go mod verify      校验依赖
			go mod why         解释为什么需要依赖

	*/
}