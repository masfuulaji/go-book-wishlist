package models

import (
	"time"

	"github.com/lib/pq"
)

type Director struct {
	Birthday     time.Time
	UpdatedAt    time.Time   `db:"updated_at" json:"updated_at"`
	CreatedAt    time.Time   `db:"created_at" json:"created_at"`
	DeletedAt    pq.NullTime `db:"deleted_at" json:"deleted_at"`
	Name         string
	PlaceOfBirth string `db:"place_of_birth"`
	ID           int
}
