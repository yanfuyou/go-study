package src

import (
	"fmt"
	"testing"
)

func TestClose(t *testing.T) {
	ch := make(chan int)
	go func() {
		a := <-ch
		fmt.Println(a)
	}()
	ch <- 1
	// 关闭channel
	close(ch)
	// 判断channel是否关闭
	v, ok := <-ch
	// 返回ok表示channel已经关闭,V表示该管道中的值,没有则为该值的0值
	fmt.Println(v, ok)
}
