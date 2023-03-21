package model

type User struct {
	// gorm.Model
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// 创建关联表名函数
func (User) TableName() string {
	// 返回表名
	return "user"
}
