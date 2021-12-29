package simpe_mm

import (
	"encoding/json"
	"time"

	"github.com/fatihkahveci/simple-matchmaking/rules"
	"github.com/fatihkahveci/simple-matchmaking/server"
	"github.com/fatihkahveci/simple-matchmaking/store"
	"github.com/rs/zerolog/log"
)

type Matchmaking struct {
	Name          string
	Store         store.Store
	Server        server.Server
	Timeout       time.Duration
	Rule          rules.MatchRule
	SearchTimeOut int
}

type Options struct {
	Name    string
	Store   store.Store
	Server  server.Server
	Rule    rules.MatchRule
	Timeout time.Duration
}

func NewMatchmaking(opts *Options) *Matchmaking {
	return &Matchmaking{
		Name:    opts.Name,
		Store:   opts.Store,
		Timeout: opts.Timeout,
		Rule:    opts.Rule,
		Server:  opts.Server,
	}
}

func (m *Matchmaking) Start() {

	log.Info().
		Str("store", m.Store.GetName()).
		Str("rule", m.Rule.GetName()).
		Msg("matchmaking_start")

	for true {
		m.RunLoop()
		time.Sleep(time.Millisecond * 10)
	}

}

func (m *Matchmaking) AddUser(user store.User) {
	m.Store.Add(user)
}

func (m *Matchmaking) RemoveUser(user store.User) {
	m.Store.Remove(user)
}

func (m *Matchmaking) RunLoop() {
	allUsers := m.Store.GetAll()
	if len(allUsers) > 0 {
		for _, user := range allUsers {
			if m.isUserExtendTime(user) {
				m.RemoveUser(user)
				allUsers = m.Store.GetAll()
				timeOutResponse, err := NewMatchTimeOutResponse(m.Rule.GetName(), user)

				if err != nil {
					errorResponse, _ := NewErrorResponse(err)
					m.Server.Publish(m.Name, errorResponse)
				}

				m.Server.Publish(m.Name, timeOutResponse)
				log.Info().
					Str("user_id", user.ID).
					Msg("user_timeout")
			}
			for _, otherUser := range allUsers {
				if m.CanMatch(user, otherUser) {
					m.SendMatch(user, otherUser)
					allUsers = m.Store.GetAll()
					break
				}
			}
		}
	}
}

func (m *Matchmaking) CanMatch(user1, user2 store.User) bool {
	if user1.ID == user2.ID {
		return false
	}
	return m.Rule.Match(user1, user2)
}

func (m *Matchmaking) SendMatch(user1, user2 store.User) {
	m.RemoveUser(user1)
	m.RemoveUser(user2)

	matchResponse := NewMatchFinishResponse(m.Rule.GetName(), user1, user2)

	jsonData, err := json.Marshal(matchResponse)

	if err != nil {
		errorResponse, _ := NewErrorResponse(err)
		m.Server.Publish(m.Name, errorResponse)
	}

	m.Server.Publish(m.Name, jsonData)

	log.Info().
		Str("user1_id", user1.ID).
		Str("user2_id", user2.ID).
		Msg("match")
}

func (m *Matchmaking) isUserExtendTime(user store.User) bool {
	now := time.Now()
	extendTime := user.JoinTime.Add(m.Timeout)
	if now.Unix() > extendTime.Unix() {
		return true
	}

	return false
}
