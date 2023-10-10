package utils

import "math/rand"

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	intBytes    = "0123456789"
)

// RandString 是一个返回随机字符串的函数
func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// RandCode 是一个返回随机整数的函数
func RandCode(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = intBytes[rand.Intn(len(intBytes))]
	}
	return string(b)
}
