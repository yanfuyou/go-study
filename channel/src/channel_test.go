package src

import (
	"fmt"
	"testing"
	"time"
)

func Count(ch chan int) {
	ch <- 1
	fmt.Println("counting ")
}

func TestChannel(t *testing.T) {
	chs := make([]chan int, 10)
	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}

	for _, ch := range chs {
		<-ch
	}
}

// TestTimeout 使用select 避免永久等待
func TestTimeout(t *testing.T) {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9) // 等待一秒
		timeout <- true
	}()
	ch := make(chan int, 1)
	// 若不能从ch中读取到数据,也会在1秒后结束等等
	select {
	case <-ch:
	case <-timeout:
		fmt.Println("read timeout")
	}
}
