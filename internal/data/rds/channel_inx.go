package rds

import (
	"context"
	"time"

	api "github.com/DOGTT/dm-api-server/api/base"
	"gorm.io/gorm"
)

func init() {
	dbModelList = append(dbModelList, &UserChannelIxnEvent{})
	dbModelList = append(dbModelList, &UserChannelIxnState{})
}

// 爱宠频道互动事件
type UserChannelIxnEvent struct {
	UId       uint64              `gorm:"index;column:uid"`
	PId       uint64              `gorm:"index;column:pid"`
	ChannelId uint64              `gorm:"index"`
	IxnEvent  api.ChannelIxnEvent `gorm:"type:int;default:0"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// 用户频道互动状态
type UserChannelIxnState struct {
	UId       uint64 `gorm:"index:ix_user_channel_inx,unique;column:uid"`
	ChannelId uint64 `gorm:"index:ix_user_channel_inx,unique"`
	//
	IxnState api.ChannelIxnState `gorm:"type:int;default:0;index:ix_user_channel_inx,unique"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (c *RDSClient) BatchCreateUserPetChannelIxnEvent(ctx context.Context, d []*UserChannelIxnEvent) error {
	return c.db.WithContext(ctx).CreateInBatches(d, 10).Error
}

func (c *RDSClient) CreateUserChannelIxnState(ctx context.Context, d *UserChannelIxnState) error {
	return c.db.WithContext(ctx).Create(d).Error
}

func (c *RDSClient) DeleteUserChannelIxnState(ctx context.Context, d *UserChannelIxnState) error {
	res := c.db.WithContext(ctx).Where(d).Delete(d)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}
