package repository

import (
	"context"
)

const deleteQuery = `
UPDATE knives
SET deleted_at = NOW()
WHERE id = $1 AND deleted_at IS NULL`

func (r *KnifeRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Exec(ctx, deleteQuery, id)
	return err
}
