package main

import (
	"bufio"
	"fmt"
	"github.com/go19/game"
	"os"
	"strconv"
	"strings"
)

var cc *game.CenterClient

func startCenterService() {
	cs := &game.CenterServer{}

	server := game.NewIPCServer(cs)

	c := game.NewIPCClient(server)

	cc = &game.CenterClient{
		c,
	}
}

func GetCmdHandlers() map[string]func(args []string) int {
	return map[string]func([]string) int{
		"help":   help,
		"quit":   quit,
		"login":  login,
		"logout": logout,
		"list":   list,
		"send":   send,
	}
}

func help(args []string) int {
	fmt.Println(`
cmd:
login <name><level><exp>
logout <name>
send <msg>
list
quit
help


`)
	return 0
}

func quit(args []string) int {
	return 0
}

func list(args []string) int {
	ps, err := cc.ListPlayer()
	if err != nil {
		return 1
	}
	if len(ps) == 0 {
		fmt.Println("no online player...")
		return 0
	}
	for i, v := range ps {
		fmt.Println("index:", i+1, " , value:", v)
	}
	return 0
}

func logout(args []string) int {
	if len(args) != 2 {
		fmt.Println("usage:logout <name>")
		return 1
	}
	cc.RemovePlayer(args[1])
	return 0
}

func login(args []string) int {
	if len(args) != 4 {
		fmt.Println("usage:login <name><level><exp>")
		return 1
	}
	level, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("level should be a integer")
		return 1
	}

	exp, err2 := strconv.Atoi(args[3])
	if err2 != nil {
		fmt.Println("exp should be a integer")
		return 1
	}
	player := &game.GamePlayer{
		Name:  args[1],
		Level: level,
		Exp:   exp,
	}

	err = cc.AddPlayer(player)
	if err != nil {
		fmt.Println("failed to add player", err)
		return 1
	}
	return 0
}

func send(args []string) int {

	message := strings.Join(args[1:], " ")

	err := cc.Broadcast(message)
	if err != nil {
		fmt.Println("broadcast error", err)
		return 1
	}
	return 0
}

func main() {
	fmt.Println("Game server solution")
	startCenterService()
	r := bufio.NewReader(os.Stdin)

	handlers := GetCmdHandlers()
	for {
		b, _, _ := r.ReadLine()
		line := string(b)
		tokens := strings.Split(line, " ")
		fmt.Println(tokens)

		if handler, ok := handlers[tokens[0]]; ok {
			ret := handler(tokens)
			if ret != 0 {
				break
			}
		}
	}
}
