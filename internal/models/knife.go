package models

import "time"

type Knife struct {
	ID          string     `json:"id"`
	Name        *string    `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	Price       *int       `json:"price,omitempty"`
	Material    *string    `json:"material,omitempty"`
	BladeLength *float64   `json:"blade_length,omitempty"`
	Handle      *string    `json:"handle,omitempty"`
	Brand       *string    `json:"brand,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
