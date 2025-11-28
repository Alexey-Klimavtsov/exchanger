package car_repo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/asaipov/gorenda/internal/it/model/car_model"
	"github.com/asaipov/gorenda/internal/service/car_service"
	"time"
)

type CarRepo struct {
	db *sql.DB
}

func NewCarRepo(db *sql.DB) *CarRepo {
	return &CarRepo{db: db}
}

func (r *CarRepo) CreateNewCar(ctx context.Context, c *car_model.CarModel) (*car_model.CarModel, error) {
	now := time.Now().UTC()

	const q = `INSERT INTO cars (brand, model, year, rental_price, image_url, created_at)
	           VALUES (?, ?, ?, ?, ?, ?)`

	var imageUrl = c.ImageUrl

	row, err := r.db.ExecContext(ctx, q,
		c.Brand,
		c.Model,
		c.Year,
		c.RentalPrice,
		imageUrl,
		now,
	)
	if err != nil {
		return nil, car_service.ErrInternal
	}
	id, _ := row.LastInsertId()
	c.ID = id
	c.CreatedAt = now
	return c, nil
}

func (r *CarRepo) UpdateCar(ctx context.Context, c *car_model.CarModel, id int64) (*car_model.CarModel, error) {
	now := time.Now().UTC()

	const q = `UPDATE cars
	           SET brand=?, model=?, year=?, rental_price=?, image_url=?, updated_at=?
	           WHERE id=?`

	res, err := r.db.ExecContext(ctx, q,
		c.Brand,
		c.Model,
		c.Year,
		c.RentalPrice,
		c.ImageUrl,
		now,
		id,
	)
	if err != nil {
		return nil, car_service.ErrInternal
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		return nil, car_service.ErrNotFound
	}

	c.ID = id
	c.UpdatedAt = &now
	return c, nil
}

func (r *CarRepo) GetCarById(ctx context.Context, id int64) (*car_model.CarModel, error) {
	const q = `SELECT brand, model, year, rental_price, image_url, created_at, updated_at, id
	           FROM cars WHERE id = ?`

	row := r.db.QueryRowContext(ctx, q, id)

	var c car_model.CarModel
	var imageUrl *string
	var updatedAt *time.Time

	err := row.Scan(&c.Brand, &c.Model, &c.Year, &c.RentalPrice, &imageUrl, &c.CreatedAt, &updatedAt, &c.ID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, car_service.ErrNotFound
		}
		return nil, car_service.ErrDataReading
	}

	c.ImageUrl = imageUrl
	c.UpdatedAt = updatedAt

	return &c, nil
}

func (r *CarRepo) DeleteCar(ctx context.Context, id int64) (deletedId int64, err error) {
	const q = `DELETE FROM cars WHERE id = ?`

	res, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return 0, car_service.ErrDelete
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		return 0, car_service.ErrNotFound
	}

	return id, nil
}

func (r *CarRepo) Exists(ctx context.Context, id int64) (bool, error) {
	const q = `SELECT COUNT(*) from cars WHERE id = ?`

	var count int

	row := r.db.QueryRowContext(ctx, q, id)

	err := row.Scan(&count)
	if err != nil {
		return false, car_service.ErrDataReading
	}

	return count == 0, nil
}
