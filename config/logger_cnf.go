package config

// LoggerConfig 日志配置
type LoggerConfig struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show_line"`      // 显示行
	LogInConsole bool   `yaml:"log_in_console"` // 是否将日志输出到控制台
}
