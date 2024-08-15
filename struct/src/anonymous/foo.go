package anonymous

import (
	"fmt"
)

// 定义一个Foo类

// Foo 继承Base类
type Foo struct {
	Base
	// ...
}

// Bar 重写从Base中继承的方法
func (foo *Foo) Bar() {
	// 先调用了Base类的bar方法
	foo.Base.bar()
	fmt.Println("foo`s bar")
}
