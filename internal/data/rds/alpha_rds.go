package rds

import (
	"github.com/DOGTT/dm-api-server/internal/conf"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbModelList = []any{}

type RDSClient struct {
	db *gorm.DB
}

func New(conf *conf.RDSConfig) (c *RDSClient, err error) {
	c = &RDSClient{}
	c.db, err = gorm.Open(postgres.Open(conf.ConnectionString()), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return
	}
	err = c.db.AutoMigrate(dbModelList...)
	return
}
