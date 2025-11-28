package user_repo

import (
	"context"
	"database/sql"
	"errors"
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
	const q = `
	INSERT INTO users(first_name, last_name, surname, email, birthday, created_at)
	VALUES (?, ?, ?, ?, ?, ?)
	`

	now := time.Now().UTC()

	var surname sql.NullString
	if u.Surname != nil {
		surname = sql.NullString{String: *u.Surname, Valid: true}
	}

	res, err := r.db.ExecContext(ctx, q,
		u.FirstName,
		u.LastName,
		surname,
		u.Email,
		u.Birthday,
		now,
	)
	if err != nil {
		return nil, user_service.ErrInternal
	}

	id, _ := res.LastInsertId()
	u.ID = id
	u.CreatedAt = now

	return u, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, u *user_model.UserModel, id int64) (*user_model.UserModel, error) {
	const q = `
	UPDATE users
	SET first_name=?, last_name=?, surname=?, email=?, birthday=?, updated_at=?
	WHERE id=? AND deleted_at IS NULL
	`

	now := time.Now().UTC()

	var surname sql.NullString
	if u.Surname != nil {
		surname = sql.NullString{String: *u.Surname, Valid: true}
	}

	res, err := r.db.ExecContext(ctx, q,
		u.FirstName,
		u.LastName,
		surname,
		u.Email,
		u.Birthday,
		now,
		id,
	)
	if err != nil {
		return nil, user_service.ErrInternal
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return nil, user_service.ErrNotFound
	}

	u.ID = id
	u.UpdatedAt = &now
	return u, nil
}

func (r *UserRepo) GetUserById(ctx context.Context, id int64) (*user_model.UserModel, error) {

	const q = `SELECT id, first_name, last_name, surname, email, birthday, created_at, updated_at
               FROM users
               WHERE id = ? AND deleted_at IS NULL`

	var (
		u       user_model.UserModel
		surname sql.NullString
		updAt   sql.NullTime
	)

	row := r.db.QueryRowContext(ctx, q, id)
	if err := row.Scan(
		&u.ID, &u.FirstName, &u.LastName, &surname,
		&u.Email, &u.Birthday, &u.CreatedAt, &updAt,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, user_service.ErrNotFound
		}
		return nil, user_service.ErrInternal
	}

	if surname.Valid {
		u.Surname = &surname.String
	}
	if updAt.Valid {
		u.UpdatedAt = &updAt.Time
	}

	return &u, nil
}

func (r *UserRepo) DeleteUser(ctx context.Context, id int64) (int64, error) {
	const q = `DELETE FROM users WHERE id = ?`

	res, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return 0, user_service.ErrInternal
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		return 0, user_service.ErrNotFound
	}

	return id, nil
}

func (r *UserRepo) Exists(ctx context.Context, id int64) (bool, error) {
	const q = `SELECT COUNT(*) from users WHERE id = ?`

	var count int

	row := r.db.QueryRowContext(ctx, q, id)

	err := row.Scan(&count)
	if err != nil {
		return false, user_service.ErrInternal
	}

	return count == 0, nil
}
