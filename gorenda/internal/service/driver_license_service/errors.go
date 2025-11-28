package driver_license_service

import "errors"

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("driver license not found")
	ErrCreate       = errors.New("couldn't create a driver license")
	ErrUpdate       = errors.New("couldn't update a driver license")
	ErrDelete       = errors.New("couldn't delete a driver license")
	ErrDataReading  = errors.New("data reading error")
	ErrInternal     = errors.New("server drivier license internal error")
)
