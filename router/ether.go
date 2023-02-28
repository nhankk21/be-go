package router

import (
	"godk/core"
	"godk/handler/ether"
)

func setupEtherRoute(server *core.Server) {
	server.SetHandler("/ether/get_balance", ether.GetEtherBalance, core.APIMethodEnum.GET)
}
