package simpe_mm_test

import (
	"testing"
	"time"

	simpe_mm "github.com/fatihkahveci/simple-matchmaking"
	"github.com/fatihkahveci/simple-matchmaking/rules"
	"github.com/fatihkahveci/simple-matchmaking/server"
	"github.com/fatihkahveci/simple-matchmaking/store"
	"github.com/stretchr/testify/assert"
)

func TestDirectMatchRule(t *testing.T) {
	duration, err := time.ParseDuration("5s")
	inMemory := store.NewInMemoryStore()
	respServer := server.NewRespServer(inMemory, ":1111")

	rule := rules.NewDirectMatchRule()
	if err != nil {
		t.Error(err)
	}

	opts := &simpe_mm.Options{
		Name:    "test",
		Store:   inMemory,
		Server:  respServer,
		Rule:    rule,
		Timeout: duration,
	}

	mm := simpe_mm.NewMatchmaking(opts)

	u1 := store.User{
		ID:    "1",
		Score: 11,
		Fields: map[string]interface{}{
			"level": 10,
		},
	}
	u2 := store.User{
		ID:    "2",
		Score: 4,
		Fields: map[string]interface{}{
			"level": 10,
		},
	}
	u3 := store.User{
		ID:    "3",
		Score: 3,
		Fields: map[string]interface{}{
			"level": 10,
		},
	}

	mm.AddUser(u1)
	mm.AddUser(u2)
	mm.AddUser(u3)

	assert.Equal(t, "Direct", rule.GetName())

	mm.RunLoop()

}

func TestScoreMatchRule(t *testing.T) {

	duration, err := time.ParseDuration("5s")

	inMemory := store.NewInMemoryStore()
	respServer := server.NewRespServer(inMemory, ":2222")

	rule := rules.NewScoreMatchRule(5, 5)
	if err != nil {
		t.Error(err)
	}

	opts := &simpe_mm.Options{
		Name:    "test",
		Store:   inMemory,
		Server:  respServer,
		Rule:    rule,
		Timeout: duration,
	}

	mm := simpe_mm.NewMatchmaking(opts)

	u1 := store.User{
		ID:    "1",
		Score: 11,
		Fields: map[string]interface{}{
			"level": 10,
		},
	}
	u2 := store.User{
		ID:    "2",
		Score: 4,
		Fields: map[string]interface{}{
			"level": 10,
		},
	}
	u3 := store.User{
		ID:    "3",
		Score: 3,
		Fields: map[string]interface{}{
			"level": 10,
		},
	}

	mm.AddUser(u1)
	mm.AddUser(u2)
	mm.AddUser(u3)

	assert.Equal(t, "Score", rule.GetName())

	mm.RunLoop()
}

func TestIsExtendTime(t *testing.T) {
	duration, err := time.ParseDuration("1s")
	inMemory := store.NewInMemoryStore()
	respServer := server.NewRespServer(inMemory, ":4444")

	rule := rules.NewDirectMatchRule()
	if err != nil {
		t.Error(err)
	}

	opts := &simpe_mm.Options{
		Name:    "test",
		Store:   inMemory,
		Server:  respServer,
		Rule:    rule,
		Timeout: duration,
	}

	mm := simpe_mm.NewMatchmaking(opts)

	u1 := store.User{
		ID:    "1",
		Score: 11,
		Fields: map[string]interface{}{
			"level": 10,
		},
	}

	mm.AddUser(u1)

	assert.Equal(t, "Direct", rule.GetName())

	go mm.Start()

	time.Sleep(time.Second * 2)
}
