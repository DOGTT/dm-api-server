package rds

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

var (
	channelSetModel = &ChannelStats{}
)

func init() {
	dbModelList = append(dbModelList, channelSetModel)
}

// 频道配置 TODO
type ChannelConfig struct {
	MaxUserCount int `json:"max_user_count"`
}

type ChannelSet struct {
	Id uint64 `gorm:"primaryKey;autoIncrement"`

	// 频道配置
	Config ChannelConfig `gorm:"type:jsonb"`
	// 个性标签列表
	CustomTags pq.StringArray `gorm:"type:text[]"`

	// --- 基础字段
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	// 软删除字段
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
