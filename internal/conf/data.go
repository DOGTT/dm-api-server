package conf

import "fmt"

// RDSConfig 表示 PostgreSQL 的连接配置
type RDSConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"db_name"`
	SSLMode  string `yaml:"ssl_mode"`
}

// 生成 GORM 的连接字符串
func (c *RDSConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DbName, c.SSLMode)
}

type FDSConfig struct {
	Endpoint  string `yaml:"endpoint"`
	AccessKey string `yaml:"access_key"`
	SecretKey string `yaml:"secret_key"`
}

type MapDataConfig struct {
	Endpoint string `yaml:"endpoint"`
	Key      string `yaml:"key"`
}

type DataConfig struct {
	RDS     *RDSConfig     `yaml:"rds"`
	FDS     *FDSConfig     `yaml:"fds"`
	MapData *MapDataConfig `yaml:"map_data"`
}
