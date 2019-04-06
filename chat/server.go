package chat

import (
	"time"
	"fmt"
	"net"
	"log"
)

type ChatServer struct {

	Rooms map[string]*Room
	Bind2 string
}

func (s *ChatServer) reportServerStatus()  {

	for {
		time.Sleep(5 * time.Second)
		for _,r := range s.Rooms {
			fmt.Printf(" room name:%s , client count:%d\n",r.Name, len(r.Clients))
		}
	}
}

func (s *ChatServer)  ListenAndServe()  {

	listener, err := net.Listen("tcp",s.Bind2)
	if err != nil {
		log.Fatal("serve listen error:",err)
	}


	defer listener.Close()
	go s.reportServerStatus()


	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("listener accept error", err)
		}

		go func(conn net.Conn, s *ChatServer) {

			c := &Client{

				Server:s,
				Name:fmt.Sprintf("%s", conn.RemoteAddr()),
				Conn:conn,
				In:make(chan  *Message),
				Out:make(chan  *Message),
				Quit:make(chan bool),
				Rooms: map[string]*Room{},
			}

			go c.Listen()
			go c.Receive()

		}(conn, s)
	}
}