package repository

import (
	"context"
	"knives/internal/models"

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

func (r *KnifePhotoRepository) GetByKnifeID(ctx context.Context, knifeID string) ([]models.KnifePhoto, error) {
	var photos []models.KnifePhoto
	if err := pgxscan.Select(ctx, r.db, &photos, getPhotosByKnifeIDQuery, knifeID); err != nil {
		return nil, err
	}
	return photos, nil
}
