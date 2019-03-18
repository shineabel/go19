package main

import "fmt"

type user struct {
	Name string
	Email string
}

type admin struct {
	u user
	level string
}


type notifier interface {
	notify()
}

func (u *user )notify()  {
	fmt.Printf("send user email to:%s<%s>",u.Name,u.Email)
}

func sendNotification(u notifier)  {
	u.notify()
}
func main() {

	ad := admin{
		u:user{
			Name:"test",
			Email:"shine@gmail.com",
		},
		level:"super",
	}

	ad.u.notify()
	sendNotification(&ad.u)

}
