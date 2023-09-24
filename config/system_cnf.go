package config

import "fmt"

// SystemConfig 系统配置
type SystemConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

// GetAddr 获取系统地址
func (s *SystemConfig) GetAddr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

// GetEnv 获取系统环境
func (s *SystemConfig) GetEnv() string {
	return s.Env
}
