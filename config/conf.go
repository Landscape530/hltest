package config

import (
	"github.com/spf13/viper"
)

var GlobalConfig *Conf

type Conf struct {
	Val    string
	Repo   *Repo
	Server *Server
}

type Server struct {
	HttpPort   string
	ServerAddr string
	Env        string
}

type Repo struct {
	Connection string
}

func InitConfig(path string) {
	configVip := viper.New()
	configVip.SetConfigFile(path)

	// 读取配置
	if err := configVip.ReadInConfig(); err != nil {
		panic(err)
	}

	// 配置映射到结构体
	GlobalConfig = &Conf{}
	if err := configVip.Unmarshal(GlobalConfig); err != nil {
		panic(err)
	}

}
