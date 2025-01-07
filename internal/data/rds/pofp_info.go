package rds

import (
	"context"
	"fmt"
	"time"

	base_api "github.com/DOGTT/dm-api-server/api/base"
	"github.com/lib/pq"
	"gorm.io/gorm/clause"
)

func init() {
	dbModelList = append(dbModelList, &PofpInfo{})
}

// 足迹点基本信息
type PofpInfo struct {
	UUID   string `gorm:"type:varchar(22);primaryKey;"`
	TypeId uint   `gorm:"index"`
	// - 归属的PetID
	PId uint64 `gorm:"index"`
	// - 关键内容
	Title   string         `gorm:"type:text;size:50;not null"`
	LngLat  string         `gorm:"type:geometry(Point,4326);not null"`
	Photos  pq.StringArray `gorm:"type:text[]"`
	Content string         `gorm:"type:text;size:1024;"`
	// Tags       pq.StringArray `gorm:"type:text[]"`
	// - 附属坐标信息
	PoiId   string                 `gorm:"type:varchar(32)"`
	Address string                 `gorm:"type:text;size:256"`
	PoiData map[string]interface{} `gorm:"type:jsonb"`
	// - 互动信息
	ViewsCnt    int       `gorm:"default:0"`
	LikesCnt    int       `gorm:"default:0"`
	MarksCnt    int       `gorm:"default:0"`
	CommentsCnt int       `gorm:"default:0"`
	LastView    time.Time `gorm:"autoCreateTime"`
	LastMark    time.Time `gorm:"autoCreateTime"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
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

func (c *RDSClient) CreatePofpInfo(ctx context.Context, info *PofpInfo) error {
	return c.db.WithContext(ctx).Create(info).Error
}

func (c *RDSClient) UpdatePofpInfo(ctx context.Context, info *PofpInfo) error {
	if info.UUID == "" {
		return fmt.Errorf("uuid is empty")
	}
	updateField := []string{}
	if info.Content != "" {
		updateField = append(updateField, "content")
	}
	if info.Title != "" {
		updateField = append(updateField, "title")
	}
	return c.db.WithContext(ctx).Model(&PofpInfo{}).Where("uuid = ?", info.UUID).
		Select(updateField).
		Updates(info).Error
}

func (c *RDSClient) DeletePofpInfo(ctx context.Context, uuid string, pid uint64) error {
	return c.db.WithContext(ctx).Where(&PofpInfo{UUID: uuid, PId: pid}).Delete(PofpInfo{}).Error
}

func (c *RDSClient) GetPofpInfo(ctx context.Context, uuid string) (*PofpInfo, error) {
	var info PofpInfo
	err := c.db.WithContext(ctx).Where("uuid = ?", uuid).First(&info).Error
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (c *RDSClient) ExistPofpInfo(ctx context.Context, uuid string) error {
	var count int64
	err := c.db.WithContext(ctx).Model(&PofpInfo{}).Where("uuid = ?", uuid).Count(&count).Error
	if err != nil {
		return err
	}
	if count < 1 {
		return fmt.Errorf("not found")
	}
	return nil
}

func (c *RDSClient) BatchQueryPofpInfoListByBound(ctx context.Context, typeIDs []uint, bc *base_api.BoundCoord) ([]*PofpInfo, error) {

	var results []*PofpInfo
	query := c.db.WithContext(ctx).Model(&PofpInfo{}).
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

func (c *RDSClient) IncreasePofpViewCount(ctx context.Context, uuid string, t InxType) (int, error) {
	var info PofpInfo
	field := InxTypeFieldName[t]
	// 锁定行
	if err := c.db.WithContext(ctx).Model(&info).Where("uuid = ?", uuid).
		Select(field).Clauses(clause.Locking{Strength: "UPDATE"}).First(&info).Error; err != nil {
		return 0, err
	}

	// 更新操作
	info.UUID = uuid
	resCount := 0
	switch t {
	case InxTypeView:
		info.ViewsCnt++
		resCount = info.ViewsCnt
	case InxTypeLike:
		info.LikesCnt++
		resCount = info.LikesCnt
	case InxTypeMark:
		info.MarksCnt++
		resCount = info.MarksCnt
	case InxTypeComment:
		info.CommentsCnt++
		resCount = info.CommentsCnt
	}
	if err := c.db.WithContext(ctx).Select(field).Save(&info).Error; err != nil {
		return 0, err
	}
	return resCount, nil
}
