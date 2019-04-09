package chat

import "time"

const (
	_ = iota
	NORMAL
	JOIN
	DISMISS
	QUIT
	KICK
	PAUSE
)

type Message struct {
	sender   *Client
	time     time.Time
	receiver string
	command  int
	content  interface{}
}
