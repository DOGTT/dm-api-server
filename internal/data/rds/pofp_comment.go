package rds

import (
	"time"

	"github.com/lib/pq"
)

func init() {
	dbModelList = append(dbModelList, &PofpComment{})
}

// 足迹点评论
type PofpComment struct {
	UUID string `gorm:"type:varchar(22);primaryKey;"`
	// - 归属的PetID
	PId uint `gorm:"index"`
	//
	Photos  pq.StringArray `gorm:"type:text[]"` // 使用 PostgreSQL 的数组类型
	Content string         `gorm:"type:text"`
	Likes   int            `gorm:"default:0"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
