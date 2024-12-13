package utils

import (
	"encoding/base64"
	"fmt"

	"github.com/google/uuid"
)

func GenUUID() string {
	return uuid.New().String()
}

func GenShortenUUID() string {
	u := uuid.New()
	return base64.RawURLEncoding.EncodeToString(u[:])
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
