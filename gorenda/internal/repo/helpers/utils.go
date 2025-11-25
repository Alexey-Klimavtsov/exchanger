package helpers

import (
	"errors"
	"fmt"
	"strings"

	"modernc.org/sqlite"
)

var (
	ErrDuplicate   = errors.New("duplicate")
	ErrForeignKey  = errors.New("foreign_key")
	ErrNotNull     = errors.New("not_null")
	ErrCheckFailed = errors.New("check_failed")
)

func MapSQLiteError(err error) error {
	var sqliteErr *sqlite.Error
	if !errors.As(err, &sqliteErr) {
		return err
	}

	msg := err.Error()

	switch {
	case strings.Contains(msg, "UNIQUE"):
		return fmt.Errorf("%w: record already exists", ErrDuplicate)
	case strings.Contains(msg, "FOREIGN KEY"):
		return fmt.Errorf("%w: related record not found", ErrForeignKey)
	case strings.Contains(msg, "NOT NULL"):
		return fmt.Errorf("%w: missing required field", ErrNotNull)
	case strings.Contains(msg, "CHECK"):
		return fmt.Errorf("%w: check constraint failed", ErrCheckFailed)
	default:
		return fmt.Errorf("sqlite error: %v", msg)
	}
}
