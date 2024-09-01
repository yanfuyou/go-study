package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Student struct {
	*Person
	Class string
}

func main() {
	per := &Person{Name: "小明", Age: 18}
	jStr, _ := json.Marshal(per)
	personStr := string(jStr)
	fmt.Println(personStr)
	var b Student
	json.Unmarshal(jStr, &b)
	b.Class = "终极一班"
	fmt.Println("解析json字符串： ", b.Name, b.Age, b.Class)
}
