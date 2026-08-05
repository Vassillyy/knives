package postgres

import (
	"context"
	"knives/internal/domain"
)

const createQuery = `
INSERT INTO knives (
	id,
	name,
	description,
	price,
	material,
	blade_length,
	handle,
	brand
	)
VALUES (gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7)
RETURNING id, created_at, updated_at`

func (r *KnifeRepository) Create(ctx context.Context, knife *domain.Knife) error {
	return r.db.QueryRow(ctx, createQuery,
		knife.Name, knife.Description, knife.Price, knife.Material,
		knife.BladeLength, knife.Handle, knife.Brand).
		Scan(&knife.ID, &knife.CreatedAt, &knife.UpdatedAt)
}
