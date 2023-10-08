package config

// Config 是一个全局变量，用来保存配置信息
type Config struct {
	Mysql       MysqlConfig  `yaml:"mysql"`
	Logger      LoggerConfig `yaml:"logger"`
	System      SystemConfig `yaml:"system"`
	QQConfig    QQConfig     `yaml:"qq"`
	JwtConfig   JwtConfig    `yaml:"jwt"`
	EmailConfig EmailConfig  `yaml:"email"`
}
