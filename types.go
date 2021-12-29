package matchmaker

import (
	"encoding/json"
	"github.com/fatihkahveci/simple-matchmaking/store"
	"time"
)

type MatchFinishResponse struct {
	User1         store.User `json:"user_1"`
	User2         store.User `json:"user_2"`
	MatchRuleName string     `json:"match_rule_name"`
	Time          time.Time  `json:"time"`
	ActionType    string     `json:"action_type"`
}

type TimeOutResponse struct {
	User          store.User `json:"user"`
	Time          time.Time  `json:"time"`
	ActionType    string     `json:"action_type"`
	MatchRuleName string     `json:"match_rule_name"`
}

type ErrorResponse struct {
	ActionType string `json:"action_type"`
	Error      error  `json:"error"`
}

func NewMatchFinishResponse(matchRule string, user1, user2 store.User) MatchFinishResponse {
	return MatchFinishResponse{
		Time:          time.Now(),
		MatchRuleName: matchRule,
		User1:         user1,
		User2:         user2,
		ActionType:    "match",
	}
}

func NewMatchTimeOutResponse(matchRule string, user store.User) ([]byte, error) {
	timeoutResponse := TimeOutResponse{
		Time:          time.Now(),
		MatchRuleName: matchRule,
		ActionType:    "timeout",
		User:          user,
	}

	return json.Marshal(timeoutResponse)
}

func NewErrorResponse(err error) ([]byte, error) {
	errorResponse := ErrorResponse{
		ActionType: "error",
		Error:      err,
	}

	return json.Marshal(errorResponse)
}
