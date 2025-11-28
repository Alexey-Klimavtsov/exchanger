package user_service

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidInput = fmt.Errorf("invalid input")
	ErrNotFound     = fmt.Errorf("user not found")
	ErrCreate       = fmt.Errorf("cannot create user")
	ErrUpdate       = fmt.Errorf("cannot update user")
	ErrDelete       = fmt.Errorf("cannot delete user")
	ErrDataReading  = fmt.Errorf("data reading error")
	ErrInternal     = errors.New("server user internal error")
)
