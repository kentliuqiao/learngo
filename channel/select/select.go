package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(
				rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	// for {
	// 	n, ok := <-c
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Printf("worker %d receiverd %d\n", id, n)
	// }

	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d receiverd %d\n", id, n)
	}
}

// channel 作为返回值
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)
	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select { // 非阻塞式从管道中获取数据
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond): // 在没两次select之间如果超过800毫秒还未产生新的数据，则提示timeout
			fmt.Println("timeout")
		case <-tick:
			fmt.Println("len of values =", len(values))
		case <-tm: // 让该循环10秒之后结束
			fmt.Println("bye")
			return
		}
	}
}
