package main

// elasticsearch 学习  参考文档： https://www.liwenzhou.com/posts/Go/go_elasticsearch/
// elasticsearch 下载
// https://www.elastic.co/cn/downloads/elasticsearch
// 注意：elasticsearch 与 kibana 版本要一致

// elasticsearch 启动
// 执行 elasticsearch.bat

// elasticsearch 配置文件
// G:\Share\elasticsearch-7.6.2-windows-x86_64\elasticsearch-7.6.2\config\elasticsearch.yml
/*
	#network.host: 192.168.0.1
	#discovery.seed_hosts: ["host1", "host2"]          // 集群，可以配置多台服务器
 */

// 常用命令：
/*
	查看健康状态
	curl -X GET 127.0.0.1:9200/_cat/health?v

	查询当前es集群中所有的indices
	curl -X GET 127.0.0.1:9200/_cat/indices?v

	创建索引
	curl -X PUT 127.0.0.1:9200/www
		输出：
		{"acknowledged":true,"shards_acknowledged":true,"index":"www"}

	删除索引
	curl -X DELETE 127.0.0.1:9200/www
		输出：
		{"acknowledged":true}

	插入记录
	curl -H "ContentType:application/json" -X POST 127.0.0.1:9200/user/person -d '
	{
		"name": "dsb",
		"age": 9000,
		"married": true
	}'
		输出：
		{
			"_index": "user",
			"_type": "person",
			"_id": "MLcwUWwBvEa8j5UrLZj4",
			"_version": 1,
			"result": "created",
			"_shards": {
				"total": 2,
				"successful": 1,
				"failed": 0
			},
			"_seq_no": 3,
			"_primary_term": 1
		}

		也可以使用PUT方法，但是需要传入id

		curl -H "ContentType:application/json" -X PUT 127.0.0.1:9200/user/person/4 -d '
		{
			"name": "sb",
			"age": 9,
			"married": false
		}'


	检索
	Elasticsearch的检索语法比较特别，使用GET方法携带JSON格式的查询条件。

	全检索：

	curl -X GET 127.0.0.1:9200/user/person/_search
	按条件检索：

	curl -H "ContentType:application/json" -X PUT 127.0.0.1:9200/user/person/4 -d '
	{
		"query":{
			"match": {"name": "sb"}
		}
	}'
	ElasticSearch默认一次最多返回10条结果，可以像下面的示例通过size字段来设置返回结果的数目。

	curl -H "ContentType:application/json" -X PUT 127.0.0.1:9200/user/person/4 -d '
	{
		"query":{
			"match": {"name": "sb"},
			"size": 2
		}
	}'

*/

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

type Person struct {
	Name string
	Age int
	Married bool
}
// 我们使用第三方库https://github.com/olivere/elastic来连接ES并进行操作。
// 注意下载与你的ES相同版本的client，例如我们这里使用的ES是7.2.1的版本，那么我们下载的client也要与之对应为github.com/olivere/elastic/v7。
/*
	启动 ES
    执行 elasticsearch.bat

	配置 kibana
		config\kibana.yml

	elasticsearch.hosts: ["http://localhost:9200"]     // 配置ES IP与端口
	i18n.locale: "en"                                  // 配置 kibana 主页默认的语言


	启动 kibana 用于展示 ES 中的数据，提供图形化可检索的界面功能
	执行 kibana.bat

	访问 kibana
	默认地址：http://localhost:5601

	执行以下代码，将指定数据插入到 es 中的指定表中，
	然后去 kibana 主页点击左下角 "齿轮图标"
	点击  kibana 索引模式 --->  右上角 创建索引模式  ---> 会自动展示 es 中所有的表  ---> 索引模式 框内手动输入表名

 */
func main() {

	// 创建连接 elastic 的客户端
	client, err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		fmt.Println("elastic.NewClient err : ", err)
		return
	}
	fmt.Println("elastic.NewClient success!")

	p1 := Person{
		Name:    "杨",
		Age:     25,
		Married: false,
	}

	p2 := Person{
		Name:    "andy",
		Age:     29,
		Married: true,
	}

	p3 := Person{
		Name:    "刘德华",
		Age:     59,
		Married: true,
	}

	// index 在 elastic 中代表 table(表)
	put1, err := client.Index().Index("user").BodyJson(p1).Do(context.Background())   // Do 代表执行
	if err != nil {
		fmt.Println("put1 err : ", err)
		return
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)


	put2, err := client.Index().Index("user").BodyJson(p2).Do(context.Background())
	if err != nil {
		fmt.Println("put2 err : ", err)
		return
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put2.Id, put2.Index, put2.Type)

	put3, err := client.Index().Index("user").BodyJson(p3).Do(context.Background())
	if err != nil {
		fmt.Println("put3 err : ", err)
		return
	}
	fmt.Printf("Indexed user %s to index %s, type %s\n", put3.Id, put3.Index, put3.Type)

}
