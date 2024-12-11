package rds

import (
	"time"

	"github.com/google/uuid"
)

func init() {
	dbModelList = append(dbModelList, &UserPofpRecord{})
}

const (
	IntTypeView    = iota
	IntTypeLike    // 喜欢
	IntTypeMark    // 标记过
	IntTypeComment // 评论
)

// 足迹点个互动记录 喜欢/踩过/评论
type UserPofpRecord struct {
	UUID     uuid.UUID `gorm:"type:uuid;primaryKey;"`
	UId      uint      `gorm:"index"`
	PId      uint      `gorm:"index"`
	PofpUUID string    `gorm:"index"`
	//
	IntType int `gorm:"type:int;default:0"` // 0-点赞,1-收藏

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
