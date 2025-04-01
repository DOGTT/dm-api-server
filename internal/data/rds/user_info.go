// user_info.go
package rds

import (
	"context"
	"time"

	"gorm.io/gorm"
)

func init() {
	dbModelList = append(dbModelList, &UserInfo{})
}

// 用户信息
type UserInfo struct {
	Id       uint64 `gorm:"primaryKey;autoIncrement"`
	WeChatId string `gorm:"type:varchar(32);unique"`
	Phone    string `gorm:"type:varchar(16);"`

	Pets []*PetInfo `gorm:"many2many:user_pets;"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	// 软删除字段
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *UserInfo) TableName() string {
	return "user_info"
}

func (c *RDSClient) CreateUserInfo(ctx context.Context, userInfo *UserInfo) error {
	res := c.db.WithContext(ctx).Create(userInfo)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *RDSClient) GetUserInfoByID(ctx context.Context, id uint64) (*UserInfo, error) {
	var userInfo UserInfo
	res := c.db.WithContext(ctx).
		Where(sqlEqualId, id).
		First(&userInfo)
	if res.Error != nil {
		return nil, res.Error
	}
	return &userInfo, nil
}

func (c *RDSClient) GetUserInfoByWeChatID(ctx context.Context, wxID string) (*UserInfo, error) {
	var userInfo UserInfo
	res := c.db.WithContext(ctx).
		Where(sqlEqual(sqlFieldWeChatId), wxID).
		First(&userInfo)
	if res.Error != nil {
		return nil, res.Error
	}
	return &userInfo, nil
}

func (c *RDSClient) UpdateUserInfo(ctx context.Context, userInfo *UserInfo) error {
	res := c.db.WithContext(ctx).Save(userInfo)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *RDSClient) DeleteUserInfo(ctx context.Context, userInfo *UserInfo) error {
	res := c.db.WithContext(ctx).Delete(userInfo)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
