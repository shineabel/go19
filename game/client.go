package game

import "encoding/json"

type IPCClient struct {
	con chan  string
}

func NewIPCClient(s *IPCServer) *IPCClient  {

	c := s.Connect()
	return &IPCClient{
		con:c,
	}
}

func (c *IPCClient) Call(method, params string)(resp *Response, err error)  {
	req := &Request{
		Method:method,
		Params:params,
	}
	var b []byte
	b, err  = json.Marshal(req)
	if err != nil {
		return
	}

	c.con <- string(b)

	str := <-c.con

	var resp1 Response
	err = json.Unmarshal([]byte(str), &resp1)
	resp = &resp1


	return resp, nil
}

func (c *IPCClient) Close()  {

	c.con <- "close"
}
