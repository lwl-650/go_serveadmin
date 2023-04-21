package model

type WebRouter struct {
	// gorm.Model
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// 创建关联表名函数
func (WebRouter) TableName() string {
	// 返回表名
	return "WebRouter"
}
