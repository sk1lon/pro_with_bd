package tasks

import (
	"time"
)

type Tasks struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool
	Created_at  time.Time
}
