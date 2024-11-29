// user_info.go
package rds

import (
	"time"

	"github.com/lib/pq"
)

func init() {
	dbModelList = append(dbModelList, &UserInfo{})
}

// 用户信息
type UserInfo struct {
	ID        string `gorm:"type:varchar(22);primaryKey;"`
	WeiChatID string `gorm:"type:varchar(32)"`
	Name      string
	PIDMain   string
	PIDList   pq.StringArray

	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
