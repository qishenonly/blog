package cmd

import "flag"

type Options struct {
	DB bool
}

// ParseOptions 解析命令行参数
func ParseOptions() *Options {
	opt := &Options{}
	flag.BoolVar(&opt.DB, "db", false, "初始化数据库")
	flag.Parse()
	return opt
}

// IsStopWebServer 是否停止 Web 服务
func IsStopWebServer(opt Options) bool {
	return opt.DB
}

// InitOptions 初始化命令行参数
func InitOptions(opt Options) {
	if opt.DB {
		MakeMigrations()
	}
}
