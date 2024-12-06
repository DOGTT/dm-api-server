package rds

import (
	"time"

	"github.com/lib/pq"
)

func init() {
	dbModelList = append(dbModelList, &PetInfo{})
}

// 宠信息
type PetInfo struct {
	ID        string         `gorm:"type:varchar(22);primaryKey;"`
	UID       string         `gorm:"type:varchar(22);"`
	Name      string         `gorm:"type:varchar(20);"`
	Gender    uint8          `gorm:"type:smallint;"`
	Avatar    string         `gorm:"type:varchar(255);"`
	Specie    string         `gorm:"type:varchar(20);"`
	Breed     string         `gorm:"type:varchar(20);"`
	Weight    int            `gorm:"type:smallint;"`
	BrithDate string         `gorm:"type:date;"`
	PhotoList pq.StringArray `gorm:"type:text[]"`

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
