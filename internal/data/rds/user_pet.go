package rds

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	userPetModel = &UserPet{}
)

func init() {
	dbModelList = append(dbModelList, userPetModel)
}

type UserPet struct {
	UId uint64 `gorm:"primaryKey;column:uid"`
	PId uint64 `gorm:"primaryKey;column:pid"`
	// 宠物对用户的称呼 aa的bb
	PetTitle string `gorm:"type:varchar(16);"`
	// 用户宠物的关系状态
	PetStatus uint8 `gorm:"type:smallint;"`
	// 关联模型（方便查询）
	User *UserInfo `gorm:"foreignKey:UId"`
	Pet  *PetInfo  `gorm:"foreignKey:PId"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	// 软删除字段
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (t *UserPet) TableName() string {
	return "user_pets"
}

func (c *RDSClient) CreateUserPet(ctx context.Context, userPet *UserPet) error {
	res := c.db.WithContext(ctx).
		Create(userPet)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *RDSClient) GetPetIdsFromUserId(ctx context.Context, uid uint64) (petIds []uint64, err error) {
	userPets := []*UserPet{}
	res := c.db.WithContext(ctx).
		Select(sqlFieldPId).
		Where(sqlEqual(sqlFieldUId), uid).
		Scan(&userPets)
	if res.Error != nil {
		return nil, res.Error
	}
	if len(userPets) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	petIds = make([]uint64, 0)
	for _, up := range userPets {
		petIds = append(petIds, up.PId)
	}
	return petIds, nil
}

func (c *RDSClient) ListUserPetByUId(ctx context.Context, uid uint64) ([]*UserPet, error) {
	userPets := []*UserPet{}
	res := c.db.WithContext(ctx).
		Preload(clause.Associations).
		Where(sqlEqual(sqlFieldUId), uid).
		Scan(&userPets)
	if res.Error != nil {
		return nil, res.Error
	}
	if len(userPets) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return userPets, nil
}

func (c *RDSClient) ListUserPetByPId(ctx context.Context, pid uint64) ([]*UserPet, error) {
	userPets := []*UserPet{}
	res := c.db.WithContext(ctx).
		Preload(clause.Associations).
		Where(sqlEqual(sqlFieldPId), pid).
		Scan(&userPets)
	if res.Error != nil {
		return nil, res.Error
	}
	if len(userPets) == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	return userPets, nil
}

func (c *RDSClient) UpdateUserPet(ctx context.Context, userPet *UserPet) error {
	if userPet.UId == 0 || userPet.PId == 0 {
		return fmt.Errorf("id is empty")
	}
	updateField := []string{}
	if userPet.PetTitle != "" {
		updateField = append(updateField, "pet_title")
	}
	if userPet.PetStatus != 0 {
		updateField = append(updateField, "pet_status")
	}
	res := c.db.WithContext(ctx).Model(userPetModel).
		Where(sqlEqual(sqlFieldUId), userPet.UId).
		Where(sqlEqual(sqlFieldUId), userPet.UId).
		Select(updateField).
		Updates(userPet)
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return res.Error
}

func (c *RDSClient) DeleteUserPet(ctx context.Context, userPet *UserPet) error {
	res := c.db.WithContext(ctx).Delete(userPet)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
