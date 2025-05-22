package rds

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm/clause"

	api "github.com/DOGTT/dm-api-server/api/base"
	"gorm.io/gorm"
)

const (
	channelMaxPageSize = 1000
)

var (
	channelModel = &ChannelInfo{}
)

func init() {
	dbModelList = append(dbModelList, channelModel)
}

// 足迹频道基本信息
type ChannelInfo struct {
	Id uint64 `gorm:"primaryKey;autoIncrement"`
	// - 创建者的 Uid
	UId uint64 `gorm:"index;column:uid"`
	// 类型id
	TypeId uint32 `gorm:"index"`
	// - 关键内容
	Title string `gorm:"type:text;size:50;not null"`
	// 频道头像
	AvatarId string `gorm:"type:text"`
	// 简介
	Intro string `gorm:"type:text;size:1024;"`
	// 位置坐标
	LngLat string `gorm:"type:geometry(Point,4326);not null"`
	// 位置的关键兴趣点详情
	PoiDetail PoiDetail `gorm:"type:jsonb"`

	// 配置和状态子表
	Set ChannelSet `gorm:"foreignKey:Id;constraint:OnDelete:CASCADE"`
	// 互动统计子表
	Stats ChannelStats `gorm:"foreignKey:Id;constraint:OnDelete:CASCADE"`

	CommonTableTails
}

type PoiDetail struct {
	// - 附属坐标信息
	PoiId   string `json:"poi_id"`
	Address string `json:"address"`
}

// Implement the Scanner interface for PoiDetail
func (p *PoiDetail) Scan(value any) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, p)
}

// Implement the Valuer interface for PoiDetail
func (p PoiDetail) Value() (driver.Value, error) {
	return json.Marshal(p)
}

type ChannelFilter struct {
	TypeIDs    []int32
	BoundCoord *api.BoundCoord

	OrderByUpdateDesc bool

	Offset          uint32
	Limit           uint32
	orderAscColumn  []string
	orderDescColumn []string
}

func (f *ChannelFilter) check() error {
	// 限制单次查询最大条数
	if f.Limit > channelMaxPageSize {
		return fmt.Errorf("invalid limit: %d", f.Limit)
	}
	if f.Limit == 0 {
		f.Limit = limitDefault
	}
	if f.OrderByUpdateDesc {
		f.orderDescColumn = append(f.orderDescColumn, sqlFieldUpdatedAt)
	}
	return nil
}

// 生成查询条件，必须按照索引的顺序排列
func (f *ChannelFilter) generate(db *gorm.DB) error {
	if err := f.check(); err != nil {
		return err
	}
	if len(f.TypeIDs) != 0 {
		db = db.Where(sqlIn(sqlFieldTypeId), f.TypeIDs)
	}
	if f.BoundCoord != nil {
		db = db.Where(sqlWhereLngLatContain,
			f.BoundCoord.Sw.Lng, f.BoundCoord.Sw.Lat,
			f.BoundCoord.Ne.Lng, f.BoundCoord.Ne.Lat)
	}
	if f.Offset != 0 {
		db = db.Offset(int(f.Offset))
	}
	if f.Limit != 0 {
		db = db.Limit(int(f.Limit))
	}
	for _, col := range f.orderAscColumn {
		db = db.Order(sqlOrderAsc(col))
	}
	for _, col := range f.orderDescColumn {
		db = db.Order(sqlOrderDesc(col))
	}
	return nil
}

func (c *RDSClient) CreateChannelInfo(ctx context.Context, info *ChannelInfo) error {
	return c.db.WithContext(ctx).
		Preload(clause.Associations).
		Create(info).Error
}

func (c *RDSClient) UpdateChannelInfo(ctx context.Context, info *ChannelInfo) error {
	if info.Id == 0 {
		return fmt.Errorf("id is empty")
	}
	updateField := []string{}
	if info.Title != "" {
		updateField = append(updateField, sqlFieldTitle)
	}
	if info.Intro != "" {
		updateField = append(updateField, sqlFieldIntro)
	}
	if info.AvatarId != "" {
		updateField = append(updateField, sqlFieldAvatarId)
	}
	res := c.db.WithContext(ctx).Model(channelModel).
		Where(sqlEqualId, info.Id).
		Select(updateField).
		Updates(info)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}

func (c *RDSClient) DeleteChannelInfo(ctx context.Context, id uint64) error {
	channelModel := &ChannelInfo{
		Id: id,
	}
	res := c.db.WithContext(ctx).
		Select(clause.Associations).
		Delete(channelModel)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}

func (c *RDSClient) GetChannelInfo(ctx context.Context, id uint64) (result *ChannelInfo, err error) {
	err = c.db.WithContext(ctx).
		Preload(clause.Associations).
		Where(sqlEqualId, id).
		First(&result).Error
	if err != nil {
		return
	}
	return
}

func (c *RDSClient) GetChannelFullInfo(ctx context.Context, Id uint64) (result *ChannelInfo, err error) {
	err = c.db.WithContext(ctx).
		Preload(clause.Associations).
		Select([]string{sqlFieldAll, sqlSelectLngLat}).
		First(&result, Id).Error
	if err != nil {
		return
	}
	return
}

func (c *RDSClient) GetChannelCreatorId(ctx context.Context, id uint64) (uid uint64, err error) {
	result := &ChannelInfo{}
	err = c.db.WithContext(ctx).
		Select(sqlFieldUId).
		First(&result, id).Error
	if err != nil {
		return
	}
	uid = result.UId
	return
}

func (c *RDSClient) CountChannelInfo(ctx context.Context, f *ChannelFilter) (count int64, err error) {
	dbIns := c.db.WithContext(ctx).Model(channelModel)
	if err = f.generate(dbIns); err != nil {
		return
	}
	err = dbIns.Count(&count).Error
	return
}

func (c *RDSClient) ListChannelInfo(ctx context.Context, f *ChannelFilter) (results []*ChannelInfo, err error) {
	dbIns := c.db.WithContext(ctx).Model(channelModel)
	dbIns.Select([]string{sqlFieldAll, sqlSelectLngLat})
	if err = f.generate(dbIns); err != nil {
		return
	}
	err = dbIns.Scan(&results).Error
	return
}
