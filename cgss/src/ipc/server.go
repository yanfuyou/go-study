package ipc

// 实现一个简单的进程间通信框架

import (
	"encoding/json"
	"fmt"
)

type Request struct {
	Method string "method"
	Params string "params"
}

type Response struct {
	Code string "code"
	Body string "body"
}

// Server 通信接口
type Server interface {
	Name() string
	Handle(method, params string) *Response
}

// 接口实现
type IpcServer struct {
	Server
}

func NewIpcServer(server Server) *IpcServer {
	return &IpcServer{server}

}

// 创建连接
func (server *IpcServer) Connect() chan string {
	session := make(chan string, 0)
	go func(c chan string) {
		for {
			request := <-c
			if request == "CLOSE" {
				// 关闭连接
				break
			}
			var req Request
			err := json.Unmarshal([]byte(request), &req)
			if err != nil {
				fmt.Println("错误的请求格式:", request)
			}
			resp := server.Handle(req.Method, req.Params)
			b, err := json.Marshal(resp)
			// 返回结果
			c <- string(b)
		}
	}(session)
	fmt.Println("session closed")
	return session
}
