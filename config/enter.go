package config

// Config 是一个全局变量，用来保存配置信息
type Config struct {
	Mysql     MysqlConfig     `yaml:"mysql"`
	Cache     CacheConfig     `yaml:"flydb"`
	Logger    LoggerConfig    `yaml:"logger"`
	System    SystemConfig    `yaml:"system"`
	QQ        QQConfig        `yaml:"qq"`
	Jwt       JwtConfig       `yaml:"jwt"`
	Email     EmailConfig     `yaml:"email"`
	UploadImg UploadImgConfig `yaml:"upload"`
}
