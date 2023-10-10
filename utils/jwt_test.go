package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	//InitCnf()

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
	//InitCnf()

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

func TestValidToken(t *testing.T) {
	//InitCnf()

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

	ok1, err := ValidToken(token)
	assert.Nil(t, err)
	assert.False(t, ok1)

	// 休眠5秒，让token过期
	//time.Sleep(5 * time.Second)

	// 验证token是否有效
	//ok2, err := ValidToken(token)
	//assert.NotNil(t, err)
	//assert.True(t, ok2)
}
