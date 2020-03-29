package main

import (
	"image"
	"sync"
)

var icons map[string]image.Image

var loadIconsOnce sync.Once //

// sync.Once 使用场景
/*
sync 包提供了针对一次性初始化问题的特优解决方案：sync.Once
从概念上来讲， Once 包含一个 布尔变量和一个互斥变量，布尔变量记录初始化是否已经完成，
互斥量则负责保护这个布尔变量和客户端的数据结构。

*/
func main() {

	for i := 0; i < 100000; i++ {
		// go Icon("left")    // 当多个 goroutine 调用 Icon 方法时，不是并发安全的。
	}

}

// 现代的编译器和 CPU 可能会在保证每个 goroutine 都满足串行一致的基础上自由地重排访问内存的顺序。 loadIcons 函数可能会被重排为以下结果：
/*
	func loadIcons(){
		icons = make(map[string]image.Image)     // 执行到这里，代表 icons 已经实例化了，不为  nil ,但是 map 里面的初始化还未完成
		icons["left"] = loadIcon("left.png")
		icons["up"] = loadIcon("up.png")
		icons["right"] = loadIcon("right.png")
		icons["down"] = loadIcon("down.png")
	}

	// 在这种情况下就会出现 即使判断了 icons 不是 nil 也不意味着初始化完了。 考虑到这种情况，我们能想到的办法就是 添加互斥锁，
	// 保证初始化 icons 的时候不会被其它的 goroutine 操作，但是这样做又引发性能问题。

*/
func loadIcons() {
	icons = map[string]image.Image{
		"left":  loadIcon("left.png"),
		"up":    loadIcon("up.png"),
		"right": loadIcon("right.png"),
		"down":  loadIcon("down.png"),
	}
}

// Icon 是一个根据名字获取图片的方法
func Icon(name string) image.Image {
	/*
		if icons == nil {
			loadIcons() // 重点在这里。
		}
	*/

	loadIconsOnce.Do(loadIcons)
	// 使用 sync.Once (包含一个布尔变量和一个互斥量，布尔变量记录初始化是否已经完成，互斥量则负责保护这个布尔变量和客户端的数据结构)

	// 每次调用 Do(loadIcons) 时会先锁定互斥量并检查里边的布尔变量。
	// 在第一次调用时，这个布尔变量为假， Do 会调用 loadIcons 然后把变量设置为真。
	// 后续的调用相当于空操作，只是通过 互斥量的同步来保证 loadIcons 对内存产生的效果 (在这里就是 icons 变量)对所有的 groutine 可见。

	// 详见 sync.Once 源码 (源码很简单)

	return icons[name]
}

func loadIcon(name string) image.Image {
	return nil // 模拟返回了 Image 对象
}

// 详见 sync.Once 源码 (源码很简单)

/*
package sync

import (
	"sync/atomic"
)

// Once is an object that will perform exactly one action.
type Once struct {
	// done indicates whether the action has been performed.
	// It is first in the struct because it is used in the hot path.
	// The hot path is inlined at every call site.
	// Placing done first allows more compact instructions on some architectures (amd64/x86),
	// and fewer instructions (to calculate offset) on other architectures.
	done uint32
	m    Mutex
}


func (o *Once) Do(f func()) {
			....


	if atomic.LoadUint32(&o.done) == 0 {   // 查看 标识位，默认为 0  （第一个进入的可以正常通过，但是其它的 goroutine 进来则会被挡在门外）
		// Outlined slow-path to allow inlining of the fast-path.
		o.doSlow(f)
	}
}

func (o *Once) doSlow(f func()) {
	o.m.Lock()    // 加锁
	defer o.m.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)  // 如果第一次进来，则将 值置为 1， 并调用 f() (这个函数就是 上面的 loadIcons)
		f()
	}
}


*/
