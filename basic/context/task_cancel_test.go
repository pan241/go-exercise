package main

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
	"time"
)

func TestCtxCancel(t *testing.T) {
	// 级联取消
	ctx := context.Background()
	pCtx, pcancel := context.WithCancel(ctx)

	assert.Equal(t, 2, runtime.NumGoroutine())

	for i := 0; i < 3; i++ {
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					fmt.Println("parent is working")
					time.Sleep(time.Millisecond * 500)
				}
			}
		}(pCtx)
	}

	// 创建子上下文
	cCtx, _ := context.WithCancel(pCtx)
	// goroutine 监听
	for i := 0; i < 5; i++ {
		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done():
					return
				default:
					fmt.Printf("child%d is working\n", i)
					time.Sleep(time.Millisecond * 500)
				}
			}
		}(cCtx)
	}
	time.Sleep(time.Millisecond * 500)
	pcancel()
	time.Sleep(time.Millisecond * 1000)

	assert.Equal(t, 2, runtime.NumGoroutine())
	assert.Equal(t, context.Canceled, pCtx.Err())

}
