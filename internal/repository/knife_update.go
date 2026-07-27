package repository

import (
	"context"
	"knives/internal/models"
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
WHERE id = $8 AND deleted_at IS NULL`

func (r *KnifeRepository) Update(ctx context.Context, knife *models.Knife) error {
	_, err := r.db.Exec(ctx, updateQuery,
		knife.Name, knife.Description, knife.Price, knife.Material,
		knife.BladeLength, knife.Handle, knife.Brand, knife.ID)
	return err
}
