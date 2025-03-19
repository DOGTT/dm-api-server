package rds

import (
	"context"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

func init() {
	dbModelList = append(dbModelList, &ChannelPostInfo{})
}

// 足迹频道帖子
type ChannelPostInfo struct {
	// 评论唯一id
	Id uint64 `gorm:"primaryKey;autoIncrement"`
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

	// --- 基础字段
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	// 软删除字段
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (c *RDSClient) CreateChannelComment(ctx context.Context, d *ChannelPostInfo) error {
	return c.db.WithContext(ctx).Create(d).Error
}
