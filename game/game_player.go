package game

import "fmt"

type GamePlayer struct {
	Name string
	Level int
	Exp int
	Mq chan  GameMessage
}

func NewGamePlayer() *GamePlayer  {
	m := make(chan GameMessage, 1024)

	p:= &GamePlayer{
		Name:"",
		Level:0,
		Exp:0,
		Mq:m,
	}
	go func(p *GamePlayer) {

		for {
			me := <- p.Mq
			fmt.Printf(p.Name , " receive msg:", me.Content)
		}
	}(p)
	return p
}


type GameMessage struct {
	Content string `json:"message"`
	From string `json:"from"`
	To string `json:"to"`
}