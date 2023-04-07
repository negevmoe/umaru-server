package server

import (
	"umaru-server/application/server/restful"
)

func HttpServer() error {
	return restful.Run()
}
