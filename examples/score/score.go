package main

import (
	"time"

	simpe_mm "github.com/fatihkahveci/simple-matchmaking"
	"github.com/fatihkahveci/simple-matchmaking/rules"
	"github.com/fatihkahveci/simple-matchmaking/server"
	"github.com/fatihkahveci/simple-matchmaking/store"
)

func main() {
	inMemory := store.NewInMemoryStore()
	dur, _ := time.ParseDuration("10s")

	r := rules.NewScoreMatchRule(10, 15)

	respServer := server.NewRespServer(inMemory, ":1234")

	opts := &simpe_mm.Options{
		Name:    "score",
		Store:   inMemory,
		Server:  respServer,
		Rule:    r,
		Timeout: dur,
	}

	matcher := simpe_mm.NewMatchmaking(opts)

	matcher.Start()
}
