package rds

import (
	"context"
	"time"

	"gorm.io/gorm"
)

var (
	PostReactModel = &PostReact{}
)

func init() {
	dbModelList = append(dbModelList, PostReactModel)
}

// 用户帖子回应 如: 有表情，文字类型
type PostReact struct {
	PostId uint64 `gorm:"index"`
	UId    uint64 `gorm:"index;column:uid"`
	// 反应内容，文字或者表情id
	ReactId string `gorm:"index;type:varchar(32)"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func (c *RDSClient) CreatePostReact(ctx context.Context, d *PostReact) error {
	return c.db.WithContext(ctx).Create(d).Error
}

func (c *RDSClient) DeletePostReact(ctx context.Context, d *PostReact) error {
	res := c.db.WithContext(ctx).Where(d).Delete(d)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}

func (c *RDSClient) ListPostReactByReactId(ctx context.Context, postId uint64, reactId string) (results []*PostReact, err error) {
	dbIns := c.db.WithContext(ctx).Model(PostReactModel)
	dbIns = dbIns.Select([]string{sqlFieldAll}).
		Where(sqlIn(sqlFieldPostId), postId).
		Where(sqlIn(sqlFieldReactId), reactId)
	err = dbIns.Scan(&results).Error
	return
}
