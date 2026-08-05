package postgres

import (
	"time"

	"knives/internal/domain"
)

type knifePhotoRow struct {
	ID        string     `db:"id"`
	KnifeID   string     `db:"knife_id"`
	S3Key     string     `db:"s3_key"`
	Filename  string     `db:"filename"`
	CreatedAt time.Time  `db:"created_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}

func (r knifePhotoRow) toDomain() *domain.KnifePhoto {
	return &domain.KnifePhoto{
		ID:        r.ID,
		KnifeID:   r.KnifeID,
		S3Key:     r.S3Key,
		Filename:  r.Filename,
		CreatedAt: r.CreatedAt,
		DeletedAt: r.DeletedAt,
	}
}

func knifePhotoRowsToDomain(rows []knifePhotoRow) []domain.KnifePhoto {
	out := make([]domain.KnifePhoto, len(rows))
	for i, r := range rows {
		out[i] = *r.toDomain()
	}
	return out
}