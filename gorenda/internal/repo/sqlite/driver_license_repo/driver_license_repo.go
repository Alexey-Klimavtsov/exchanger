package driver_license_repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/asaipov/gorenda/internal/it/model/driver_license_model"
	"github.com/asaipov/gorenda/internal/service/driver_license_service"
	"github.com/asaipov/gorenda/internal/service/user_service"
	"time"
)

type DriverLicenseRepo struct {
	db *sql.DB
}

func NewDriverLicenseRepo(db *sql.DB) *DriverLicenseRepo {
	return &DriverLicenseRepo{db: db}
}

func (r *DriverLicenseRepo) CreateLicense(ctx context.Context, dl *driver_license_model.DriverLicenseModel) (*driver_license_model.DriverLicenseModel, error) {
	const q = `INSERT INTO driver_licenses (user_id, number, issued_at, expires_at, created_at) VALUES(?,?,?,?,?)`

	var dlCopy = dl
	now := time.Now().UTC()

	res, err := r.db.ExecContext(ctx, q, dlCopy.UserID, dlCopy.Number, dlCopy.IssuedAt, dlCopy.ExpiresAt, now)
	if err != nil {
		fmt.Println(err)
		return nil, driver_license_service.ErrInternal
	}

	dlCopy.CreatedAt = now
	dlCopy.ID, _ = res.LastInsertId()

	return dlCopy, nil
}

func (r *DriverLicenseRepo) UpdateLicense(ctx context.Context, dl *driver_license_model.DriverLicenseModel, id int64) (*driver_license_model.DriverLicenseModel, error) {
	const q = `UPDATE driver_licenses
	SET number = ?,
    issued_at = ?,
    expires_at = ?,
    updated_at = ?
	WHERE id = ?
`
	var dlCopy = dl
	now := time.Now().UTC()

	res, err := r.db.ExecContext(ctx, q, dlCopy.Number, dlCopy.IssuedAt, dlCopy.ExpiresAt, now, id)
	if err != nil {
		return nil, driver_license_service.ErrInternal
	}
	affected, _ := res.RowsAffected()

	if affected == 0 {
		return nil, driver_license_service.ErrNotFound
	}
	dlCopy.UpdatedAt = &now
	dlCopy.ID = id
	return dlCopy, nil
}

func (r *DriverLicenseRepo) GetLicensesByUserId(ctx context.Context, id int64) ([]*driver_license_model.DriverLicenseModel, error) {
	const q = `
        SELECT id, number, issued_at, expires_at, created_at, updated_at, user_id
        FROM driver_licenses
        WHERE user_id = ?
    `

	rows, queryErr := r.db.QueryContext(ctx, q, id)
	if queryErr != nil {
		return nil, user_service.ErrInternal
	}
	defer rows.Close()

	var rightsCategory []*driver_license_model.DriverLicenseModel

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
			&dl.UserID,
		); scanErr != nil {
			return nil, driver_license_service.ErrInternal
		}
		if upd.Valid {
			dl.UpdatedAt = &upd.Time
		}
		rightsCategory = append(rightsCategory, &dl)
	}

	err := rows.Err()
	if err != nil {
		return nil, driver_license_service.ErrInternal
	}
	return rightsCategory, nil
}
