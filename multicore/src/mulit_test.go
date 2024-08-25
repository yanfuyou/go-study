package src

import (
	"fmt"
	"runtime"
	"testing"
)

type Vector []float64

func (v Vector) Op(a float64) float64 {
	return a
}

const NCPU = 10 // CPU 为16核

// DoSome 分配给每个CPU的计算任务
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
	c <- 1 // 发信号给任务管理器,表示任务完成
}

func (v Vector) DoAll(u Vector) {
	c := make(chan int, NCPU)
	for i := 0; i < NCPU; i++ {
		go v.DoSome(i*len(v)/NCPU, (i+1)*len(v)/NCPU, u, c)
	}
	// 等待所有任务执行完成
	for i := 0; i < NCPU; i++ {
		<-c // 读取操作是阻塞的,因此每读取到一个值就表示一个任务完成
	}
	// 表示所有任务以执行完成

}

// 模拟一批整数进行分批累加
func TestMulticoreTask(t *testing.T) {
	// go 目前还不能很智能的发现和利用多核的优势
	// 可以通过设置环境变量 [GOMAXPROCS]的值来控制使用多少个CPU核心
	// 或者在代码中启动goroutine前先调用 runtime.GOMAXPROCS(16) 来设置使用多少个核心
	nCpu := runtime.NumCPU()
	runtime.GOMAXPROCS(nCpu)
	// 可以通过 runtime.NumCPU() 函数来获取当前机器的核心数
	println("当前核心数: ")
	v1 := Vector{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	v1.DoAll(v1)
	for _, v := range v1 {
		fmt.Printf("%f ", v)
	}
}
