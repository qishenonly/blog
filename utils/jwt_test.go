package utils

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"blog/config"
	"blog/global"
)

func InitCnf() {
	const ConfigFile = "../settings.yaml"
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

	// 将配置信息保存到全局变量中
	global.Config = c
}

func TestGenerateToken(t *testing.T) {
	InitCnf()

	token, err := GenerateToken(JwtPayload{
		UserId:   1,
		Username: "admin",
		Nickname: "管理员",
		Role:     "1",
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, token)
}

func TestParseToken(t *testing.T) {
	InitCnf()

	token, err := GenerateToken(JwtPayload{
		UserId:   1,
		Username: "admin",
		Nickname: "管理员",
		Role:     "1",
	})

	assert.Nil(t, err)
	assert.NotEmpty(t, token)

	claims, err := ParseToken(token)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), claims.UserId)
	assert.Equal(t, "admin", claims.Username)
	assert.Equal(t, "管理员", claims.Nickname)
	assert.Equal(t, "1", claims.Role)
}
