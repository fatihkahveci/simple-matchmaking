package store

import "time"

type User struct {
	ID       string                 `json:"id"`
	Score    float64                `json:"score"`
	JoinTime time.Time              `json:"join_time"`
	Fields   map[string]interface{} `json:"fields"`
}
