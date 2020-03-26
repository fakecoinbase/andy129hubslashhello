package math_pkg

// math_pk 是一个计算的包

// Student 是一个结构体
// Student 首字母大写， 字段 Name, Age 首字母都大写，可让其他包访问
type Student struct {
	Name string
	Age  int
}

// Add 是一个求和的函数，可供外包访问
func Add(x, y int) int {
	return x + y
}
