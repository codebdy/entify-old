package config

import (
	"strconv"

	"github.com/spf13/viper"
	"rxdrag.com/entify/consts"
)

type EnvConfig struct {
	v *viper.Viper
}

const (
	TRUE  = "true"
	FALSE = "false"
)

func newEnvConfig() *EnvConfig {
	var e EnvConfig
	e.v = viper.New()
	e.v.SetEnvPrefix(consts.CONFIG_PREFIX)
	e.v.BindEnv(consts.DB_DRIVER)
	e.v.BindEnv(consts.DB_USER)
	e.v.BindEnv(consts.DB_PASSWORD)
	e.v.BindEnv(consts.DB_HOST)
	e.v.BindEnv(consts.DB_PORT)
	e.v.BindEnv(consts.DB_DATABASE)
	e.v.BindEnv(consts.ID)

	e.v.SetDefault(consts.SERVICE_ID, "1")
	e.v.SetDefault(consts.DB_DRIVER, "mysql")
	return &e
}

func (e *EnvConfig) getString(key string) string {
	return e.v.Get(key).(string)
}

func (e *EnvConfig) getBool(key string) bool {
	return e.v.Get(key) == TRUE
}

func (e *EnvConfig) getInt(key string) int {
	value := e.getString(key)
	i, err := strconv.ParseInt(value, 0, 32)
	if err != nil {
		return int(i)
	}
	return 0
}

func (e *EnvConfig) getDbConfig() DbConfig {
	var cfg DbConfig
	cfg.Driver = e.getString(consts.DB_DRIVER)
	cfg.Database = e.getString(consts.DB_DATABASE)
	cfg.Host = e.getString(consts.DB_HOST)
	cfg.Port = e.getString(consts.DB_PORT)
	cfg.User = e.getString(consts.DB_USER)
	cfg.Password = e.getString(consts.DB_PASSWORD)
	return cfg
}
