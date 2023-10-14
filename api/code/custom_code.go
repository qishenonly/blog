package code

type CustomCode int

const (
	// 获取参数错误，一般是因为参数不符合要求（空值、长度不符合要求等）
	ParamError CustomCode = 41000
)
