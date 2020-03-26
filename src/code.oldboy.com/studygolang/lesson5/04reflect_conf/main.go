package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// Config 是一个 日志文件的配置文件
// 注意1： 解析config 文件得到的数据 会赋值给这个结构体里的字段，所以首字母要大写
// 注意2：通过反射设置字段时，如果在进行 string.ParseInt() 转换时，可能会出现类型转换异常，
// 例如： max_size= ,  或者 max_size=20兆,  通过反射查找字段时 是根据 结构体中定义的顺序来找的
// 一旦 MaxSize 这个字段在设置值的时候出现异常，如果不做其他处理，则程序中断，其后面的其他字段 的值也无法设置成功。
type Config struct {
	FilePath string `conf:"file_path"`
	FileName string `conf:"file_name"`
	MaxSize  int64  `conf:"max_size"`
}

// 默认日志文件最大容量为 200M
// 我尝试设置一个日志文件的默认值 (在string.ParseInt() 解析报错时)
// 但是呢，这样程序就不灵活了，万一传入的 结构体中还有其他 int 类型的字段，
// 我人为的设置一个 日志最大存储量，就显的想 模拟考试前已经知道了答案一样。
// const logFileMaxSize = 1024 * 1024 * 200

//  parseConfig() 与 parseConfig2() 对比测试 请看 lesson5/test/main.go
func main() {

	test2()
}

func test2() {
	var c = &Config{}
	err := parseConfig2("logconfig.conf", c)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(c.FilePath)
	fmt.Println(c.FileName)
	fmt.Println(c.MaxSize)
}

// 此方法经过测试
// 优点：比 parseConfig() 循环嵌套的方法 快 4倍。
// 缺点：新申请了 map 保存 key,value 的值，加大了内存的占用
// 不循环嵌套， 解析文件放入到map 中为一个循环， 反射结构体字段与map 比较为一个循环， 两个分开的循环，区别在于： 新申请了 map 空间
func parseConfig2(configName string, result interface{}) (err error) {

	// 0, 前提条件, result 必须是一个 ptr, 才能修改值
	t, v, err := checkObjType(result)
	if err != nil {
		return err
	}

	// 1, 读取config 文件，将解析出来的 key 和 value 存放到一个 configMap
	configMap, err := readConfig(configName)
	if err != nil {
		return
	}

	// 2, 将 configMap 与 结构体中的字段一 一对比，匹配则赋值
	err = setValueToObj(t, v, configMap)

	return
}

// 将解析文件得到的map 与 需要反射的对象进行 key 比较，一旦匹配则赋值
func setValueToObj(t reflect.Type, v reflect.Value, configMap map[string]string) (err error) {
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

// 检查需要反射的对象的类型
func checkObjType(result interface{}) (t reflect.Type, v reflect.Value, err error) {
	t = reflect.TypeOf(result)
	v = reflect.ValueOf(result)

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
	return
}

// 根据文件名读取配置文件信息，解析出 key 和 value 的值放入到 map 中并返回
func readConfig(configName string) (configMap map[string]string, err error) {

	// 1, 打开文件，读取内容
	data, err := ioutil.ReadFile(configName) // 返回值为 []byte
	if err != nil {
		err = fmt.Errorf("打开配置文件%s失败!", configName)
		return
	}

	// 2, 将 ReadFile 返回的字节数组转换为一个 string 类型
	configStr := string(data)
	configStr = strings.TrimSpace(configStr) // 可以去除文本内容的首行和尾行 的多余空格

	// 3. 将读取的文件数据按照行分割，得到一个行的切片

	lineSlice := strings.Split(configStr, "\r\n") // 注意： "\r\n" 是window 下 换行符， 项目部署到服务器上之后需要更改为 unix 换行符 "\n"
	fmt.Println("len(lineSlice) : ", len(lineSlice))
	configMap = make(map[string]string, len(lineSlice)) // 根据解析了多少行来分配 map

	// 一行一行的解析
	for index, line := range lineSlice {

		line = strings.TrimSpace(line)
		// 去除 line 首和尾连续的空格符,  示例： "  file_path    =  /etc/log/   "  ---> TrimSpace()之后:  "file_path    =  /etc/log/"

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

	return

}
