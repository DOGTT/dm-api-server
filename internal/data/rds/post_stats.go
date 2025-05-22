package rds

import (
	"context"
	"time"

	"gorm.io/gorm"
)

var (
	postStatsModel = &PostStats{}
)

func init() {
	dbModelList = append(dbModelList, postStatsModel)
}

type PostStatsType int

const (
	PostStatsView PostStatsType = iota
	PostStatsStar
	PostStatsPost
	PostStatsPee
)

var (
	PostStatsFieldCntList  = []string{"views_cnt", "stars_cnt", "posts_cnt", "pee_cnt"}
	PostStatsFieldTimeList = []string{"", "last_star_at", "last_post_at", "last_pee_at"}
)

type PostStats struct {
	Id uint64 `gorm:"primaryKey;autoIncrement"`

	// -- 反应互动信息
	Reactions map[string]int `gorm:"type:jsonb"`
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

func (c *RDSClient) PostStatsIncrease(ctx context.Context, chanId uint64, st PostStatsType) error {
	updateField := map[string]any{}
	fieldName := PostStatsFieldCntList[st]
	updateField[fieldName] = gorm.Expr(fieldName + " + 1")
	if timeField := PostStatsFieldTimeList[st]; timeField != "" {
		updateField[timeField] = time.Now()
	}
	res := c.db.WithContext(ctx).Model(postStatsModel).
		Where(sqlEqualId, chanId).
		Updates(updateField)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}

func (c *RDSClient) PostStatsDecrease(ctx context.Context, chanId uint64, st PostStatsType) error {
	updateField := map[string]any{}
	fieldName := PostStatsFieldCntList[st]
	updateField[fieldName] = gorm.Expr(fieldName + " - 1")
	if timeField := PostStatsFieldTimeList[st]; timeField != "" {
		updateField[timeField] = time.Now()
	}
	res := c.db.WithContext(ctx).Model(postStatsModel).
		Where(sqlEqualId, chanId).
		Updates(updateField)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}
