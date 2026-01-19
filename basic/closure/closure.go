package main

import "fmt"

// 闭包由方法和运行时上下文组成
func genClosure(param string) func() string {
	return func() string {
		return param
	}
}

func main() {
	f := genClosure("hello")
	fmt.Println(f())

	// 1.22+版本中，for会为每个i创建单独的变量，每次闭包捕获的都是单独的i
	for i := 0; i < 10; i++ {
		defer func() {
			println(i)
		}()
	}
}
