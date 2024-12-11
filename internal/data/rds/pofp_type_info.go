// poi_type_info.go
package rds

import (
	"time"
)

func init() {
	dbModelList = append(dbModelList, &PofpTypeInfo{})
}

// 足迹点类型信息
type PofpTypeInfo struct {
	Id uint `gorm:"primaryKey;autoIncrement"`
	// 足迹点类型名称, 如: 探险, 小憩, 溜溜
	Name string `gorm:"type:varchar(12)"`
	// 覆盖半径, 以米为单位
	CoverageRadius int `gorm:"type:smallint;"`
	// 主题色, 16进制
	// 如: #FF0000
	ThemeColor string `gorm:"type:varchar(12)"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
