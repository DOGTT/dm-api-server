package rds

import (
	"context"
	"fmt"
	"time"

	grpc_api "github.com/DOGTT/dm-api-server/api/grpc"
	"github.com/lib/pq"
	"gorm.io/gorm/clause"
)

func init() {
	dbModelList = append(dbModelList, &PofpInfo{})
}

// 足迹点基本信息
type PofpInfo struct {
	UUID   string `gorm:"type:varchar(32);primaryKey;"`
	TypeId uint   `gorm:"index"`
	// - 归属的PetID
	PId uint `gorm:"index"`
	// - 关键内容
	Title   string         `gorm:"type:text;size:50;not null"`
	LatLng  string         `gorm:"type:geometry(Point,4326);not null"`
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

func PointCoordToGeometry(p *grpc_api.PointCoord) string {
	return fmt.Sprintf("SRID=4326;POINT(%f %f)", p.Lng, p.Lat)
}

func PointCoordFromGeometry(s string) *grpc_api.PointCoord {
	var lat, lng float32
	_, _ = fmt.Sscanf(s, "SRID=4326;POINT(%f %f)", &lng, &lat) // 经度在前，纬度在后
	return &grpc_api.PointCoord{
		Lat: lat,
		Lng: lng,
	}
}

func (c *RDSClient) CreatePofpInfo(ctx context.Context, info *PofpInfo) error {
	return c.db.WithContext(ctx).Create(info).Error
}

func (c *RDSClient) UpdatePofpInfo(ctx context.Context, info *PofpInfo) error {
	return c.db.WithContext(ctx).Model(&PofpInfo{}).Where("uuid = ?", info.UUID).
		Select("title", "content").
		Updates(info).Error
}

func (c *RDSClient) DeletePofpInfo(ctx context.Context, uuid string) error {
	return c.db.WithContext(ctx).Where("uuid = ?", uuid).Delete(PofpInfo{}).Error
}

func (c *RDSClient) GetPofpInfo(ctx context.Context, uuid string) (*PofpInfo, error) {
	var info PofpInfo
	err := c.db.WithContext(ctx).Where("uuid = ?", uuid).First(&info).Error
	if err != nil {
		return nil, err
	}
	return &info, nil
}

func (c *RDSClient) BatchQueryPofpInfoListByBound(ctx context.Context, typeIDs []uint, bc *grpc_api.BoundCoord) ([]*PofpInfo, error) {

	var results []*PofpInfo
	query := c.db.WithContext(ctx).Model(&PofpInfo{}).
		Select("uuid, type_id, p_id, lat_lng, title").
		Where("ST_Contains(ST_MakeEnvelope(?, ?, ?, ?, 4326), lat_lng)",
			bc.Sw.Lng, bc.Sw.Lat, bc.Ne.Lng, bc.Ne.Lng)
	// 如果 typeIDs 不为 nil，则添加筛选条件
	if typeIDs != nil {
		query = query.Where("type_id IN ?", typeIDs)
	}
	err := query.Limit(100).Scan(&results).Error
	return results, err
}

func (c *RDSClient) IncreasePofpViewCount(ctx context.Context, uuid string) (int, error) {
	var info PofpInfo
	// 锁定行
	if err := c.db.WithContext(ctx).Model(&info).Where("uuid = ?", uuid).
		Select(InxTypeFieldName[InxTypeLike]).Clauses(clause.Locking{Strength: "UPDATE"}).First(&info).Error; err != nil {
		return 0, err
	}

	// 更新操作
	info.ViewsCnt++
	if err := c.db.WithContext(ctx).Save(&info).Error; err != nil {
		return 0, err
	}
	return info.ViewsCnt, nil
}
