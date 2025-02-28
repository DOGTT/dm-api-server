package rds

import (
	"context"
	"time"

	"github.com/lib/pq"
)

func init() {
	dbModelList = append(dbModelList, &ChannelPostInfo{})
}

// 足迹频道帖子
type ChannelPostInfo struct {
	// 评论唯一id
	UUID string `gorm:"type:varchar(22);primaryKey;"`
	// - 创建者UId
	UId uint64 `gorm:"index"`
	// 底层足迹id
	ChannelUUID string `gorm:"index"`
	// 关联的上级帖子id, 空则为根帖子
	ParentUUID string
	// 帖子内容
	Content string `gorm:"type:text"`
	// 帖子图片
	Photos pq.StringArray `gorm:"type:text[]"`
	// 添加的标签
	Tags []string `gorm:"type:text[]"`

	// -- 动态信息
	Likes int `gorm:"default:0"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (c *RDSClient) CreatePofpComment(ctx context.Context, d *ChannelPost) error {
	return c.db.WithContext(ctx).Create(d).Error
}
