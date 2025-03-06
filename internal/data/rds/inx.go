package rds

import (
	"context"
	"time"
)

func init() {
	dbModelList = append(dbModelList, &UserPetChannelIxnEvent{})
	dbModelList = append(dbModelList, &UserChannelIxnState{})
}

// 爱宠频道互动事件
type UserPetChannelIxnEvent struct {
	UId      uint64           `gorm:"index"`
	PId      uint64           `gorm:"index"`
	IxnEvent UserIxnEventType `gorm:"type:int;default:0"`

	ChannelUUID string `gorm:"index"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// 用户频道互动状态
type UserChannelIxnState struct {
	UId uint64 `gorm:"index"`
	//
	IxnState    UserIxnStateType `gorm:"type:int;default:0"`
	ChannelUUID string           `gorm:"index"`
	CreatedAt   time.Time        `gorm:"autoCreateTime"`
}

// 用户帖子互动状态
type UserPostIxnState struct {
	UId uint64 `gorm:"index"`
	//
	IxnState UserIxnStateType `gorm:"type:int;default:0"`
	// 帖子类型
	PostUUID  string    `gorm:"index"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// 事件型 互动
type UserIxnEventType uint

const (
	UserIxnEventTypeDefault UserIxnEventType = iota // 默认
	UserIxnEventTypeLand                            // 到访打卡
)

// 状态型 互动
type UserIxnStateType uint

const (
	UserIxnStateTypeDefault UserIxnStateType = iota // 默认
	UserIxnStateTypeStar                            // 收藏
	UserIxnStateTypeJoin                            // 加入
)

var (
	InxTypeFieldName = []string{"views_cnt", "likes_cnt", "marks_cnt", "comments_cnt"}
)

func (c *RDSClient) CreateUserPetChannelIxnEvent(ctx context.Context, d *UserPetChannelIxnEvent) error {
	return c.db.WithContext(ctx).Create(d).Error
}

func (c *RDSClient) CreateUserChannelIxnState(ctx context.Context, d *UserChannelIxnState) error {
	return c.db.WithContext(ctx).Create(d).Error
}
