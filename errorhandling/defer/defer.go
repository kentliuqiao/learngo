package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kentliuqiao/learngo/functional/fib"
)

func tryDefer() {
	defer fmt.Println(1)
	fmt.Println(2)
}

func writeFile(fileName string) {
	// file, err := os.Create(fileName)
	file, err := os.OpenFile(fileName,
		os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op,
				pathError.Path,
				pathError.Err)
		}
		return
	}
	defer file.Close()

	f := fib.Fibonacci()
	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("fib.txt")
}
