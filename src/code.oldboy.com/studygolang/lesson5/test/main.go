package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Config 是一个 日志文件的配置文件
// 解析config 文件得到的数据 会赋值给这个结构体里的字段，所以首字母要大写
type Config struct {
	FilePath string `conf:"file_path"`
	FileName string `conf:"file_name"`
	MaxSize  int64  `conf:"max_size"`
}

// 解析 conf文件
func main() {

	now := time.Now()
	test1() // 循环嵌套测试
	fmt.Println(time.Since(now))

	fmt.Println("--------------------------------------测试运行时间---------------------------------")
	fmt.Println("----------------------------------------------------------------------------------")

	now2 := time.Now()
	test2() // 分开两个循环测试
	fmt.Println(time.Since(now2))

	// test1() 与 test2() 运行的时间，  test1() 的运行时间是 test2() 的 4倍左右
	/*
			1, test1() 为循环嵌套： 每解析一行 conf 文件 都要和 结构体中的几个字段进行比较。
			2, test2() 为分开的两个循环： 解析 conf 文件为一个循环，把key,value 存放到 map 中 ,
		 反射结构体字段为一个循环， 通过 map 根据key 找 value 的特性，  在本次循环中找到结构体中的字段并设置值.

	*/
}

func test2() {
	var c = &Config{}
	fmt.Println("解析前：", c)
	err := parseConfig2("logconfig.conf", c)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(c.FilePath)
	fmt.Println(c.FileName)
	fmt.Println(c.MaxSize)
	// fmt.Println("解析后：", c)
}

// 此方法经过测试
// 优点：比 parseConfig() 循环嵌套的方法 快 4倍。
// 缺点：新申请了 map 保存 key,value 的值，加大了内存的占用
// 不循环嵌套， 解析文件放入到map 中为一个循环， 反射结构体字段与map 比较为一个循环， 两个分开的循环，区别在于： 新申请了 map 空间
func parseConfig2(configName string, result interface{}) (err error) {

	t := reflect.TypeOf(result)
	v := reflect.ValueOf(result)

	// 0, 前提条件, result 必须是一个 ptr, 才能修改值
	if t.Kind() != reflect.Ptr {
		err = errors.New("result 必须是一个指针类型")
		return
	}
	// 确认是指针类型，我们可以获取指针所指向的值，方便后面操作反射的方法以及设置值
	t = t.Elem()
	v = v.Elem()
	// 获取指针所指向值的 类型是不是 结构体类型
	if t.Kind() != reflect.Struct {
		err = errors.New("result 必须是一个结构体指针类型")
		return
	}

	// 1, 打开文件，读取内容
	data, err := ioutil.ReadFile(configName) // 返回值为 []byte
	if err != nil {
		err = fmt.Errorf("打开配置文件%s失败!", configName)
		return
	}

	// 2, 将 ReadFile 返回的字节数组转换为一个 string 类型
	configStr := string(data)

	/*
			fmt.Println("configStr : ", configStr)

			示例打印信息如下：

				configStr :     file_path     =     /etc/log/
		 		 file_name   =       test.log.err
				max_size   =      5120

	*/

	// 3. 将读取的文件数据按照行分割，得到一个行的切片
	configStr = strings.TrimSpace(configStr)      // 可以去除文本内容的首行和尾行 的多余空格
	lineSlice := strings.Split(configStr, "\r\n") // 注意： "\r\n" 是window 下 换行符， 项目部署到服务器上之后需要更改为 unix 换行符 "\n"
	fmt.Println("len(lineSlice) : ", len(lineSlice))
	var configMap = make(map[string]string, len(lineSlice)) // 根据解析了多少行来分配 map (实际上，如果有太多的注释也被包含了进去，至于最后一行内容下面多余的空行怎么处理？)

	// 一行一行的解析
	for index, line := range lineSlice {
		// fmt.Printf("%d , str : %s \n", index, line)
		line = strings.TrimSpace(line) // 去除 line 首和尾连续的空格符,  示例： "  file_path    =  /etc/log/   "  ---> TrimSpace()之后:  "file_path    =  /etc/log/"

		/*
			for i, v := range lineSlice {
				fmt.Printf("%d , str : %s \n", i, v)
			}

			示例打印信息如下：

			0 , str :    file_path     =     /etc/log/
			1 , str :  file_name   =       test.log.err
			2 , str : max_size   =      5120

		*/

		// 去除空格之后，line 的长度依然为0，那代表没有任何内容，则排除
		// 再判断如果是以 # 为前缀，则代表注释，也排除
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			// 忽略空行 与 注释行
			// fmt.Println("空格或注释")
			continue
		}

		// 开始解析有内容的行
		equalIndex := strings.Index(line, "=") // 获取 = 的位置
		if equalIndex == -1 {
			err = fmt.Errorf("第 %d 行语法错误", index+1) // 索引加1 代表真实行号
			return
		}
		// 拿到 "=" 的下标， 针对 line 字符串进行切片操作，从而达到 以 "=" 分割的目的，最终拿到 key 和 value
		// 示例： line : "file_path    =  /etc/log/" ,  进行切片操作之后， key : "file_path    ",  value : "  /etc/log/"
		key := line[:equalIndex]
		value := line[equalIndex+1:]

		// 拿到的 key 和 value ， 前后可能还存在空格，所以还需要处理一下
		// 示例：key : "file_path    "    ---> TrimSpace()之后： "file_path"
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		// 去除空格之后，如果长度为0，则代表没有key 值，所以语法错误
		if len(key) == 0 {
			err = fmt.Errorf("第 %d 行语法错误", index+1) // 索引加1 代表真实行号
			return
		}
		// fmt.Printf("key : %s, value : %s\n", key, value)

		// 将解析得到 key 和 value 保存在 configMap 中
		configMap[key] = value
	}

	fmt.Println("configMap==> len : ", len(configMap))
	if len(configMap) == 0 {
		err = errors.New("没有解析到任何有用配置信息")
		return
	}

	// 利用反射 给 result interface{} 赋值
	// 遍历结构体的每一个字段和 configMap 里面的key 进行比较，匹配上了就赋值
	for i := 0; i < t.NumField(); i++ {
		structField := t.Field(i)
		fieldName := structField.Name // 结构体中字段的名字
		fileType := structField.Type
		fileTypeKind := fileType.Kind()

		tagName := structField.Tag.Get("conf") // 获取 结构体中字段对应的 tag 信息 ( conf )

		value, ok := configMap[tagName]

		// fmt.Printf("fieldName : %s , fileType : %s, fileTypeKind : %s, tagName : %s\n", fieldName, fileType, fileTypeKind, tagName)
		if ok {
			// 字段名字匹配了之后，还需要判断结构体中字段的类型，然后根据类型 调用不同的方法
			switch fileTypeKind {
			case reflect.Int64:
				vl, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return err
				}
				fmt.Println("setInt")
				// v.Field(i).SetInt(vl)   // 两种方式都行
				v.FieldByName(fieldName).SetInt(vl)
			case reflect.String:
				fmt.Println("setString")
				// v.Field(i).SetString(confValue)
				v.FieldByName(fieldName).SetString(value)

			}
		}
	}

	return
}

func test1() {
	var c = &Config{}
	fmt.Println("解析前：", c)
	err := parseConfig("logconfig.conf", c)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(c.FilePath)
	fmt.Println(c.FileName)
	fmt.Println(c.MaxSize)
	// fmt.Println("解析后：", c)
}

// 循环嵌套的方式
func parseConfig(configName string, result interface{}) (err error) {

	t := reflect.TypeOf(result)
	v := reflect.ValueOf(result)

	// 0, 前提条件, result 必须是一个 ptr, 才能修改值
	if t.Kind() != reflect.Ptr {
		err = errors.New("result 必须是一个指针类型")
		return
	}
	// 确认是指针类型，我们可以获取指针所指向的值，方便后面操作反射的方法以及设置值
	t = t.Elem()
	v = v.Elem()
	// 获取指针所指向值的 类型是不是 结构体类型
	if t.Kind() != reflect.Struct {
		err = errors.New("result 必须是一个结构体指针类型")
		return
	}

	// 1, 打开文件，读取内容
	data, err := ioutil.ReadFile(configName) // 返回值为 []byte
	if err != nil {
		err = fmt.Errorf("打开配置文件%s失败!", configName)
		return
	}

	// 2, 将 ReadFile 返回的字节数组转换为一个 string 类型
	configStr := string(data)

	/*
			fmt.Println("configStr : ", configStr)

			示例打印信息如下：

				configStr :     file_path     =     /etc/log/
		 		 file_name   =       test.log.err
				max_size   =      5120

	*/

	// 3. 将读取的文件数据按照行分割，得到一个行的切片
	configStr = strings.TrimSpace(configStr)
	lineSlice := strings.Split(configStr, "\r\n") // 注意： "\r\n" 是window 下 换行符， 项目部署到服务器上之后需要更改为 unix 换行符 "\n"
	fmt.Println("len(lineSlice) : ", len(lineSlice))
	// 一行一行的解析
	for index, line := range lineSlice {
		// fmt.Printf("%d , str : %s \n", index, line)
		line = strings.TrimSpace(line) // 去除 line 首和尾连续的空格符,  示例： "  file_path    =  /etc/log/   "  ---> TrimSpace()之后:  "file_path    =  /etc/log/"

		/*
			for i, v := range lineSlice {
				fmt.Printf("%d , str : %s \n", i, v)
			}

			示例打印信息如下：

			0 , str :    file_path     =     /etc/log/
			1 , str :  file_name   =       test.log.err
			2 , str : max_size   =      5120

		*/

		// 去除空格之后，line 的长度依然为0，那代表没有任何内容，则排除
		// 再判断如果是以 # 为前缀，则代表注释，也排除
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			// 忽略空行 与 注释行
			// fmt.Println("空格或注释")
			continue
		}

		// 开始解析有内容的行
		equalIndex := strings.Index(line, "=") // 获取 = 的位置
		if equalIndex == -1 {
			err = fmt.Errorf("第 %d 行语法错误", index+1) // 索引加1 代表真实行号
			return
		}
		// 拿到 "=" 的下标， 针对 line 字符串进行切片操作，从而达到 以 "=" 分割的目的，最终拿到 key 和 value
		// 示例： line : "file_path    =  /etc/log/" ,  进行切片操作之后， key : "file_path    ",  value : "  /etc/log/"
		key := line[:equalIndex]
		value := line[equalIndex+1:]

		// 拿到的 key 和 value ， 前后可能还存在空格，所以还需要处理一下
		// 示例：key : "file_path    "    ---> TrimSpace()之后： "file_path"
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		// 去除空格之后，如果长度为0，则代表没有key 值，所以语法错误
		if len(key) == 0 {
			err = fmt.Errorf("第 %d 行语法错误", index+1) // 索引加1 代表真实行号
			return
		}
		// fmt.Printf("key : %s, value : %s\n", key, value)

		// 利用反射 给 result interface{} 赋值
		// 遍历结构体的每一个字段和 key 进行比较，匹配上了就赋值
		for i := 0; i < t.NumField(); i++ {
			structField := t.Field(i)
			fieldName := structField.Name // 结构体中字段的名字
			fileType := structField.Type
			fileTypeKind := fileType.Kind()

			tagName := structField.Tag.Get("conf") // 获取 结构体中字段对应的 tag 信息 ( conf )

			// fmt.Printf("fieldName : %s , fileType : %s, fileTypeKind : %s, tagName : %s\n", fieldName, fileType, fileTypeKind, tagName)

			// 如果配置文件里面的 key 与 结构体中字段的 tag 信息匹配，则可以进行赋值操作
			if key == tagName {
				// 字段名字匹配了之后，还需要判断结构体中字段的类型，然后根据类型 调用不同的方法
				switch fileTypeKind {
				case reflect.Int64:
					vl, err := strconv.ParseInt(value, 10, 64)
					if err != nil {
						return err
					}
					fmt.Println("setInt")
					// v.Field(i).SetInt(vl)   // 两种方式都行
					v.FieldByName(fieldName).SetInt(vl)
				case reflect.String:
					fmt.Println("setString")
					// v.Field(i).SetString(confValue)
					v.FieldByName(fieldName).SetString(value)

				}
			}
		}

	}

	return
}

/*
func parseConfig(configName string, x interface{}) error {

	f, err := os.Open(configName)
	if err != nil {

		// return errors.New("文件打开失败")  // 创建一个新的错误
		return fmt.Errorf("打开配置文件 %s 失败", configName)
	}

	defer f.Close() // 注意关闭

	bufReader := bufio.NewReader(f)
	var configMap = make(map[string]string, 8)
	for {
		str, err := bufReader.ReadString('\n') // 这里有漏洞，如果配置文件里的最后一行结尾 没有换行符的话，它就读不出来。
		if err == io.EOF {                     // 注意判断文件是否读完
			fmt.Println("文件已读完 : str : ", str) // 没有换行符的最后一行的内容，会在 io.EOF 之后，获取。
			break
			// return err
		}
		if err != nil { // err != nil 可能会等于 io.EOF
			fmt.Println("err == nil")
			return err
		}
		fmt.Printf("每一行数据：%s , 字符长度：%d\n", str, len(str))

		splitArr := strings.Split(str, "=")
		if len(splitArr) == 2 {
			splitArr[0] = strings.TrimSpace(splitArr[0])
			splitArr[1] = strings.TrimSpace(splitArr[1])
			fmt.Println("splitArr[0] , len ", splitArr[0], len(splitArr[0]))
			fmt.Println("splitArr[1] , len ", splitArr[1], len(splitArr[1]))

			configMap[splitArr[0]] = splitArr[1] // 将成功解析的每一行数据放入到 map 中
		}

		// str = strings.TrimSpace(str)
		// 针对还没解析的字符串 进行去除空格的操作，它只能去除前后空格，而不能去除字符串中间的空格，
		// 所以我们需要先把字符串分割之后，再针对分割出来的字符串做前后空格消除处理
		// fmt.Printf("每一行trim()数据：%s , 字符长度：%d\n", str, len(str))
	}

	if len(configMap) != 0 {
		t := reflect.TypeOf(x)
		v := reflect.ValueOf(x)
		if t != nil {

			if t.Kind() == reflect.Ptr {
				t = t.Elem()
				v = v.Elem()
				fmt.Println("..........Elem()........")
			}

			if t.Kind() == reflect.Struct {
				fmt.Println("..........Struct........")
				fmt.Println("filed 数量：", t.NumField())
				for i := 0; i < t.NumField(); i++ {
					structField := t.Field(i)
					fieldName := structField.Name // 结构体中字段的名字
					fileType := structField.Type
					fileTypeKind := fileType.Kind()

					tagName := structField.Tag.Get("conf") // 获取 结构体中字段对应的 tag 信息 ( conf )
					confValue, ok := configMap[tagName]    // 通过tagName 去找对应的值

					fmt.Printf("fieldName : %s , fileType : %s, fileType : %s, tagName : %s\n", fieldName, fileType, fileTypeKind, tagName)

					if ok {
						if fileTypeKind == reflect.Int64 {
							vl, err := strconv.ParseInt(confValue, 10, 64)
							if err != nil {
								return err
							}
							fmt.Println("setInt")
							// v.Field(i).SetInt(vl)   // 两种方式都行
							v.FieldByName(fieldName).SetInt(vl)
						}
						if fileTypeKind == reflect.String {
							fmt.Println("setString")
							// v.Field(i).SetString(confValue)
							v.FieldByName(fieldName).SetString(confValue)
						}
					}

				}
			}
		}

	}

	return fmt.Errorf("解析文件成功")
}
*/
