package rules

import "github.com/fatihkahveci/simple-matchmaking/store"

type MatchRule interface {
	Match(user1, user2 store.User) bool
	GetName() string
}
