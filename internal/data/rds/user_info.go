// user_info.go
package rds

import (
	"context"
	"time"
)

// var validate = validator.New()

func init() {
	dbModelList = append(dbModelList, &UserInfo{})
}

// 用户信息
type UserInfo struct {
	Id       uint64 `gorm:"primaryKey;autoIncrement"`
	WeChatId string `gorm:"type:varchar(32);unique"`
	Phone    string `gorm:"type:varchar(16);"`
	// 宠物的称呼 aa的bb
	PetTitle string `gorm:"type:varchar(6);"`

	// 绑定的宠物列表
	Pets []PetInfo `gorm:"many2many:user_pets;"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// func (u *UserInfo) BeforeCreate(tx *gorm.DB) (err error) {
// 	if err := validate.Struct(u); err != nil {
// 		return err
// 	}
// 	return nil
// }

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

func (c *RDSClient) GetUserInfoByID(ctx context.Context, id string) (*UserInfo, error) {
	var userInfo UserInfo
	res := c.db.WithContext(ctx).Where("id = ?", id).First(&userInfo)
	if res.Error != nil {
		return nil, res.Error
	}
	return &userInfo, nil
}

func (c *RDSClient) GetUserInfoByWeChatID(ctx context.Context, wxID string) (*UserInfo, error) {
	var userInfo UserInfo
	res := c.db.WithContext(ctx).Preload("Pets").Where("we_chat_id = ?", wxID).First(&userInfo)
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
