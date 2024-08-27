package ipc

import (
	"fmt"
	"testing"
	"time"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(request, params string) *Response {
	//return "ECHO:" + request
	return &Response{"200", params}
}

func (server *EchoServer) Name() string {
	return "EchoServer"
}

func TestIpc(t *testing.T) {
	server := NewIpcServer(&EchoServer{})
	c1 := NewIpcClient(server)
	c2 := NewIpcClient(server)
	r1, _ := c1.Call("post", "from c1")
	r2, _ := c2.Call("post", "from c2")
	fmt.Println(r1.Body)
	fmt.Println(r2.Body)
	time.Sleep(5 * time.Second)
	c1.Close()
	c2.Close()
}
