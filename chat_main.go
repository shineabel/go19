package main

import (
	"fmt"
	"github.com/go19/chat"
)

func main() {
	fmt.Printf("start chat room")

	server := &chat.ChatServer{

		Bind2: ":12345",
		Rooms: make(map[string]*chat.Room),
	}

	server.ListenAndServe()
}
