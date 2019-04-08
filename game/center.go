package game

import (
	"sync"
	"encoding/json"
	"fmt"
	"errors"
)

type CenterServer struct {
	servers map[string] Server
	players []*GamePlayer
	room []*GameRoom
	mutex sync.Mutex

}

func NewCenterServer() *CenterServer  {

	return &CenterServer{
		servers:make(map[string] Server),
		players:make([]*GamePlayer, 0),
	}
}

func (cs *CenterServer) AddPlayer(params string) error {
	player := NewGamePlayer()
	err := json.Unmarshal([]byte(params), &player)
	if err != nil {
		fmt.Println("params unmarshal error", err)
		return err
	}

	cs.mutex.Lock()
	defer cs.mutex.Unlock()
	cs.players = append(cs.players,player)
	return nil
}

func (cs *CenterServer) RemovePlayer(params string) error  {


	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	for index , p := range  cs.players {
		if p.Name == params {
			if len(cs.players) == 1 {
				cs.players = make([]*GamePlayer, 0)

			} else {

				if index == len(cs.players)-1 {

					cs.players = cs.players[: index]
				} else if index == 0 {
					cs.players = cs.players[ 1 :]
				} else {
					cs.players = append(cs.players[:index - 1],cs.players[index + 1:]...)
				}
			}
		}
	}
	return errors.New("not found player")
}

func (cs *CenterServer) ListPlayer() (players string, err error)  {
	cs.mutex.Lock()

	defer cs.mutex.Unlock()
	if len(cs.players) > 0 {

		b , _ := json.Marshal(cs.players)
		players = string(b)

	} else {
		err = errors.New("no player online")
	}

	return
}

func (cs *CenterServer) broadcast(msg string) error  {



	cs.mutex.Lock()
	defer cs.mutex.Unlock()

	var message GameMessage
	err := json.Unmarshal([]byte(msg), &message)
	if err != nil {
		return err
	}

	if len(cs.players) > 0 {
		for _, p := range  cs.players {
			p.Mq <- message
		}

	} else {
		err = errors.New("no player online")
	}
	return err
}

func (cs *CenterServer) Handle(method , params string) *Response  {

	switch method {

	case "add":
		err := cs.AddPlayer(params)
		if err != nil {
			return &Response{
				Code:err.Error(),
			}
		}
	case "remove":
		err := cs.RemovePlayer(params)
		if err != nil {
			return &Response{
				Code:err.Error(),
			}
		}
	case "list":
		players , err := cs.ListPlayer()
		if err != nil {
			return &Response{
				Code:err.Error(),
			}
		}
		return &Response{
			Code:"200",
			Body:players,

		}
	case "broadcast":
		err := cs.broadcast(params)
		if err != nil {
			return &Response{
				Code:err.Error(),
			}
		}
		return &Response{
			Code:"200",
		}

	default:
		return &Response{
			Code:"404",
			Body:method + "-" + params,
		}
	}
	return &Response{
		Code:"200",
	}
}


func (cs *CenterServer) Name() string  {

	return "centerserver"
}

type GameRoom struct {

}
