package game

import (
	"encoding/json"
	"errors"
	"fmt"
)

type CenterClient struct {
	*IPCClient
}

func (cc *CenterClient) AddPlayer(player *GamePlayer) error {
	fmt.Println("===============>>>", *player)
	b, err := json.Marshal(*player)
	if err != nil {
		fmt.Println("===", err)
		return err
	}
	resp, err := cc.Call("add", string(b))
	if err == nil && resp.Code == "200" {
		return nil
	}
	return err
}

func (cc *CenterClient) RemovePlayer(name string) error {

	resp, _ := cc.Call("remove", name)
	if resp.Code == "200" {
		return nil
	}
	return errors.New(resp.Code)
}

func (cc *CenterClient) ListPlayer() (ps []*GamePlayer, err error) {
	resp, _ := cc.Call("list", "")
	if resp.Code != "200" {
		err = errors.New(resp.Code)
		return
	}

	err = json.Unmarshal([]byte(resp.Body), &ps)
	return
}

func (cc *CenterClient) Broadcast(message string) error {

	msg := &GameMessage{
		Content: message,
	}

	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	resp, _ := cc.Call("broadcast", string(b))
	if resp.Code == "200" {
		return nil
	}
	return errors.New(resp.Code)

}
