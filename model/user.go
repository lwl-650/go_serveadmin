package model

import "fmt"

type User struct {
	// gorm.Model
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Pass  string `json:"pass"`
	Token string `json:"token"`
}

func init() {
	fmt.Println("vi+++++++++++++++++++++++++++: ------------------------------------------")
}

// 创建关联表名函数
func (User) TableName() string {
	// 返回表名
	return "user"
}
