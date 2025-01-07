package rds

import (
	"context"
	"time"

	"github.com/lib/pq"
)

func init() {
	dbModelList = append(dbModelList, &PofpComment{})
}

// 足迹点评论
type PofpComment struct {
	// - 评论PetID
	PId  uint   `gorm:"index"`
	UUID string `gorm:"type:varchar(22);primaryKey;"`
	// 底层足迹id
	PofpUUID string `gorm:"index"`
	// 上级id
	ParentUUID string
	//
	Photos pq.StringArray `gorm:"type:text[]"`
	// 评论内容
	Content string `gorm:"type:text"`
	Likes   int    `gorm:"default:0"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (c *RDSClient) CreatePofpComment(ctx context.Context, d *PofpComment) error {
	return c.db.WithContext(ctx).Create(d).Error
}
