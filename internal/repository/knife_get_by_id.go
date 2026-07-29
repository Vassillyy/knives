package repository

import (
	"context"
	"errors"
	"knives/internal/models"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

const getByIDQuery = `
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
WHERE id = $1 AND deleted_at IS NULL`

func (r *KnifeRepository) GetByID(ctx context.Context, id string) (*models.Knife, error) {
	var k models.Knife
	if err := pgxscan.Get(ctx, r.db, &k, getByIDQuery, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &k, nil
}
