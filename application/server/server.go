package server

import (
	"umaru/application/server/restful"
)

func HttpServer() error {
	return restful.Run()
}
