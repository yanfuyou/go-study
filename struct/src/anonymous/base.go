package anonymous

import "fmt"

// 匿名组合

// Base 定义一个Base类,并声明foo 和bar两个方法
type Base struct {
	Name string
}

func (base *Base) foo() {

}

func (base *Base) bar() {
	fmt.Println("base`s bar")
}
