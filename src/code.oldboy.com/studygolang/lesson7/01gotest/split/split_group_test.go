package split

import (
	"reflect"
	"testing"
)

// 注意：多个测试用例，函数名不能相同

func TestGroupSplit(t *testing.T) {

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

	for name, tc := range testGroup {

		// fmt.Println("--------------进入for------------")

		//  传统测试形式：
		/*
			ret := Split(tc.str, tc.sep)
			if ok := reflect.DeepEqual(ret, tc.want); !ok {
				t.Errorf("测试用例：%v 出错！ 期望得到: %v, 实际得到：%v\n", name, tc.want, ret)
				// 当某一项测试用例失败了，则可以打印出 name ，明确知道是哪个测试用例出现了问题
				// 并且一个测试组，如果第一个测试用例出现了错误，则不会影响后续的执行，结果会打印出 成功的和 失败的。
			}
		*/

		// 传统测试打印信息：(go test -run Group -v)
		/*
			=== RUN   TestGroupSplit
			--- FAIL: TestGroupSplit (0.00s)
				split_group_test.go:30: 测试用例：multi 出错！ 期望得到: [23], 实际得到：[ 23]
			FAIL
			exit status 1
			FAIL    code.oldboy.com/studygolang/lesson7/01gotest/split      0.058s
		*/

		// go 语言内置函数：

		t.Run(name, func(t *testing.T) {
			ret := Split(tc.str, tc.sep)
			// fmt.Println("------------进入t.Run--------------")
			if ok := reflect.DeepEqual(ret, tc.want); !ok {
				t.Errorf("期望得到: %v, 实际得到：%v\n", tc.want, ret)
			}
		})

		// 说明1：
		// t.Run() 打印信息：
		// go test -run Group -v ,  打印信息如下：
		/*
			=== RUN   TestGroupSplit
			=== RUN   TestGroupSplit/multi
			=== RUN   TestGroupSplit/multi2
			=== RUN   TestGroupSplit/normal
			=== RUN   TestGroupSplit/none
			--- FAIL: TestGroupSplit (0.00s)
				--- FAIL: TestGroupSplit/multi (0.00s)
					split_group_test.go:56: 期望得到: [23], 实际得到：[ 23]
				--- PASS: TestGroupSplit/multi2 (0.00s)
				--- PASS: TestGroupSplit/normal (0.00s)
				--- PASS: TestGroupSplit/none (0.00s)
			FAIL
			exit status 1
			FAIL    code.oldboy.com/studygolang/lesson7/01gotest/split      0.037s

		*/

		// 说明2：
		// 使用 t.Run() 的好处，是可以使用命令行执行 某个子单元测试:
		// go test -v -run TestGroupSplit/none  ,  打印信息如下：
		/*
			=== RUN   TestGroupSplit
			=== RUN   TestGroupSplit/none
			--- PASS: TestGroupSplit (0.00s)
			    --- PASS: TestGroupSplit/none (0.00s)
			PASS
			ok      code.oldboy.com/studygolang/lesson7/01gotest/split      0.037s
		*/

		// 说明3：
		// 使用 t.Run() 指定子单元 名称， 子单元也可以进行 模糊匹配，例如：
		// go test -v -run TestGroupSplit/multi      (匹配上了 multi 和 multi2 两个子单元)
		/*
			=== RUN   TestGroupSplit
			=== RUN   TestGroupSplit/multi
			=== RUN   TestGroupSplit/multi2
			--- FAIL: TestGroupSplit (0.00s)
				--- FAIL: TestGroupSplit/multi (0.00s)
					split_group_test.go:56: 期望得到: [23], 实际得到：[ 23]
				--- PASS: TestGroupSplit/multi2 (0.00s)
			FAIL
			exit status 1
			FAIL    code.oldboy.com/studygolang/lesson7/01gotest/split      0.039s

		*/

		// 说明4：
		// 查看子单元测试项，即使某个单元测试测试不通过，也不影响你想要测试的指定名称的 子单元。
	}

}
