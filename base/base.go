package main

import "fmt"

type user struct {
	Name string
	Email string
}

func (u user )Notify()  {
	fmt.Printf("send user email to:%s<%s>",u.Name,u.Email)
}
func main() {

	bill := user{"Bill","Bill@gmail.com"}
	bill.Notify()

	lily := &user{"Lily","Lily@gmail.com"}
	lily.Notify()

	(*lily).Notify()
}
