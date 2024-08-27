package ipc

import (
	"encoding/json"
)

// IpcClient ipc客户端
type IpcClient struct {
	conn chan string
}

func NewIpcClient(server *IpcServer) *IpcClient {
	c := server.Connect()
	return &IpcClient{c}
}

// Call 发送请求
func (client *IpcClient) Call(method, params string) (resp *Response, err error) {
	req := &Request{method, params}
	var b []byte
	b, err = json.Marshal(req)
	if err != nil {
		return
	}
	client.conn <- string(b)
	// 等待返回结果
	str := <-client.conn
	var resp1 Response
	err = json.Unmarshal([]byte(str), &resp1)
	resp = &resp1
	return
}

// Close 关闭连接
func (client *IpcClient) Close() {
	client.conn <- "CLOSE"
}
