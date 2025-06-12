package event

import "time"

const (
	HighPriority   = "high"
	MiddlePriority = "middle"
	LowPriority    = "low"
)

type Push struct {
	ID         string    `json:"id"`
	Priority   string    `json:"priority"`
	Username   string    `json:"username"`
	Repository string    `json:"repository"`
	Diff       string    `json:"diff"`
	Time       time.Time `json:"time"`
}
