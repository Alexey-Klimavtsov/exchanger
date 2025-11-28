package booking_repo

import (
	"context"
	"database/sql"
	"github.com/asaipov/gorenda/internal/it/model/booking_model"
	"github.com/asaipov/gorenda/internal/service/booking_service"
	"time"
)

type BookingRepo struct {
	db *sql.DB
}

func NewBookingRepo(db *sql.DB) *BookingRepo {
	return &BookingRepo{db: db}
}

func (r *BookingRepo) CreateBooking(ctx context.Context, b *booking_model.BookingModel) (*booking_model.BookingModel, error) {
	const q = `
        INSERT INTO bookings (user_id, car_id, date_from, date_to, price_per_day, status, created_at)
        VALUES (?, ?, ?, ?, ?, ?, ?)
    `

	now := time.Now().UTC()
	b.CreatedAt = now

	res, err := r.db.ExecContext(
		ctx,
		q,
		b.UserID,
		b.CarID,
		b.DateFrom,
		b.DateTo,
		b.PricePerDay,
		b.Status,
		now,
	)
	if err != nil {
		return nil, booking_service.ErrInternal
	}

	b.ID, _ = res.LastInsertId()

	return b, nil
}

func (r *BookingRepo) GetBookings(ctx context.Context) ([]*booking_model.BookingModel, error) {
	const q = `
		SELECT id, user_id, car_id, date_from, date_to, price_per_day, status, created_at
		FROM bookings
		ORDER BY created_at DESC
	`

	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, booking_service.ErrInternal
	}
	defer rows.Close()

	var list []*booking_model.BookingModel

	for rows.Next() {
		var b booking_model.BookingModel
		if scanErr := rows.Scan(
			&b.ID,
			&b.UserID,
			&b.CarID,
			&b.DateFrom,
			&b.DateTo,
			&b.PricePerDay,
			&b.Status,
			&b.CreatedAt,
		); scanErr != nil {
			return nil, booking_service.ErrInternal
		}

		list = append(list, &b)
	}
	if rows.Err() != nil {
		return nil, booking_service.ErrInternal
	}

	return list, nil
}

func (r *BookingRepo) IsCarAvailable(ctx context.Context, carID int64, from, to time.Time) (bool, error) {
	const q = `SELECT COUNT(*) 
FROM bookings
WHERE car_id = ?
  AND status = 'done'
  AND NOT (
        date_to   <  ?   
     OR date_from >  ?   
  );`

	var count int
	row := r.db.QueryRowContext(ctx, q, carID, from, to)

	scanErr := row.Scan(&count)
	if scanErr != nil {
		return false, booking_service.ErrInternal
	}
	return count == 0, nil
}
