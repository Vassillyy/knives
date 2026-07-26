package models

import "time"

type Knife struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	Price       int        `json:"price"`
	Material    *string    `json:"material,omitempty"`
	BladeLength *float64   `json:"blade_length,omitempty"`
	Handle      *string    `json:"handle,omitempty"`
	Brand       *string    `json:"brand,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

func (k *Knife) ScanFields() []any {
	return []any{
		&k.ID, &k.Name, &k.Description, &k.Price, &k.Material,
		&k.BladeLength, &k.Handle, &k.Brand, &k.CreatedAt, &k.UpdatedAt, &k.DeletedAt,
	}
}

func (k *Knife) baseArgs() []any {
	return []any{k.Name, k.Description, k.Price, k.Material, k.BladeLength, k.Handle, k.Brand}
}

func (k *Knife) CreateArgs() []any {
	return k.baseArgs()
}

func (k *Knife) UpdateArgs() []any {
	return append(k.baseArgs(), k.ID)
}
