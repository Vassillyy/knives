package dto

import (
	"time"

	"knives/internal/domain"
)

type KnifeRequest struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Price       *int     `json:"price"`
	Material    *string  `json:"material"`
	BladeLength *float64 `json:"blade_length"`
	Handle      *string  `json:"handle"`
	Brand       *string  `json:"brand"`
}

func (r KnifeRequest) ToDomain() *domain.Knife {
	return &domain.Knife{
		Name:        r.Name,
		Description: r.Description,
		Price:       r.Price,
		Material:    r.Material,
		BladeLength: r.BladeLength,
		Handle:      r.Handle,
		Brand:       r.Brand,
	}
}

type KnifeResponse struct {
	ID          string    `json:"id"`
	Name        *string   `json:"name,omitempty"`
	Description *string   `json:"description,omitempty"`
	Price       *int      `json:"price,omitempty"`
	Material    *string   `json:"material,omitempty"`
	BladeLength *float64  `json:"blade_length,omitempty"`
	Handle      *string   `json:"handle,omitempty"`
	Brand       *string   `json:"brand,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func KnifeFromDomain(k *domain.Knife) KnifeResponse {
	return KnifeResponse{
		ID:          k.ID,
		Name:        k.Name,
		Description: k.Description,
		Price:       k.Price,
		Material:    k.Material,
		BladeLength: k.BladeLength,
		Handle:      k.Handle,
		Brand:       k.Brand,
		CreatedAt:   k.CreatedAt,
		UpdatedAt:   k.UpdatedAt,
	}
}

func KnivesFromDomain(knives []domain.Knife) []KnifeResponse {
	out := make([]KnifeResponse, len(knives))
	for i, k := range knives {
		out[i] = KnifeFromDomain(&k)
	}
	return out
}
