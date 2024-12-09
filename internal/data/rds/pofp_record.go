package rds

import (
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func init() {
	dbModelList = append(dbModelList, &POFPRecord{})
}

// 足迹点互动记录
type POFPRecord struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;"`
	PID       string         `gorm:"type:varchar(32)"`
	PhotoList pq.StringArray `gorm:"type:text[]"` // 使用 PostgreSQL 的数组类型
	Title     string         `gorm:"size:255;not null"`
	Content   string         `gorm:"type:text"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
