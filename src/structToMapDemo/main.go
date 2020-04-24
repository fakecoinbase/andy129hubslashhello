package main

import (
	"encoding/json"
	"fmt"
)

type UserInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Profile `json:"profile"`
}

// Profile 配置信息
type Profile struct {
	Hobby string `json:"hobby"`
}

// 嵌套结构体转 map， 其实是转为 map[string]interface{} 嵌套 map[string]interface{} 的
// 还能使用第三方库 structs , 地址： https://github.com/fatih/structs  （但是作者目前设置为 只读）
// 参考文档：https://www.liwenzhou.com/posts/Go/struct2map/
func main() {

	u1 := UserInfo{Name: "q1mi", Age: 18, Profile: Profile{"双色球"}}

	data, err := json.Marshal(u1)
	if err != nil {
		fmt.Println("json marshal err : ", err)
		return
	}

	var m1 = make(map[string]interface{}, 8)

	json.Unmarshal(data, &m1)

	fmt.Printf("%#v\n", m1)
}
