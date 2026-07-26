package repository

import (
	"context"
	"knifes/internal/models"
	"knifes/internal/repository/queries"

	"github.com/jackc/pgx/v5/pgxpool"
)

var knifeQueries = queries.MustLoad("knife.sql")

type KnifeRepository struct {
	db *pgxpool.Pool
}

func NewKnifeRepository(db *pgxpool.Pool) *KnifeRepository {
	return &KnifeRepository{db: db}
}

func (r *KnifeRepository) GetAll(ctx context.Context) ([]models.Knife, error) {
	rows, err := r.db.Query(ctx, knifeQueries.Get("GetAll"))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var knives []models.Knife
	for rows.Next() {
		var k models.Knife
		err := rows.Scan(k.ScanFields()...)
		if err != nil {
			return nil, err
		}
		knives = append(knives, k)
	}
	return knives, nil
}

func (r *KnifeRepository) GetByID(ctx context.Context, id string) (*models.Knife, error) {
	var k models.Knife
	err := r.db.QueryRow(ctx, knifeQueries.Get("GetByID"), id).
		Scan(k.ScanFields()...)
	if err != nil {
		return nil, err
	}
	return &k, nil
}

func (r *KnifeRepository) Create(ctx context.Context, knife *models.Knife) error {
	return r.db.QueryRow(ctx, knifeQueries.Get("Create"),
		knife.CreateArgs()...).
		Scan(&knife.ID, &knife.CreatedAt, &knife.UpdatedAt)
}

func (r *KnifeRepository) Update(ctx context.Context, knife *models.Knife) error {
	_, err := r.db.Exec(ctx, knifeQueries.Get("Update"),
		knife.UpdateArgs()...)
	return err
}

func (r *KnifeRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, knifeQueries.Get("Delete"), id)
	return err
}
