package models

import (
	"time"

	"github.com/lib/pq"
)

type Category struct {
	CreatedAt   time.Time   `db:"created_at"`
	UpdatedAt   time.Time   `db:"updated_at"`
	DeletedAt   pq.NullTime `db:"deleted_at"`
	Name        string
	Description string
	ID          int
}
