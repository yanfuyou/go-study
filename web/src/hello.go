package main

import (
	"io"
	"log"
	"net/http"
)

// ResponseWriter用于包装处理HTTP服务端的响应消息;Request表示此次请求的一个数据结构体(客户端)
func helloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello go web")
}

func main() {
	// 分发请求,定义web资源路径
	// p1,路径;p2,回调方法
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err.Error())
	}
}
