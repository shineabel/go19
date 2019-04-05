package main

import "fmt"

type API2 interface {
	Test() string
}



type AModuleAPI interface {
	TestA() string
}

type AMoudleAPIImplement struct {

}

func (*AMoudleAPIImplement) TestA() string  {
	return "A module api working"
}

type BModuleAPI interface {
	TestB() string
}

type BModuleAPIImplement struct {

}

func (*BModuleAPIImplement) TestB() string  {
	return "b module api working"
}

func NewAModuleAPI() AModuleAPI  {
	return &AMoudleAPIImplement{}
}

func NewBModuleAPI() BModuleAPI  {
	return &BModuleAPIImplement{}
}

type APIImplement struct {
	a AModuleAPI
	b BModuleAPI
}

func (api *APIImplement) Test() string  {

	a := api.a.TestA()
	b := api.b.TestB()

	return fmt.Sprintf("%s:%s",a,b)
}

func NewAPI2() API2  {
	return &APIImplement{
		a:NewAModuleAPI(),
		b:NewBModuleAPI(),
	}
}