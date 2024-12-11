package rds

import (
	"context"
	"time"
)

func init() {
	dbModelList = append(dbModelList, &PetInfo{})
}

// 宠信息
type PetInfo struct {
	Id  uint `gorm:"primaryKey;autoIncrement"`
	UId uint `gorm:"index"` // 外键字段，指向 UserInfo 的 ID

	Name      string `gorm:"type:varchar(20);"`
	Gender    uint8  `gorm:"type:smallint;"`
	AvatarId  string `gorm:"type:varchar(255);"`
	Specie    string `gorm:"type:varchar(20);"`
	Breed     string `gorm:"type:varchar(20);"`
	Weight    int    `gorm:"type:smallint;"`
	BirthDate string `gorm:"type:date;"`
	// PhotoList pq.StringArray `gorm:"type:text[]"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (p *PetInfo) TableName() string {
	return "pet_info"
}

func (c *RDSClient) GetPetInfoByID(ctx context.Context, id uint) (*PetInfo, error) {
	var pet PetInfo
	err := c.db.WithContext(ctx).Where("id = ?", id).First(&pet).Error
	if err != nil {
		return nil, err
	}
	return &pet, nil
}

func (c *RDSClient) GetFirstPetInfoByUID(ctx context.Context, uid uint) (*PetInfo, error) {
	var pet PetInfo
	err := c.db.WithContext(ctx).Where("uid = ?", uid).First(&pet).Error
	if err != nil {
		return nil, err
	}
	return &pet, nil
}
