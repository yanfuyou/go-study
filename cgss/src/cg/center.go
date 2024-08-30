package cg

import (
	"encoding/json"
	"errors"
	"go-study/cgss/src/ipc"
	"sync"
)

type Message struct {
	From    string "from"
	To      string "to"
	Content string "content"
}
type CenterServer struct {
	services map[string]ipc.Server
	players  []*Player
	rooms    []*int
	mutex    sync.RWMutex
}

func (server *CenterServer) Name() string {
	//TODO implement me
	panic("implement me")
}

func NewCenterServer() *CenterServer {
	services := make(map[string]ipc.Server)
	players := make([]*Player, 0)
	return &CenterServer{services: services, players: players}
}

func (server *CenterServer) addPlayer(params string) error {
	player := new(Player)
	err := json.Unmarshal([]byte(params), &player)
	if err != nil {
		return err
	}
	server.mutex.Lock()
	defer server.mutex.Unlock()
	// 省略重复登录校验
	server.players = append(server.players, player)
	return nil
}

func (server *CenterServer) removePlayer(params string) error {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	for i, v := range server.players {
		if v.Name == params {
			if len(server.players) == 1 {
				// 仅有一个用户,清空
				server.players = make([]*Player, 0)
			} else if i == len(server.players)-1 {
				server.players = server.players[:i-1]
			} else if i == 0 {
				server.players = server.players[1:]
			} else {
				server.players = append(server.players[:i-1], server.players[:i+1]...)
			}
			return nil
		}
	}
	return errors.New("player not found")
}

func (server *CenterServer) listPlayers(params string) (players string, err error) {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	if len(server.players) > 0 {
		b, _ := json.Marshal(server.players)
		players = string(b)
	} else {
		err = errors.New("no players online")
	}
	return
}

// board 广播
func (server *CenterServer) broadcast(params string) error {
	var message Message
	err := json.Unmarshal([]byte(params), &message)
	if err != nil {
		return err
	}

	server.mutex.Lock()
	defer server.mutex.Unlock()

	if len(server.players) > 0 {
		for _, player := range server.players {
			player.mq <- &message
		}
	} else {
		err = errors.New("no players online")
	}
	return err
}

func (server *CenterServer) Handle(method, params string) *ipc.Response {
	switch method {
	case "addplayer":
		err := server.addPlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	case "removeplayer":
		err := server.removePlayer(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
	case "listplayers":
		players, err := server.listPlayers(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
		return &ipc.Response{Code: "200", Body: players}
	case "broadcast":
		err := server.broadcast(params)
		if err != nil {
			return &ipc.Response{Code: err.Error()}
		}
		return &ipc.Response{Code: "200"}
	default:
		return &ipc.Response{Code: "404", Body: method + ":" + params}
	}
	return &ipc.Response{Code: "200"}
}

func Name() string {
	return "CenterServer"
}
