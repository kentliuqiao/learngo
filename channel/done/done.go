package main

import (
	"fmt"
	"sync"
)

// channel 作为参数
// 接受方有两种方法判断channel是否已被close：ok 或者 range
func doWork(id int, w worker) {
	// for {
	// 	n, ok := <-c
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Printf("worker %d receiverd %d\n", id, n)
	// }

	for n := range w.in {
		fmt.Printf("worker %d receiverd %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

// channel 作为返回值
func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}

func chanDemo() {
	// var c chan int // c == nil
	var workers [10]worker
	var wg sync.WaitGroup
	// for i := 0; i < 10; i++ {
	// 	channels[i] = make(chan int)
	// 	go worker(i, channels[i])
	// }
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, w := range workers {
		wg.Add(1)
		w.in <- 'a' + i
		// <-workers[i].done
	}
	for i, w := range workers {
		wg.Add(1)
		w.in <- 'A' + i
		// <-workers[i].done
	}

	wg.Wait()

	// wait for all dones
	// for _, w := range workers {
	// 	<-w.done
	// 	<-w.done
	// }
	// time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()
}
