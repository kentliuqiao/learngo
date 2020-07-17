package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"strings"
// )

// // func fibonacci() intGen {
// // 	a, b := 0, 1
// // 	return func() int {
// // 		a, b = b, a+b
// // 		return a
// // 	}
// // }

// // 定义一个类型，用来实现Read接口
// // type intGen func() int

// // 实现Read接口
// // 函数是一等公民，可以作为参数，返回值，当然也可以作为一个接收者
// func (g intGen) Read(p []byte) (n int, err error) {
// 	next := g()
// 	s := fmt.Sprintf("%d\n", next)
// 	if next > 10000 {
// 		return 0, io.EOF
// 	}

// 	// TODO: incorrect if p is too small!
// 	return strings.NewReader(s).Read(p)
// }

// func printFileContents(reader io.Reader) {
// 	scanner := bufio.NewScanner(reader)
// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}
// }

// func main() {
// 	f := fibonacci()
// 	// fmt.Println(f()) // 1
// 	// fmt.Println(f()) // 1
// 	// fmt.Println(f()) // 2
// 	// fmt.Println(f()) // 3
// 	// fmt.Println(f()) // 5
// 	// fmt.Println(f()) // 8
// 	// fmt.Println(f()) // 13
// 	// fmt.Println(f()) // 21

// 	printFileContents(f)
// }
