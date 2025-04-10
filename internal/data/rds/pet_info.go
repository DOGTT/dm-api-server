package rds

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
)

var (
	petModel = &PetInfo{}
)

func init() {
	dbModelList = append(dbModelList, petModel)
}

// 宠信息
type PetInfo struct {
	Id uint64 `gorm:"primaryKey;autoIncrement"`
	// 创建者用户id
	UId uint64 `gorm:"type:bigint;;column:uid"`
	// 状态
	Status uint8 `gorm:"type:smallint;"`
	// 物种
	Specie string `gorm:"type:varchar(5);"`
	// 名字
	Name string `gorm:"type:varchar(20);"`
	// 简介
	Intro string `gorm:"type:varchar(128);"`
	// 性别
	Gender uint8 `gorm:"type:smallint;"`
	// 生日
	BirthDate string `gorm:"type:varchar(10);"`
	// 头像
	AvatarId string `gorm:"type:varchar(255);"`
	// 体型
	Size string `gorm:"type:varchar(3);"`
	// 品种
	Breed string `gorm:"type:varchar(20);"`
	// 体重
	Weight int `gorm:"type:smallint;"`

	// 关联
	Users []*UserInfo `gorm:"many2many:user_pets;"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	// 软删除字段
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (p *PetInfo) TableName() string {
	return "pet_info"
}

func (c *RDSClient) CreatePetInfo(ctx context.Context, in *PetInfo) error {
	res := c.db.WithContext(ctx).Create(in)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *RDSClient) UpdatePetInfo(ctx context.Context, pet *PetInfo) error {
	if pet.Id == 0 {
		return fmt.Errorf("id is empty")
	}
	res := c.db.WithContext(ctx).Save(pet)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}

func (c *RDSClient) GetPetInfoByID(ctx context.Context, id uint) (*PetInfo, error) {
	var pet PetInfo
	err := c.db.WithContext(ctx).First(&pet, id).Error
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

func (c *RDSClient) GetPetCreatorId(ctx context.Context, id uint64) (uid uint64, err error) {
	result := &PetInfo{}
	err = c.db.WithContext(ctx).
		Select(sqlFieldUId).
		First(&result, id).Error
	if err != nil {
		return
	}
	uid = result.UId
	return
}
