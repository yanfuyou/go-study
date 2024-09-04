package main

import "fmt"
//下面这个C语法的头文件引用注释是有意义的 
// 也可以使用单行注释导入头文件

/*
#include <stdlib.h>
#include <stdio.h>
void hello(){
	printf("Hello from Cgo\n");
}
*/
import "C"
import "unsafe"
func Random() int{
	return int(C.random())
}

func Seed(i int){
	C.srandom(C.uint(i))
}

func show(){
	C.hello()
}

func main(){
	Seed(100)
	fmt.Println("Random:", Random())
	var gostr string
	gostr = "Hello from Go"
	cstr := C.CString(gostr)
	defer C.free(unsafe.Pointer(cstr))
	fmt.Println("goStr:", gostr)
	// C.sprintf(cstr,"content is: %d",123)
	show()
}