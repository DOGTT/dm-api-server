package rds

/*
* 1. 所有数据库字段在rds层收敛
* 2. 对常用数据操作方法进行原子化合理抽象
 */

import (
	"context"

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

// TransactionClient 结构体（继承 RDSClient）
type TransactionClient struct {
	RDSClient
	tx *gorm.DB
}

// NewTransaction 开启事务，并返回 TransactionClient 实例
func (c *RDSClient) NewTransaction(ctx context.Context) (*TransactionClient, error) {
	tx := c.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &TransactionClient{
		RDSClient: RDSClient{db: tx},
		tx:        tx,
	}, nil
}

// Commit 提交事务
func (tc *TransactionClient) Commit() error {
	return tc.tx.Commit().Error
}

// Rollback 回滚事务
func (tc *TransactionClient) Rollback() error {
	return tc.tx.Rollback().Error
}
