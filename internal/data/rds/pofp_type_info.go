// poi_type_info.go
package rds

import (
	"time"
)

func init() {
	dbModelList = append(dbModelList, &POFPTypeInfo{})
}

// 足迹点类型信息
type POFPTypeInfo struct {
	ID   string `gorm:"type:varchar(12);primaryKey;"`
	Name string `gorm:"type:varchar(12)"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
