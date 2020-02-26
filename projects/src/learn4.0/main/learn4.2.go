package main

import "fmt"

// 学习第四章-复合数据类型--4.2-slice
/*
	slice 表示一个拥有相同类型元素的可变长度的序列。 slice 通常写成 []T, 其中元素的类型都是 T；
		它看上去像没有长度的数组类型。
	数组和 slice 是紧密关联的。 slice 是一种轻量级的数据结构，可以用来访问数组的部分或者全部的元素，
		而这个数组称为 slice 的底层数组。slice 有三个属性：指针、长度和容量。
		指针指向数组的第一个可以从 slice 中访问的元素，这个元素并不一定是数组的第一个元素。
		长度是指 slice 中的元素个数，它不能超过slice 的容量。容量的大小通常是从 slice 的起始元素
		到底层数组的最后一个元素间元素的个数。Go 的内置函数 len 和 cap 用来返回 slice 的长度和容量。
 */
func main() {
	fmt.Println("learn4.2")

	sliceFunc()
}

/*
	一个底层数组可以对应多个 slice， 这些slice 可以引用数组的任何位置，彼此之间的元素还可以重叠。
 */
func sliceFunc(){
	months:= [...]string{1:"January", 2:"February",3:"March",4:"April",5:"May",6:"June",7:"July",
							8:"August",9:"September",10:"October",11:"November",12:"December"}
	/*	以上省略了 数组的第一个元素，months[0] 默认为""
		所以 January 就是 months[1], December 是 months[12]。一般来讲，数组中索引 0 的位置存放数组的第一个元素，
		但是由于月份总是从 1 开始，因此我们可以不设置索引为 0 的元素，这样它的值就是空字符串。
	 */

	/*
		slice 操作符 s[i:j] (其中 0<=i<=j<=cap(s) ) 创建了一个新的 slice, 这个新的 slice 引用了序列 s 中
		从 i 到 j-1 索引位置的所有元素，这里的 s 既可以是数组或者 指向数组的指针，也可以是 slice。
		新的 slice 的元素个数是 j-i 个。 如果上面的表达式中省略了 i , 那么新 slice 的起始索引位置就是 0，即 i=0;
		如果省略了 j , 那么新  slice 的结束索引位置是 len(s)-1，即 j=len(s)。
		因此 slice months[1:13] 引用了所有的有效月份，同样的写法可以是 months[1:]。
		slice months[:] 引用了整个数组。接下来，我们定义元素重叠的 slice，分别用来表示第二季度的月份和北半球的夏季月份：
	 */

	Q2:= months[4:7]
	summer:= months[6:9]
	fmt.Println(Q2)        // "[April May June]"
	fmt.Println(summer)    // "[June July August]"

	for _, s:= range summer {
		for _, q:= range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)   // "June appears in both"
			}
		}
	}

	/*
		如果 slice 的引用超过了被引用对象的容量，即 cap(s)，那么会导致程序宕机；
		但是如果 slice 的引用超出了被引用对象的长度，即 len(s)，那么最终 slice 会比原 slice 长：
	 */
	//fmt.Println(summer[:20])     // 运行出错： panic: runtime error: slice bounds out of range [:20] with capacity 7
	endlessSummer:= summer[:5]   // 在 slice 容量范围内扩展了 slice
	fmt.Println(endlessSummer)   // "[June July August September October]"

	str := endlessSummer[0][1:]
	fmt.Printf("slice endlessSummer[0][1:] ：%s\n", str)   // "une"
	// 取 endlessSummer 里面下标为0的元素的值，也就是 "June",
	//	然后截取这个字符串：(下标为1 开始到最后的 字符) 得到 "une"
}



