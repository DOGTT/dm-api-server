// poi_table.go
package rdb

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Location struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"type:varchar(100)"`
	Coordinates string `gorm:"type:geometry(Point,4326)"` // 使用 PostGIS 的 geometry 类型
}

func main() {
	dsn := "host=localhost user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移
	db.AutoMigrate(&Location{})
}
