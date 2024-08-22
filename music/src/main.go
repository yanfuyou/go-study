package main

import (
	"bufio"
	"fmt"
	"go-study/music/library"
	"go-study/music/src/entity"
	"go-study/music/src/mp"
	"os"
	"strconv"
	"strings"
)

var lib *library.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}
	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&entity.MusicEntity{Id: strconv.Itoa(id), Name: tokens[2], Artist: tokens[3],
				Source: tokens[4], Type: tokens[5]})
		} else {
			fmt.Println("USAGE: lib add <name><artist><source><type>")
		}
	case "remove":
		if len(tokens) == 3 {
			idx, _ := strconv.Atoi(tokens[2])
			lib.Remove(idx)
			// 缺一个按名称移除的方法
		} else {
			fmt.Println("USAGE: lib remove <id>")
		}
	default:
		fmt.Println("error command!!!,", tokens[1])
	}
}
func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}
	m := lib.Find(tokens[2])
	if m == nil {
		fmt.Println("music ", tokens[2], " not found")
		return
	}
	mp.Play(m.Source, m.Type)
}

func main() {
	lib = library.NewMusicManager()
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command -> ")
		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)
		if line == "q" || line == "e" {
			break
		}
		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)
		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("error command!!!", tokens[0])
		}
	}
}
