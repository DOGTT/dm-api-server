package rds

import (
	"time"
)

func init() {
	// dbModelList = append(dbModelList, &UserPostIxnEvent{})
}

// 用户帖子互动状态 TODO
type UserPostIxnEvent struct {
	UId uint64 `gorm:"index"`
	// 帖子类型
	PostId uint64 `gorm:"index"`
	//
	// IxnState  api.ChannelIxnState `gorm:"type:int;default:0"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
