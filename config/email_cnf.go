package config

type EmailConfig struct {
	Host             string `json:"host" yaml:"host"`
	Port             int    `json:"port" yaml:"port"`
	User             string `json:"user" yaml:"user"`
	Password         string `json:"password" yaml:"password"`
	DefaultFromEmail string `json:"default_from_email" yaml:"default_from_email"` // 默认发件人邮箱
	UseSSL           bool   `json:"use_ssl" yaml:"use_ssl"`
	UseTLS           bool   `json:"use_tls" yaml:"use_tls"`
}
