package rds

import (
	"errors"
	"strings"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

var ()

func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func IsDuplicateErr(err error) bool {
	if err == nil {
		return false
	}
	pgErr := new(pq.Error)
	if errors.As(err, &pgErr) &&
		pgErr.Code == "23505" {
		return true
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return true
	}
	if strings.Contains(err.Error(), "duplicate key") {
		return true
	}
	return false
}
