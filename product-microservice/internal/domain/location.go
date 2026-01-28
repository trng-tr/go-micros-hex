package domain

import "time"

type Location struct {
	ID          int64
	Ville       string
	Description *string //not mendatory
	CreatedAt   time.Time
	UpdatedAt   *time.Time //not mendatory
}
