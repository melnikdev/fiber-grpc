package model

import (
	"database/sql"
	"time"
)

type Event struct {
	ID          uint
	Name        string
	Active      bool
	ActivatedAt sql.NullTime
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
