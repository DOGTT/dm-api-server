package rds

import (
	"time"

	"github.com/lib/pq"
)

var (
	modelChannelStats = &ChannelStats{}
)

func init() {
	dbModelList = append(dbModelList, modelChannelStats)
}

// 频道配置
type ChannelConfig struct {
	MaxUserCount int `json:"max_user_count"`
}

type ChannelStats struct {
	Id uint64 `gorm:"primaryKey;autoIncrement"`

	// 频道配置
	Config ChannelConfig `gorm:"type:jsonb"`
	// 个性标签列表
	CustomTags pq.StringArray `gorm:"type:text[]"`

	// - 互动信息
	ViewsCnt    int `gorm:"default:0"`
	LikesCnt    int `gorm:"default:0"`
	MarksCnt    int `gorm:"default:0"`
	CommentsCnt int `gorm:"default:0"`

	LastView time.Time `gorm:"autoCreateTime"`
	LastMark time.Time `gorm:"autoCreateTime"`

	// --- 基础字段
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
