package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

/*
	纵观count变量、trigger()以及改造后的for语句和go函数
	就是让count变量成为一个信号，它的值总是下一个可以调用打印函数的go函数的序号
	这个序号就是启用goroutine时当次迭代的序号。正因为如此，go函数执行顺序与go语句执行顺序完全一致
	此外trigger()实现了一种自旋(spinning)。除非发现条件已满足，否则它会不断地进行检查

	异步发起的go函数得到了同步(按照既定顺序)执行
*/
func main() {
	var count uint32
	/*
		func()代表既无参数声明也无结果声明的函数类型
		原子操作函数对被操作的数值的类型有约束，因为int类型长度会根据计算架构(32位或64位)而改变
		所以对count以及相关的变量和参数的类型进行了统一变更(由int变为uint32)
	*/
	trigger := func(i uint32, fn func()) {
		for {
			//不断获取count变量的值，并判断该值是否与参数i的值相同
			if n := atomic.LoadUint32(&count); n == i {
				fn()
				/*
					操作变量count使用的都是原子操作
					这是由于trigger()会被多个goroutine并发调用，所以它用到的非本地变量count，就被多个用户级线程共用了
					因此对它的操作就产生了竞态条件(race condition)，破坏了程序的并发安全性
					应该对这样的操作加以保护，在sync/atomic包中声明了很多用于原子操作的函数
				*/
				atomic.AddUint32(&count, 1)
				break
			}
			//不相等则让当前goroutine睡眠一纳秒再进入下一个迭代
			time.Sleep(time.Nanosecond)
		}
	}
	
	//
	for i := uint32(0); i < 10; i++ {
		go func(i uint32) {
			//声明一个匿名函数
			fn := func() {
				fmt.Println(i)
			}
			trigger(i, fn)
		}(i)
	}
	/*
		想让主goroutine最后一个运行完毕
		当所有手动启用的goroutine都运行完毕后，count的值一定会是10
	*/
	trigger(10, func() {})
}