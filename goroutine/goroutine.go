package main

import (
	"fmt"
	"time"
)

// 协程 Coroutine
// 轻量级线程
// 非抢占式多任务处理，有协程主动交出控制权（1.14加入异步抢占）
// 编译器/解释器/虚拟机层面的多任务，非OS层面
// 多个协程可能在一个或多个线程上运行

// subroutines are special cases of more general program components, called coroutines. --By Donald Knuth
// 子程序（函数）时协程的一个特例
// 普通函数：main-->dowork（main函数将控制权交由dowork函数，dowork完成之后才会将控制权交回给main）
// 协程：main<-->dowork（main函数和dowork之间存在一个双向通道，双方可以互相通信，且控制权可以来回切换）

// 其他语言中的协程
// C++: Boos.Corotine
// Java: 不支持（一些特殊JVM支持）
// Python: 使用yield关键字实现协程；Python 3.5加入了async def对协程原生支持

// goroutine
// 任何函数只需要加上go就能够送给调度器运行
// 不需要在定义时区分是否是异步函数（优于Python）
// 调度器在合适的点进行切换（不需要特意程序员关心何时进行切换）
// 使用 -race 进行数据访问冲突检测

// goroutine可能的切换点
// I/O，select
// channel
// 等待锁
// 函数调用（有时，具体由调度器决定）
// runtime.Goshced()（手动切换）
// 以上只是参考，不能保证一定切换，不能保证在除了以上的其他地方不切换

func main() {
	var a [10]int
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				// 在1.14版本以前，goroutine调度器采用非抢占式策略，
				// 对于没有函数调用的循环会造成调度器死锁，
				// 下面的程序如果去掉下面第二行代码，在1.14以前会造成死锁，
				// 程序无法退出，跑满cpu
				// 该问题在1.14得到了解决，在这个版本中，goroutine调度器
				// 实现了异步抢占的策略
				// a[i]++
				// runtime.Gosched() // 手动交出控制权

				fmt.Printf("hello from goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
	fmt.Println(a)
}
