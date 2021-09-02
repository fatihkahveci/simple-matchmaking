package main

import (
	"github.com/fatihkahveci/simple-matchmaking"
	"github.com/fatihkahveci/simple-matchmaking/rules"
	"github.com/fatihkahveci/simple-matchmaking/server"
	"github.com/fatihkahveci/simple-matchmaking/store"
	"time"
)

func main() {
	inMemory := store.NewInMemoryStore()
	dur, _ := time.ParseDuration("10s")

	r := rules.NewScoreMatchRule(10, 15)

	respServer := server.NewRespServer(inMemory, ":1234")

	matcher := simpe_mm.NewMatchmaking("score", respServer, inMemory, r, dur)

	matcher.Start()
}
