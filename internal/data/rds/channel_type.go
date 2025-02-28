// poi_type_info.go
package rds

import (
	"context"
	"time"
)

func init() {
	dbModelList = append(dbModelList, &ChannelTypeInfo{})
}

// 足迹频道类型信息
type ChannelTypeInfo struct {
	Id int16 `gorm:"primaryKey;autoIncrement"`
	// 足迹频道类型名称, 如: 探险, 小憩, 溜溜
	Name string `gorm:"type:varchar(12)"`
	// 覆盖半径, 以米为单位
	CoverageRadius int `gorm:"type:smallint;"`
	// 主题色, 16进制
	// 如: #FF0000
	ThemeColor string `gorm:"type:varchar(12)"`
	// 创建时间
	CreatedAt time.Time `gorm:"autoCreateTime"`
	// 更新时间
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (c *RDSClient) ListPofpTypeInfo(ctx context.Context) (list []*ChannelTypeInfo, err error) {
	// 全表查询
	list = make([]*ChannelTypeInfo, 0)
	err = c.db.Model(&ChannelTypeInfo{}).Find(&list).Error
	return
}
