package main

import (
	"awesomeProject/config"
	"awesomeProject/server/base/db"
	"awesomeProject/server/base/router"
)

var (
	configFilePath = "doc/hltest.yaml"
)

func main() {
	config.InitConfig(configFilePath)
	conf := config.GlobalConfig

	db.InitDB(conf.Repo.Connection)
	router.InitRouter(conf.Server.HttpPort)
}
