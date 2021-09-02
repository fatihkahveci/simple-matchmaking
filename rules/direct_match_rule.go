package rules

import (
	"github.com/fatihkahveci/simple-matchmaking/store"
)

type DirectMatchRule struct{}

func (r DirectMatchRule) Match(user1, user2 store.User) bool {
	return true
}

func (r DirectMatchRule) GetName() string {
	return "Direct"
}

func NewDirectMatchRule() DirectMatchRule {
	return DirectMatchRule{}
}
