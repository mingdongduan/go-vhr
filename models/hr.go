package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

var cache map[string][]Hr

type Hr struct {
	Model
	Name string `json:"name"`
	Phone string `json:"phone"`
	Telephone string `json:"telephone"`
	Address string `json:"address"`
	Enabled int `json:"enabled"`
	Username string `json:"username"`
	Password string `json:"password"`
	Userface string `json:"userface"`
	Remark string `json:"remark"`
}

func init() {
	cache = make(map[string][]Hr)
}

func MatchPassword(name, password string) bool {
	if len(name) == 0 || len(password) == 0 {
		return false
	}
	hr := GetHrByName(name)
	if len(hr.Password) == 0 {
		return false
	}
	return CheckPasswordHash(password, hr.Password)
}

func GetHrByName(name string) Hr {
	var hr Hr
	
	DB.Select("*").Where(" username = ?", name).Find(&hr)
	return hr
}

func ListHr() []Hr {
	var hrs []Hr
	if cache["hrs"] != nil {
		return cache["hrs"]
	}

	DB.Select("*").Where(" enabled = 1").Find(&hrs)
	cache["hrs"] = hrs
	return hrs
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}