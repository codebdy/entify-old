package config

import (
	"strconv"
	"strings"

	"github.com/spf13/viper"
	"rxdrag.com/entify/consts"
)

type EnvConfig struct {
	v      *viper.Viper
	values map[string]interface{}
}

const (
	TRUE  = "true"
	FALSE = "false"
)

func newEnvConfig() *EnvConfig {
	var e EnvConfig
	e.v = viper.New()
	e.v.BindEnv(consts.PARAMS)
	e.v.BindEnv(consts.DB_USER)
	e.v.BindEnv(consts.DB_PASSWORD)
	e.v.BindEnv(consts.DB_HOST)
	e.v.BindEnv(consts.DB_PORT)
	e.v.BindEnv(consts.DB_DATABASE)
	e.v.BindEnv(consts.SERVICE_ID)

	params := e.v.Get(consts.PARAMS)

	if params != nil {
		e.parseParams(params.(string))
	}
	return &e
}

func (e *EnvConfig) parseParams(paramsStr string) {
	items := strings.Split(paramsStr, "&")
	for _, item := range items {
		elements := strings.Split(item, "=")
		if len(elements) > 1 {
			e.values[strings.Trim(elements[0], " ")] = strings.Trim(elements[1], " ")
		}
	}
}

func (e *EnvConfig) getString(key string) string {
	str := e.values[key]
	if str == nil {
		str = e.v.Get(key)
	}

	if str != nil {
		return str.(string)
	}
	return ""
}

func (e *EnvConfig) getBool(key string) bool {
	bl := e.values[key]
	if bl == nil {
		return e.v.Get(key) == TRUE
	}

	return bl == TRUE
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
