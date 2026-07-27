package repository

import (
	"context"
	"knives/internal/models"

	"github.com/georgysavva/scany/v2/pgxscan"
)

const getAllQuery = `
SELECT
	id,
	name,
	description,
	price,
	material,
	blade_length,
	handle,
	brand,
	created_at,
	updated_at,
	deleted_at
FROM knives
WHERE deleted_at IS NULL`

func (r *KnifeRepository) GetAll(ctx context.Context) ([]models.Knife, error) {
	var knives []models.Knife
	if err := pgxscan.Select(ctx, r.db, &knives, getAllQuery); err != nil {
		return nil, err
	}
	return knives, nil
}
