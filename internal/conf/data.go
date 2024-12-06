package conf

import "fmt"

// RDSConfig 表示 PostgreSQL 的连接配置
type RDSConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

// 生成 GORM 的连接字符串
func (c *RDSConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.Dbname, c.SSLMode)
}

type DataConfig struct {
	RDS *RDSConfig `yaml:"rds"`
}
