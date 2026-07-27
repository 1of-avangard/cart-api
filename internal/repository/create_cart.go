package repository

import "context"

func (r *Repository) CreateCart(ctx context.Context) (int, error) {
	var id int
	query := "INSERT INTO carts DEFAULT VALUES RETURNING id"
	err := r.db.QueryRowContext(ctx, query).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
