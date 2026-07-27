package repository

import (
	"cart_api/internal/errorsx"
	"context"
)

func (r *Repository) RemoveCartItem(ctx context.Context, cartID int, itemID int) error {

	query := "DELETE FROM cart_items WHERE cart_id = $1 AND id = $2"
	res, err := r.db.ExecContext(ctx, query, cartID, itemID)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errorsx.ErrItemNotFound

	}

	return nil
}
