package mp

import (
	"fmt"
	"time"
)

type MP3Plyer struct {
	stat     int
	progress int
}

func (p *MP3Plyer) Play(source string) {
	fmt.Println("playing music", source)
	p.progress = 0
	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Print(".")
		p.progress += 10
	}
	fmt.Println("\nplaying done", source)
}
