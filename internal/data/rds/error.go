package rds

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

var ()

func IsDuplicateErr(err error) bool {
	if err == nil {
		return false
	}
	pgErr := new(pq.Error)
	if errors.As(err, &pgErr) &&
		pgErr.Code == "23505" {
		fmt.Printf("Error bbb: %v\n", pgErr)
		return true
	}
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return true
	}
	if strings.Contains(err.Error(), "duplicate key") {
		return true
	}
	// fmt.Printf("Error type: %T\n", err) // 打印错误类型
	return false
}
