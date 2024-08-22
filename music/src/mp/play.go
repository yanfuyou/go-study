package mp

import "fmt"

type Player interface {
	Play(source string)
}

func Play(source, mtype string) {
	var p Player
	p = &MP3Plyer{}
	switch mtype {
	case "MP3":
		p = &MP3Plyer{}
	case "WAV":

	default:
		fmt.Println("Unsupported music type")
		return
	}
	p.Play(source)
}
