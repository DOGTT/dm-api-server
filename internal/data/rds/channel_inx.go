package rds

import (
	"context"
	"time"

	api "github.com/DOGTT/dm-api-server/api/base"
	"gorm.io/gorm"
)

func init() {
	dbModelList = append(dbModelList, &ChannelIxnEvent{})
	dbModelList = append(dbModelList, &ChannelPetIxnEvent{})
	dbModelList = append(dbModelList, &ChannelIxnState{})
}

// 爱宠频道互动事件
type ChannelIxnEvent struct {
	ChannelId uint64                  `gorm:"index"`
	UId       uint64                  `gorm:"index;column:uid"`
	IxnEvent  api.ChannelUserIxnEvent `gorm:"type:int;default:0"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// 爱宠频道互动事件
type ChannelPetIxnEvent struct {
	ChannelId uint64                 `gorm:"index"`
	UId       uint64                 `gorm:"index;column:uid"`
	PId       uint64                 `gorm:"index;column:pid"`
	IxnEvent  api.ChannelPetIxnEvent `gorm:"type:int;default:0"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// 用户频道互动状态
type ChannelIxnState struct {
	ChannelId uint64                  `gorm:"index:ix_channel_inx,unique"`
	UId       uint64                  `gorm:"index:ix_channel_inx,unique;column:uid"`
	IxnState  api.ChannelUserIxnState `gorm:"type:int;default:0;index:ix_channel_inx,unique"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index,unique"`
}

func (c *RDSClient) BatchCreateChannelIxnEvent(ctx context.Context, d []*ChannelIxnEvent) error {
	return c.db.WithContext(ctx).CreateInBatches(d, batchCreateSize).Error
}

func (c *RDSClient) BatchCreateChannelPetIxnEvent(ctx context.Context, d []*ChannelPetIxnEvent) error {
	return c.db.WithContext(ctx).CreateInBatches(d, batchCreateSize).Error
}

func (c *RDSClient) CreateChannelIxnState(ctx context.Context, d *ChannelIxnState) error {
	return c.db.WithContext(ctx).Create(d).Error
}

func (c *RDSClient) DeleteChannelIxnState(ctx context.Context, d *ChannelIxnState) error {
	res := c.db.WithContext(ctx).Where(d).Delete(d)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}
