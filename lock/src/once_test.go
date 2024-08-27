package src

import (
	"sync"
	"testing"
)

// 对于全局角度只需要执行一次的代码,例如全局初始化动作
// go提供了一个Once类型来保证全局的唯一性操作

var once sync.Once

var a string

func setup() {
	println("init a")
	a = "hallow go"
}

func doPrint() {
	once.Do(setup)
	println(a)
}

// 当所有其他的goroutine调用到使用once.Do语句时,均会被阻塞,知道once.Do调用结束再继续执行
func TestPrint(t *testing.T) {
	go doPrint()
	go doPrint()
}

// sync.atomic中提供了一些对于基础类型的原子操作
