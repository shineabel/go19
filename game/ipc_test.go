package game

import (
	"fmt"
	"testing"
)

type EchoServer struct {
}

func (s *EchoServer) Handle(method, params string) *Response {

	return &Response{
		Code: "ok",
		Body: "echo:" + method + "-" + params,
	}
}

func (s *EchoServer) Name() string {
	return "echoserver"

}

func TestIPCClient_Call(t *testing.T) {
	es := &EchoServer{}
	server := NewIPCServer(es)

	c1 := NewIPCClient(server)
	c2 := NewIPCClient(server)

	resp1, _ := c1.Call("foo", "from client1")
	resp2, _ := c2.Call("bar", "from client2")

	fmt.Println(resp1.Body)

	fmt.Println(resp2.Body)

	c1.Close()
	c2.Close()
}
