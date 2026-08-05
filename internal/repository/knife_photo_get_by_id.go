package repository

import (
	"context"
	"errors"
	"knives/internal/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
)

const getPhotoByIDQuery = `
SELECT
	id,
	knife_id,
	s3_key,
	filename,
	created_at,
	deleted_at
FROM knife_photos
WHERE id = $1 AND deleted_at IS NULL`

func (r *KnifePhotoRepository) GetByID(ctx context.Context, id string) (*domain.KnifePhoto, error) {
	var row knifePhotoRow
	if err := pgxscan.Get(ctx, r.db, &row, getPhotoByIDQuery, id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return row.toDomain(), nil
}