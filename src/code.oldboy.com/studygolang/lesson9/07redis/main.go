package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

var redisDb *redis.Client

// redis -- 初步学习
func main() {

	err := initClient()
	if err != nil {
		fmt.Println("connect redis failed, err : ", err)
		return
	}

	value := get("name")
	fmt.Println("value : ", value)

	// 设置中文， windows 下命令行查询 中文会显示乱码，我们需要设置一下 命令行的编码 (GBK--> UTF-8)
	set("name", "召唤神龙 #都发大水氜与")

}

func set(key string, value string) {
	// Last argument is expiration. Zero means the key has no
	// expiration time.  （没有到期时间）
	/*
		ret := redisDb.Set(key, value, 0).Val() // .Val() 可以查看操作完成后的结果
		fmt.Println("ret : ", ret)              // ret : OK
	*/

	// 严谨一些的写法：
	err := redisDb.Set(key, value, 0).Err()
	if err != nil {
		fmt.Println("set value failed, err : ", err)
		return
	}
	fmt.Println("set success!")

}

func get(key string) string {
	name := redisDb.Get(key).Val()
	return name
}

func initClient() (err error) {
	redisDb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // user default DB
	})

	// 连接 redis 服务器
	_, err = redisDb.Ping().Result()

	if err != nil {
		return err
	}
	return nil
}
