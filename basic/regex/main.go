package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is kentliuqiao@gmail.com@abc.com
email1 is abc@def.org
email2 is    kk@qq.com
email3 is  dddd@abc.com.cn
`

func main() {
	re := regexp.MustCompile(
		`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	// match := re.FindString(text)
	// matches := re.FindAllString(text, -1)// 无法提取出括号中的内容
	subMatch := re.FindAllStringSubmatch(text, -1)
	// fmt.Println(match)
	// fmt.Println(matches)
	for _, m := range subMatch {
		fmt.Println(m)
	}
}
