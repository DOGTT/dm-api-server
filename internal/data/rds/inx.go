package rds

import (
	"context"
	"time"

	"gorm.io/gorm"
)

func init() {
	dbModelList = append(dbModelList, &PetChannelIxnEvent{})
	dbModelList = append(dbModelList, &UserChannelIxnState{})
}

// 爱宠频道互动记录
type PetChannelIxnEvent struct {
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

func (c *RDSClient) CreateChannelIxnRecord(ctx context.Context, d *UserChannelIxnRecord) error {
	return c.db.WithContext(ctx).Create(d).Error
}

func (c *RDSClient) CreateChannelIxnRecordWithCount(ctx context.Context, d *UserChannelIxnRecord) error {

	return c.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var info ChannelInfo
		field := InxTypeFieldName[d.IntType]
		if err := tx.Model(&ChannelInfo{}).Where("uuid = ?", d.ChannelUUID).
			Select(field).First(&info).Error; err != nil {
			return err
		}
		if err := tx.Create(d).Error; err != nil {
			return err
		}
		info.UUID = d.ChannelUUID
		switch d.IntType {
		case InxTypeLike:
			info.LikesCnt++
		case InxTypeMark:
			info.MarksCnt++
		case InxTypeComment:
			info.CommentsCnt++
		}
		if err := tx.Select(field).Save(&info).Error; err != nil {
			return err
		}
		return nil
	})
}
