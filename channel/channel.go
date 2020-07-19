package main

import (
	"fmt"
	"time"
)

// channel 作为参数
// 接受方有两种方法判断channel是否已被close：ok 或者 range
func worker(id int, c chan int) {
	// for {
	// 	n, ok := <-c
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Printf("worker %d receiverd %d\n", id, n)
	// }

	for n := range c {
		fmt.Printf("worker %d receiverd %d\n", id, n)
	}
}

// channel 作为返回值
func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	// var c chan int // c == nil
	var channels [10]chan<- int
	// for i := 0; i < 10; i++ {
	// 	channels[i] = make(chan int)
	// 	go worker(i, channels[i])
	// }
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c) // 必须是发送方close channel
	time.Sleep(time.Millisecond)
}

func main() {
	// chanDemo()
	// bufferedChannel()
	channelClose()
}
