package api

import (
	"github.com/husniramdani/lets-go/api/controllers"
)

var server = controllers.Server{}

func Run() {
	server.Initialize()
	server.Run(":8080")
}
