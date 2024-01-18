package response

import "time"

type MovieResponse struct {
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ID          int       `json:"id"`
	Year        string    `json:"year"`
	DirectorID  int       `json:"director_id"`
}
