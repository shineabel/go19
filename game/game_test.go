package game

import (
	"encoding/json"
	"fmt"
	"testing"
)

type T struct {
	//mq chan string
	name string
}

func TestCenterClient_AddPlayer(t *testing.T) {

	//player := &GamePlayer{
	//	Name:"tom",
	//	Level:1,
	//	Exp:1,
	//	//Mq:make(chan *GameMessage, 1024),
	//
	//}
	t1 := &T{
		name: "hi",
		//mq:make(chan  string, 1024),

	}
	b, err := json.Marshal(t1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
}
