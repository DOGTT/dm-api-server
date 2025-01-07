package rds

import (
	"context"
	"time"

	"gorm.io/gorm"
)

func init() {
	dbModelList = append(dbModelList, &UserPofpIxnRecord{})
}

type InxType uint

const (
	InxTypeView InxType = iota
	InxTypeLike         // 喜欢 -- 状态型
	InxTypeMark         // 标记过 -- 可重复
)

var (
	InxTypeFieldName = []string{"views_cnt", "likes_cnt", "marks_cnt", "comments_cnt"}
)

// 足迹点个互动记录 喜欢/踩过/评论
type UserPofpIxnRecord struct {
	PId      uint64 `gorm:"index"`
	PofpUUID string `gorm:"index"`
	//
	IntType InxType `gorm:"type:int;default:0"` // 0-点赞,1-收藏

	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (c *RDSClient) CreatePofpIxnRecord(ctx context.Context, d *UserPofpIxnRecord) error {
	return c.db.WithContext(ctx).Create(d).Error
}

func (c *RDSClient) CreatePofpIxnRecordWithCount(ctx context.Context, d *UserPofpIxnRecord) error {

	return c.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var info PofpInfo
		field := InxTypeFieldName[d.IntType]
		if err := tx.Model(&PofpInfo{}).Where("uuid = ?", d.PofpUUID).
			Select(field).First(&info).Error; err != nil {
			return err
		}
		if err := tx.Create(d).Error; err != nil {
			return err
		}
		info.UUID = d.PofpUUID
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
