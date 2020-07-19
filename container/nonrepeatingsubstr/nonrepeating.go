package main

import "fmt"

var lastOccurred = make([]int, 0xffff) // 放在这，仅需要make一次即可

func lengthOfNonRepeatingSubStr(s string) int {
	// lastOccurred := make(map[rune]int) // mapassign和mapget都是比较耗时的操作，值得优化
	// lastOccurred := make([]int, 0xffff) // 这里在benchmark测试时会循环调用，也值得优化
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}

	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI := lastOccurred[ch]; lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("aaaaa"))
}
