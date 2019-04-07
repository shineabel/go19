package music

import (
	"fmt"
	"time"
)

type MusicPlayer interface {

	play(source string)
}


type Mp3Player struct {
	start int
	progress int

}

type WAVPlayer struct {

}

func (mp3 *Mp3Player) play(source string)  {

	fmt.Println("start to play mp3...")
	mp3.progress = 0
	for mp3.progress < 100{
		time.Sleep(500 * time.Millisecond)
		fmt.Println(".")
		mp3.progress += 10
	}

	fmt.Printf("play mp3 end....")
}

func (wav *WAVPlayer) play(source string)  {

}





func PlayMusic( source, t string)  {
	var  player MusicPlayer

	switch t {

		case "mp3":
			player = &Mp3Player{}
	case "wav":
		player = &WAVPlayer{}
	default:
		fmt.Printf("not support music type")
		return
	}

	player.play(source)


}
