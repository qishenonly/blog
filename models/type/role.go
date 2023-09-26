package _type

import "encoding/json"

type Role int

const (
	// 管理员
	RoleAdmin Role = 1

	// 普通用户
	RoleUser Role = 2

	// 游客
	RoleTourist Role = 3

	// 被封禁的用户
	RoleBannedUser Role = 4
)

// String 是一个实现了 Stringer 接口的方法，用来将 Role 类型转换为字符串
func (r Role) String() string {
	switch r {
	case RoleAdmin:
		return "管理员"
	case RoleUser:
		return "普通用户"
	case RoleTourist:
		return "游客"
	case RoleBannedUser:
		return "被封禁的用户"
	default:
		return "未知"
	}
}

// MarshalJSON 是一个实现了 json.Marshaler 接口的方法，用来将 Role 类型转换为 JSON 格式
func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}
