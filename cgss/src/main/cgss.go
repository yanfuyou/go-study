package main

import (
	"bufio"
	"fmt"
	"go-study/cgss/src/cg"
	"go-study/cgss/src/ipc"
	"os"
	"strconv"
	"strings"
)

var centerClient *cg.CenterClient

func startCenterService() error {
	server := ipc.NewIpcServer(&cg.CenterServer{})
	client := ipc.NewIpcClient(server)
	centerClient = &cg.CenterClient{client}
	return nil
}

func Help(args []string) int {
	fmt.Println(`
		Commands:
			login <username><level><exp>
			logout <username>
			send <message>
			listplayer
			quit(q)	
			help(h)
		`)
	return 0
}

func Quit(args []string) int {
	return 1
}

func Logout(args []string) int {
	if len(args) != 2 {
		fmt.Println("USAGE: logout <username>")
		return 0
	}
	centerClient.RemovePlayer(args[1])
	return 0
}

func Login(args []string) int {
	if len(args) != 4 {
		fmt.Println("fail")
		return 0
	}
	level, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("level must be int")
		return 0
	}
	exp, err := strconv.Atoi(args[3])
	if err != nil {
		fmt.Println("exp must be int")
		return 0
	}

	player := cg.NewPlayer()
	player.Name = args[1]
	player.Level = level
	player.Exp = exp
	err = centerClient.AddPlayer(player)
	if err != nil {
		fmt.Println("add Fail")
	}
	return 0
}

func ListPlayers(args []string) int {
	ps, err := centerClient.ListPlayer("")
	if err != nil {
		fmt.Println("list Fail")
	} else {
		for i, v := range ps {
			fmt.Println(i+1, ":", v)
		}
	}
	return 0
}

func Send(args []string) int {
	message := strings.Join(args[1:], "")
	err := centerClient.Broadcast(message)
	if err != nil {
		fmt.Println("send Fail")
	}
	return 0
}

func GetCommandHandles() map[string]func(args []string) int {
	return map[string]func(args []string) int{
		"help":       Help,
		"h":          Help,
		"quit":       Quit,
		"login":      Login,
		"logout":     Logout,
		"listplayer": ListPlayers,
		"send":       Send,
	}
}

func main() {
	fmt.Println("Game Server Solution")
	startCenterService()
	Help(nil)
	r := bufio.NewReader(os.Stdin)
	handlers := GetCommandHandles()
	// 循环读取用户输入
	for {
		fmt.Println("Command> ")
		b, _, _ := r.ReadLine()
		line := string(b)
		tokens := strings.Split(line, " ")
		if handle, ok := handlers[tokens[0]]; ok {
			ret := handle(tokens)
			if ret != 0 {
				break
			}
		} else {
			fmt.Println("Unknown command: ", tokens[0])
		}
	}
}
