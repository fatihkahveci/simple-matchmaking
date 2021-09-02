package server

import "github.com/fatihkahveci/simple-matchmaking/store"

type Server interface {
	Publish(channelName string, response []byte)
	AddUser(user store.User)
	RemoveUser(user store.User)
}
