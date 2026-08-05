package repository

import (
	"context"
	"knives/internal/domain"
)

const createPhotoQuery = `
INSERT INTO knife_photos (id, knife_id, s3_key, filename)
VALUES (gen_random_uuid(), $1, $2, $3)
RETURNING id, created_at`

func (r *KnifePhotoRepository) Create(ctx context.Context, photo *domain.KnifePhoto) error {
	return r.db.QueryRow(ctx, createPhotoQuery, photo.KnifeID, photo.S3Key, photo.Filename).
		Scan(&photo.ID, &photo.CreatedAt)
}
