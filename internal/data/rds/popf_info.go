// poi_table.go
package rds

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func init() {
	dbModelList = append(dbModelList, &POFPInfo{})
}

// 足迹点基本信息
type POFPInfo struct {
	ID         uuid.UUID      `gorm:"type:uuid;primaryKey;"`
	TypeID     string         `gorm:"type:varchar(4)"`
	CreaterUID string         `gorm:"type:varchar(32)"`
	PhotoList  pq.StringArray `gorm:"type:text[]"` // 使用 PostgreSQL 的数组类型
	Name       string         `gorm:"type:text;size:50;not null"`
	Content    string         `gorm:"type:text;size:500;"`
	Geometry   string         `gorm:"type:geometry(Point,4326);not null"` // 使用 geometry 类型
	// TagList    pq.StringArray `gorm:"type:text[]"`

	// FreshScore   int       `gorm:"not null"`
	// PopularScore int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
