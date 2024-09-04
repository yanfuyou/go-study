package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	// Kind() 返回类型的具体信息
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("type:", v.Type(), "value:", v.Float())
	// Type中有一个CanSet()函数,用于判断是否是一个可设属性
	fmt.Println("settability of v:", v.CanSet())
	// v.Set(4.0) // 将会报错

	p := reflect.ValueOf(&x) //传入x的地址
	fmt.Println("type:", p.Type())
	fmt.Println("settability of p:", p.CanSet())

	vv := p.Elem()
	fmt.Println("settability of vv:", vv.CanSet())
	vv.SetFloat(100.0)
	fmt.Println("has changed x:", x)
	fmt.Println("------------reflect of struct-----------")
	t := T{203, "hello reflect"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

}
