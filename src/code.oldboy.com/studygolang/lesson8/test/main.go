package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func login(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("./login.html") // 创建含有 模板语法的 模板文件
	if err != nil {
		panic(err)
	}

	methodStr := r.Method // 请求类型:  POST 或 GET
	if methodStr == "GET" {
		// 通过 template 包处理完后的  .html 文件，不会将 {{}} 模板语法像  文字一样显式在 页面上，而是进行了一次转化
		t.Execute(w, nil) // 将结果返回，只返回页面不返回 数据对象之类的， 可以传入 nil
	}
	if methodStr == "POST" {
		err := r.ParseForm() // 解析请求的 form 表单
		if err != nil {
			fmt.Println("解析form 表单出错, err : ", err)
			return
		}

		// 从表单中获取 用户的输入信息 1
		username := r.Form.Get("username")
		pwd := r.Form.Get("pwd")

		// 从表单获取 用户的输入信息 2
		/*
			name := r.FormValue("username")
			p := r.FormValue("pwd")
			fmt.Println("name, p : ", name, p)
		*/

		// 创建模板文件，然后将结果输出到 模板上

		if err != nil {
			fmt.Println("模板文件解析失败, err : ", err)
		}

		if username == "yang" && pwd == "129" {

			http.Redirect(w, r, "http://www.baidu.com", 302) // 重定向网址，以及状态码
			// t.Execute(w, "ok")

		} else {
			t.Execute(w, "failed") // 将结果输出，最终在页面上拿到返回的对象值
		}
	}

}

func main() {

	http.HandleFunc("/login", login)

	err := http.ListenAndServe("127.0.0.1:9090", nil)
	if err != nil {
		fmt.Println("服务器启动失败，err : ", err)
		return
	}

}
