package rds

import (
	"context"
	"time"

	"github.com/lib/pq"
)

func init() {
	dbModelList = append(dbModelList, &PofpComment{})
}

// 足迹频道评论
type PofpComment struct {
	// 评论唯一id
	UUID string `gorm:"type:varchar(22);primaryKey;"`
	// - 评论者PetID
	PId uint `gorm:"index"`
	// 底层足迹id
	PofpUUID string `gorm:"index"`
	// 上级id
	ParentUUID string
	// 是否为根消息
	IsRoot bool `gorm:"default:true"`
	// 评论内容
	Content string `gorm:"type:text"`
	//
	Photos pq.StringArray `gorm:"type:text[]"`
	//
	Likes int `gorm:"default:0"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (c *RDSClient) CreatePofpComment(ctx context.Context, d *PofpComment) error {
	return c.db.WithContext(ctx).Create(d).Error
}
