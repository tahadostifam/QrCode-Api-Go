package configs

import (
	"github.com/spf13/viper"
	lib "github.com/tahadostifam/gingo/lib"
)

var default_configs string = lib.CWD() + "/web_server/configs/conf.yml"

func InitConf() *viper.Viper {
	vi := viper.New()
	vi.SetConfigFile(default_configs)
	vi.ReadInConfig()
	return vi
}

func GetConfigs() *viper.Viper {
	return InitConf()
}