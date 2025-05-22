package rds

import (
	"context"
	"time"

	api "github.com/DOGTT/dm-api-server/api/base"
	"gorm.io/gorm"
)

func init() {
	dbModelList = append(dbModelList, &PostIxnEvent{})
	dbModelList = append(dbModelList, &PostIxnState{})
}

// 用户帖子互动事件 如: 转发
type PostIxnEvent struct {
	UId      uint64               `gorm:"index;column:uid"`
	PostId   uint64               `gorm:"index"`
	IxnEvent api.PostUserIxnEvent `gorm:"type:int;default:0"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// 用户帖子互动状态 如: 收藏
type PostIxnState struct {
	UId      uint64               `gorm:"index;column:uid"`
	PostId   uint64               `gorm:"index"`
	IxnState api.PostUserIxnState `gorm:"type:int;default:0;index:ix_channel_inx,unique"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index,unique"`
}

func (c *RDSClient) BatchCreatePostIxnEvent(ctx context.Context, d []*PostIxnEvent) error {
	return c.db.WithContext(ctx).CreateInBatches(d, batchCreateSize).Error
}

func (c *RDSClient) CreatePostIxnState(ctx context.Context, d *PostIxnState) error {
	return c.db.WithContext(ctx).Create(d).Error
}

func (c *RDSClient) DeletePostIxnState(ctx context.Context, d *PostIxnState) error {
	res := c.db.WithContext(ctx).Where(d).Delete(d)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}
