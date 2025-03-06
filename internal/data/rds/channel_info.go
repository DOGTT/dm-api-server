package rds

import (
	"context"
	"fmt"
	"time"

	base_api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/lib/pq"
)

func init() {
	dbModelList = append(dbModelList, &ChannelInfo{})
}

// 足迹频道基本信息
type ChannelInfo struct {
	UUID string `gorm:"type:varchar(22);primaryKey;"`
	// 类型id
	TypeId uint16 `gorm:"index"`
	// - 创建者的 Uid
	UId uint64 `gorm:"index"`
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

	Stats ChannelStats `gorm:"foreignKey:UUID"`
	// --- 以上是基础静态信息
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type ChannelStats struct {
	UUID string `gorm:"type:varchar(22);primaryKey;"`

	UserIds Uint64Array `gorm:"type:bigint[]"`
	PetIds  Uint64Array `gorm:"type:bigint[]"`

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

	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type PoiDetail struct {
	// - 附属坐标信息
	PoiId   string `json:"poi_id"`
	Address string `json:"address"`
}

// 频道配置
type ChannelConfig struct {
}

func PointCoordToGeometry(p *base_api.PointCoord) string {
	if p == nil {
		return ""
	}
	return fmt.Sprintf("SRID=4326;POINT(%f %f)", p.Lng, p.Lat)
}

func PointCoordFromGeometry(s string) *base_api.PointCoord {
	// 将十六进制字符串转换为字节切片
	var lat, lng float32
	_, _ = fmt.Sscanf(string(s), "POINT(%f %f)", &lng, &lat) // 经度在前，纬度在后
	return &base_api.PointCoord{
		Lat: lat,
		Lng: lng,
	}
}

func (c *RDSClient) CreateChannelInfo(ctx context.Context, info *ChannelInfo) error {
	return c.db.WithContext(ctx).Create(info).Error
}

func (c *RDSClient) UpdateChannelInfo(ctx context.Context, info *ChannelInfo) error {
	if info.UUID == "" {
		return fmt.Errorf("uuid is empty")
	}
	updateField := []string{}
	if info.Intro != "" {
		updateField = append(updateField, "intro")
	}
	if info.Title != "" {
		updateField = append(updateField, "title")
	}
	return c.db.WithContext(ctx).Model(&ChannelInfo{}).Where("uuid = ?", info.UUID).
		Select(updateField).
		Updates(info).Error
}

func (c *RDSClient) DeleteChannelInfo(ctx context.Context, uuid string) error {
	return c.db.WithContext(ctx).Where(&ChannelInfo{UUID: uuid}).Delete(ChannelInfo{}).Error
}

func (c *RDSClient) GetChannelInfo(ctx context.Context, uuid string) (*ChannelInfo, error) {
	var info ChannelInfo
	err := c.db.WithContext(ctx).Where("uuid = ?", uuid).First(&info).Error
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (c *RDSClient) ExistChannelInfo(ctx context.Context, uuid string) error {
	var count int64
	err := c.db.WithContext(ctx).Model(&ChannelInfo{}).Where("uuid = ?", uuid).Count(&count).Error
	if err != nil {
		return err
	}
	if count < 1 {
		return fmt.Errorf("not found")
	}
	return nil
}

func (c *RDSClient) BatchQueryChannelInfoListByBound(ctx context.Context, typeIDs []uint, bc *base_api.BoundCoord) ([]*ChannelInfo, error) {

	var results []*ChannelInfo
	query := c.db.WithContext(ctx).Model(&ChannelInfo{}).
		Select("uuid, type_id, p_id, ST_AsText(lng_lat) AS lng_lat, title").
		Where("ST_Contains(ST_MakeEnvelope(?, ?, ?, ?, 4326), lng_lat)",
			bc.Sw.Lng, bc.Sw.Lat, bc.Ne.Lng, bc.Ne.Lat)
	// 如果 typeIDs 不为 nil，则添加筛选条件
	if len(typeIDs) > 0 {
		query = query.Where("type_id IN ?", typeIDs)
	}
	err := query.Limit(100).Scan(&results).Error
	return results, err
}
