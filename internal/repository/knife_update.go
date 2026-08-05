package repository

import (
	"context"
	"errors"
	"knives/internal/domain"
	apperrors "knives/internal/errors"

	"github.com/jackc/pgx/v5"
)

const updateQuery = `
UPDATE knives
SET
  name = COALESCE($1, name),
  description = COALESCE($2, description),
  price = COALESCE($3, price),
  material = COALESCE($4, material),
  blade_length = COALESCE($5, blade_length),
  handle = COALESCE($6, handle),
  brand = COALESCE($7, brand),
  updated_at = NOW()
WHERE id = $8 AND deleted_at IS NULL
RETURNING id, name, description, price, material, blade_length, handle, brand, created_at, updated_at`

func (r *KnifeRepository) Update(ctx context.Context, knife *domain.Knife) (*domain.Knife, error) {
	var updated domain.Knife
	err := r.db.QueryRow(ctx, updateQuery,
		knife.Name, knife.Description, knife.Price, knife.Material,
		knife.BladeLength, knife.Handle, knife.Brand, knife.ID,
	).Scan(
		&updated.ID,
		&updated.Name,
		&updated.Description,
		&updated.Price,
		&updated.Material,
		&updated.BladeLength,
		&updated.Handle,
		&updated.Brand,
		&updated.CreatedAt,
		&updated.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}

	return &updated, nil
}
