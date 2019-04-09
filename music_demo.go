package main

import (
	"bufio"
	"fmt"
	"github.com/go19/music"
	"os"
	"strings"
)

var m *music.MusicManager

func main() {

	m = music.NewMusicManager()
	r := bufio.NewReader(os.Stdin)

	for {
		rawLine, _, _ := r.ReadLine()

		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")
		fmt.Println("receive cmd: ", tokens)

		fmt.Println("token[0]", tokens[0])
		if tokens[0] == "mge" {
			go handlerManageCmd(tokens)
		} else if tokens[0] == "play" {
			go handlePlayCmd(tokens)
		} else {
			fmt.Printf("not support cmd")
		}
	}
}

func handlerManageCmd(tokens []string) {
	switch tokens[1] {

	case "list":
		for i := 0; i < m.Len(); i++ {
			mu, _ := m.Get(i)
			fmt.Println(mu.Name, mu.Source, mu.Type)

		}
	case "add":
		if len(tokens) == 5 {
			m.Add(&music.Music{
				Name:   tokens[2],
				Source: tokens[3],
				Type:   tokens[4],
			})

		} else {
			fmt.Printf("usage:lib add <name> <source> <type>")
		}

	case "delete":
		if len(tokens) == 3 {

		} else {
			fmt.Println("usage:lib delete <name>")
		}
	default:
		fmt.Println("not support cmd...haha")

	}

}

func handlePlayCmd(tokens []string) {

	if len(tokens) != 2 {
		fmt.Println("usage :play <name>")
		return
	}

	mu := m.Find(tokens[1])
	if mu == nil {
		fmt.Printf("not found music:%s\n", tokens[1])
		return
	}

	music.PlayMusic(mu.Source, mu.Type)
}
