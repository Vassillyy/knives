package postgres

import (
	"time"

	"knives/internal/domain"
)

type knifeRow struct {
	ID          string     `db:"id"`
	Name        *string    `db:"name"`
	Description *string    `db:"description"`
	Price       *int       `db:"price"`
	Material    *string    `db:"material"`
	BladeLength *float64   `db:"blade_length"`
	Handle      *string    `db:"handle"`
	Brand       *string    `db:"brand"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
	DeletedAt   *time.Time `db:"deleted_at"`
}

func (r knifeRow) toDomain() *domain.Knife {
	return &domain.Knife{
		ID:          r.ID,
		Name:        r.Name,
		Description: r.Description,
		Price:       r.Price,
		Material:    r.Material,
		BladeLength: r.BladeLength,
		Handle:      r.Handle,
		Brand:       r.Brand,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
		DeletedAt:   r.DeletedAt,
	}
}

func knifeRowsToDomain(rows []knifeRow) []domain.Knife {
	out := make([]domain.Knife, len(rows))
	for i, r := range rows {
		out[i] = *r.toDomain()
	}
	return out
}