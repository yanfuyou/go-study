package src

import (
	"fmt"
	"testing"
)

var ch1 chan int       // 非单向的channel
var ch2 chan<- float64 // 单向channel,只用于写入float64数据
var ch3 <-chan int     // 单向channel,只用于读取int数据

// channel是一个原生数据类型,可以被传递,也可以进行转换.
func TestOneWay(t *testing.T) {
	ch4 := make(chan int)
	ch5 := <-chan int(ch4) // 单向的读channel
	fmt.Println(ch5)
	ch6 := chan<- int(ch4) // 单向的写channel
	fmt.Println(ch6)
}

func Parse(ch <-chan int) {
	for value := range ch {
		fmt.Println("Parsing value:", value)
	}
}
