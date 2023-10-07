package _type

type TagsArray []string

// Scan 是一个实现了 sql.Scanner 接口的方法，用来将数据库中的 tags 字段的值转换为 TagsArray 类型
func (t *TagsArray) Scan(src interface{}) error {
	*t = TagsArray{}
	if src == nil {
		return nil
	}
	*t = append(*t, src.(string))
	return nil
}
