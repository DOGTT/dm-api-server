package conf

import "fmt"

// PostgresConfig 表示 PostgreSQL 的连接配置
type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

// 生成 GORM 的连接字符串
func (config *PostgresConfig) ConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.Dbname, config.SSLMode)
}

type RDSConfig struct {
	*PostgresConfig
}

type DataConfig struct {
	RDS *RDSConfig `yaml:"rds"`
}
