package main

import (
	"context"
	"fmt"
	"time"
)

// context用于数据透传，任务级联取消

func main() {
	// context.Background()作为根节点
	ctx := context.Background()

	// 给ctx添加值，可以连续添加
	child1 := context.WithValue(ctx, "name", "child1")
	fmt.Printf("name: %s\n", child1.Value("name"))

	// 超时取消
	child2, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	//select {
	//case <-child2.Done():
	//	err := child2.Err()
	//	fmt.Println(err)
	//}

	time.Sleep(time.Second * 1)

	// timeout继承以最短的为准
	// 此时child3的超时时间为2s，child4的超时时间为500ms
	child3, cancel2 := context.WithTimeout(child2, time.Second*3)
	defer cancel2()
	fmt.Println(child3)
	child4, _ := context.WithTimeout(child2, time.Millisecond*500)
	fmt.Println(child4)

	// 显式关闭
	_, cancel3 := context.WithCancel(ctx)
	time.Sleep(time.Millisecond * 500)
	cancel3()
}
