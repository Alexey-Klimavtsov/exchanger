package driver_license_repo

import (
	"context"
	"database/sql"
	"github.com/asaipov/gorenda/internal/it/model/driver_license_model"
	"github.com/asaipov/gorenda/internal/service/driver_license_service"
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
		// todo - тоже стрем проверка беспонтовая
		return nil, driver_license_service.ErrCreate
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
		return nil, driver_license_service.ErrUpdate
	}
	affected, _ := res.RowsAffected()
	// todo - кажется стрем проверка
	if affected == 0 {
		return nil, driver_license_service.ErrNotFound
	}
	dlCopy.UpdatedAt = sql.NullTime{Time: now, Valid: true}
	dlCopy.ID = id
	return dlCopy, nil
}
