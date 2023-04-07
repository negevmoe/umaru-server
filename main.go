package main

import (
	"umaru/application/global"
	"umaru/application/handler"
	"umaru/application/repository"
	"umaru/application/server"
)

func main() {
	global.Init()
	handler.Init()
	repository.Init()
	go server.HttpServer()
	select {}
}
