package store

import "l0/service/models"

type Repository struct {
	database *Database
}

func (r *Repository) Create(o *models.Order) error {
	if err := r.database.db.QueryRow(
		"INSERT INTO orders (order_uid, data) VALUES ($1, $2) RETURNING order_uid", o.Uid, o.Data).
		Scan(&o.Uid); err != nil {
		return err
	}
	return nil
}

func (r *Repository) Find(uid string) (*models.Order, error) {
	o := &models.Order{}
	if err := r.database.db.QueryRow(
		"SELECT order_uid, data FROM orders WHERE order_uid = $1", uid).
		Scan(&o.Uid, &o.Data); err != nil {
		return nil, err
	}
	return o, nil
}

func (r *Repository) FindAll() ([]models.Order, error) {
	result := make([]models.Order, 0)
	rows, err := r.database.db.Query("SELECT order_uid, data FROM orders")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		o := models.Order{}
		if err := rows.Scan(&o.Uid, &o.Data); err != nil {
			return nil, err
		}
		result = append(result, o)
	}
	return result, nil
}
