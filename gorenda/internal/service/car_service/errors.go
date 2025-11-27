package car_service

import "errors"

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrNotFound     = errors.New("car not found")
	ErrCreate       = errors.New("couldn't create a car")
	ErrUpdate       = errors.New("couldn't update a car")
	ErrDelete       = errors.New("couldn't delete a car")
	ErrDataReading  = errors.New("data reading error")
)
