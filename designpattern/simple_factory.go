package main

import "fmt"

type API interface {
	say(word string) string
}

func NewAPI(t int) API  {
	if t == 1 {
		return &HelloAPI{}
	} else if t == 2{
		return &WorldAPI{}
	}
	return nil
}

type HelloAPI struct {

}

type WorldAPI struct {

}

func (api *HelloAPI) say(word string) string  {

	return fmt.Sprintf("Hello,%s", word)
}

func (api *WorldAPI) say(word string) string  {

	return fmt.Sprintf("World,%s", word)
}