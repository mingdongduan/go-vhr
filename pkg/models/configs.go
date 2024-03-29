package models

type Config struct {
	ID        uint   `gorm:"primary_key" json:"id"`
	ConfName  string `json:"conf_name"`
	ConfKey   string `json:"conf_key"`
	ConfValue string `json:"conf_value"`
}

func InitConfig() {
	_ = FindConfigs()
}

func FindConfigs() []Config {
	var config []Config
	DB.Find(&config)
	return config
}
