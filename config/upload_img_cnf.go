package config

import (
	"fmt"
)

type UploadImgConfig struct {
	Url string `yaml:"url"`
}

// GetAddr 获取地址
func (s *UploadImgConfig) GetAddr() string {
	return fmt.Sprintf("http://%s", s.Url)
}
