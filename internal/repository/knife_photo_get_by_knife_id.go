package repository

import (
	"context"

	"knives/internal/domain"

	"github.com/georgysavva/scany/v2/pgxscan"
)

const getPhotosByKnifeIDQuery = `
SELECT 
	id,
	knife_id,                                                                                                                                                                                                                                   
	s3_key,                                                                                                                                                                                                                                     
	filename,                                                                                                                                                                                                                                   
	created_at,                                                                                                                                                                                                                                 
	deleted_at
FROM knife_photos
WHERE knife_id = $1 AND deleted_at IS NULL`

func (r *KnifePhotoRepository) GetByKnifeID(ctx context.Context, knifeID string) ([]domain.KnifePhoto, error) {
	var rows []knifePhotoRow
	if err := pgxscan.Select(ctx, r.db, &rows, getPhotosByKnifeIDQuery, knifeID); err != nil {
		return nil, err
	}
	return knifePhotoRowsToDomain(rows), nil
}
