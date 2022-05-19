package config

import (
	"github.com/spf13/viper"
	"rxdrag.com/entify/consts"
)

const (
	PATH        = "."
	CONFIG_TYPE = "yaml"
	CONFIG_NAME = "config"
)

type FileConfig struct {
	v *viper.Viper
}

func newFileConfig() *FileConfig {
	var f FileConfig
	f.v = viper.New()
	f.v.SetConfigName(CONFIG_NAME) // name of config file (without extension)
	f.v.SetConfigType(CONFIG_TYPE) // REQUIRED if the config file does not have the extension in the name
	f.v.AddConfigPath(PATH)

	err := f.v.ReadInConfig() // Find and read the config file
	if err != nil {
		panic("Env not set and can not find config file")
	}

	f.v.SetDefault(consts.SERVICE_ID, 1)
	f.v.SetDefault(consts.DB_DRIVER, "mysql")
	return &f
}

func (f *FileConfig) getString(key string) string {
	return f.v.GetString(key)
}

func (f *FileConfig) getBool(key string) bool {
	return f.v.GetBool(key)
}

func (f *FileConfig) getInt(key string) int {
	return f.v.GetInt(key)
}
