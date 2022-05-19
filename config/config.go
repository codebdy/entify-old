package config

import "rxdrag.com/entify/consts"

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
	getDbConfig() DbConfig
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
	return c.getDbConfig()
}

func ServiceId() int {
	return c.getInt(consts.SERVICE_ID)
}

func init() {
	c = newEnvConfig()
	if c.getString(consts.DB_HOST) == "" {
		c = newFileConfig()
	}
}
