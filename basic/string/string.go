package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// go中string是常量，不能修改
	var s string = "Hello World"
	// b := []byte(s)

	// 字符串长度
	fmt.Println(len(s))
	// 中文字符和英语字符大小不一样，len返回字节长度，
	fmt.Println(utf8.RuneCountInString("中文"))

	// 字符串操作
	strings.Split(s, " ")
	strings.Join([]string{"a", "b", "c"}, "|")
}
