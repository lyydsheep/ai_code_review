package event

import "time"

type Push struct {
	Username   string    `json:"username"`
	Repository string    `json:"repository"`
	Diff       string    `json:"diff"`
	Time       time.Time `json:"time"`
}
