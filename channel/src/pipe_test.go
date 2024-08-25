package src

import (
	"fmt"
	"testing"
)

// PipeData 学习channel 的可被传递性
type PipeData struct {
	value   int
	handler func(int) int
	next    chan int
}

// 流式处理
func handler(queue chan *PipeData) {
	for data := range queue {
		data.next <- data.handler(data.value)
	}
}

func TestPipe(t *testing.T) {
	// 创建处理函数
	addTwo := func(x int) int {
		return x + 2
	}

	multiplyByThree := func(x int) int {
		return x * 3
	}

	subtractFour := func(x int) int {
		return x - 4
	}

	// 创建通道
	step1 := make(chan int)
	step2 := make(chan int)
	step3 := make(chan int)

	// 创建 PipeData 实例并启动 goroutine 进行处理
	go func() {
		queue := make(chan *PipeData)

		// 启动流式处理器
		go handler(queue)
		go handler(queue)
		go handler(queue)

		// 将数据放入队列
		queue <- &PipeData{value: 1, handler: addTwo, next: step1}
		queue <- &PipeData{value: <-step1, handler: multiplyByThree, next: step2}
		queue <- &PipeData{value: <-step2, handler: subtractFour, next: step3}

		// 关闭队列
		close(queue)
	}()

	// 获取最终结果
	result := <-step3
	fmt.Println("Final result:", result)
}
