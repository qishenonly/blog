package utils

import (
	"blog/global"
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

// jwt payload
type JwtPayload struct {
	UserId   int64  `json:"user_id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	// 权限  1-管理员 2-普通用户 3-游客
	Role string `json:"role"`
}

// jwt claims
type Claims struct {
	JwtPayload
	jwt.StandardClaims
}

var MySecret []byte

// GenerateToken 生成token
func GenerateToken(payload JwtPayload) (string, error) {
	MySecret = []byte(global.Config.Jwt.Secret)

	claims := Claims{
		payload,
		jwt.StandardClaims{
			// 过期时间  默认2小时
			ExpiresAt: jwt.At(time.Now().Add(time.Duration(global.Config.Jwt.Expire) * time.Hour)),
			Issuer:    global.Config.Jwt.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(MySecret)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*Claims, error) {
	MySecret = []byte(global.Config.Jwt.Secret)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		global.Logger.Errorf("解析token失败: %v", err)
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("解析token失败")
}
