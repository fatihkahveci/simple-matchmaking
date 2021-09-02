package server

import (
	"encoding/json"
	"github.com/fatihkahveci/simple-matchmaking/store"
	"github.com/tidwall/redcon"
	"strings"
)

type RespServer struct {
	PubSub *redcon.PubSub
	store  store.Store
}

func NewRespServer(storage store.Store, addr string) RespServer {
	var ps redcon.PubSub
	handler := func(conn redcon.Conn, cmd redcon.Command) {
		switch strings.ToLower(string(cmd.Args[0])) {
		default:
			conn.WriteError("ERR unknown command '" + string(cmd.Args[0]) + "'")
		case "ping":
			conn.WriteString("PONG")
		case "add":
			var user store.User
			err := json.Unmarshal(cmd.Args[1], &user)
			if err != nil {
				conn.WriteError(err.Error())
				return
			}
			storage.Add(user)
		case "remove":
			var user store.User
			err := json.Unmarshal(cmd.Args[1], &user)
			if err != nil {
				conn.WriteError(err.Error())
				return
			}
			storage.Remove(user)
		case "subscribe", "psubscribe":
			if len(cmd.Args) < 2 {
				conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
				return
			}
			command := strings.ToLower(string(cmd.Args[0]))
			for i := 1; i < len(cmd.Args); i++ {
				if command == "psubscribe" {
					ps.Psubscribe(conn, string(cmd.Args[i]))
				} else {
					ps.Subscribe(conn, string(cmd.Args[i]))
				}
			}
		}
	}

	go func() {
		err := redcon.ListenAndServe(addr, handler,

			func(conn redcon.Conn) bool {
				// Use this function to accept or deny the connection.
				// log.Printf("accept: %s", conn.RemoteAddr())
				return true
			},
			func(conn redcon.Conn, err error) {
				// This is called when the connection has been closed
				// log.Printf("closed: %s, err: %v", conn.RemoteAddr(), err)
			},
		)

		if err != nil {
			panic(err)
		}
	}()

	return RespServer{
		PubSub: &ps,
		store:  storage,
	}

}

func (r RespServer) Publish(channelName string, response []byte) {
	r.PubSub.Publish(channelName, string(response))
}

func (r RespServer) AddUser(user store.User) {
	r.store.Add(user)
}

func (r RespServer) RemoveUser(user store.User) {
	r.store.Remove(user)
}
