package core

import (
	"fmt"
	"log"
	"os"

	"github.com/qishenonly/blog/config"
	"github.com/qishenonly/blog/global"

	"gopkg.in/yaml.v2"
)

// InitCnf 用来初始化配置信息
func InitCnf() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}

	// 读取配置文件
	yamlCnf, err := os.ReadFile(ConfigFile)
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败: %s", err))
	}

	// 解析配置文件
	err = yaml.Unmarshal(yamlCnf, c)
	if err != nil {
		log.Fatal("解析配置文件失败: ", err)
	}
	log.Println("配置文件解析成功")

	// 将配置信息保存到全局变量中
	global.Config = c
}
