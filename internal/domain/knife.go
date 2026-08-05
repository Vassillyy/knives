package domain

import "time"

type Knife struct {
	ID          string
	Name        *string
	Description *string
	Price       *int
	Material    *string
	BladeLength *float64
	Handle      *string
	Brand       *string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}