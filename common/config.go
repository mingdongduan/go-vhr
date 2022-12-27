package common

import (
	"encoding/json"
	"go-vhr/tools"
	"io/ioutil"
)

type Mysql struct {
	Server string
	Port string
	Database string
	Username string
	Password string
}

func GetMysqlConf() *Mysql {
	var mysql = &Mysql{}

	exist, _ := tools.IsFileExist(MysqlConf)
	if !exist {
		tools.Logger().Errorln("mysql config file is not exist")
		return mysql
	}

	info, err := ioutil.ReadFile(MysqlConf)
	if err != nil {
		tools.Logger().Errorln("read mysql config file failed")
		return mysql
	}
	err = json.Unmarshal(info, mysql)
	if err != nil {
		tools.Logger().Errorln("unmarshal mysql config json file failed!")
	}

	return mysql
}