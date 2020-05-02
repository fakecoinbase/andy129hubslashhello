package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

func indexHandleFunc(w http.ResponseWriter, r *http.Request) {

	logrus.Infof("=====indexHandleFunc , host : %v\n", r.Host)
	fmt.Printf("=====indexHandleFunc , context : %#v\n", r.Context())

	info := fmt.Sprintf("host : %s , context : %#v", r.Host, r.Context())
	fmt.Fprintf(w, info)

}

//
func main() {

	http.HandleFunc("/", indexHandleFunc)
	logrus.Info("handle / ")

	err := http.ListenAndServe("0.0.0.0:80", nil)
	if err != nil {
		logrus.Error("ListenAndServe err : ", err)
		return
	}
}


/*  部署项目时遇到的问题：



将 golang 中编写的go 代码编译成 linux 下执行的 脚本文件后，部署到 linux 项目上时，需要给这个脚本添加 可执行的权限
chmod +x main


编写 Dockerfile

FROM golang:1.13.6
WORKDIR /go/src/
COPY . .
EXPOSE 80
CMD ["./main"]


以下命令错误：/bin/bash 执行的是 .sh 脚本文件，不能执行 go 的脚本文件 (目前测试结果如此)
记住  CMD ["/bin/bash", "/go/src/main"]


FROM golang:1.13.6   会从官网上拉取 golang 版本的镜像

WORKDIR /go/src/    进入 /go/src/ 目录下 (如果没有则创建，目前测试结果发现 默认存在)
COPY . .            将主机上当前目录下的文件 全部拷贝到 golang 镜像中的 /go/src/目录下
EXPOSE 80           暴露端口为 80
CMD ["./main"]      紧接着 WORKDIR /go/src/ , 在 /go/src/ 目录下执行 main 程序


如果没有 WORKDIR /go/src/ , 发现结果： 它会将本地文件拷贝到 /go 目录下，与 src 同级






docker-compose.yml 文件：

version: '3'

services:
  web:
    container_name: web
    image: go-web
    ports:
    - 8080:80
    networks:
      - basic

networks:
  basic:



问题1：docker-compose.yml 指定容器名的问题( container_name: web )

当指定 容器名字之后，就无法执行 以下扩容命令了

docker-compose up --scale web=3

因为，当启动三个容器，三个容器的名字都是 web ,则会冲突，所以需要扩容时，不能在里面指定名字


问题2：docker-compose.yml 端口映射的问题


当对web容器进行扩容时,必须去掉 Dockerfile 里端口映射的 定义，因为当启动多个 docker时，也会去进行80端口的映射， 但是发现已经有容器 映射了，所以其它docker 就会创建失败

解决办法：

    第一步：保证 Dockerfile 里定义了 EXPOSE 80 ,   端口号要与 go 代码中监听的端口号一致
    第一步：修改 docker-compose.yml , 去除 service 下 ports 定义， 并添加负载均衡器



修改前：

version: '3'

services:
  web:
    image: go-web
    ports:
    - 8080:80
    networks:
      - basic

networks:
  basic:




修改后：


version: '3'

services:
  web:
    image: go-web
    networks:
      - basic

  lb:
    image: dockercloud/haproxy
    ports:
      - 8080:80
    links:
      - web
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    networks:
      - basic

networks:
  basic:


修改完成之后，执行：docker-compose up --scale web=3   （发现它去官网拉取 go-web镜像，而不是用本地生成好的 go-web 镜像，所以我们在这里不能使用 ）

Pulling web (go-web:)...


以上是什么问题呢？ 原来找不到 go-web 镜像 (因为在 docker-compose.yml中定义了 image: go-web ，但是go-web 又不存在，也就是没有手动的创建，那么我们有什么办法可以让它自己创建呢？)

修改后：（去掉 image: go-web, 添加 build: .  意思是指 在当前目录下创建，它会自动找Dockerfile 文件）

version: '3'

services:
  web:
    build: .
    networks:
      - basic

  lb:
    image: dockercloud/haproxy
    ports:
      - 8080:80
    links:
      - web
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    networks:
      - basic

networks:
  basic:
    driver: bridge


执行 docker-compose 命令进行扩容
docker-compose up --scale web=3

查看自动创建的 image (发现它自动创建了一个 web-web的镜像)
docker image ls


REPOSITORY            TAG                 IMAGE ID            CREATED              SIZE
web_web               latest              b0883fb79dcc        About a minute ago   810MB


查看 docker 进程  (会发现已经启动了 三个 web docker(自动创建的容器名：web_web_1, web_web_2,web_web_3 ) 和 一个 负载均衡器的 docker)
docker ps

CONTAINER ID        IMAGE                 COMMAND                  CREATED              STATUS              PORTS                                     NAMES
7c159fb370ff        dockercloud/haproxy   "/sbin/tini -- docke…"   About a minute ago   Up 59 seconds       443/tcp, 1936/tcp, 0.0.0.0:8080->80/tcp   web_lb_1
a17b79897ed5        web_web               "./main"                 About a minute ago   Up 59 seconds       80/tcp                                    web_web_2
741dccd4fb1e        web_web               "./main"                 About a minute ago   Up About a minute   80/tcp                                    web_web_3
7caff2ea6188        web_web               "./main"                 About a minute ago   Up About a minute   80/tcp                                    web_web_1



尝试减少 docker 的数量 (减容)
docker-compose up --scale web=1

关闭 docker (docker-compose down 命令可将所有的 docker 停掉并且删除，至于network, volume可自己手动删除)
docker-compose down

Stopping web_lb_1  ... done
Stopping web_web_2 ... done
Stopping web_web_3 ... done
Stopping web_web_1 ... done
Removing web_lb_1  ... done
Removing web_web_2 ... done
Removing web_web_3 ... done
Removing web_web_1 ... done
Removing network web_basic




docker-compose.yml 添加 volumes (创建可持续化目录，外部目录映射容器内部目录)


services:
  web:
    build: .
    volumes:
    - ./app/go/web/:/go/src/:rw
    networks:
      - basic


当使用 docker-compose 命令时，结果报错，单独运行却没有问题。最后查验资料，发现是 main 程序包含在了 /go/src/ 目录下, 然后这个目录又被映射出去了，所以有问题。

解决办法:

   第一步：修改 Dockerfile   (添加 WORKDIR /go/app/ ， 创建一个 /go/app/ 目录，让main 程序拷贝到 /go/src/ 目录下, /go/app/ 目录暴露给外面)

FROM golang:1.13.6
WORKDIR /go/app/
WORKDIR /go/src/
COPY . .
EXPOSE 80
CMD ["./main"]


   第二步：修改 docker-compose.yml (修改 volumes 如下)

 services:
  web:
    build: .
    volumes:
    - ./app/go/web/:/go/app/:rw
    networks:
      - basic




volumes 命令解析： （以下意思就是 把主机当前目录下的 /app/go/web/ 目录与容器内部 /go/app/目录做映射，就可以做到 内外文件互通，保持同步性）

volumes:
    - ./app/go/web/:/go/app/:rw





测试请求时遇到的问题：

浏览器会发两次请求？？？

（目前在 windows 系统下 chrome, windows edge 浏览器都发现了这个问题, 而虚拟机中的firefox 没有这个问题，项目也是部署在 虚拟机的 docker 里。）



1， 以下哪种http状态下，浏览器会产生两次http请求？（）

a, 400
b, 404
c, 302
d, 304

选c，302，

四种状态码的意思分别为
400：请求报文存在语法错误；
404：请求的资源没有找到；
302：临时重定向；（302临时重定向，会产生两次http请求，软件开发中一般用于跨域请求先请求跨域凭证，再访问跨域后的网络资源。）
304：服务器资源没有改变，意为可直接使用客户端的缓存









*/