package postgres

import "github.com/jackc/pgx/v5/pgxpool"

type KnifePhotoRepository struct {
	db *pgxpool.Pool
}

func NewKnifePhotoRepository(db *pgxpool.Pool) *KnifePhotoRepository {
	return &KnifePhotoRepository{db: db}
}
