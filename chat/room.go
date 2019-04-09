package chat

import (
	"fmt"
	"log"
)

type Room struct {
	Server  *ChatServer
	Name    string
	Clients map[string]*Client
	In      chan *Message
	Quit    chan bool
}

func kickName(message *Message) string {
	return fmt.Sprintf("%s", message.content)
}

func (r *Room) Listen() {

	log.Printf("chat room %s opened", r.Name)

	for {
		select {

		case msg := <-r.In:
			switch msg.command {
			case QUIT:
				delete(r.Clients, msg.sender.Name)

				go r.broadcast(msg)
			case JOIN:
				log.Printf("%s joined\n", msg.sender.Name)
				r.Clients[msg.sender.Name] = msg.sender
				go r.broadcast(msg)
			case KICK:
				name := kickName(msg)
				if _, ok := r.Clients[name]; ok {
					delete(r.Clients, name)
					go r.broadcast(msg)
				}

			case DISMISS:
				go r.broadcast(msg)
				r.Quit <- true

			default:
				go r.broadcast(msg)
			}

		case <-r.Quit:

			delete(r.Server.Rooms, r.Name)
			for k := range r.Clients {
				delete(r.Clients, k)
			}
			log.Printf("chat room %s closed\n", r.Name)
			return

		}
	}
}

func (r *Room) broadcast(message *Message) {

	for _, c := range r.Clients {
		c.In <- message
	}
}
