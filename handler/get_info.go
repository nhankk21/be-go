package handler

import (
	"godk/core"
)

type Raw struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetInfo(req core.Request) *core.Response {
	// data := Raw{}
	// err := req.GetBody(&data)
	// err := common.GetBodyContent(req.Request.Body, &data)
	return &core.Response{
		Status:  core.APIStatusEnum.Ok,
		Message: "Get info successfully!, OKE",
		Data:    "Get more info later!",
	}
}
