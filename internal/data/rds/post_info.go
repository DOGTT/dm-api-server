package rds

import (
	"context"
	"fmt"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

const (
	postMaxPageSize = 1000
)

var (
	postModel = &PostInfo{}
)

func init() {
	dbModelList = append(dbModelList, postModel)
}

// 足迹频道帖子
type PostInfo struct {
	// 评论唯一id
	Id uint64 `gorm:"primaryKey;autoIncrement"`
	// 创建者UId
	UId uint64 `gorm:"index"`
	// 底层频道id
	RootId uint64 `gorm:"index"`
	// 关联的上级帖子id, 空则为根帖子
	ParentId uint64 `gorm:"default:0"`
	// 帖子内容
	Content string `gorm:"type:text"`
	// 帖子图片
	PhotoIds pq.StringArray `gorm:"type:text[]"`

	// -- 动态互动信息
	// 添加的标签
	Tags []string `gorm:"type:text[]"`
	// 喜欢的数量
	LikesCnt int `gorm:"default:0"`

	// --- 基础字段
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	// 软删除字段
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (c *RDSClient) CreatePostInfo(ctx context.Context, info *PostInfo) error {
	return c.db.WithContext(ctx).
		Create(info).Error
}

func (c *RDSClient) UpdatePostInfo(ctx context.Context, info *PostInfo) error {
	if info.Id == 0 {
		return fmt.Errorf("id is empty")
	}
	updateField := []string{}
	if info.Content != "" {
		updateField = append(updateField, sqlFieldContent)
	}
	return c.db.WithContext(ctx).Model(postModel).
		Where(sqlEqualId, info.Id).
		Select(updateField).
		Updates(info).Error
}

func (c *RDSClient) DeletePostInfo(ctx context.Context, id uint64) error {
	return c.db.WithContext(ctx).
		Where(sqlEqualId, id).Delete(postModel).Error
}

func (c *RDSClient) GetPostInfo(ctx context.Context, id uint64) (result *PostInfo, err error) {
	err = c.db.WithContext(ctx).
		Where(sqlEqualId, id).
		First(&result).Error
	if err != nil {
		return
	}
	return
}

type PostFilter struct {
	IdFrom        uint64
	OrderByIdDesc bool
	RootId        uint64

	Offset          uint32
	Limit           uint32
	orderAscColumn  []string
	orderDescColumn []string
}

func (f *PostFilter) check() error {
	// 限制单次查询最大条数
	if f.Limit > postMaxPageSize {
		return fmt.Errorf("invalid limit: %d", f.Limit)
	}
	if f.Limit == 0 {
		f.Limit = limitDefault
	}
	if f.OrderByIdDesc {
		f.orderDescColumn = append(f.orderDescColumn, sqlFieldId)
	} else {
		f.orderAscColumn = append(f.orderAscColumn, sqlFieldId)
	}
	return nil
}

// 生成查询条件，必须按照索引的顺序排列
func (f *PostFilter) generate(db *gorm.DB) error {
	if err := f.check(); err != nil {
		return err
	}
	if f.IdFrom != 0 {
		db = db.Where(sqlMore(sqlFieldId), f.IdFrom)
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

func (c *RDSClient) ListPostInfo(ctx context.Context, f *PostFilter) (results []*PostInfo, err error) {

	dbIns := c.db.WithContext(ctx).Model(postModel)
	if err = f.generate(dbIns); err != nil {
		return
	}
	err = dbIns.Scan(&results).Error
	return
}

func (c *RDSClient) GetPostCreatorId(ctx context.Context, id uint64) (uid uint64, err error) {
	result := &PostInfo{}
	err = c.db.WithContext(ctx).
		Select(sqlFieldUId).
		Where(sqlEqualId, id).First(&result).Error
	if err != nil {
		return
	}
	uid = result.UId
	return
}
