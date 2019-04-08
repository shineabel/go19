package game

import "fmt"

type GamePlayer struct {
	Name string
	Level int
	Exp int
	mq chan  *GameMessage
}

func NewGamePlayer() *GamePlayer  {
	m := make(chan *GameMessage, 1024)

	p:= &GamePlayer{
		Name:"",
		Level:0,
		Exp:0,
		mq:m,
	}
	go func(p *GamePlayer) {

		for {
			me := <- p.mq
			fmt.Printf("%s receive msg %s\n", p.Name , me.Content)
		}
	}(p)
	return p
}


type GameMessage struct {
	Content string `json:"message"`
	From string `json:"from"`
	To string `json:"to"`
}