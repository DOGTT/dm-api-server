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
