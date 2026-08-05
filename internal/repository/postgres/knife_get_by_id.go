package postgres

import (
	"context"
	"errors"
	"knives/internal/domain"
	apperrors "knives/internal/errors"

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

func (r *KnifeRepository) GetByID(ctx context.Context, id string) (*domain.Knife, error) {
	var row knifeRow
	if err := pgxscan.Get(ctx, r.db, &row, getByIDQuery, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, apperrors.ErrNotFound
		}
		return nil, err
	}
	return row.toDomain(), nil
}
