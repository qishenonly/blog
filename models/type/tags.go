package _type

import "errors"

type TagsArray []string

// Scan 是一个实现了 sql.Scanner 接口的方法，用来将数据库中的 tags 字段的值转换为 TagsArray 类型
func (t *TagsArray) Scan(src interface{}) error {
	if src == nil {
		*t = []string{} // 如果数据库字段为 null，将 TagsArray 设置为空切片
		return nil
	}

	// 检查 src 是否为 []uint8 类型
	byteSrc, ok := src.([]byte)
	if !ok {
		return errors.New("Type assertion to []byte failed")
	}

	// 将 []byte 转换为字符串
	*t = TagsArray{string(byteSrc)}
	return nil
}
