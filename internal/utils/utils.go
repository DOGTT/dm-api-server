package utils

import (
	"encoding/base64"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func GenUUID() string {
	// use v7
	uuid, _ := uuid.NewV7()
	return uuid.String()
}

func GenUUIDPrefix(prefix interface{}) string {
	return fmt.Sprintf("%v-%v", prefix, GenUUID())
}

func Base64ToBytes(base64Str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(base64Str)
}

// ConvertToUintSlice 将各种整数切片转换为 []uint
func ConvertToUintSlice[T int | int8 | int16 | int32 | int64](input []T) []uint {
	result := make([]uint, len(input))
	for i, v := range input {
		result[i] = uint(v)
	}
	return result
}

func ConvertToUint64(s string) uint64 {
	num, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		zap.L().Warn("parse int error", zap.Error(err))
	}
	return num
}
