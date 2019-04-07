package game

import (
	"fmt"
	"encoding/json"
)

type Request struct {
	Method string `json:"method"`
	Params string `json:"params"`
}

type Response struct {
	Code string `json:"code"`
	Body string `json:"body"`
}

type Server interface {
	Name() string
	Handle(method, params string) *Response
}

type IPCServer struct {
	Server
}

func NewIPCServer(s Server) *IPCServer  {
	return &IPCServer{
		s,
	}
}

func (s *IPCServer) Connect() chan string  {

	session := make(chan  string,0)

	go func(c chan string) {

		for {
			request := <- c
			if request == "close" {
				break
			}

			var req Request

			err := json.Unmarshal([]byte(request),&req)
			if err != nil {
				fmt.Println("invalid request error:", err)
				return
			}

			resp := s.Handle(req.Method, req.Params)
			b , err := json.Marshal(resp)
			c <- string(b)
		}
	}(session)



	fmt.Printf("a new session has been create successfully\n")
	return session
}

//func (s *IPCServer) Handle(method , params string) *Response  {
//
//	return &Response{
//		Code:"ok",
//		Body:"server:"+ method + "-" + params,
//	}
//}
