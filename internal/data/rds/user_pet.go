package rds

import (
	"context"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func init() {
	dbModelList = append(dbModelList, &UserPet{})
}

type UserPet struct {
	UId uint64 `gorm:"primaryKey;column:uid"`
	PId uint64 `gorm:"primaryKey;column:pid"`
	// 宠物的称呼 aa的bb
	PetTitle string `gorm:"type:varchar(16);"`

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
	res := c.db.WithContext(ctx).Save(userPet)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (c *RDSClient) DeleteUserPet(ctx context.Context, userPet *UserPet) error {
	res := c.db.WithContext(ctx).Delete(userPet)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
