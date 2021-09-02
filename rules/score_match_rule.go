package rules

import (
	"github.com/fatihkahveci/simple-matchmaking/store"
)

type ScoreMatchRule struct {
	MinThreshold float64
	MaxThreshold float64
}

func NewScoreMatchRule(minThreshold, maxThreshold float64) ScoreMatchRule {
	return ScoreMatchRule{
		MinThreshold: minThreshold,
		MaxThreshold: maxThreshold,
	}
}

func (r ScoreMatchRule) Match(user1, user2 store.User) bool {
	minScore := user1.Score - r.MinThreshold
	maxScore := user1.Score + r.MaxThreshold

	if user2.Score >= minScore && user2.Score <= maxScore {
		return true
	}

	return false
}

func (r ScoreMatchRule) GetName() string {
	return "Score"
}
