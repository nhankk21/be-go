package router

import (
	"godk/core"
	"godk/handler"
)

func SetupRouter(server core.Server) {
	server.SetHandler("/info", handler.GetInfo, core.APIMethodEnum.GET)
	setupEtherRoute(&server)
}
