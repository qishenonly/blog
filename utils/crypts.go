package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// EncryptPassword 加密密码
func EncryptPassword(pwd string) (string, error) {
	encrypt, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(encrypt), nil
}

// ComparePasswords 比较密码
func ComparePasswords(hashedPwd, plainPwd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
}
