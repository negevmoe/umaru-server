package main

import (
	"umaru-server/application/global"
	"umaru-server/application/handler"
	"umaru-server/application/repository"
	"umaru-server/application/server"
)

func main() {
	global.Init()
	handler.Init()
	repository.Init()
	go server.HttpServer()
	select {}
}
