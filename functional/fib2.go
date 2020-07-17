package main

// import (
// 	"bufio"
// 	"bytes"
// 	"fmt"
// 	"io"
// )

// func fibonacci() intGen {
// 	a, b := 0, 1
// 	return func() int {
// 		a, b = b, a+b
// 		return a
// 	}
// }

// type intGen func() int

// // 定义一个类型，用来实现Read接口
// type bufIntGen struct {
// 	g   intGen
// 	buf bytes.Buffer
// }

// // 实现Read接口
// // 函数是一等公民，可以作为参数，返回值，当然也可以作为一个接收者
// // func (g intGen) Read(p []byte) (n int, err error) {
// // 	next := g()
// // 	s := fmt.Sprintf("%d\n", next)
// // 	if next > 10000 {
// // 		return 0, io.EOF
// // 	}

// // 	// TODO: incorrect if p is too small!
// // 	return strings.NewReader(s).Read(p)
// // }
// func (b *bufIntGen) Read(p []byte) (n int, err error) {
// 	if b.buf.Len() == 0 {
// 		next := b.g()
// 		if next > 10000 {
// 			return 0, io.EOF
// 		}
// 		_, err := fmt.Fprintf(&b.buf, "%d\n", next)
// 		if err != nil {
// 			return 0, err
// 		}
// 	}
// 	return b.buf.Read(p)
// }

// func printFileContents(reader io.Reader) {
// 	scanner := bufio.NewScanner(reader)
// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}
// }

// func main() {
// 	f := fibonacci()
// 	r := bufIntGen{g: f}
// 	p := make([]byte, 2)
// 	for {
// 		n, err := r.Read(p)
// 		if err != nil {
// 			break
// 		}
// 		fmt.Printf("%s", p[:n])
// 	}

// }
