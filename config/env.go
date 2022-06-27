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
	e.v.BindEnv(consts.DB_USER)
	e.v.BindEnv(consts.DB_DRIVER)
	e.v.BindEnv(consts.DB_PASSWORD)
	e.v.BindEnv(consts.DB_HOST)
	e.v.BindEnv(consts.DB_PORT)
	e.v.BindEnv(consts.DB_DATABASE)
	e.v.BindEnv(consts.SERVICE_ID)
	e.v.BindEnv(consts.AUTH_URL)
	e.v.BindEnv(consts.UPLOAD)
	return &e
}

func (e *EnvConfig) getString(key string) string {
	str := e.v.Get(key)
	if str != nil {
		return str.(string)
	}
	return ""
}

func (e *EnvConfig) getBool(key string) bool {
	return e.v.Get(key) == TRUE
}

func (e *EnvConfig) getInt(key string) int {
	value := e.getString(key)
	if value == "" {
		return 0
	}
	i, err := strconv.ParseInt(value, 0, 32)
	if err != nil {
		return int(i)
	}
	return 0
}
