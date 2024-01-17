package request

import "time"

type Director struct {
	Birthday     time.Time
	Name         string
	PlaceOfBirth string
	ID           int
}
