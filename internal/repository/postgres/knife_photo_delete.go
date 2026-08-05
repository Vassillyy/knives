package postgres

import (
	"context"
	apperrors "knives/internal/errors"
)

const deletePhotoQuery = `DELETE FROM knife_photos WHERE id = $1`

func (r *KnifePhotoRepository) Delete(ctx context.Context, id string) error {
	result, err := r.db.Exec(ctx, deletePhotoQuery, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return apperrors.ErrNotFound
	}
	return nil
}
