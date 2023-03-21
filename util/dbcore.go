package util

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var err error

type Sql struct {
	root     string
	password string
	local    string
	sqlname  string
}

// host: 127.0.0.1
//
//	user: root
//	dbname: goserve
//	pwd: 123456
var getsql = &Sql{
	root:     ReadeYaml("mysql.user"),
	password: ReadeYaml("mysql.pwd"),
	local:    ReadeYaml("mysql.host"),
	sqlname:  ReadeYaml("mysql.dbname"),
}

func init() {

	// 用户名:密码@tcp(ip:port)/数据库?charset=utf8mb4&parseTime=True&loc=Local
	dsn := getsql.root + ":" + getsql.password + "@tcp(" + getsql.local + ")/" + getsql.sqlname + "?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		fmt.Println(err)
	}
}
