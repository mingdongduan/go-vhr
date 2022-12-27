package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-vhr/common"
	"go-vhr/tools"
)

var DB *gorm.DB

type Model struct {
	ID	uint	`gorm:"primary_key" json:"id"`
}

func init() {
	Connect()
}

func Connect() error {
	mysql := common.GetMysqlConf()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", mysql.Username, mysql.Password, mysql.Server, mysql.Port, mysql.Database)
	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		tools.Logger().Errorln(err)
		panic("数据库连接失败!")
	}
	DB.SingularTable(true)
	DB.LogMode(true)
	DB.SetLogger(tools.Logger())
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	InitConfig()
	return nil
}

func Execute(sql string) error {
	return DB.Exec(sql).Error
}
func CloseDB() {
	defer DB.Close()
}
