package api

import (
	"lets-go/api/controllers"
)

var server = controllers.Server{}

func Run() {
	server.Initialize()
	server.Run(":8090")
}
