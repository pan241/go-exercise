package main

import (
	"fmt"
	"sync"
)

func main() {

	var (
		// 通过waitgroup管理协程
		wg sync.WaitGroup
		// 通过锁同步
		mu sync.Mutex
	)
	// 通过channel通信和同步
	ch := make(chan int, 10)
	go func() {
		ch <- 1
	}()

	mp := map[int]int{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			defer func() {
				wg.Done()
			}()
			mu.Lock()
			defer func() {
				mu.Unlock()
			}()
			mp[idx] = <-ch
			ch <- mp[idx] + 1
		}(i)
	}

	wg.Wait()
	fmt.Printf("%+v", mp)
}
