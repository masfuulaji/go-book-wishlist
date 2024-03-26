package models

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Movie struct {
	CreatedAt   time.Time   `db:"created_at"`
	UpdatedAt   time.Time   `db:"updated_at"`
	DeletedAt   pq.NullTime `db:"deleted_at"`
	Title       string
	Description string
	Year        string
	DirectorID  int `db:"director_id"`
	ID          int
	Rating      sql.NullFloat64
}
