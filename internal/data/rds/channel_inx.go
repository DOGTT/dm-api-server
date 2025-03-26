package rds

import (
	"context"
	"time"

	api "github.com/DOGTT/dm-api-server/api/base"
)

func init() {
	dbModelList = append(dbModelList, &UserPetChannelIxnEvent{})
	dbModelList = append(dbModelList, &UserChannelIxnState{})
}

// 爱宠频道互动事件
type UserPetChannelIxnEvent struct {
	UId       uint64              `gorm:"index"`
	PId       uint64              `gorm:"index"`
	ChannelId uint64              `gorm:"index"`
	IxnEvent  api.ChannelIxnEvent `gorm:"type:int;default:0"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// 用户频道互动状态
type UserChannelIxnState struct {
	UId       uint64 `gorm:"index"`
	ChannelId uint64 `gorm:"index"`
	//
	IxnState api.ChannelIxnState `gorm:"type:int;default:0;index"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (c *RDSClient) BatchCreateUserPetChannelIxnEvent(ctx context.Context, d []*UserPetChannelIxnEvent) error {
	return c.db.WithContext(ctx).Create(d).Error
}

func (c *RDSClient) CreateUserChannelIxnState(ctx context.Context, d *UserChannelIxnState) error {
	return c.db.WithContext(ctx).Create(d).Error
}

func (c *RDSClient) DeleteUserChannelIxnState(ctx context.Context, d *UserChannelIxnState) error {
	return c.db.WithContext(ctx).Delete(d).Error
}
