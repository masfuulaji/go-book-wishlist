package response

import "time"

type DirectorResponse struct {
	Birthday     time.Time `json:"birthday"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	Name         string    `json:"name"`
	PlaceOfBirth string    `db:"place_of_birth" json:"place_of_birth"`
	ID           int       `json:"id"`
}
