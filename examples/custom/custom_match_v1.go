package main

import (
	"time"

	simpe_mm "github.com/fatihkahveci/simple-matchmaking"
	"github.com/fatihkahveci/simple-matchmaking/server"
	"github.com/fatihkahveci/simple-matchmaking/store"
)

type CustomFieldMatchRule struct {
	Field        string
	MinThreshold int
	MaxThreshold int
}

func NewCustomFieldMatchRule(field string, minThreshold, maxThreshold int) CustomFieldMatchRule {
	return CustomFieldMatchRule{
		Field:        field,
		MinThreshold: minThreshold,
		MaxThreshold: maxThreshold,
	}
}

func (r CustomFieldMatchRule) Match(user1, user2 store.User) bool {
	user1Level := user1.Fields[r.Field].(int)
	user2Level := user2.Fields[r.Field].(int)

	minLevel := user1Level - r.MinThreshold
	maxLevel := user1Level + r.MaxThreshold

	if user2Level >= minLevel && user2Level <= maxLevel {
		return true
	}

	return false
}

func (r CustomFieldMatchRule) GetName() string {
	return "CustomField"
}

func main() {
	inMemory := store.NewInMemoryStore()
	dur, _ := time.ParseDuration("10s")

	r := NewCustomFieldMatchRule("level", 10, 20)

	respServer := server.NewRespServer(inMemory, ":1234")

	opts := &simpe_mm.Options{
		Name:    "custom",
		Store:   inMemory,
		Server:  respServer,
		Rule:    r,
		Timeout: dur,
	}

	matcher := simpe_mm.NewMatchmaking(opts)

	matcher.Start()
}
