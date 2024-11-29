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
	ID         uuid.UUID      `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	TypeID     string         `gorm:"type:varchar(4)"`
	CreaterUID string         `gorm:"type:varchar(32)"`
	PhotoList  pq.StringArray `gorm:"type:text[]"` // 使用 PostgreSQL 的数组类型
	Title      string         `gorm:"size:255;not null"`
	Content    string         `gorm:"type:text"`
	Geometry   string         `gorm:"type:geometry(Point,4326);not null"` // 使用 geometry 类型
	TagList    pq.StringArray `gorm:"type:text[]"`

	// FreshScore   int       `gorm:"not null"`
	// PopularScore int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
