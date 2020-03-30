package split

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

// 注意：
/*
1, split_test.go  测试文件，必须以 _test.go 结尾

2, Test函数名一定是以 Test 开头，但是后面不一定要与 需要测试的函数名一样，但是 一般建议一样。
// func TestAa(t *testing.T) {

3, (t *testing.T)   必须传入的参数

4, 多个测试用例，函数名不能相同

5, 测试驱动开发
由go test 单元测试出来的问题，返过去修改Split() 方法， 称之为： 测试驱动开发，由测试结果 去驱动 代码的健壮性
// 修复了问题之后，为了严谨 ，还需要将 其他 测试用例一起测试一下，以免 为了修复这个问题，而带来其他问题

6, 测试组   ---> 详见： split_group_test.go

7, 测试覆盖率  (查看 函数里面的代码是否都被跑到了)

	go test -cover

	// 将覆盖率结果信息 输出到外部文件
	go test -cover -coverprofile=c.out

	// 将结果使用 浏览器打开
	go tool cover -html=c.out

8，基准测试

	// 基准测试的前提是，前面不能有其它 faild 的测试
	// 根据基准测试，我们可以看内存使用情况，申请情况，函数执行时间， 针对性的去优化

9, 性能比较，详见：fib/fib_test.go

10, 重置时间
	b.ResetTimer

11, 并行测试

	RunParallel会创建出多个goroutine，并将b.N分配给这些goroutine执行， 其中goroutine数量的默认值为GOMAXPROCS。
	用户如果想要增加非CPU受限（non-CPU-bound）基准测试的并行性， 那么可以在RunParallel之前调用SetParallelism 。
	RunParallel通常会与-cpu标志一同使用。

12, Setup 与 Tear Down

	测试程序有时需要在测试之前进行额外的设置（setup）或在测试之后进行拆卸（teardown）。

13, 示例函数

*/

// 单元测试
func TestSplit(t *testing.T) {

	t.Log("进来测试") // 使用 go test -v 可以显式测试的详细信息 (例如 这条log 信息)
	got := Split("a:b:c", ":")
	want := []string{"a", "b", "c"}
	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望结果：%v, 实际结果：%v\n", want, got)
	}
	fmt.Println("fmt 输出结束")
	// go test 时，它把 fmt 语句也当成是代码的一部分，
	// 所以不用 go test -v 也能正常显式
	// 注意与 t.Log() 的区别，以及调用顺序

	/*  打印结果如下：  (go test)

	fmt 输出结束
	PASS
	ok      code.oldboy.com/studygolang/lesson7/01gotest/split      0.037s

	*/

	/*  打印结果如下：  (go test -v)

	=== RUN   TestSplit
	fmt 输出结束
	--- PASS: TestSplit (0.00s)
	    split_test.go:14: 进来测试
	PASS
	ok      code.oldboy.com/studygolang/lesson7/01gotest/split      0.042s

	*/
}

// Test 中的T 必须大写
func TestSplit2(t *testing.T) {

	t.Log("进来测试22")
	got := Split("a:b:c", "*")
	want := []string{"a:b:c"}

	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望结果：%v, 实际结果：%v\n", want, got)
	}
	fmt.Println("fmt 输出结束2")
}

func TestNoneSplit(t *testing.T) {
	t.Log("进来测试 None")
	got := Split("a:b:c", "*")
	want := []string{"a:b:c"}

	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望结果：%v, 实际结果：%v\n", want, got)
	}
	fmt.Println("fmt 输出结束3")
}

func TestNoneSplit2(t *testing.T) {
	t.Log("进来测试 None2")
	got := Split("a:b:c", "*")
	want := []string{"a:b:c"}

	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望结果：%v, 实际结果：%v\n", want, got)
	}
	fmt.Println("fmt 输出结束4")
}

/*  常用指令：

// 默认测试所有
go test
go test -v

// 指定函数名 测试
go test -run="TestSplit2"
go test -run="TestSplit2" -v

go test -run TestSplit2
go test -run TestSplit2 -v

// 指定 包含特殊名称的 函数 进行测试,
// 例如：指定 None ---> 则会执行 TestNoneSplit,  TestNoneSplit2
// 例如：指定 Split ---> 则会执行 TestSplit, TestSplit2, TestNoneSplit, TestNoneSplit2
go test -run TestNoneSplit -v
go test -run None -v
go test -run="None" -v

*/

// 当分隔符不是一个 字符时，例如：以 "fa" 为分隔， 分隔字符串 "fsdfafdfa"
func TestMultSepSplit(t *testing.T) {
	t.Log("进来测试 MultSep")
	got := Split("sdfafdfa", "fa")
	want := []string{"sd", "fd"}

	if ok := reflect.DeepEqual(got, want); !ok {
		t.Fatalf("期望结果：%v, 实际结果：%v\n", want, got)
	}
	fmt.Println("fmt 输出结束5")
}

// 发现分隔符的长度大于1时， 之前写的 Split() 方法测试不通过。
// 由go test 单元测试出来的问题，返过去修改Split() 方法， 称之为： 测试驱动开发，由测试结果 去驱动 代码的健壮性
// 修复了问题之后，为了严谨 ，还需要将 其他 测试用例一起测试一下，以免 为了修复这个问题，而带来其他问题

// 基准测试
// 基准测试的前提是，前面不能有其它 faild 的测试
func BenchmarkSplit(b *testing.B) {
	b.Log("这是一个基准测试")
	for i := 0; i < b.N; i++ {
		Split("a:b:c", ":")
	}
}

// 输入指令：go test -bench=Split
//  信息解析：
// BenchmarkSplit-8         5157609               226 ns/op
// -8 , 8个数字代表几个CPU，默认跑满,  5157609 代表你要测试的这个函数，在一秒的时间内，执行了多少次，  226 ns/op  表示每次运行花了 226 纳秒

//  打印信息：
/*
goos: windows
goarch: amd64
pkg: code.oldboy.com/studygolang/lesson7/01gotest/split
BenchmarkSplit-8         5157609               226 ns/op
--- BENCH: BenchmarkSplit-8
    split_test.go:147: 这是一个基准测试
    split_test.go:147: 这是一个基准测试
    split_test.go:147: 这是一个基准测试
    split_test.go:147: 这是一个基准测试
    split_test.go:147: 这是一个基准测试
PASS
ok      code.oldboy.com/studygolang/lesson7/01gotest/split      1.441s

*/

// 输入指令: go test -bench=Split -cpu=1    (指定 cpu 数量)
// 打印信息：
/*
goos: windows
goarch: amd64
pkg: code.oldboy.com/studygolang/lesson7/01gotest/split
BenchmarkSplit   4996182               230 ns/op
--- BENCH: BenchmarkSplit
    split_test.go:148: 这是一个基准测试
    split_test.go:148: 这是一个基准测试
    split_test.go:148: 这是一个基准测试
    split_test.go:148: 这是一个基准测试
    split_test.go:148: 这是一个基准测试
PASS
ok      code.oldboy.com/studygolang/lesson7/01gotest/split      1.431s

*/

// 输入指令：go test -bench=Split -benchmem  (查看更详细的内存信息)
// 信息解析：
// BenchmarkSplit-8         5551309               217 ns/op             112 B/op          3 allocs/op
// 前面三个就不说了，说后面两个代表的意思：    112 B/op,  每次使用了 112 个byte,   3 allocs/op 表示 每次有 3 次申请内存的操作
// 打印信息：
/*
goos: windows
goarch: amd64
pkg: code.oldboy.com/studygolang/lesson7/01gotest/split
BenchmarkSplit-8         5551309               217 ns/op             112 B/op          3 allocs/op
--- BENCH: BenchmarkSplit-8
    split_test.go:148: 这是一个基准测试
    split_test.go:148: 这是一个基准测试
    split_test.go:148: 这是一个基准测试
    split_test.go:148: 这是一个基准测试
    split_test.go:148: 这是一个基准测试
PASS
ok      code.oldboy.com/studygolang/lesson7/01gotest/split      1.466s

*/

// 可以根据内存的打印信息，我们继续优化代码

// 优化后 (对比上面的信息，内存的申请次数，以及占用字节，执行性能 都有很大的提升)
// 输入指令： go test -bench=Split -benchmem
/*
goos: windows
goarch: amd64
pkg: code.oldboy.com/studygolang/lesson7/01gotest/split
BenchmarkSplit-8        10205478               112 ns/op              48 B/op          1 allocs/op
--- BENCH: BenchmarkSplit-8
    split_test.go:148: 这是一个基准测试
    split_test.go:148: 这是一个基准测试
    split_test.go:148: 这是一个基准测试
    split_test.go:148: 这是一个基准测试
    split_test.go:148: 这是一个基准测试
PASS
ok      code.oldboy.com/studygolang/lesson7/01gotest/split      1.308s

*/

// 性能比较 ，详见： fib/fib_test.go

// 重置时间  (去除一些耗时的无关操作，例如：连接数据库，网络请求等，去掉这些耗时的操作，重新计算 功能函数的耗时时间)
// b.ResetTimer之前的处理不会放到执行时间里，也不会输出到报告中，所以可以在之前做一些不计划作为测试报告的操作。例如：

/*
	func BenchmarkSplit(b *testing.B) {
	time.Sleep(5 * time.Second) // 假设需要做一些耗时的无关操作
	b.ResetTimer()              // 重置计时器
	for i := 0; i < b.N; i++ {
		Split("沙河有沙又有河", "沙")
	}
}

*/

// 并行测试  (通常在后面加 Parallel)
func BenchmarkSplitParallel(b *testing.B) {
	// b.SetParallelism(1) // 设置使用的CPU数  (貌似没有作用,  直接在命令行里加入 -cpu=1 指定)
	b.RunParallel(func(pb *testing.PB) { // 关键是这里，开启多个 gotoutine 并发执行
		for pb.Next() {
			Split("a:b:c", ":")
		}
	})
}

// 输入指令：go test -bench=Parallel
// 打印信息：
/*
	goos: windows
	goarch: amd64
	pkg: code.oldboy.com/studygolang/lesson7/01gotest/split
	BenchmarkSplitParallel-8        39922276                30.1 ns/op
	PASS
	ok      code.oldboy.com/studygolang/lesson7/01gotest/split      2.258s

*/

// 输入指令：go test -bench=Parallel -cpu=1
// 打印信息：
/*
	goos: windows
	goarch: amd64
	pkg: code.oldboy.com/studygolang/lesson7/01gotest/split
	BenchmarkSplitParallel  10802042               107 ns/op
	PASS
	ok      code.oldboy.com/studygolang/lesson7/01gotest/split      1.309s

*/

// 在整个测试用例之前 和 之后做什么事情
func TestMain(m *testing.M) {
	fmt.Println("write setup code here...") // 测试用例之前准备做什么

	retCode := m.Run() // 执行整个 测试用例后 返回的结果。

	fmt.Println("write teardown code here...") // 测试完了之后做什么
	os.Exit(retCode)

	// 打印信息：
	/*
		write setup code here...
		fmt 输出结束
		fmt 输出结束2
		fmt 输出结束3
		fmt 输出结束4
		fmt 输出结束5
		PASS
		write teardown code here...
		ok      code.oldboy.com/studygolang/lesson7/01gotest/split      0.037s

	*/
}

// 在 子测试用例 之前和之后做什么事情
// 有时候我们可能需要为每个测试集设置Setup与Teardown，也有可能需要为每个子测试设置Setup与Teardown。下面我们定义两个函数工具函数如下：

// 测试集的 Setup 与 Teardown
func setupTestCase(t *testing.T) func(t *testing.T) {
	fmt.Println("如有需要在此执行：测试之前的setup")
	return func(t *testing.T) {
		fmt.Println("如有需要在此执行：测试之后的teardown")
	}
}

// 子测试的 Setup 与 Teardown
func setupSubTest(t *testing.T) func(t *testing.T) {
	t.Log("如有需要在此执行：子测试之前的setup")
	return func(t *testing.T) {
		t.Log("如有需要在此执行：子测试之后的teardown")
	}
}

func TestGroupSplit2(t *testing.T) {

	// 创建一个结构体，定个 传入参数 和 期望结果
	type testcase struct {
		str  string
		sep  string
		want []string
	}

	var testGroup = map[string]testcase{
		"multi":  testcase{"231", "1", []string{"23"}},
		"multi2": testcase{"231", "1", []string{"23"}},
		"normal": testcase{"a:b:c", ":", []string{"a", "b", "c"}},
		"none":   testcase{"a:b:c", "*", []string{"a:b:c"}},
	}

	// 测试集的 setup 与 teardown
	teardownTestCase := setupTestCase(t)
	defer teardownTestCase(t)

	for name, tc := range testGroup {

		t.Run(name, func(t *testing.T) {

			// 子测试的 setup 与 teardown
			teardownSubTest := setupSubTest(t)
			defer teardownSubTest(t)

			ret := Split(tc.str, tc.sep)
			// fmt.Println("------------进入t.Run--------------")
			if ok := reflect.DeepEqual(ret, tc.want); !ok {
				t.Errorf("期望得到: %v, 实际得到：%v\n", tc.want, ret)
			}
		})
	}
}

// 测试集的 setup , teardown
// 输入指令：go test -v -run=TestGroupSplit2
/*
	write setup code here...
	=== RUN   TestGroupSplit2
	如有需要在此执行：测试之前的setup
	=== RUN   TestGroupSplit2/multi
	=== RUN   TestGroupSplit2/multi2
	=== RUN   TestGroupSplit2/normal
	=== RUN   TestGroupSplit2/none
	如有需要在此执行：测试之后的teardown
	--- PASS: TestGroupSplit2 (0.00s)
		--- PASS: TestGroupSplit2/multi (0.00s)
		--- PASS: TestGroupSplit2/multi2 (0.00s)
		--- PASS: TestGroupSplit2/normal (0.00s)
		--- PASS: TestGroupSplit2/none (0.00s)
	PASS
	write teardown code here...
	ok      code.oldboy.com/studygolang/lesson7/01gotest/split      0.046s
*/

// 子测试的 setup, teardown
// 输入指令：go test -v -run=TestGroupSplit2
/*
	write setup code here...
	=== RUN   TestGroupSplit2
	=== RUN   TestGroupSplit2/none
	=== RUN   TestGroupSplit2/multi
	=== RUN   TestGroupSplit2/multi2
	=== RUN   TestGroupSplit2/normal
	--- PASS: TestGroupSplit2 (0.00s)
		--- PASS: TestGroupSplit2/none (0.00s)
			split_test.go:343: 如有需要在此执行：子测试之前的setup
			split_test.go:345: 如有需要在此执行：子测试之后的teardown
		--- PASS: TestGroupSplit2/multi (0.00s)
			split_test.go:343: 如有需要在此执行：子测试之前的setup
			split_test.go:345: 如有需要在此执行：子测试之后的teardown
		--- PASS: TestGroupSplit2/multi2 (0.00s)
			split_test.go:343: 如有需要在此执行：子测试之前的setup
			split_test.go:345: 如有需要在此执行：子测试之后的teardown
		--- PASS: TestGroupSplit2/normal (0.00s)
			split_test.go:343: 如有需要在此执行：子测试之前的setup
			split_test.go:345: 如有需要在此执行：子测试之后的teardown
	PASS
	write teardown code here...
	ok      code.oldboy.com/studygolang/lesson7/01gotest/split      0.040s

*/

// 示例函数  ( // OutPut: 只能有一个，  当有多个示例，多个结果时，结果必须另起一行)
/*
	// OutPut: [a b c]
	// [f efe]

	或者：

	// OutPut:
	// [a b c]
	// [f efe]

*/
func ExampleSplit() {
	fmt.Println(Split("a:b:c", ":"))
	fmt.Println(Split("fsdefesd", "sd"))
	/* 错误格式
	// OutPut: [a b c]
	// OutPut: [f efe]
	*/

	// OutPut: [a b c]
	// [f efe]
}

// 输入指令：go test -v -run=Example
// 打印如下：
/*
	=== RUN   ExampleSplit
	--- PASS: ExampleSplit (0.00s)
	PASS
	ok      code.oldboy.com/studygolang/lesson7/01gotest/split      0.036s

*/
