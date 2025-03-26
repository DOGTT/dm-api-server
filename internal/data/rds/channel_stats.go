package rds

import (
	"context"
	"time"

	"gorm.io/gorm"
)

var (
	channelStatsModel = &ChannelStats{}
)

func init() {
	dbModelList = append(dbModelList, channelStatsModel)
}

type ChannelStats struct {
	Id uint64 `gorm:"primaryKey;autoIncrement"`

	// - 互动信息
	// 查看数
	ViewsCnt int `gorm:"default:0"`
	// 收藏数量
	StarsCnt int `gorm:"default:0"`
	// 帖子数量
	PostsCnt int `gorm:"default:0"`
	// 宠物到访次数
	PeeCnt int `gorm:"default:0"`

	// 收藏数量
	LastStarAt time.Time `gorm:"autoCreateTime"`
	// 最新互动时间
	LastPostAt time.Time `gorm:"autoCreateTime"`
	// 最新宠物到访时间
	LastPeeAt time.Time `gorm:"autoCreateTime"`

	// --- 基础字段
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	// 软删除字段
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ChannelStatsType int

const (
	ChannelStatsView ChannelStatsType = iota
	ChannelStatsStar
	ChannelStatsPost
	ChannelStatsPee
)

var (
	channelStatsFieldCntList  = []string{"views_cnt", "stars_cnt", "posts_cnt", "pee_cnt"}
	channelStatsFieldTimeList = []string{"", "last_star_at", "last_post_at", "last_pee_at"}
)

func (c *RDSClient) ChannelStatsIncrease(ctx context.Context, chanId uint64, st ChannelStatsType) error {
	updateField := map[string]any{}
	fieldName := channelStatsFieldCntList[st]
	updateField[fieldName] = gorm.Expr(fieldName + " + 1")
	if timeField := channelStatsFieldTimeList[st]; timeField != "" {
		updateField[timeField] = time.Now()
	}
	return c.db.WithContext(ctx).Model(channelStatsModel).
		Where(sqlEqualId, chanId).
		Updates(updateField).Error
}

func (c *RDSClient) ChannelStatsDecrease(ctx context.Context, chanId uint64, st ChannelStatsType) error {
	updateField := map[string]any{}
	fieldName := channelStatsFieldCntList[st]
	updateField[fieldName] = gorm.Expr(fieldName + " - 1")
	if timeField := channelStatsFieldTimeList[st]; timeField != "" {
		updateField[timeField] = time.Now()
	}
	return c.db.WithContext(ctx).Model(channelStatsModel).
		Where(sqlEqualId, chanId).
		Updates(updateField).Error
}
