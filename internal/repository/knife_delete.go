package repository

import (
	"context"
	apperrors "knives/internal/errors"
)

const deleteQuery = `
UPDATE knives
SET deleted_at = NOW()
WHERE id = $1 AND deleted_at IS NULL`

func (r *KnifeRepository) Delete(ctx context.Context, id string) error {
	result, err := r.db.Exec(ctx, deleteQuery, id)
	if err != nil {
		return err
	}
	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return apperrors.ErrNotFound
	}

	return nil
}
