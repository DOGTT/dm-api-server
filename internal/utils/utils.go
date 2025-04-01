package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"slices"

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

// ConvertToUintSlice 将各种整数切片转换为 []uint
func ConvertToUintSlice[T int | int8 | int16 | int32 | int64](input []T) []uint {
	result := make([]uint, len(input))
	for i, v := range input {
		result[i] = uint(v)
	}
	return result
}

func StrToUint64(in string) uint64 {
	if in == "" {
		return 0
	}
	num, err := strconv.ParseUint(in, 10, 64)
	if err != nil {
		zap.L().Warn("parse int error", zap.Error(err))
	}
	return num
}

func Uint64ToStr(in uint64) string {
	if in == 0 {
		return ""
	}
	return strconv.FormatUint(in, 10)
}

func CopyFromPtr(in *string) string {
	if in == nil {
		return ""
	} else {
		return *in
	}
}

func Base64ToBytes(in string) []byte {
	data, err := base64.StdEncoding.DecodeString(in)
	if err != nil {
		zap.L().Warn("base64 decode error", zap.Error(err))
	}
	return data
}

func ValidateUsername(username string) (string, error) {
	// 去除前后空格
	username = strings.TrimSpace(username)

	// 检查长度
	if utf8.RuneCountInString(username) < 3 {
		return "", errors.New("用户名太短，至少需要1个字符")
	}
	if utf8.RuneCountInString(username) > 20 {
		return "", errors.New("用户名太长，最多20个字符")
	}

	// 检查非法字符
	if matched, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+$`, username); !matched {
		return "", errors.New("用户名只能包含字母、数字、下划线和连字符")
	}

	// 检查保留用户名
	reservedNames := []string{"admin", "root", "system"}
	if slices.Contains(reservedNames, strings.ToLower(username)) {
		return "", errors.New("该用户名是保留名称，请选择其他用户名")
	}

	return username, nil
}
