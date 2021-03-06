package main

import (
	"log-collector/config"
	"log-collector/infra"
	"log-collector/routes"
)

func main() {
	app := routes.GetEngine()
	app.Run(config.AppHost + ":" + config.AppPort)
}

func init() {
	infra.InitKafka()
}

func closeResource() {
	infra.CloseKafka()
}
