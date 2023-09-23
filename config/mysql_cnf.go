package config

import "fmt"

// MysqlConfig mysql配置
type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DB       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	LogLevel string `yaml:"log_level"`
}

// DSN是Data Source Name的缩写，即数据源名称。它是一个包含了数据库连接信息的字符串，它包含了以下几个部分：
// 数据库类型（如mysql，postgres，sqlite等）
// 登录用户名
// 登录密码
// 主机地址
func (m *MysqlConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.User, m.Password, m.Host, m.Port, m.DB)
}
