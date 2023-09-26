package _type

import "encoding/json"

type RegisterSource int

const (
	// 从邮箱注册
	RegisterFromEmail RegisterSource = 1

	// 从 QQ 注册
	RegisterFromQQ RegisterSource = 2

	// 从 GitHub 注册
	RegisterFromGitHub RegisterSource = 3

	// 从 Gitee 注册
	RegisterFromGitee RegisterSource = 4
)

// String 是一个实现了 Stringer 接口的方法，用来将 RegisterSource 类型转换为字符串
func (r RegisterSource) String() string {
	switch r {
	case RegisterFromEmail:
		return "邮箱"
	case RegisterFromQQ:
		return "QQ"
	case RegisterFromGitHub:
		return "GitHub"
	case RegisterFromGitee:
		return "Gitee"
	default:
		return "未知"
	}
}

// MarshalJSON 是一个实现了 json.Marshaler 接口的方法，用来将 RegisterSource 类型转换为 JSON 格式
func (r RegisterSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}
