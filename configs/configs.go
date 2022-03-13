package configs

import (
	"github.com/spf13/viper"
	lib "$1/lib"
)

var default_configs string = lib.CWD() + "/configs/conf.yml"

func InitConf() *viper.Viper {
	vi := viper.New()
	vi.SetConfigFile(default_configs)
	vi.ReadInConfig()
	return vi
}

func GetConfigs() *viper.Viper {
	return InitConf()
}