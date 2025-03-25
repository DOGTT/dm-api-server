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

	// - 互动信息
	// 查看数
	ViewsCnt int `gorm:"default:0"`
	// 收藏数量
	StarsCnt int `gorm:"default:0"`
	// 到访宠物数量
	PetMarksCnt int `gorm:"default:0"`
	// 帖子数量
	PostsCnt int `gorm:"default:0"`
	// 最新查看时间
	LastView time.Time `gorm:"autoCreateTime"`
	// 最新宠物到访时间
	LastPetMark time.Time `gorm:"autoCreateTime"`
	// 最新互动时间
	LastInx time.Time `gorm:"autoCreateTime"`

	// 频道配置
	Config ChannelConfig `gorm:"type:jsonb"`
	// 个性标签列表
	CustomTags pq.StringArray `gorm:"type:text[]"`

	// --- 基础字段
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
