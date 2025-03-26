// nolint:unused
package rds

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

const limitDefault = 100

// common field define
const (
	sqlFieldAll = "*"

	sqlFieldId       = "id"
	sqlFieldUId      = "uid"
	sqlFieldName     = "name"
	sqlFieldTitle    = "title"
	sqlFieldIntro    = "intro"
	sqlFieldContent  = "content"
	sqlFieldTypeId   = "type_id"
	sqlFieldAvatarId = "avatar_id"
	sqlFieldWeChatId = "we_chat_id"

	sqlFieldCreatedAt = "created_at"
	sqlFieldUpdatedAt = "updated_at"
)

// equal
const (
	sqlEqualId = "id = ?"
)

// select
const (
	sqlSelectLngLat = "ST_AsText(lng_lat) AS lng_lat"
)

// where
const (
	sqlWhereLngLatContain = "ST_Contains(ST_MakeEnvelope(?, ?, ?, ?, 4326), lng_lat)"
)

// 生成等于条件
func sqlEqual(fieldKey string) string {
	return fieldKey + " = ?"
}

// 生成不等于条件
func sqlNotEqual(fieldKey string) string {
	return fieldKey + " != ?"
}

// 生成小于条件
func sqlLess(fieldKey string) string {
	return fieldKey + " < ?"
}

// 生成小于等于条件
func sqlLessEqual(fieldKey string) string {
	return fieldKey + " <= ?"
}

// 生成大于条件
func sqlMore(fieldKey string) string {
	return fieldKey + " > ?"
}

// 生成大于等于条件
func sqlMoreEqual(fieldKey string) string {
	return fieldKey + " >= ?"
}

// 生成 LIKE 条件
func sqlLike(fieldKey string) string {
	return fieldKey + " LIKE ?"
}

// 生成 IN 条件
func sqlIn(fieldKey string) string {
	return fieldKey + " IN ?"
}

// 生成子查询 IN 条件
func sqlSubIn(fieldKey string) string {
	return fieldKey + " IN (?)"
}

// 生成子查询 NOT IN 条件
func sqlSubNotIn(fieldKey string) string {
	return fieldKey + " NOT IN (?)"
}

func sqlOrderAsc(fieldKey string) string {
	return fieldKey + " ASC"
}

func sqlOrderDesc(fieldKey string) string {
	return fieldKey + " DESC"
}

type Uint64Array []uint64

// 实现 driver.Valuer 接口
func (a Uint64Array) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// 实现 sql.Scanner 接口
func (a *Uint64Array) Scan(value interface{}) error {
	if value == nil {
		*a = nil
		return nil
	}
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, a)
}
