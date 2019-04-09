package chat

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

type Client struct {
	Server *ChatServer
	Name   string
	Conn   net.Conn
	In     chan *Message
	Out    chan *Message
	Quit   chan bool
	Rooms  map[string]*Room
}

func (c *Client) Listen() {

	fmt.Sprintf("new client %s joined\n", c.Name)

	for {

		select {
		case msg := <-c.In:
			go c.Write(msg)

		case msg := <-c.Out:

			switch msg.command {

			case QUIT:
				for _, r := range c.Rooms {
					r.In <- msg
				}
				c.Quit <- true
			case JOIN:
				name := msg.receiver
				var r *Room
				if _, ok := c.Server.Rooms[name]; !ok {

					r = &Room{
						Server:  c.Server,
						Name:    name,
						In:      make(chan *Message),
						Clients: make(map[string]*Client, 0),
						Quit:    make(chan bool),
					}
					c.Server.Rooms[name] = r
					go r.Listen()

				}

				c.Rooms[name] = c.Server.Rooms[name]
				c.Rooms[name].In <- msg

			default:

				c.Rooms[msg.receiver].In <- msg

			}

		case <-c.Quit:
			return
		}
	}
}

func (c *Client) Write(msg *Message) {
	s := fmt.Sprintf("%s %s:%s\n", msg.time.Format(time.RFC3339), msg.sender.Name, msg.content)

	c.Conn.Write([]byte(s))
}

func (c *Client) Receive() {

	buf := bufio.NewReader(c.Conn)
	var msg *Message
	for {
		line, err := buf.ReadString('\n')
		if err != nil || len(line) == 0 {

			if err == io.EOF || len(line) == 0 {
				log.Printf(" %s remote closed\n", c.Name)
				msg = &Message{

					time:     time.Now(),
					command:  QUIT,
					content:  fmt.Sprintf("%s lefted\n", c.Name),
					sender:   c,
					receiver: "",
				}
			} else {
				log.Printf("%s lefted\n", c.Conn.RemoteAddr())
				msg = &Message{

					time:     time.Now(),
					command:  QUIT,
					content:  fmt.Sprintf("%s disconnected\n", c.Name),
					sender:   c,
					receiver: "",
				}
			}

			c.Out <- msg
			break
		} else {

			data := strings.Split(strings.TrimSpace(line), " ")

			room, content := data[0], data[1]

			if _, ok := c.Rooms[room]; ok {
				msg = &Message{
					time:     time.Now(),
					command:  NORMAL,
					content:  content,
					sender:   c,
					receiver: room,
				}
			} else {
				msg = &Message{
					time:     time.Now(),
					content:  content,
					command:  JOIN,
					sender:   c,
					receiver: room,
				}
			}
		}
		c.Out <- msg
	}
}
