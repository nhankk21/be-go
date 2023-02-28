package ether

import (
	"godk/business/token"
	"godk/core"
)

func GetEtherBalance(req core.Request) *core.Response {
	address := req.GetParam("address")
	if address == "" {
		return &core.Response{
			Status:  core.APIStatusEnum.Invalid,
			Message: "Empty address!",
		}
	}
	ether, err := token.GetTokenBalances(address)
	if err != nil {
		return &core.Response{
			Status:  core.APIStatusEnum.Error,
			Message: err.Error(),
		}
	}

	return &core.Response{
		Status:  core.APIStatusEnum.Ok,
		Message: "Get ether successfully!",
		Data:    ether,
	}
}
