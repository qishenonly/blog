package utils

import (
	"errors"
	"time"

	"github.com/qishenonly/blog/global"

	"github.com/dgrijalva/jwt-go/v4"
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

var jwtSecret []byte

// GenerateToken 生成token
func GenerateToken(payload JwtPayload) (string, error) {
	jwtSecret = []byte(global.Config.Jwt.Secret)

	claims := Claims{
		payload,
		jwt.StandardClaims{
			// 过期时间  默认2小时
			ExpiresAt: jwt.At(time.Now().Add(time.Duration(global.Config.Jwt.Expire) * time.Hour)),
			Issuer:    global.Config.Jwt.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*Claims, error) {
	jwtSecret = []byte(global.Config.Jwt.Secret)

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
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

// ValidToken 校验token是否过期
func ValidToken(tokenString string) (bool, error) {
	jwtSecret = []byte(global.Config.Jwt.Secret)

	tokenClaims, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			expiresTime := (claims.ExpiresAt).Unix()
			now := time.Now().Unix()
			if now > expiresTime {
				//token过期了
				return true, nil
			}
			return false, nil
		}
	}
	return true, err
}

// GetUserIdFromToken 从token中获取用户id
func GetUserIdFromToken(tokenString string) (int64, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return 0, err
	}
	return claims.UserId, nil
}

// GetUsernameFromToken 从token中获取用户名
func GetUsernameFromToken(tokenString string) (string, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return "", err
	}
	return claims.Username, nil
}

// GetNicknameFromToken 从token中获取昵称
func GetNicknameFromToken(tokenString string) (string, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return "", err
	}
	return claims.Nickname, nil
}

// GetRoleFromToken 从token中获取权限
func GetRoleFromToken(tokenString string) (string, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return "", err
	}
	return claims.Role, nil
}
