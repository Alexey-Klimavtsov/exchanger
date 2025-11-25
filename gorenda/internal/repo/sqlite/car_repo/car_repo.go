package car_repo

import (
	"context"
	"database/sql"
	"github.com/asaipov/gorenda/internal/it/model/car_model"
	"github.com/asaipov/gorenda/internal/repo/helpers"
	"time"
)

type CarRepo struct {
	db *sql.DB
}

func NewCarRepo(db *sql.DB) *CarRepo {
	return &CarRepo{db: db}
}

func (r *CarRepo) CreateNewCar(ctx context.Context, car *car_model.CarModel) (*car_model.CarModel, error) {
	now := time.Now().UTC()
	var c = car
	const q = `INSERT INTO cars(brand, model, year, rental_price, image_url, created_at) VALUES(?,?,?,?,?,?)`
	var imageUrl string
	if c.ImageUrl.Valid {
		imageUrl = c.ImageUrl.String
	}
	res, err := r.db.ExecContext(ctx, q, c.Brand, c.Model, c.Year, c.RentalPrice, imageUrl, now)

	if err != nil {
		return nil, helpers.MapSQLiteError(err)
	}

	c.CreatedAt = now
	c.ID, _ = res.LastInsertId()

	return c, nil
}
func (r *CarRepo) UpdateCar(ctx context.Context, car *car_model.CarModel, id int64) (*car_model.CarModel, error) {
	now := time.Now().UTC()
	var c = car
	const q = `UPDATE cars
	SET brand = ?,
    model = ?,
    year = ?,
    rental_price = ?,
    image_url = ?,
    updated_at = ?
WHERE id = ?
`
	var imageUrl string
	if c.ImageUrl.Valid {
		imageUrl = c.ImageUrl.String
	}
	_, err := r.db.ExecContext(ctx, q, c.Brand, c.Model, c.Year, c.RentalPrice, imageUrl, now, id)

	if err != nil {
		return nil, helpers.MapSQLiteError(err)
	}

	c.UpdatedAt = now

	return c, nil
}
func (r *CarRepo) GetCarById(ctx context.Context, id int64) (*car_model.CarModel, error) {
	const q = `SELECT brand, model, year, rental_price, image_url, created_at, updated_at, id FROM cars WHERE id = ?`
	row := r.db.QueryRowContext(ctx, q, id)
	var c car_model.CarModel
	var imageUrl sql.NullString
	if err := row.Scan(&c.Brand, &c.Model, &c.Year, &c.RentalPrice, &imageUrl, &c.CreatedAt, &c.UpdatedAt, &c.ID); err != nil {
		return nil, helpers.MapSQLiteError(err)
	}
	return &c, nil
}
func (r *CarRepo) DeleteCar(ctx context.Context, id int64) (deletedId int64, err error) {
	const q = `DELETE FROM cars WHERE id = ?`

	res, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return 0, err
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		return 0, sql.ErrNoRows
	}

	return id, nil
}
