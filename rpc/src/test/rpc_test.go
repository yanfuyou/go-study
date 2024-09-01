package test

import (
	"fmt"
	"go-study/rpc/src/server"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"testing"
)

func TestRpc(t *testing.T) {
	arith := new(server.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal()
	}
	go http.Serve(l, nil)

	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing: ", err)
	}
	args := &server.Args{7, 8}
	var replay int
	err = client.Call("Arith.Multiply", arith, replay)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d =%d", args.A, args.B, replay)
	quotient := new(server.Quotinet)
	divCall := client.Go("Arith.Device", args, &quotient, nil)
	replayCall := <-divCall.Done
	fmt.Println(replayCall)
}
