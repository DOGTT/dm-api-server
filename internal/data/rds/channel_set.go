package rds

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

var (
	channelSetModel = &ChannelSet{}
)

func init() {
	dbModelList = append(dbModelList, channelSetModel)
}

// 频道配置 TODO
type ChannelConfig struct {
	MaxUserCount int `json:"max_user_count"`
}

// Implement the Scanner interface for ChannelConfig
func (p *ChannelConfig) Scan(value any) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, p)
}

// Implement the Valuer interface for ChannelConfig
func (p ChannelConfig) Value() (driver.Value, error) {
	return json.Marshal(p)
}

type ChannelSet struct {
	Id uint64 `gorm:"primaryKey;autoIncrement"`

	// 频道配置
	Config ChannelConfig `gorm:"type:jsonb"`
	// 个性标签列表
	CustomTags pq.StringArray `gorm:"type:text[]"`

	// --- 基础字段
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	// 软删除字段
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
