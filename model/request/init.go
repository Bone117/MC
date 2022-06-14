package request

import "fmt"

type InitDB struct {
	DBType   string `json:"dbType"` // 数据库类型
	Host     string `json:"host"`   // 服务器地址
	Port     string `json:"port"`   // 端口
	UserName string `json:"username" binding:"required"`
	Password string `json:"password"`
	DBName   string `json:"db_name" binding:"required"`
}

func (i *InitDB) MysqlEmptyDsn() string {
	if i.Host == "" {
		i.Host = "127.0.0.1"
	}
	if i.Port == "" {
		i.Port = "3306"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", i.UserName, i.Password, i.Host, i.Port)
}
