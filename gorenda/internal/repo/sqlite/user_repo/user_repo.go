package user_repo

import (
	"context"
	"database/sql"
	"errors"
	"github.com/asaipov/gorenda/internal/it/model/driver_license_model"
	"github.com/asaipov/gorenda/internal/it/model/user_model"
	"github.com/asaipov/gorenda/internal/service/user_service"
	"time"
)

type UserRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(ctx context.Context, u *user_model.UserModel) (*user_model.UserModel, error) {
	const q = `INSERT INTO users 
        (first_name, last_name, surname, birthday, created_at) 
        VALUES (?, ?, ?, ?, ?)`

	now := time.Now().UTC()

	var surname string
	if u.Surname.Valid {
		surname = u.Surname.String
	} else {
		surname = ""
	}

	res, err := r.db.ExecContext(ctx, q,
		u.FirstName,
		u.LastName,
		surname,
		u.Birthday,
		now,
	)

	if err != nil {
		return nil, user_service.ErrCreate
	}

	id, _ := res.LastInsertId()

	u.ID = id
	u.CreatedAt = now

	return u, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, u *user_model.UserModel, id int64) (*user_model.UserModel, error) {
	const q = `UPDATE users
	           SET first_name = ?,
	               last_name  = ?,
	               surname    = ?,
	               birthday   = ?,
	               updated_at = ?
	           WHERE id = ? 
	             AND deleted_at IS NULL`

	now := time.Now().UTC()

	// TODO - поправить тип
	var surname interface{}
	if u.Surname.Valid {
		surname = u.Surname.String
	} else {
		surname = nil
	}

	res, err := r.db.ExecContext(ctx, q,
		u.FirstName,
		u.LastName,
		surname,
		u.Birthday,
		now,
		id,
	)

	if err != nil {
		return nil, user_service.ErrUpdate
	}

	affected, _ := res.RowsAffected()

	if affected == 0 {
		return nil, user_service.ErrNotFound
	}

	u.ID = id
	u.UpdatedAt = sql.NullTime{Time: now, Valid: true}

	return u, nil
}

func (r *UserRepo) GetUserById(ctx context.Context, id int64) (*user_model.UserModel, error) {
	const q = `
        SELECT id, first_name, last_name, surname, birthday, created_at, updated_at
        FROM users
        WHERE id = ? AND deleted_at IS NULL
    `

	var (
		u       user_model.UserModel
		surname sql.NullString
		updAt   sql.NullTime
	)

	row := r.db.QueryRowContext(ctx, q, id)

	if err := row.Scan(
		&u.ID,
		&u.FirstName,
		&u.LastName,
		&surname,
		&u.Birthday,
		&u.CreatedAt,
		&updAt,
	); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, user_service.ErrNotFound
		}
		return nil, user_service.ErrDataReading
	}

	u.Surname = surname
	u.UpdatedAt = updAt

	const dlQ = `
        SELECT id, number, issued_at, expires_at, created_at, updated_at
        FROM driver_licenses
        WHERE user_id = ?
    `

	rows, queryErr := r.db.QueryContext(ctx, dlQ, u.ID)
	if queryErr != nil {
		return nil, user_service.ErrDataReading
	}
	defer rows.Close()

	for rows.Next() {
		var dl driver_license_model.DriverLicenseModel
		var upd sql.NullTime

		if scanErr := rows.Scan(
			&dl.ID,
			&dl.Number,
			&dl.IssuedAt,
			&dl.ExpiresAt,
			&dl.CreatedAt,
			&upd,
		); scanErr != nil {
			return nil, user_service.ErrDataReading
		}

		dl.UpdatedAt = upd
		u.RightsCategory = append(u.RightsCategory, &dl)
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		return nil, user_service.ErrDataReading
	}

	return &u, nil
}

func (r *UserRepo) DeleteUser(ctx context.Context, id int64) (int64, error) {
	const q = `DELETE FROM users WHERE id = ?`

	res, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		// TODO - проверить корректность проверки
		if errors.Is(err, sql.ErrNoRows) {
			return 0, user_service.ErrNotFound
		}
		return 0, user_service.ErrDelete
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		return 0, user_service.ErrNotFound
	}

	return id, nil

}
