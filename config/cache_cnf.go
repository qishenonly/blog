package config

type CacheConfig struct {
	DirPath      string `yaml:"dir_path"`
	DataFileSize int64  `yaml:"data_file_size"`
	SyncWrite    bool   `yaml:"sync_write"`
}
