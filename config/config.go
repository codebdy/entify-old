package config

import (
	"fmt"

	"rxdrag.com/entify/consts"
)

const TABLE_NAME_MAX_LENGTH = 64

var c Config

type DbConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

type Config interface {
	getString(key string) string
	getBool(key string) bool
	getInt(key string) int
}

func GetString(key string) string {
	return c.getString(key)
}
func GetBool(key string) bool {
	return c.getBool(key)
}
func GetInt(key string) int {
	return c.getInt(key)
}
func GetDbConfig() DbConfig {
	var cfg DbConfig
	cfg.Driver = GetString(consts.DB_DRIVER)
	cfg.Database = GetString(consts.DB_DATABASE)
	cfg.Host = GetString(consts.DB_HOST)
	cfg.Port = GetString(consts.DB_PORT)
	cfg.User = GetString(consts.DB_USER)
	cfg.Password = GetString(consts.DB_PASSWORD)
	if cfg.Driver == "" {
		cfg.Driver = "mysql"
	}
	return cfg
}

func ServiceId() int {
	serviceId := c.getInt(consts.SERVICE_ID)
	if serviceId == 0 {
		return 1
	}
	return serviceId
}

func init() {
	c = newEnvConfig()
	fmt.Println("哈哈", c.getString(consts.DB_HOST))
	//if c.getString(consts.DB_HOST) == "" {
	//	c = newFileConfig()
	//}
}
