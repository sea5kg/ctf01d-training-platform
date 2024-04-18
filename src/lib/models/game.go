package models

import (
	"time"
)

type Game struct {
	ID          string    `db:"id"`
	StartTime   time.Time `db:"start_time"`
	EndTime     time.Time `db:"end_time"`
	Description string    `db:"description"`
}
