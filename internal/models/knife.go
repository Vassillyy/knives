package models

import "time"

type Knife struct {
	ID          string     `db:"id"          json:"id"`
	Name        *string    `db:"name"        json:"name,omitempty"`
	Description *string    `db:"description" json:"description,omitempty"`
	Price       *int       `db:"price"       json:"price,omitempty"`
	Material    *string    `db:"material"    json:"material,omitempty"`
	BladeLength *float64   `db:"blade_length" json:"blade_length,omitempty"`
	Handle      *string    `db:"handle"      json:"handle,omitempty"`
	Brand       *string    `db:"brand"       json:"brand,omitempty"`
	CreatedAt   time.Time  `db:"created_at"  json:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"  json:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"  json:"-"`
}
